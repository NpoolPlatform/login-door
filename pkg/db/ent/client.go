// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/login-door/pkg/db/ent/migrate"

	"github.com/NpoolPlatform/login-door/pkg/db/ent/empty"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Empty is the client for interacting with the Empty builders.
	Empty *EmptyClient
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
	c.Empty = NewEmptyClient(c.config)
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
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Empty:  NewEmptyClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config: cfg,
		Empty:  NewEmptyClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Empty.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
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
	c.Empty.Use(hooks...)
}

// EmptyClient is a client for the Empty schema.
type EmptyClient struct {
	config
}

// NewEmptyClient returns a client for the Empty from the given config.
func NewEmptyClient(c config) *EmptyClient {
	return &EmptyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `empty.Hooks(f(g(h())))`.
func (c *EmptyClient) Use(hooks ...Hook) {
	c.hooks.Empty = append(c.hooks.Empty, hooks...)
}

// Create returns a create builder for Empty.
func (c *EmptyClient) Create() *EmptyCreate {
	mutation := newEmptyMutation(c.config, OpCreate)
	return &EmptyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Empty entities.
func (c *EmptyClient) CreateBulk(builders ...*EmptyCreate) *EmptyCreateBulk {
	return &EmptyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Empty.
func (c *EmptyClient) Update() *EmptyUpdate {
	mutation := newEmptyMutation(c.config, OpUpdate)
	return &EmptyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EmptyClient) UpdateOne(e *Empty) *EmptyUpdateOne {
	mutation := newEmptyMutation(c.config, OpUpdateOne, withEmpty(e))
	return &EmptyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EmptyClient) UpdateOneID(id int) *EmptyUpdateOne {
	mutation := newEmptyMutation(c.config, OpUpdateOne, withEmptyID(id))
	return &EmptyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Empty.
func (c *EmptyClient) Delete() *EmptyDelete {
	mutation := newEmptyMutation(c.config, OpDelete)
	return &EmptyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *EmptyClient) DeleteOne(e *Empty) *EmptyDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *EmptyClient) DeleteOneID(id int) *EmptyDeleteOne {
	builder := c.Delete().Where(empty.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EmptyDeleteOne{builder}
}

// Query returns a query builder for Empty.
func (c *EmptyClient) Query() *EmptyQuery {
	return &EmptyQuery{
		config: c.config,
	}
}

// Get returns a Empty entity by its id.
func (c *EmptyClient) Get(ctx context.Context, id int) (*Empty, error) {
	return c.Query().Where(empty.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EmptyClient) GetX(ctx context.Context, id int) *Empty {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *EmptyClient) Hooks() []Hook {
	return c.hooks.Empty
}
