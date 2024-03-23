package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/fasterci/rules_gitops/testing/it_sidecar/client"
)

var setup client.K8STestSetup

func TestMain(m *testing.M) {
	setup = client.K8STestSetup{
		PortForwardServices: map[string]int{
			"myapp": 80,
		},
	}
	setup.TestMain(m)
}

func TestSimpleServer(t *testing.T) {
	appServerPort := setup.GetServiceLocalPort("myapp")
	response, err := http.Get(fmt.Sprintf("http://localhost:%d", appServerPort))
	if err != nil {
		t.FailNow()
	}
	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}
}
