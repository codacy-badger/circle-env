package domain

type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Envs []*Env
