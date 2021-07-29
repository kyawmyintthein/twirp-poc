package main

import (
	"context"
	"log"
	"net/http"

	userpbs "github.com/kyawmyintthein/twirp-poc/protopbs/protos/user/v1"
	"github.com/twitchtv/twirp"
)

func main() {

	client := userpbs.NewUserServceProtobufClient("http://localhost:8080", &http.Client{}, twirp.WithClientPathPrefix("/rz"))
	for i := 0; i < 5; i++ {
		loginSession, err := client.Onboarding(context.Background(), &userpbs.UserOnboardingRequest{})
		if err != nil {
			// We got an error. Is it a twirp Error?
			if twerr, ok := err.(twirp.Error); ok {
				// Twirp errors support custom, arbitrary metadata. For example, a
				// server could inform a client that a particular error is retryable.
				if twerr.Meta("retryable") != "" {
					log.Printf("got error %q, retrying", twerr)
					continue
				}
			}
			log.Fatal(err)
		} else {
			log.Println("onboarding response", loginSession)
			break
		}
	}

	for i := 0; i < 5; i++ {
		empty, err := client.UpdateProfile(context.Background(), &userpbs.UpdateProfileRequest{})
		if err != nil {
			// We got an error. Is it a twirp Error?
			if twerr, ok := err.(twirp.Error); ok {
				// Twirp errors support custom, arbitrary metadata. For example, a
				// server could inform a client that a particular error is retryable.
				if twerr.Meta("retryable") != "" {
					log.Printf("got error %q, retrying", twerr)
					continue
				}
			}
			log.Fatal(err)
		} else {
			log.Println("update response", empty)
			break
		}
	}
}
