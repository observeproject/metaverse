package storage

import (
	"github.com/observeproject/metaverse/pkg/historical"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/pkg/labels"
)

type ResourceContent interface {
	Type() string
	SecondaryTypes() []string
	Attributes() map[string]string
}

type ResourceRelationSet interface {
	ResourceSet() ResourceSet
	RelationSet() RelationSet
}

type ResourceSet interface {
	Next() bool
	At() Resource
	// Err The error that iteration as failed with. When an error occurs, set cannot continue to iterate.
	Err() error
	// Warnings A collection of warnings for the whole set, could be return even iteration has not failed with error.
	Warnings() Warnings
}

type RelationSet interface {
	Next() bool
	At() Relation
	Err() error
	Warning() Warnings
}

type Relation interface {
	Name() string
	Source() string
	Target() string
	Records() historical.TimePairs
}

type Resource interface {
	Type() string
	Urn() string
	SecondaryTypes() []string
	CreatedAt() model.Time
	DeletedAt() model.Time
	Attributes() AttributeSet
	State() State
}

type Attribute interface {
	// Name return the attribute name.
	Name() string
	// Records returns the complete set of resource attributes. Including the identifying of resource.
	Records() historical.RecordStrings
}

type AttributeSet interface {
	Next() bool
	At() Attribute
	// Err The error that iteration as failed with. When an error occurs, set cannot continue to iterate.
	Err() error
	// Warnings A collection of warnings for the whole set, could be return even iteration has not failed with error.
	Warnings() Warnings
}

type StateSet interface {
	Next() bool
	At() State
	// Err The error that iteration as failed with. When an error occurs, set cannot continue to iterate.
	Err() error
	// Warnings A collection of warnings for the whole set, could be return even iteration has not failed with error.
	Warnings() Warnings
}

type State interface {
	Resource() string
	States() labels.Labels
	Name() string
	Records() historical.RecordNumber
}

type Warnings []error
