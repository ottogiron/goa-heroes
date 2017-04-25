package services

import (
	"github.com/ottogiron/goa-heroes/app"
)

//HeroService bussiness logic
type HeroService struct {
}

//List return a list of heroes
func (h HeroService) List() []app.Hero {
	return nil
}
