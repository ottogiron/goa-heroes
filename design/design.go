package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("comics", func() { // API defines the microservice endpoint and
	Title("Comics Heroes API")          // other global properties. There should be one
	Description("A simple goa service") // and exactly one API definition appearing in
	Scheme("http")                      // the design.
	Host("localhost:8080")
})

var _ = Resource("hero", func() { // Resources group related API endpoints
	BasePath("/heroes")     // together. They map to REST resources for REST
	DefaultMedia(HeroMedia) // services.

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all heroes")
		Response(OK, CollectionOf(HeroMedia))
	})

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get hero by id") // with its path, parameters (both path
		Routing(GET("/:heroID"))      // parameters and querystring values) and payload
		Params(func() {               // (shape of the request body).
			Param("heroID", Integer, "Bottle ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new hero")
		Payload(func() {
			Member("name")
			Required("name")
		})
		Response(Created, "/hero/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

})

// HeroMedia defines the media type used to render heroes.
var HeroMedia = MediaType("application/vnd.goa.example.heroe+json", func() {
	Description("A comic hero")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique heroe ID")
		Attribute("name", String, "Name of hero")
		Required("id", "href", "name")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")   // Media types may have multiple views and must
		Attribute("href") // have a "default" view.
		Attribute("name")
	})
})
