// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package things

import (
	"context"

	"github.com/mainflux/mainflux/pkg/errors"
)

var (
	// ErrMalformedEntity indicates malformed entity specification (e.g.
	// invalid username or password).
	ErrMalformedEntity = errors.New("malformed entity specification")

	// ErrNotFound indicates a non-existent entity request.
	ErrNotFound = errors.New("non-existent entity")

	// ErrConflict indicates that entity already exists.
	ErrConflict = errors.New("entity already exists")

	// ErrScanMetadata indicates problem with metadata in db
	ErrScanMetadata = errors.New("failed to scan metadata in db")

	// ErrSelectEntity indicates error while reading entity from database
	ErrSelectEntity = errors.New("select entity from db error")

	// ErrEntityConnected indicates error while checking connection in database
	ErrEntityConnected = errors.New("check thing-channel connection in database error")
)

// Metadata to be used for mainflux thing or channel for customized
// describing of particular thing or channel.
type Metadata map[string]interface{}

// Thing represents a Mainflux thing. Each thing is owned by one user, and
// it is assigned with the unique identifier and (temporary) access key.
type Thing struct {
	ID       string
	Owner    string
	Name     string
	Key      string
	Metadata Metadata
}

// Page contains page related metadata as well as list of things that
// belong to this page.
type Page struct {
	PageMetadata
	Things []Thing
}

// ThingRepository specifies a thing persistence API.
type ThingRepository interface {
	// Save persists multiple things. Things are saved using a transaction. If one thing
	// fails then none will be saved. Successful operation is indicated by non-nil
	// error response.
	Save(ctx context.Context, ths ...Thing) ([]Thing, error)

	// Update performs an update to the existing thing. A non-nil error is
	// returned to indicate operation failure.
	Update(ctx context.Context, t Thing) error

	// UpdateKey updates key value of the existing thing. A non-nil error is
	// returned to indicate operation failure.
	UpdateKey(ctx context.Context, owner, id, key string) error

	// RetrieveByID retrieves the thing having the provided identifier, that is owned
	// by the specified user.
	RetrieveByID(ctx context.Context, owner, id string) (Thing, error)

	// RetrieveByKey returns thing ID for given thing key.
	RetrieveByKey(ctx context.Context, key string) (string, error)

	// RetrieveAll retrieves the subset of things owned by the specified user.
	RetrieveAll(ctx context.Context, owner string, pm PageMetadata) (Page, error)

	// RetrieveByChannel retrieves the subset of things owned by the specified
	// user and connected or not connected to specified channel.
	RetrieveByChannel(ctx context.Context, owner, channel string, offset, limit uint64, connected bool) (Page, error)

	// Remove removes the thing having the provided identifier, that is owned
	// by the specified user.
	Remove(ctx context.Context, owner, id string) error
}

// ThingCache contains thing caching interface.
type ThingCache interface {
	// Save stores pair thing key, thing id.
	Save(context.Context, string, string) error

	// ID returns thing ID for given key.
	ID(context.Context, string) (string, error)

	// Removes thing from cache.
	Remove(context.Context, string) error
}
