package domain

type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Envs []*Env

func (es *Envs) Has(e *Env) bool {
	for _, ex := range *es {
		if ex.Name == e.Name {
			return true
		}
	}

	return false
}
