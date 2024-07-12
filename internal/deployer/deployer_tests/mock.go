package deployer_tests

import "github.com/stretchr/testify/mock"

type MockDeployer struct {
	mock.Mock
}

func (m *MockDeployer) CreatePod(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *MockDeployer) DeletePod(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *MockDeployer) GetPodList() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}
