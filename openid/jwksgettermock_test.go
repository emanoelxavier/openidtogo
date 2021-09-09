package openid

import (
	"testing"

	"github.com/go-jose/go-jose/v3"
)

type jwksGetterMock struct {
	t     *testing.T
	Calls chan Call
}

func newJwksGetterMock(t *testing.T) *jwksGetterMock {
	return &jwksGetterMock{t, make(chan Call)}
}

type getJwksCall struct {
	url string
}

type getJwksResponse struct {
	jwks jose.JSONWebKeySet
	err  error
}

func (c *jwksGetterMock) getJwkSet(url string) (jose.JSONWebKeySet, error) {
	c.Calls <- &getJwksCall{url}
	gr := (<-c.Calls).(*getJwksResponse)
	return gr.jwks, gr.err
}

func (c *jwksGetterMock) assertGetJwks(url string, jwks jose.JSONWebKeySet, err error) {
	call := (<-c.Calls).(*getJwksCall)
	if url != anything && call.url != url {
		c.t.Error("Expected getJwks with", url, "but was", call.url)
	}
	c.Calls <- &getJwksResponse{jwks, err}
}

func (c *jwksGetterMock) close() {
	close(c.Calls)
}

func (c *jwksGetterMock) assertDone() {
	if _, more := <-c.Calls; more {
		c.t.Fatal("Did not expect more calls.")
	}
}
