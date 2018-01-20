package models

type Hateoas struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

func NewHateoas(rel string, href string) Hateoas {
	return Hateoas{Rel: rel, Href: href}
}
