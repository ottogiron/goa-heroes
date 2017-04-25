package main

import (
	"github.com/goadesign/goa"
	"github.com/ottogiron/goa-heroes/app"
	"github.com/ottogiron/goa-heroes/services"
)

func main() {
	service := goa.New("heroes")

	heroService := &services.HeroService{}
	heroController := NewHeroController(service, heroService)

	app.MountHeroController(service, heroController)

	// Run service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError(err.Error())
	}
}
