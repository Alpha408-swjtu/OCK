package test

import (
	"context"
	"custom_controller/client"
	"fmt"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestCrd(t *testing.T) {
	crdClient, err := client.GetCrdClient()
	if err != nil {
		t.Fail()
	}

	s := schema.GroupVersionResource{
		Group:    "stable.example.com",
		Version:  "v1beta1",
		Resource: "redises",
	}

	result, _ := crdClient.Resource(s).Namespace("default").Get(context.TODO(), "redis-cluster", metav1.GetOptions{})
	fmt.Println(result.Object)
}
