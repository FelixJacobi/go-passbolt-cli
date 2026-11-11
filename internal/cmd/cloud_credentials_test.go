package cmd

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/haspr-dev/go-haspr-cli/internal/util"
	"github.com/haspr-dev/go-haspr/api"
	"github.com/spf13/viper"
)

func TestCloudValidationPrecedesAuthentication(t *testing.T) {
	old := util.ClientProvider
	util.ClientProvider = func(context.Context) (*api.Client, error) { t.Fatal("invalid request authenticated"); return nil, nil }
	defer func() { util.ClientProvider = old; viper.Reset(); newCommandTree() }()
	for _, test := range []struct{ protocol, id, info string }{
		{"aws", "bad-id", ""}, {"k8s", "bad-id", ""},
		{"k8s", "01234567-89ab-cdef-0123-456789abcdef", `{"apiVersion":"client.authentication.k8s.io/v1","kind":"ExecCredential","spec":{"interactive":false}}`},
		{"k8s", "01234567-89ab-cdef-0123-456789abcdef", `{"apiVersion":"wrong","kind":"ExecCredential"}`},
	} {
		t.Setenv("KUBERNETES_EXEC_INFO", test.info)
		viper.Reset()
		root := newCommandTree()
		viper.Set("agent", true)
		root.SetArgs([]string{"credential-helper", test.protocol, "get", "--id", test.id})
		if err := root.Execute(); err == nil {
			t.Fatal("invalid cloud request succeeded")
		}
	}
}
func TestHelperExecutableArguments(t *testing.T) {
	for _, brand := range []string{"haspr", "pbolt", "passbolt"} {
		for _, protocol := range []string{"git", "docker"} {
			got := HelperArguments(protocol+"-credential-"+brand+".exe", []string{"get"})
			if !reflect.DeepEqual(got, []string{"credential-helper", protocol, "get"}) {
				t.Fatalf("brand %s protocol %s: %v", brand, protocol, got)
			}
		}
	}
	if got := HelperArguments("haspr", []string{"get", "resource"}); strings.Join(got, " ") != "get resource" {
		t.Fatal(got)
	}
	if got := HelperArguments("passbolt", []string{"get", "resource"}); strings.Join(got, " ") != "get resource" {
		t.Fatal(got)
	}
}
