package kube

import (
	"fmt"

	agentdevsyclient "github.com/devsy-org/agentapi/pkg/clientset/versioned"
	devsyclient "github.com/devsy-org/api/pkg/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Interface interface {
	kubernetes.Interface
	Devsy() devsyclient.Interface
	Agent() agentdevsyclient.Interface
}

func NewForConfig(c *rest.Config) (Interface, error) {
	kubeClient, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, fmt.Errorf("create kube client: %w", err)
	}

	devsyClient, err := devsyclient.NewForConfig(c)
	if err != nil {
		return nil, fmt.Errorf("create devsy client: %w", err)
	}

	agentDevsyClient, err := agentdevsyclient.NewForConfig(c)
	if err != nil {
		return nil, fmt.Errorf("create agent client: %w", err)
	}

	return &client{
		Interface:        kubeClient,
		devsyClient:      devsyClient,
		agentDevsyClient: agentDevsyClient,
	}, nil
}

type client struct {
	kubernetes.Interface
	devsyClient      devsyclient.Interface
	agentDevsyClient agentdevsyclient.Interface
}

func (c *client) Devsy() devsyclient.Interface {
	return c.devsyClient
}

func (c *client) Agent() agentdevsyclient.Interface {
	return c.agentDevsyClient
}
