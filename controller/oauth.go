package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/goadesign/goa"
	"github.com/tikasan/goa-oauth2-practice/app"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type User struct {
	Login string `json:"login"`
	Email string `json:"email"`
}

// OauthController implements the oauth resource.
type OauthController struct {
	*goa.Controller
	*oauth2.Config
}

// NewOauthController creates a oauth controller.
func NewOauthController(service *goa.Service) *OauthController {
	conf := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Scopes:       []string{"user"},
		Endpoint:     github.Endpoint,
	}
	return &OauthController{
		Controller: service.NewController("OauthController"),
		Config:     conf,
	}
}

// Callback runs the callback action.
func (c *OauthController) Callback(ctx *app.CallbackOauthContext) error {

	// OauthController_Callback: start_implement
	tok, err := c.Config.Exchange(ctx, ctx.Code)
	if err != nil {
		log.Fatal(err)
	}

	client := c.Config.Client(ctx, tok)
	resp, err := client.Get("https://api.github.com/user?access_token=" + tok.AccessToken)

	// make the request
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

	// OauthController_Callback: end_implement
	return nil
}

// Login runs the login action.
func (c *OauthController) Login(ctx *app.LoginOauthContext) error {
	// OauthController_Login: start_implement

	// Put your logic here
	url := c.Config.AuthCodeURL("state")
	goa.LogInfo(ctx.Context, "url: this", "this", url)
	return ctx.Found()
	// OauthController_Login: end_implement
}
