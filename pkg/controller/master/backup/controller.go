package backup

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type BackupController struct {
	clientset *kubernetes.Clientset
	interval  time.Duration
}

func NewBackupController() (*BackupController, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &BackupController{
		clientset: clientset,
		interval:  5 * time.Minute,
	}, nil
}

func (c *BackupController) Run(ctx context.Context) error {
	logrus.Info("Starting backup controller")

	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := c.processBackups(); err != nil {
				logrus.Errorf("Failed to process backups: %v", err)
			}
		case <-ctx.Done():
			logrus.Info("Stopping backup controller")
			return nil
		}
	}
}

func (c *BackupController) processBackups() error {
	pvcs, err := c.clientset.CoreV1().PersistentVolumeClaims("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	logrus.Infof("Processing %d PVCs for backup", len(pvcs.Items))

	for _, pvc := range pvcs.Items {
		logrus.Infof("Backing up PVC: %s/%s", pvc.Namespace, pvc.Name)
	}

	return nil
}
