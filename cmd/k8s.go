package cmd

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func initK8sClient(kubeconfig string) (*kubernetes.Clientset, error) {
	var config *rest.Config

	if kubeconfig == "" {
		config, _ = initInClusterClient()
	} else {
		config, _ = initOutOfClusterClient(kubeconfig)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func initInClusterClient() (*rest.Config, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}

func initOutOfClusterClient(kubeconfig string) (*rest.Config, error) {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	return config, nil
}
