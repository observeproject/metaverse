package storage

import (
	"context"

	"github.com/observeproject/common/resource/model"
)

type Appendable interface {
	// Appender returns a new appender for the storage.
	Appender() Appender
}

type Queryable interface {
	Querier(ctx context.Context, mint, maxt int64) (Querier, error)
}

type RelationSwitch int

const (
	RelationOn RelationSwitch = iota << 1
	RelationOff
)

// Appender provides batched appends against a storage.
// It must be completed with a call to Commit or Rollback and must not be reused afterwards.
// Operations on the Appender interface are not goroutine-safe.
type Appender interface {

	// AppendResource adds a new copy for the given resource.
	AppendResource(urn string, t int64, content ResourceContent) error

	// AppendState adds a new value of the state for the given resource.
	AppendState(urn string, state string, t int64, value uint64) error

	// AppendRelation adds a new connection of the given resources.
	AppendRelation(name, srcRrn, targetUrn string, t int64, rs RelationSwitch) error

	// Commit submits the collected resources and purges the batch. If Commit
	// returns a non-nil error, it also rolls back all modifications made in
	// the appender so far, as Rollback would do. In any case, an Appender
	// must not be used anymore after Commit has been called.
	Commit() error

	// Rollback rolls back all modifications made in the appender so far.
	// Appender has to be discarded after rollback.
	Rollback() error
}

type baseQuerier interface {
	// AttributeNames returns all the unique attribute names of given type.
	AttributeNames(typ string) ([]string, Warnings, error)

	// Types returns all the unique schema names present in sorted order
	Types() ([]string, Warnings, error)

	// Close releases the resources of the Querier.
	Close() error
}

type Querier interface {
	baseQuerier

	// SelectResources return a set of Resources that matches the given ResourceSelector.
	// Caller can specify if it requires return resources to be sorted.
	SelectResources(sortResources bool, selectors ...model.ResourceSelector) ResourceRelationSet

	// SelectStates returns a set of resource state that matcher the given ResourceSelector.
	// Caller can specify if requires return states to be sorted
	SelectStates(sortStates bool, selectors ...model.StateSelector) StateSet
}
