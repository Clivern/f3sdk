// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package f3sdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

const (
	// BaseURL Form3 API Base
	BaseURL = "https://api.form3.tech"
)

// error modes
var (
	ErrResourceNotFound = errors.New("Specified resource does not exist")
	ErrVersionConflict  = errors.New("Specified version incorrect")
)

// Account struct
type Account struct {
	client  *HTTPClient
	baseURL string
}

// NewAccountClient creates an account client
func NewAccountClient(timeout int, baseURL string) *Account {
	return &Account{
		client:  NewHTTPClient(timeout),
		baseURL: baseURL,
	}
}

// Create creates a new account
func (a *Account) Create(ctx context.Context, data *Data) (*Data, error) {
	body, err := data.ConvertToJSON()

	if err != nil {
		return data, fmt.Errorf("Unexpected error: %s", err.Error())
	}

	response, err := a.client.Post(
		ctx,
		fmt.Sprintf("%s/v1/organisation/accounts", a.baseURL),
		body,
		map[string]string{},
		map[string]string{},
	)

	if err != nil {
		return data, fmt.Errorf("Unexpected error: %s", err.Error())
	}

	statusCode := a.client.GetStatusCode(response)
	responseBody, err := a.client.ToString(response)

	if err != nil {
		return data, fmt.Errorf("Unexpected error: %s", err.Error())
	}

	if statusCode != http.StatusCreated {
		return data, fmt.Errorf("Invalid response: statusCode [%d] and body [%s]", statusCode, responseBody)
	}

	err = data.LoadFromJSON([]byte(responseBody))

	if err != nil {
		return data, fmt.Errorf("Unexpected error: %s", err.Error())
	}

	return data, nil
}

// FetchOneByAccountID gets an account by ID
func (a *Account) FetchOneByAccountID(ctx context.Context, accountID string) (*Data, error) {
	var data *Data

	response, err := a.client.Get(
		ctx,
		fmt.Sprintf("%s/v1/organisation/accounts/%s", a.baseURL, accountID),
		map[string]string{},
		map[string]string{},
	)

	if err != nil {
		return data, fmt.Errorf("Unexpected error: %s", err.Error())
	}

	statusCode := a.client.GetStatusCode(response)

	responseBody, err := a.client.ToString(response)

	if err != nil {
		return data, fmt.Errorf("Unexpected error: %s", err.Error())
	}

	if statusCode != http.StatusOK {
		return data, fmt.Errorf("Invalid response: statusCode [%d] and body [%s]", statusCode, responseBody)
	}

	data = &Data{}

	err = data.LoadFromJSON([]byte(responseBody))

	if err != nil {
		return data, fmt.Errorf("Unexpected error: %s", err.Error())
	}

	return data, nil
}

// DeleteOneByAccountID deletes an account by ID
func (a *Account) DeleteOneByAccountID(ctx context.Context, accountID string, version int) error {
	response, err := a.client.Delete(
		ctx,
		fmt.Sprintf("%s/v1/organisation/accounts/%s?version=%d", a.baseURL, accountID, version),
		map[string]string{},
		map[string]string{},
	)

	if err != nil {
		return fmt.Errorf("Unexpected error: %s", err.Error())
	}

	statusCode := a.client.GetStatusCode(response)

	if statusCode == http.StatusNotFound {
		return ErrResourceNotFound
	}

	if statusCode == http.StatusConflict {
		return ErrVersionConflict
	}

	if statusCode == http.StatusNoContent {
		return nil
	}

	responseBody, err := a.client.ToString(response)

	if err != nil {
		return fmt.Errorf("Unexpected error: %s", err.Error())
	}

	return fmt.Errorf("Invalid response: statusCode [%d] and body [%s]", statusCode, responseBody)
}
