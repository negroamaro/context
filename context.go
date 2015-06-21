////////////////////////////////////////////////////////////////////////////////
// Copyright 2015 Negroamaro. All rights reserved.                            //
////////////////////////////////////////////////////////////////////////////////

package context

import (
	x "golang.org/x/net/context"
	"time"
)

// Context has been implemented for some of the negroamaro frameworks.
// It must be assigned one of the context instance in one of goroutine.
type Context interface {

	/**
	 * function for goroutine internal use.
	 */
	// ID returns the logically unique identifies of the goroutine.
	ID() int64
	// Close notify the end of goroutine for CloseHandler.
	Close()

	/**
	 * function for goroutine owner use.
	 */
	// Cancel
	Cancel()

	/**
	 * implementation of 'golang.org/x/net/context'.
	 */
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

// CloseHandler is goroutine end event listener that is notified by Close method.
type CloseHandler func(ctx Context)

// New is context constructor.
func New(id int64, handler CloseHandler) Context {
	ctx, cancel := x.WithCancel(x.Background())
	return &usefulContext{
		ctx,
		id,
		cancel,
		handler}
}

// context implementation.
type usefulContext struct {
	x.Context
	id     int64
	cancel x.CancelFunc
	closed CloseHandler
}

// context implementation.
func (c *usefulContext) ID() int64 {
	return c.id
}

// context implementation.
func (c *usefulContext) Close() {
	if c.closed != nil {
		c.closed(c)
	}
}

// context implementation.
func (c *usefulContext) Cancel() {
	c.cancel()
}

// EOF
