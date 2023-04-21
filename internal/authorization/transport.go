package authorization

import (
	"fmt"
	"net/http"
)

type authorizationRoundTripper struct {
	rt       http.RoundTripper
	username string
	token    string
}

func (t authorizationRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	authHeader := fmt.Sprintf("Bearer %s", t.token)
	request.Header.Set("Authorization", authHeader)
	response, err := t.rt.RoundTrip(request)
	return response, err
}

func AuthorizationTransport(transport http.RoundTripper, token string) http.RoundTripper {
	return &authorizationRoundTripper{
		rt:    transport,
		token: token,
	}
}
