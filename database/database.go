package database

// import (
// 	"context"
// 	"database/sql"
// 	"errors"
// 	"fmt"
// 	"time"
// )

// type DatabaseService interface {
// 	StartConnection() error
// 	Select(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
// 	Write(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
// 	BeginTransaction(ctx context.Context) error
// 	Commit(ctx context.Context) error
// 	Rollback(ctx context.Context) error
// }

// type Query interface {
// 	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
// 	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
// }

// type service struct {
// 	tx *sql.Tx
// 	ro *connection
// 	rw *connection
// }

// type connection struct {
// 	db               *sql.DB
// 	params           *parameters
// 	driver           string
// 	connectionString string
// 	lastError        error
// }

// type parameters struct {
// 	host            string
// 	name            string
// 	username        string
// 	password        string
// 	charset         string
// 	collation       string
// 	parseTime       bool
// 	multiStatements bool
// 	timeout         int
// 	maxOpenConns    int
// 	maxIdleConns    int
// 	maxLifetime     time.Duration
// }

// const (
// 	ContextKey        = "database_context"
// 	TxNotStartedError = "transaction not started"
// 	ConnectingDbError = "error connecting to database"
// )

// func New(driver string, roParams, rwParams *parameters) DatabaseService {
// 	return &service{
// 		ro: &connection{
// 			params: roParams,
// 			driver: driver,
// 		},
// 		rw: &connection{
// 			params: rwParams,
// 			driver: driver,
// 		},
// 	}
// }

// func NewParameters(
// 	host, name, username, password, charset, collation string,
// 	parseTime, multiStatements bool,
// 	timeout, maxOpenConns, maxIdleConns int,
// 	maxLifetime time.Duration,
// ) *parameters {
// 	return &parameters{
// 		host:            host,
// 		name:            name,
// 		username:        username,
// 		password:        password,
// 		charset:         charset,
// 		collation:       collation,
// 		parseTime:       parseTime,
// 		multiStatements: multiStatements,
// 		timeout:         timeout,
// 		maxOpenConns:    maxOpenConns,
// 		maxIdleConns:    maxIdleConns,
// 		maxLifetime:     maxLifetime,
// 	}
// }

// func FromContext(ctx context.Context) DatabaseService {
// 	c, _ := ctx.Value(ContextKey).(DatabaseService)
// 	return c
// }

// func (c *connection) Connect() {
// 	c.connectionString = fmt.Sprintf(
// 		"%s:%s@tcp(%s)/%s?charset=%s&collation=%s&parseTime=%t&multiStatements=%t&timeout=%ds",
// 		c.params.username,
// 		c.params.password,
// 		c.params.host,
// 		c.params.name,
// 		c.params.charset,
// 		c.params.collation,
// 		c.params.parseTime,
// 		c.params.multiStatements,
// 		c.params.timeout,
// 	)

// 	c.db, c.lastError = sql.Open(c.driver, c.connectionString)
// 	if c.lastError == nil {
// 		c.db.SetMaxOpenConns(c.params.maxOpenConns)
// 		c.db.SetMaxIdleConns(c.params.maxIdleConns)
// 		c.db.SetConnMaxLifetime(c.params.maxLifetime)
// 		c.db.SetConnMaxIdleTime(c.params.maxLifetime)
// 	}
// }

// func (s *service) StartConnection() error {
// 	if s.connectRo().lastError != nil {
// 		return s.ro.lastError
// 	}

// 	if s.connectRw().lastError != nil {
// 		return s.rw.lastError
// 	}

// 	return nil
// }

// func (s *service) connectRo() *connection {
// 	if s.ro.db == nil {
// 		s.ro.Connect()
// 	}

// 	return s.ro
// }

// func (s *service) connectRw() *connection {
// 	if s.rw.db == nil {
// 		s.rw.Connect()
// 	}

// 	return s.rw
// }

// func (s *service) QueryRo() Query {
// 	if s.ro.db == nil {
// 		s.connectRo()
// 	}

// 	return s.ro.db
// }

// func (s *service) QueryRw() Query {
// 	if s.tx != nil {
// 		return s.tx
// 	}

// 	if s.rw.db == nil {
// 		s.connectRw()
// 	}

// 	return s.rw.db
// }

// func (s *service) Select(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
// 	rows, err := s.QueryRo().QueryContext(ctx, query, args...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return rows, nil
// }

// func (s *service) Write(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
// 	result, err := s.QueryRw().ExecContext(ctx, query, args...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

// func (s *service) BeginTransaction(ctx context.Context) error {
// 	if s.rw.lastError != nil {
// 		return s.rw.lastError
// 	}

// 	tx, err := s.connectRw().db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	s.tx = tx
// 	return nil
// }

// func (s *service) Commit(ctx context.Context) error {
// 	if s.tx == nil {
// 		return errors.New(TxNotStartedError)
// 	}

// 	defer func() {
// 		s.tx = nil
// 	}()

// 	return s.tx.Commit()
// }

// func (s *service) Rollback(ctx context.Context) error {
// 	if s.tx == nil {
// 		return errors.New(TxNotStartedError)
// 	}

// 	defer func() {
// 		s.tx = nil
// 	}()

// 	return s.tx.Rollback()
// }
