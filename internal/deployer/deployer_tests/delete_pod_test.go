package deployer_tests

import (
	"context"
	"github.com/lilpipidron/sync-service/internal/deployer"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestDeletePod(t *testing.T) {
	clientset := fake.NewSimpleClientset(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod",
			Namespace: "default",
		},
	})
	namespace := "default"
	podDeployer := &deployer.PodDeployer{
		Client:    clientset,
		Namespace: namespace,
	}

	err := podDeployer.DeletePod("test-pod")
	assert.NoError(t, err)

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	assert.NoError(t, err)
	assert.Len(t, pods.Items, 0)
}
