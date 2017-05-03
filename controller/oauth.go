package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bradrydzewski/go.auth/oauth2"
	"github.com/goadesign/goa"
	"github.com/tikasan/goa-oauth2-practice/app"
)

const (
	Scope = "user:email"  // grant access to the `users` api
	State = "FqB4EbagQ2o" // random string to protect against CSRF attacks
)

var client = oauth2.Client{
	RedirectURL:      "http://localhost:8080/login/callback",
	AccessTokenURL:   "https://github.com/login/oauth/access_token",
	AuthorizationURL: "https://github.com/login/oauth/authorize",
	ClientId:         "",
	ClientSecret:     "",
}

type User struct {
	Login string `json:"login"`
	Email string `json:"email"`
}

// OauthController implements the oauth resource.
type OauthController struct {
	*goa.Controller
}

// NewOauthController creates a oauth controller.
func NewOauthController(service *goa.Service) *OauthController {
	return &OauthController{Controller: service.NewController("OauthController")}
}

// Callback runs the callback action.
func (c *OauthController) Callback(ctx *app.CallbackOauthContext) error {
	// OauthController_Callback: start_implement

	// Put your logic here
	accessToken, err := client.GrantToken(ctx.Code)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Access Token", accessToken)
	}

	// create the http.Request that will access a restricted resource
	req, _ := http.NewRequest("GET", "https://api.github.com/user?access_token="+accessToken.AccessToken, nil)

	// make the request
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal the body
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	user := User{}
	if err := json.Unmarshal(raw, &user); err != nil {
		log.Fatal(err)
	}

	goa.LogInfo(ctx.Context, "userinfo", "userinfo", user)

	// OauthController_Callback: end_implement
	return nil
}

// Login runs the login action.
func (c *OauthController) Login(ctx *app.LoginOauthContext) error {
	// OauthController_Login: start_implement

	// Put your logic here
	ctx.ResponseData.Header().Set("Location", client.AuthorizeRedirect(Scope, State))
	return ctx.Found()
	// OauthController_Login: end_implement
	return nil
}
