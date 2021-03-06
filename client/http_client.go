// Copyright 2016 Albert Nigmatzianov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package client

import (
	"bytes"
	"errors"
	"net/url"
	"strconv"

	"github.com/bogem/nehm/ui"
	"github.com/valyala/fasthttp"
)

const (
	apiURL   = "https://api.soundcloud.com"
	clientID = "11a37feb6ccc034d5975f3f803928a32"
)

var (
	ErrForbidden = errors.New("403 - Forbidden")
	ErrNotFound  = errors.New("404 - Not Found")

	uriBuffer = new(bytes.Buffer)
)

func resolve(params url.Values) ([]byte, error) {
	uri := formResolveURI(params)
	return get(uri)
}

func formResolveURI(params url.Values) string {
	params.Set("client_id", clientID)

	uriBuffer.Reset()
	uriBuffer.WriteString(apiURL)
	uriBuffer.WriteString("/resolve?")
	uriBuffer.WriteString(params.Encode())
	return uriBuffer.String()
}

func search(params url.Values) ([]byte, error) {
	uri := formSearchURI(params)
	return get(uri)
}

func formSearchURI(params url.Values) string {
	params.Set("client_id", clientID)

	uriBuffer.Reset()
	uriBuffer.WriteString(apiURL)
	uriBuffer.WriteString("/tracks?")
	uriBuffer.WriteString(params.Encode())
	return uriBuffer.String()
}

func getFavorites(uid string, params url.Values) ([]byte, error) {
	uri := formFavoritesURI(uid, params)
	return get(uri)
}

func formFavoritesURI(uid string, params url.Values) string {
	params.Set("client_id", clientID)

	uriBuffer.Reset()
	uriBuffer.WriteString(apiURL)
	uriBuffer.WriteString("/users/")
	uriBuffer.WriteString(uid)
	uriBuffer.WriteString("/favorites?")
	uriBuffer.WriteString(params.Encode())
	return uriBuffer.String()
}

func get(uri string) ([]byte, error) {
	statusCode, body, err := fasthttp.Get(nil, uri)
	if err != nil {
		return nil, err
	}
	if err := handleStatusCode(statusCode); err != nil {
		return nil, err
	}
	return body, nil
}

func handleStatusCode(statusCode int) error {
	switch {
	case statusCode == 403:
		return ErrForbidden
	case statusCode == 404:
		return ErrNotFound
	case statusCode >= 300 && statusCode < 500:
		ui.Term("invalid response from SoundCloud: "+strconv.Itoa(statusCode), nil)
	case statusCode >= 500:
		ui.Term("there is a problem by SoundCloud. Please wait a while", nil)
	}
	return nil
}
