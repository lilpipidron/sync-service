package deployer_tests

import (
	"github.com/lilpipidron/sync-service/internal/deployer"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestGetPodList(t *testing.T) {
	clientset := fake.NewSimpleClientset(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod-1",
			Namespace: "default",
		},
	}, &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod-2",
			Namespace: "default",
		},
	})
	namespace := "default"
	podDeployer := &deployer.PodDeployer{
		Client:    clientset,
		Namespace: namespace,
	}

	podList, err := podDeployer.GetPodList()
	assert.NoError(t, err)
	assert.Len(t, podList, 2)
	assert.Contains(t, podList, "test-pod-1")
	assert.Contains(t, podList, "test-pod-2")
}
