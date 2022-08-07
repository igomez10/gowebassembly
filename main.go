package main

import (
	"fmt"
	"net/http"
)

func main() {
	println("Hello, world!")

	// make a get http request to retrieve an image
	for {
		go func() {
			req, err := http.NewRequest("GET", "https://api.github.com", nil)
			if err != nil {
				fmt.Println(err)
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()

			// println(resp.StatusCode)
		}()
	}

}
