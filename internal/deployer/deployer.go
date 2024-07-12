package deployer

import (
	"context"
	"github.com/lilpipidron/sync-service/internal/config"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Deployer interface {
	CreatePod(name string) error
	DeletePod(name string) error
	GetPodList() ([]string, error)
}

type PodDeployer struct {
	Client    kubernetes.Interface
	Namespace string
}

func NewPodDeployer(cfg *config.Config) *PodDeployer {
	kubernetesConfig, err := rest.InClusterConfig()
	if err != nil {
		kubernetesConfig, err = clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
		if err != nil {
			panic(err)
		}
	}

	clientset, err := kubernetes.NewForConfig(kubernetesConfig)
	if err != nil {
		panic(err)
	}

	return &PodDeployer{
		Client:    clientset,
		Namespace: cfg.Namespace,
	}
}

// CreatePod creates a new pod in Kubernetes with the specified name.
// It configures the pod with a basic container running a specified image,
// and sets resource limits for CPU and memory.
// Returns an error if pod creation fails.
func (p *PodDeployer) CreatePod(name string) error {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  name,
					Image: "golang:latest",
					Resources: corev1.ResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("50m"),
							corev1.ResourceMemory: resource.MustParse("512Mi"),
						},
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("100m"),
							corev1.ResourceMemory: resource.MustParse("1024Mi"),
						},
					},
				},
			},
		},
	}

	_, err := p.Client.CoreV1().Pods(p.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	return err
}

// DeletePod deletes a pod with the specified name from Kubernetes.
// Returns an error if pod deletion fails.
func (p *PodDeployer) DeletePod(name string) error {
	return p.Client.CoreV1().Pods(p.Namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

// GetPodList retrieves a list of pod names currently running in the specified namespace in Kubernetes.
// Returns a slice of pod names and an error if retrieval fails.
func (p *PodDeployer) GetPodList() ([]string, error) {
	pods, err := p.Client.CoreV1().Pods(p.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	podNames := make([]string, len(pods.Items))
	for i, pod := range pods.Items {
		podNames[i] = pod.Name
	}
	return podNames, nil
}
