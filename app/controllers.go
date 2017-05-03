// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/goa-oauth2-practice/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-oauth2-practice
// --version=v1.1.0
//
// API "auth": Application Controllers
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// OauthController is the controller interface for the Oauth actions.
type OauthController interface {
	goa.Muxer
	Callback(*CallbackOauthContext) error
	Login(*LoginOauthContext) error
}

// MountOauthController "mounts" a Oauth resource controller on the given service.
func MountOauthController(service *goa.Service, ctrl OauthController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/login/callback", ctrl.MuxHandler("preflight", handleOauthOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/login", ctrl.MuxHandler("preflight", handleOauthOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCallbackOauthContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Callback(rctx)
	}
	h = handleOauthOrigin(h)
	service.Mux.Handle("GET", "/login/callback", ctrl.MuxHandler("Callback", h, nil))
	service.LogInfo("mount", "ctrl", "Oauth", "action", "Callback", "route", "GET /login/callback")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLoginOauthContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Login(rctx)
	}
	h = handleOauthOrigin(h)
	service.Mux.Handle("GET", "/login", ctrl.MuxHandler("Login", h, nil))
	service.LogInfo("mount", "ctrl", "Oauth", "action", "Login", "route", "GET /login")
}

// handleOauthOrigin applies the CORS response headers corresponding to the origin.
func handleOauthOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
