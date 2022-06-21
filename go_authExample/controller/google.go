package controller

import (
	"context"
	"net/http"
	"io/ioutil"
	"github.com/Shaviaditya/BasicGo/go_authExample/config"
	"github.com/gofiber/fiber/v2"
);
func Login(c *fiber.Ctx) error {
	googleConfig := config.ConfigSetup()
	url := googleConfig.AuthCodeURL("state")
	err := c.Redirect(url,fiber.StatusSeeOther); if err!= nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}
	return c.Redirect(url,fiber.StatusSeeOther)
}

/* func Login(res http.ResponseWriter, req *http.Request){
	googleConfig := config.ConfigSetup()
	url := googleConfig.AuthCodeURL("state")
	http.Redirect(res,req,url,http.StatusSeeOther)
} */

func Callback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "state" {
		return c.Status(503).SendString("States Invalid")
	}
	code := c.Query("code")
	googleConfig := config.ConfigSetup()
	token,err := googleConfig.Exchange(context.Background(),code)
	if err!=nil {
		return c.Status(503).SendString("Code Token exchange failed")
	}
	// resp := c.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token="+token.AccessToken)
	// fmt.Println(resp)
	resp,err:= http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token="+token.AccessToken)
	if err!=nil {
		return c.Status(503).SendString("Fetch Failed")
	}
	userData,err:= ioutil.ReadAll(resp.Body)
	if err!=nil {
		return c.Status(503).SendString("Parse Failed")
	}
	return c.SendString(string(userData))
}

/* func Callback(res http.ResponseWriter, req *http.Request){
	state := req.URL.Query()["state"][0]
	if state != "state" {
		fmt.Fprintln(res, "States Invalid")
		return
	}

	code := req.URL.Query()["code"][0]
	googleConfig := config.ConfigSetup()
	token,err := googleConfig.Exchange(context.Background(),code)
	if err!=nil {
		fmt.Fprintln(res, "Code-Token Exchange Failed")
		return
	}

	resp,err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token="+token.AccessToken)
	if err!=nil {
		fmt.Fprintln(res, "User Data fetch Failed")
		return 
	}

	userData,err := ioutil.ReadAll(resp.Body)
	if err!=nil {
		fmt.Fprintln(res, "User Data parsing Failed")
		return
	}

	fmt.Fprintln(res, string(userData))
} */