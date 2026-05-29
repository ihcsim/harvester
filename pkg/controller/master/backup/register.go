package backup

import (
	"context"

	"github.com/harvester/harvester/pkg/config"
)

func Register(ctx context.Context, management *config.Management, options config.Options) error {
	controller, err := NewBackupController()
	if err != nil {
		return err
	}

	go func() {
		if err := controller.Run(ctx); err != nil {
			panic(err)
		}
	}()

	return nil
}
