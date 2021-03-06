// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/wipha26/app/ent/migrate"

	"github.com/wipha26/app/ent/dispense"
	"github.com/wipha26/app/ent/drug"
	"github.com/wipha26/app/ent/patient"
	"github.com/wipha26/app/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Dispense is the client for interacting with the Dispense builders.
	Dispense *DispenseClient
	// Drug is the client for interacting with the Drug builders.
	Drug *DrugClient
	// Patient is the client for interacting with the Patient builders.
	Patient *PatientClient
	// User is the client for interacting with the User builders.
	User *UserClient
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
	c.Dispense = NewDispenseClient(c.config)
	c.Drug = NewDrugClient(c.config)
	c.Patient = NewPatientClient(c.config)
	c.User = NewUserClient(c.config)
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
		ctx:      ctx,
		config:   cfg,
		Dispense: NewDispenseClient(cfg),
		Drug:     NewDrugClient(cfg),
		Patient:  NewPatientClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
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
		config:   cfg,
		Dispense: NewDispenseClient(cfg),
		Drug:     NewDrugClient(cfg),
		Patient:  NewPatientClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Dispense.
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
	c.Dispense.Use(hooks...)
	c.Drug.Use(hooks...)
	c.Patient.Use(hooks...)
	c.User.Use(hooks...)
}

// DispenseClient is a client for the Dispense schema.
type DispenseClient struct {
	config
}

// NewDispenseClient returns a client for the Dispense from the given config.
func NewDispenseClient(c config) *DispenseClient {
	return &DispenseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dispense.Hooks(f(g(h())))`.
func (c *DispenseClient) Use(hooks ...Hook) {
	c.hooks.Dispense = append(c.hooks.Dispense, hooks...)
}

// Create returns a create builder for Dispense.
func (c *DispenseClient) Create() *DispenseCreate {
	mutation := newDispenseMutation(c.config, OpCreate)
	return &DispenseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Dispense.
func (c *DispenseClient) Update() *DispenseUpdate {
	mutation := newDispenseMutation(c.config, OpUpdate)
	return &DispenseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DispenseClient) UpdateOne(d *Dispense) *DispenseUpdateOne {
	mutation := newDispenseMutation(c.config, OpUpdateOne, withDispense(d))
	return &DispenseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DispenseClient) UpdateOneID(id int) *DispenseUpdateOne {
	mutation := newDispenseMutation(c.config, OpUpdateOne, withDispenseID(id))
	return &DispenseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Dispense.
func (c *DispenseClient) Delete() *DispenseDelete {
	mutation := newDispenseMutation(c.config, OpDelete)
	return &DispenseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DispenseClient) DeleteOne(d *Dispense) *DispenseDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DispenseClient) DeleteOneID(id int) *DispenseDeleteOne {
	builder := c.Delete().Where(dispense.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DispenseDeleteOne{builder}
}

// Create returns a query builder for Dispense.
func (c *DispenseClient) Query() *DispenseQuery {
	return &DispenseQuery{config: c.config}
}

// Get returns a Dispense entity by its id.
func (c *DispenseClient) Get(ctx context.Context, id int) (*Dispense, error) {
	return c.Query().Where(dispense.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DispenseClient) GetX(ctx context.Context, id int) *Dispense {
	d, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return d
}

// QueryDrug queries the drug edge of a Dispense.
func (c *DispenseClient) QueryDrug(d *Dispense) *DrugQuery {
	query := &DrugQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dispense.Table, dispense.FieldID, id),
			sqlgraph.To(drug.Table, drug.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dispense.DrugTable, dispense.DrugColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPatient queries the patient edge of a Dispense.
func (c *DispenseClient) QueryPatient(d *Dispense) *PatientQuery {
	query := &PatientQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dispense.Table, dispense.FieldID, id),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dispense.PatientTable, dispense.PatientColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Dispense.
func (c *DispenseClient) QueryUser(d *Dispense) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dispense.Table, dispense.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dispense.UserTable, dispense.UserColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DispenseClient) Hooks() []Hook {
	return c.hooks.Dispense
}

// DrugClient is a client for the Drug schema.
type DrugClient struct {
	config
}

// NewDrugClient returns a client for the Drug from the given config.
func NewDrugClient(c config) *DrugClient {
	return &DrugClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `drug.Hooks(f(g(h())))`.
func (c *DrugClient) Use(hooks ...Hook) {
	c.hooks.Drug = append(c.hooks.Drug, hooks...)
}

// Create returns a create builder for Drug.
func (c *DrugClient) Create() *DrugCreate {
	mutation := newDrugMutation(c.config, OpCreate)
	return &DrugCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Drug.
func (c *DrugClient) Update() *DrugUpdate {
	mutation := newDrugMutation(c.config, OpUpdate)
	return &DrugUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DrugClient) UpdateOne(d *Drug) *DrugUpdateOne {
	mutation := newDrugMutation(c.config, OpUpdateOne, withDrug(d))
	return &DrugUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DrugClient) UpdateOneID(id int) *DrugUpdateOne {
	mutation := newDrugMutation(c.config, OpUpdateOne, withDrugID(id))
	return &DrugUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Drug.
func (c *DrugClient) Delete() *DrugDelete {
	mutation := newDrugMutation(c.config, OpDelete)
	return &DrugDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DrugClient) DeleteOne(d *Drug) *DrugDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DrugClient) DeleteOneID(id int) *DrugDeleteOne {
	builder := c.Delete().Where(drug.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DrugDeleteOne{builder}
}

// Create returns a query builder for Drug.
func (c *DrugClient) Query() *DrugQuery {
	return &DrugQuery{config: c.config}
}

// Get returns a Drug entity by its id.
func (c *DrugClient) Get(ctx context.Context, id int) (*Drug, error) {
	return c.Query().Where(drug.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DrugClient) GetX(ctx context.Context, id int) *Drug {
	d, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return d
}

// QueryDispenses queries the dispenses edge of a Drug.
func (c *DrugClient) QueryDispenses(d *Drug) *DispenseQuery {
	query := &DispenseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(drug.Table, drug.FieldID, id),
			sqlgraph.To(dispense.Table, dispense.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, drug.DispensesTable, drug.DispensesColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DrugClient) Hooks() []Hook {
	return c.hooks.Drug
}

// PatientClient is a client for the Patient schema.
type PatientClient struct {
	config
}

// NewPatientClient returns a client for the Patient from the given config.
func NewPatientClient(c config) *PatientClient {
	return &PatientClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `patient.Hooks(f(g(h())))`.
func (c *PatientClient) Use(hooks ...Hook) {
	c.hooks.Patient = append(c.hooks.Patient, hooks...)
}

// Create returns a create builder for Patient.
func (c *PatientClient) Create() *PatientCreate {
	mutation := newPatientMutation(c.config, OpCreate)
	return &PatientCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Patient.
func (c *PatientClient) Update() *PatientUpdate {
	mutation := newPatientMutation(c.config, OpUpdate)
	return &PatientUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PatientClient) UpdateOne(pa *Patient) *PatientUpdateOne {
	mutation := newPatientMutation(c.config, OpUpdateOne, withPatient(pa))
	return &PatientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PatientClient) UpdateOneID(id int) *PatientUpdateOne {
	mutation := newPatientMutation(c.config, OpUpdateOne, withPatientID(id))
	return &PatientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Patient.
func (c *PatientClient) Delete() *PatientDelete {
	mutation := newPatientMutation(c.config, OpDelete)
	return &PatientDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PatientClient) DeleteOne(pa *Patient) *PatientDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PatientClient) DeleteOneID(id int) *PatientDeleteOne {
	builder := c.Delete().Where(patient.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PatientDeleteOne{builder}
}

// Create returns a query builder for Patient.
func (c *PatientClient) Query() *PatientQuery {
	return &PatientQuery{config: c.config}
}

// Get returns a Patient entity by its id.
func (c *PatientClient) Get(ctx context.Context, id int) (*Patient, error) {
	return c.Query().Where(patient.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PatientClient) GetX(ctx context.Context, id int) *Patient {
	pa, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pa
}

// QueryDispenses queries the dispenses edge of a Patient.
func (c *PatientClient) QueryDispenses(pa *Patient) *DispenseQuery {
	query := &DispenseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(patient.Table, patient.FieldID, id),
			sqlgraph.To(dispense.Table, dispense.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, patient.DispensesTable, patient.DispensesColumn),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PatientClient) Hooks() []Hook {
	return c.hooks.Patient
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryDispenses queries the dispenses edge of a User.
func (c *UserClient) QueryDispenses(u *User) *DispenseQuery {
	query := &DispenseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(dispense.Table, dispense.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.DispensesTable, user.DispensesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
