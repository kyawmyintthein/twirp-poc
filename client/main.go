package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"

	userpbs "github.com/kyawmyintthein/twirp-poc/protopbs/protos/user/v1"
	"github.com/twitchtv/twirp"
	"golang.org/x/net/http2"
)

func main() {

	client := userpbs.NewUserServceProtobufClient("https://localhost:8080", &http.Client{
		Transport: transport2(),
	}, twirp.WithClientPathPrefix("/rz"))

	go func() {
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
				continue
			}
		}
	}()
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

func transport2() *http2.Transport {
	return &http2.Transport{
		TLSClientConfig:    tlsConfig(),
		DisableCompression: true,
		AllowHTTP:          false,
	}
}

func tlsConfig() *tls.Config {
	crt, err := ioutil.ReadFile("../server.crt")
	if err != nil {
		log.Fatal(err)
	}

	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(crt)

	return &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: false,
		ServerName:         "localhost",
	}
}
