package config

import (
	"github.com/cweill/gotests/gotests/process"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
);

func ConfigSetup() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID: "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:5500/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}