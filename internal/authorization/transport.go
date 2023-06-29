package authorization

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

type authorizationRoundTripper struct {
	rt       http.RoundTripper
	username string
	token    string
	isCloud  bool
}

func (t authorizationRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	var authHeader string
	if t.isCloud {
		auth := t.username + ":" + t.token
		base64Auth := base64.StdEncoding.EncodeToString([]byte(auth))
		authHeader = fmt.Sprintf("Basic %s", base64Auth)
	} else {
		authHeader = fmt.Sprintf("Bearer %s", t.token)
	}

	request.Header.Set("Authorization", authHeader)
	response, err := t.rt.RoundTrip(request)
	return response, err
}

func AuthorizationTransport(transport http.RoundTripper, token string, isCloud bool) http.RoundTripper {
	return &authorizationRoundTripper{
		rt:      transport,
		token:   token,
		isCloud: isCloud,
	}
}
