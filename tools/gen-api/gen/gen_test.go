package gen

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/iancoleman/strcase"
)

func TestGenApiName(t *testing.T) {

	tests := []struct {
		path   string
		name   string
		expect string
	}{
		{"/cluster", "index", "ClusterIndex"},
		{"/nodes/{node}/qemu/{vmid}/firewall/rules/{pos}", "update_rule", "NodesQemuFirewallRulesUpdateRule"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("path: %s, name: %s", test.path, test.name), func(t *testing.T) {
			actual := genApiName(test.path, test.name)

			if actual != test.expect {
				t.Errorf("except: %s, actual: %s", test.expect, actual)
			}
		})
	}

}

func TestToCamelTemplateFunc(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"sparse", "Sparse"},
		{"encryption-key", "EncryptionKey"},
		{"lio_tpg", "LioTpg"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("input: %s", test.input), func(t *testing.T) {
			actual := strcase.ToCamel(test.input)

			if actual != test.expect {
				t.Errorf("except: %s, actual: %s", test.expect, actual)
			}
		})
	}
}

func TestTpl(t *testing.T) {

	tp, _ := tpl.New("aaaa").Parse("{{ .Name | toCamel  }}")

	var buf bytes.Buffer

	tp.Execute(&buf, struct{ Name string }{"abc"})

	fmt.Println(buf.Bytes())

}
