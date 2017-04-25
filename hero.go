package main

import (
	"github.com/goadesign/goa"
	"github.com/ottogiron/goa-heroes/app"
)

// HeroController implements the hero resource.
type HeroController struct {
	*goa.Controller
}

// NewHeroController creates a hero controller.
func NewHeroController(service *goa.Service) *HeroController {
	return &HeroController{Controller: service.NewController("HeroController")}
}

// Create runs the create action.
func (c *HeroController) Create(ctx *app.CreateHeroContext) error {
	// HeroController_Create: start_implement

	// Put your logic here

	// HeroController_Create: end_implement
	return nil
}

// List runs the list action.
func (c *HeroController) List(ctx *app.ListHeroContext) error {
	// HeroController_List: start_implement

	// Put your logic here

	// HeroController_List: end_implement
	res := app.HeroCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *HeroController) Show(ctx *app.ShowHeroContext) error {
	// HeroController_Show: start_implement

	// Put your logic here

	// HeroController_Show: end_implement
	res := &app.Hero{}
	return ctx.OK(res)
}
