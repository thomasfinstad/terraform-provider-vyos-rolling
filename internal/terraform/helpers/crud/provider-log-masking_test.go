package crud

import (
	"bufio"
	"context"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflogtest"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"
	ipv4ResModel "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/firewall/ipv4-name/resourcemodel"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/tests/api"
)

// TestProviderLogApiKeyMasking tests to ensure api key does not leak in logs
func TestProviderLogApiKeyMasking(t *testing.T) {
	endpoint := "localhost:50009"
	apiKey := `my-test-api-key`

	logReader, logWriter, err := os.Pipe()
	defer logWriter.Close()
	defer logReader.Close()
	logScanner := bufio.NewScanner(logReader)
	if err != nil {
		t.Fatal("Could not create log r/w:", err)
	}
	ctx := tflogtest.RootLogger(context.Background(), logWriter)

	ctxMutilators := data.CtxMutilators(endpoint, apiKey)

	// Run ctx mutilators for client
	for _, fn := range ctxMutilators {
		ctx = fn(ctx)
	}

	// Client configuration for data sources and resources
	config := data.NewProviderData(
		client.NewClient(ctx, endpoint, apiKey, "TODO: add useragent with provider version", true),
	)

	// Add ctx mutilators to provider data
	config.CtxMutilatorAdd(ctxMutilators...)

	// ---------------------------------
	//  Test:
	// ---------------------------------

	// API mocking
	eList := api.NewExchangeList()

	// Delete resource API call
	exchangeCreateResource := eList.Add()
	exchangeCreateResource.Expect(
		"/configure",
		apiKey,
		`[`+
			`{"op":"delete","path":["firewall","ipv4","name","TestProviderLogApiKeyMasking"]}`+
			`]`,
	).Response(
		200,
		`{
				"success": true,
				"data": null,
				"error": null
			}`,
	)

	// From resource model
	model := &ipv4ResModel.FirewallIPvfourName{
		// SelfIdentifier:                       basetypes.NewStringValue("TestProviderLogApiKeyMasking"),
		SelfIdentifier: basetypes.NewObjectValueMust(
			map[string]attr.Type{"name": basetypes.StringType{}},
			map[string]attr.Value{"name": basetypes.NewStringValue("TestProviderLogApiKeyMasking")}),

		LeafFirewallIPvfourNameDefaultAction: basetypes.NewStringValue("reject"),
		LeafFirewallIPvfourNameDefaultLog:    basetypes.NewBoolValue(true),
		LeafFirewallIPvfourNameDescrIPtion:   basetypes.NewStringValue("Managed by terraform"),
	}

	// Server
	srv := &http.Server{
		Addr: endpoint,
	}
	api.Server(srv, eList)

	// Client
	client := client.NewClient(ctx, "http://"+endpoint, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)
	providerData.Config.CrudSkipCheckChildBeforeDelete = true

	// Execute test
	err = delete(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}
	logWriter.Close()

	// // Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
	}

	reBad := regexp.MustCompile(apiKey)

	for logScanner.Scan() {

		l := logScanner.Text()
		if reBad.MatchString(l) {
			t.Error("Found matching string that is supposed to be masked:\n", l)
		}

		if strings.Contains(l, `***`) {
			t.Log("Found masked output in:", l)
		}
	}
}
