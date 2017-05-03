//go:generate goagen bootstrap -d github.com/tikasan/goa-oauth2-practice/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tikasan/goa-oauth2-practice/app"
	"github.com/tikasan/goa-oauth2-practice/controller"
)

func main() {
	// Create service
	service := goa.New("auth")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "oauth" controller
	c := controller.NewOauthController(service)
	app.MountOauthController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
