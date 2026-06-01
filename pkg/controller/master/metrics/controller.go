package metrics

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/apis/meta/v1 as metav1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type MetricsController struct {
	clientset *kubernetes.Clientset
	namespace string
}

func NewMetricsController(kubeconfig string) (*MetricsController, error) {
	var config *rest.Config
	var err error

	if kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		config, err = rest.InClusterConfig()
	}

	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &MetricsController{
		clientset: clientset,
		namespace: "harvester-system",
	}, nil
}

func (c *MetricsController) Start(ctx context.Context) error {
	logrus.Info("Starting metrics controller")

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := c.collectMetrics(); err != nil {
				logrus.Errorf("Failed to collect metrics: %v", err)
			}
		case <-ctx.Done():
			logrus.Info("Stopping metrics controller")
			return nil
		}
	}
}

func (c *MetricsController) collectMetrics() error {
	pods, err := c.clientset.CoreV1().Pods(c.namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	logrus.Infof("Collected metrics for %d pods", len(pods.Items))
	return nil
}
