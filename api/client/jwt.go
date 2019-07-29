// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "rankdb": jwt Resource Client
//
// Command:
// $ goagen
// --design=github.com/Vivino/rankdb/api/design
// --out=$(GOPATH)/src/github.com/Vivino/rankdb/api
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// JWTJWTPath computes a request path to the jwt action of jwt.
func JWTJWTPath() string {

	return fmt.Sprintf("/jwt")
}

// JWT key generator.
// If left disabled in config, Unauthorized is returned
func (c *Client) JWTJWT(ctx context.Context, path string, scope string, expire *int, onlyElements *string, onlyLists *string) (*http.Response, error) {
	req, err := c.NewJWTJWTRequest(ctx, path, scope, expire, onlyElements, onlyLists)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewJWTJWTRequest create the request corresponding to the jwt action endpoint of the jwt resource.
func (c *Client) NewJWTJWTRequest(ctx context.Context, path string, scope string, expire *int, onlyElements *string, onlyLists *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("scope", scope)
	if expire != nil {
		tmp55 := strconv.Itoa(*expire)
		values.Set("expire", tmp55)
	}
	if onlyElements != nil {
		values.Set("only_elements", *onlyElements)
	}
	if onlyLists != nil {
		values.Set("only_lists", *onlyLists)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}