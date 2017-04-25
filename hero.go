package main

import (
	"github.com/goadesign/goa"
	"github.com/ottogiron/goa-heroes/app"
	"github.com/ottogiron/goa-heroes/services"
)

// HeroController implements the hero resource.
type HeroController struct {
	*goa.Controller
	heroService *services.HeroService
}

// NewHeroController creates a hero controller.
func NewHeroController(service *goa.Service, heroService *services.HeroService) *HeroController {
	return &HeroController{Controller: service.NewController("HeroController"), heroService: heroService}
}

// Create runs the create action.
func (c *HeroController) Create(ctx *app.CreateHeroContext) error {
	// HeroController_Create: start_implement

	// Put your logic here
	hero := &app.Hero{
		Name: ctx.Payload.Name,
	}

	c.heroService.Create(hero)
	// HeroController_Create: end_implement
	return ctx.Created()
}

// List runs the list action.
func (c *HeroController) List(ctx *app.ListHeroContext) error {
	// HeroController_List: start_implement

	// Put your logic here

	// HeroController_List: end_implement
	res := c.heroService.List()
	return ctx.OK(res)
}

// Show runs the show action.
func (c *HeroController) Show(ctx *app.ShowHeroContext) error {
	// HeroController_Show: start_implement

	// Put your logic here

	// HeroController_Show: end_implement
	res := c.heroService.Show(ctx.HeroID)
	return ctx.OK(res)
}
