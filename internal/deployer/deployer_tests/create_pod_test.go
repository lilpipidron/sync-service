package deployer_tests

import (
	"context"
	"github.com/lilpipidron/sync-service/internal/deployer"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestCreatePod(t *testing.T) {
	clientset := fake.NewSimpleClientset()
	namespace := "default"
	podDeployer := &deployer.PodDeployer{
		Client:    clientset,
		Namespace: namespace,
	}

	err := podDeployer.CreatePod("test-pod")
	assert.NoError(t, err)

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	assert.NoError(t, err)
	assert.Len(t, pods.Items, 1)
	assert.Equal(t, "test-pod", pods.Items[0].Name)
}
