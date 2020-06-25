package domain

import "sort"

type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Envs []*Env

func (es *Envs) Has(name string) bool {
	for _, e := range *es {
		if name == e.Name {
			return true
		}
	}

	return false
}

func (es *Envs) Get(name string) *Env {
	for _, e := range *es {
		if name == e.Name {
			return e
		}
	}

	return nil
}

func (es *Envs) Sort() {
	sort.SliceStable(*es, func(i, j int) bool { return (*es)[i].Name < (*es)[j].Name })
}

func (es *Envs) Compare(comp *Envs, del bool) *Diffs {
	ds := new(Diffs)

	for _, c := range *comp {
		if es.Has(c.Name) {
			*ds = append(*ds, &Diff{Changed, c.Name, es.Get(c.Name), c})
		} else {
			*ds = append(*ds, &Diff{Added, c.Name, nil, c})
		}
	}

	for _, e := range *es {
		if !comp.Has(e.Name) {
			if del {
				*ds = append(*ds, &Diff{Deleted, e.Name, e, nil})
			} else {
				*ds = append(*ds, &Diff{NotChanged, e.Name, e, nil})
			}
		}
	}

	ds.Sort()
	return ds
}

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
	Before *Env
	After  *Env
}

type Diffs []*Diff

func (ds *Diffs) Sort() {
	sort.SliceStable(*ds, func(i, j int) bool { return (*ds)[i].Name < (*ds)[j].Name })
}
