package tests

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/name/resourcemodel"
)

// MarshalVyos calls greetings.Hello with a name, checking
// for a valid return value.
func TestFirewallNameMarshalVyos(t *testing.T) {
	model := &resourcemodel.FirewallName{
		ID:                            basetypes.NewStringValue("test-id"),
		LeafFirewallNameDefaultAction: basetypes.NewStringValue("drop"),
	}

	want, err := json.Marshal(map[string]interface{}{
		"default-action": "drop",
	})
	if err != nil {
		t.Fatalf(`desired value can not be jsonified: %v`, err)
	}

	result, err := helpers.MarshalVyos(context.Background(), model)
	if !reflect.DeepEqual(result, want) {
		t.Fatalf(`model.MarshalVyos() = %q, %v, want match for %#q, nil`, result, err, want)
	}
}
