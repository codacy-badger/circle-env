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
		if e := es.Get(c.Name); e != nil {
			*ds = append(*ds, &Diff{Changed, c.Name, e.Value, c.Value})
		} else {
			*ds = append(*ds, &Diff{Added, c.Name, "", c.Value})
		}
	}

	for _, e := range *es {
		if !comp.Has(e.Name) {
			if del {
				*ds = append(*ds, &Diff{Deleted, e.Name, e.Value, ""})
			} else {
				*ds = append(*ds, &Diff{NotChanged, e.Name, e.Value, ""})
			}
		}
	}

	ds.sort()
	return ds
}
