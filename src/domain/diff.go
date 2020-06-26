package domain

import "sort"

// DiffStatus ...
type DiffStatus int

const (
	_ DiffStatus = iota

	// NotChanged ...
	NotChanged
	// Added ...
	Added
	// Changed ...
	Changed
	// Deleted ...
	Deleted
)

// Diff ...
type Diff struct {
	Status DiffStatus
	Name   string
	Before string
	After  string
}

// Diffs ...
type Diffs []*Diff

func (ds *Diffs) sort() {
	sort.SliceStable(*ds, func(i, j int) bool { return (*ds)[i].Name < (*ds)[j].Name })
}
