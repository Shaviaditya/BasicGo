package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
);

func ConfigSetup() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "320086037196-95g0ecmdrna8ov222m5oe4jdoqr26483.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:5500/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}