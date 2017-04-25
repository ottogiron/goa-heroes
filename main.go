package main

import (
	"github.com/goadesign/goa"
	"github.com/ottogiron/goa-heroes/app"
)

func main() {
	service := goa.New("heroes")

	heroController := NewHeroController(service)

	app.MountHeroController(service, heroController)

	// Run service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError(err.Error())
	}
}
