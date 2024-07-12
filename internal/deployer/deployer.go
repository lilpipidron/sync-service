package deployer

type Deployer interface {
	CreatePod(name string) error
	DeletePod(name string) error
	GetPodList() ([]string, error)
}

type PodDeployer struct{}
