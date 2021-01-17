package data

import (
	"context"
	"github.com/s4kibs4mi/newschain-cache/ent"
)

func GetLatestBlock(c *ent.LatestBlockClient) (*ent.LatestBlock, error) {
	return c.
		Query().
		Only(context.Background())
}
