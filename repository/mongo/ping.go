package mongo

import (
	"context"
	"examservice/utils"
)

func (r *Repository) PingDB(ctx context.Context) error {
	err := r.conn.Ping(ctx, nil)
	if err != nil {
		return utils.NewInternalServerError("Failed to ping MongoDB")
	}
	return nil
}
