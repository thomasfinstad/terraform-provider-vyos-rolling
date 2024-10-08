package helpers

import (
	"context"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/go-test/deep"
	"github.com/hashicorp/terraform-plugin-log/tflogtest"
)

// TestGenerateVyosOpsMultiLvl check that we can iron out differing levels of config
func TestGenerateVyosOpsMultiLvl(t *testing.T) {
	hasPath := []string{"firewall", "zone", "TF-Examples"}
	hasData := map[string]any{
		"intra-zone-filtering": map[string]any{
			"action": "accept",
			"firewall": map[string]any{
				"name": "test"}}}

	want := [][]string{
		{"firewall", "zone", "TF-Examples", "intra-zone-filtering", "action", "accept"},
		{"firewall", "zone", "TF-Examples", "intra-zone-filtering", "firewall", "name", "test"},
	}

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	got := GenerateVyosOps(ctx, hasPath, hasData)

	// Sort results
	sort.Slice(want, func(i, j int) bool {
		return strings.Join(want[i], "") < strings.Join(want[j], "")
	})
	sort.Slice(got, func(i, j int) bool {
		return strings.Join(got[i], "") < strings.Join(got[j], "")
	})

	// Diff results
	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("Want: %v", want)
		t.Errorf("Got: %v", got)
		t.Errorf("compare failed: %v", diff)
	}
}

// TestGenerateVyosOpsMultiValue check that we can iron out multi value configs
func TestGenerateVyosOpsMultiValue(t *testing.T) {
	hasPath := []string{"service", "ntp", "allow-client"}
	hasData := map[string]any{
		"address": []string{"127.0.0.0/8", "169.254.0.0/16", "10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16", "::1/128", "fe80::/10", "fc00::/7"}}

	want := [][]string{
		{"service", "ntp", "allow-client", "address", "127.0.0.0/8"},
		{"service", "ntp", "allow-client", "address", "10.0.0.0/8"},
		{"service", "ntp", "allow-client", "address", "172.16.0.0/12"},
		{"service", "ntp", "allow-client", "address", "169.254.0.0/16"},
		{"service", "ntp", "allow-client", "address", "192.168.0.0/16"},
		{"service", "ntp", "allow-client", "address", "::1/128"},
		{"service", "ntp", "allow-client", "address", "fe80::/10"},
		{"service", "ntp", "allow-client", "address", "fc00::/7"},
	}

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	got := GenerateVyosOps(ctx, hasPath, hasData)

	// Sort results
	sort.Slice(want, func(i, j int) bool {
		return strings.Join(want[i], "") < strings.Join(want[j], "")
	})
	sort.Slice(got, func(i, j int) bool {
		return strings.Join(got[i], "") < strings.Join(got[j], "")
	})

	if len(got) != len(want) {
		t.Fatalf("Wanted result of length: %d, got: %d", len(want), len(got))
	}

	// Diff results
	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("Want: %v", want)
		t.Errorf("Got: %v", got)
		t.Errorf("compare failed: %v", diff)
	}
}
