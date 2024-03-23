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

// TestGenerateVyosOps generates vyos ops to prevent a bug with ironing out the parameters
func TestGenerateVyosOps(t *testing.T) {
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
