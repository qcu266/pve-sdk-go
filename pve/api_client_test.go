package pve

import (
	"context"
	"crypto/tls"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	auth, err := AuthOptionsFromEnv()
	if err != nil {
		t.Error(err)
	}

	client := New(auth,
		WithTLSOptions(&tls.Config{
			InsecureSkipVerify: true,
		}),
	)

	response, err := client.VersionVersion(context.TODO(), nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(response)
}

func TestClientListVm(t *testing.T) {
	auth, err := AuthOptionsFromEnv()
	if err != nil {
		t.Error(err)
	}

	client := New(auth,
		WithTLSOptions(&tls.Config{
			InsecureSkipVerify: true,
		}),
	)

	// x := "qemu"
	response, err := client.ClusterResourcesResources(context.TODO(), &ClusterResourcesResourcesRequest{
		// Type: ,
	})
	x := len(*response)
	fmt.Println(len(*response))
	fmt.Println(x)
	if err != nil {
		t.Error(err)
	}
}
