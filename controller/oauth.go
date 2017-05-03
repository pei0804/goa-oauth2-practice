package controller

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/tikasan/goa-oauth2-practice/app"
	"github.com/tikasan/goa-oauth2-practice/util"
)

const (
	// OAuth2ClientID is the only authorized client ID
	OAuth2ClientID = ""
	// OAuth2ClientSecret is the only authorized client secret
	OAuth2ClientSecret = ""
)

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
	g := util.NewGitHubOAuth(OAuth2ClientID, OAuth2ClientSecret, "[“user:email”]")
	tokenStr, _ := g.GenerateToken(ctx.Code, "ABC")
	userInfo := g.GetUserInfo(tokenStr)
	githubID := fmt.Sprintf("%0.f", userInfo["id"].(float64))
	goa.LogInfo(ctx.Context, "userinfo", "userinfo", userInfo)
	goa.LogInfo(ctx.Context, "githubID", "githubID", githubID)

	// OauthController_Callback: end_implement
	return nil
}

// Login runs the login action.
func (c *OauthController) Login(ctx *app.LoginOauthContext) error {
	// OauthController_Login: start_implement

	// Put your logic here
	g := util.NewGitHubOAuth(OAuth2ClientID, OAuth2ClientSecret, "[“user:email”]")

	// OauthController_Login: end_implement
	ctx.ResponseData.Header().Set("Location", g.GetAuthorizeUrl("ABC"))
	return ctx.Found()
}
