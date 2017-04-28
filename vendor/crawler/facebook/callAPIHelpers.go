package facebook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Makes the request with the url built by our default url helper.
func defaultCallGraphAPI(path, accessToken string) ([]byte, error) {
	return callGraphAPI(defaultGraphAPIURL(path, accessToken).String())
}

// Builds an url for a call to the Facebook Graph API.
func defaultGraphAPIURL(path, accessToken string) *url.URL {
	query := url.Values{}
	query.Set("fields", strings.Join(extractedPageFields, ","))
	query.Set("access_token", accessToken)

	return &url.URL{
		Scheme:   "https",
		Host:     graphAPIHost,
		Path:     graphAPIVersion + "/" + path,
		RawQuery: query.Encode(),
	}
}

// Makes http request to the provided url and checks for Facebook error message.
func callGraphAPI(url string) ([]byte, error) {
	// Make API call.
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Facebook API call error: %s", err)
	}

	// Proccess response
	defer response.Body.Close()
	rawResponseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading content of Facebook response: %s", err)
	}

	// Check if API returned an internal error.
	var apiError graphAPIError
	if err := json.Unmarshal(rawResponseBody, &apiError); err != nil {
		return nil, fmt.Errorf("Error docoding page response as error: %s", err)
	}

	if !apiError.isEmpty() {
		return nil, apiError
	}

	// All good.
	return rawResponseBody, nil
}
