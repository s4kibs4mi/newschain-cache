// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/s4kibs4mi/newschain-cache/ent/migrate"

	"github.com/s4kibs4mi/newschain-cache/ent/latestblock"
	"github.com/s4kibs4mi/newschain-cache/ent/post"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// LatestBlock is the client for interacting with the LatestBlock builders.
	LatestBlock *LatestBlockClient
	// Post is the client for interacting with the Post builders.
	Post *PostClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.LatestBlock = NewLatestBlockClient(c.config)
	c.Post = NewPostClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		LatestBlock: NewLatestBlockClient(cfg),
		Post:        NewPostClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:      cfg,
		LatestBlock: NewLatestBlockClient(cfg),
		Post:        NewPostClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		LatestBlock.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.LatestBlock.Use(hooks...)
	c.Post.Use(hooks...)
}

// LatestBlockClient is a client for the LatestBlock schema.
type LatestBlockClient struct {
	config
}

// NewLatestBlockClient returns a client for the LatestBlock from the given config.
func NewLatestBlockClient(c config) *LatestBlockClient {
	return &LatestBlockClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `latestblock.Hooks(f(g(h())))`.
func (c *LatestBlockClient) Use(hooks ...Hook) {
	c.hooks.LatestBlock = append(c.hooks.LatestBlock, hooks...)
}

// Create returns a create builder for LatestBlock.
func (c *LatestBlockClient) Create() *LatestBlockCreate {
	mutation := newLatestBlockMutation(c.config, OpCreate)
	return &LatestBlockCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of LatestBlock entities.
func (c *LatestBlockClient) CreateBulk(builders ...*LatestBlockCreate) *LatestBlockCreateBulk {
	return &LatestBlockCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for LatestBlock.
func (c *LatestBlockClient) Update() *LatestBlockUpdate {
	mutation := newLatestBlockMutation(c.config, OpUpdate)
	return &LatestBlockUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LatestBlockClient) UpdateOne(lb *LatestBlock) *LatestBlockUpdateOne {
	mutation := newLatestBlockMutation(c.config, OpUpdateOne, withLatestBlock(lb))
	return &LatestBlockUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LatestBlockClient) UpdateOneID(id int) *LatestBlockUpdateOne {
	mutation := newLatestBlockMutation(c.config, OpUpdateOne, withLatestBlockID(id))
	return &LatestBlockUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for LatestBlock.
func (c *LatestBlockClient) Delete() *LatestBlockDelete {
	mutation := newLatestBlockMutation(c.config, OpDelete)
	return &LatestBlockDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *LatestBlockClient) DeleteOne(lb *LatestBlock) *LatestBlockDeleteOne {
	return c.DeleteOneID(lb.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *LatestBlockClient) DeleteOneID(id int) *LatestBlockDeleteOne {
	builder := c.Delete().Where(latestblock.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LatestBlockDeleteOne{builder}
}

// Query returns a query builder for LatestBlock.
func (c *LatestBlockClient) Query() *LatestBlockQuery {
	return &LatestBlockQuery{config: c.config}
}

// Get returns a LatestBlock entity by its id.
func (c *LatestBlockClient) Get(ctx context.Context, id int) (*LatestBlock, error) {
	return c.Query().Where(latestblock.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LatestBlockClient) GetX(ctx context.Context, id int) *LatestBlock {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LatestBlockClient) Hooks() []Hook {
	return c.hooks.LatestBlock
}

// PostClient is a client for the Post schema.
type PostClient struct {
	config
}

// NewPostClient returns a client for the Post from the given config.
func NewPostClient(c config) *PostClient {
	return &PostClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `post.Hooks(f(g(h())))`.
func (c *PostClient) Use(hooks ...Hook) {
	c.hooks.Post = append(c.hooks.Post, hooks...)
}

// Create returns a create builder for Post.
func (c *PostClient) Create() *PostCreate {
	mutation := newPostMutation(c.config, OpCreate)
	return &PostCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Post entities.
func (c *PostClient) CreateBulk(builders ...*PostCreate) *PostCreateBulk {
	return &PostCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Post.
func (c *PostClient) Update() *PostUpdate {
	mutation := newPostMutation(c.config, OpUpdate)
	return &PostUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PostClient) UpdateOne(po *Post) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPost(po))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PostClient) UpdateOneID(id uuid.UUID) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPostID(id))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Post.
func (c *PostClient) Delete() *PostDelete {
	mutation := newPostMutation(c.config, OpDelete)
	return &PostDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PostClient) DeleteOne(po *Post) *PostDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PostClient) DeleteOneID(id uuid.UUID) *PostDeleteOne {
	builder := c.Delete().Where(post.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PostDeleteOne{builder}
}

// Query returns a query builder for Post.
func (c *PostClient) Query() *PostQuery {
	return &PostQuery{config: c.config}
}

// Get returns a Post entity by its id.
func (c *PostClient) Get(ctx context.Context, id uuid.UUID) (*Post, error) {
	return c.Query().Where(post.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PostClient) GetX(ctx context.Context, id uuid.UUID) *Post {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PostClient) Hooks() []Hook {
	return c.hooks.Post
}
