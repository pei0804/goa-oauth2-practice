package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("auth", func() {
	Title("example")
	Description("A simple appengine example")
	Contact(func() {
		Name("goa team")
		Email("admin@goa.design")
		URL("http://goa.design")
	})
	License(func() {
		Name("MIT")
		URL("https://github.com/goadesign/goa/blob/master/LICENSE")
	})
	Docs(func() {
		Description("goa guide")
		URL("http://goa.design/getting-started.html")
	})
	Host("localhost:8080")
	Scheme("http")
	Scheme("https")
	BasePath("/")

	Origin("*", func() {
		Methods("GET")
		MaxAge(600)
		Credentials()
	})
})

var Message = MediaType("application/vnd.goa.message.json", func() {
	Attributes(func() {
		Attribute("message", String, "message")
		Required("message")
	})
	View("default", func() {
		Attribute("message")
	})
})

var _ = Resource("oauth", func() {
	Description("oauth")

	Action("login", func() {
		Routing(GET("/login"))
		Response(OK)
		Response(Found)
		Response(BadRequest)
	})
	Action("callback", func() {
		Routing(GET("/login/callback"))
		Params(func() {
			Param("code", String, "code")
			Required("code")
		})
		Response(OK)
		Response(BadRequest)
	})
})
