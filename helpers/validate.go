package helpers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"unicode/utf8"
)

func ValidateStars(stars int, err error) bool {
	if stars > 5 || stars < 1 {
		return false
	}
	return true
}

func ValidateURI(uri string) bool { // this function check if the syntax of URI is correct or not
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return false
	}

	u, err := url.Parse(uri)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
		return false
	}

	return true
}

/**
This function check if the URI is reachable and alive or not, I didnot use it
in the validation, I am only showing the idea of URI validity
*/
func urlIsReachable(web string) bool {
	response, errors := http.Get(web)

	if errors != nil {
		_, netErrors := http.Get(web)

		if netErrors != nil {
			fmt.Fprintf(os.Stderr, "no internet\n")
			os.Exit(1)
		}

		return false
	}

	if response.StatusCode == 200 {
		return true
	}

	return false
}

func ValidateNameUtf8(name string) bool {
	if utf8.ValidString(name) {
		return true
	}
	return false
}
