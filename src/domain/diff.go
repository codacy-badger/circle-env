package domain

import "sort"

type DiffStatus int

const (
	_ DiffStatus = iota
	NotChanged
	Added
	Changed
	Deleted
)

type Diff struct {
	Status DiffStatus
	Name   string
	Before string
	After  string
}

type Diffs []*Diff

func (ds *Diffs) sort() {
	sort.SliceStable(*ds, func(i, j int) bool { return (*ds)[i].Name < (*ds)[j].Name })
}
