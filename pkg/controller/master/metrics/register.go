package metrics

import (
	"context"

	"github.com/harvester/harvester/pkg/config"
)

func Register(ctx context.Context, management *config.Management, options config.Options) error {
	controller, err := NewMetricsController("")
	if err != nil {
		return err
	}

	go controller.Start(ctx)

	return nil
}
