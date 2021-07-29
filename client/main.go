package main

import (
	"context"
	"log"
	"net/http"

	objectpbs "github.com/kyawmyintthein/twirp-poc/protopbs/protos/objects"
	"github.com/twitchtv/twirp"
)

func main() {

	client := objectpbs.NewObjectServiceProtobufClient("http://localhost:8080", &http.Client{}, twirp.WithClientPathPrefix("/rz"))

	for i := 0; i < 5; i++ {
		empty, err := client.GetObject(context.Background(), &objectpbs.ObjectRequest{})
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
