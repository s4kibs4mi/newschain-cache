package data

import (
	"context"
	"github.com/google/uuid"
	"github.com/s4kibs4mi/newschain-cache/ent"
	"github.com/s4kibs4mi/newschain-cache/ent/post"
)

func ListEvents(c *ent.PostClient, from, limit int) ([]*ent.Post, error) {
	return c.
		Query().
		Order(ent.Desc(post.FieldCreatedAt)).
		Offset(from).
		Limit(limit).
		All(context.Background())
}

func SearchEvents(c *ent.PostClient, query string, from, limit int) ([]*ent.Post, error) {
	return c.
		Query().
		Where(
			post.TitleContains(query),
		).
		Order(ent.Desc(post.FieldCreatedAt)).
		Offset(from).
		Limit(limit).
		All(context.Background())
}

func GetPostByID(c *ent.PostClient, ID uuid.UUID) (*ent.Post, error) {
	return c.
		Query().
		Where(
			post.IDEQ(ID),
		).
		Only(context.Background())
}
