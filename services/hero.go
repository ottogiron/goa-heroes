package services

import (
	"github.com/ottogiron/goa-heroes/app"
)

var db = map[int]*app.Hero{
	1: &app.Hero{
		ID:   1,
		Name: "Spiderman",
	},
}
var id int

func init() {
	id = 2
}

//HeroService bussiness logic
type HeroService struct {
}

//List return a list of heroes
func (h HeroService) List() app.HeroCollection {
	var heroCollection app.HeroCollection
	for _, hero := range db {
		heroCollection = append(heroCollection, hero)
	}

	return heroCollection
}

//Create creates a new hero
func (h HeroService) Create(hero *app.Hero) {
	hero.ID = id
	id++
	db[hero.ID] = hero
}

//Show shows a hero
func (h HeroService) Show(id int) *app.Hero {
	return db[id]
}
