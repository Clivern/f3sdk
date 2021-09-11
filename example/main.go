// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"

	"github.com/clivern/f3sdk"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()

	accountClient := f3sdk.NewAccountClient(30, f3sdk.BaseURL)

	// Create Account
	data := &f3sdk.Data{
		Data: &f3sdk.AccountData{
			ID:             id.String(),
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Type:           "accounts",
			Attributes: &f3sdk.AccountAttributes{
				AccountNumber:           "41426819",
				AlternativeNames:        []string{"Sam Holder"},
				BankID:                  "400300",
				BankIDCode:              "GBDSC",
				BaseCurrency:            "GBP",
				Bic:                     "NWBKGB22",
				Iban:                    "GB11NWBK40030041426819",
				Name:                    []string{"Samantha Holder"},
				SecondaryIdentification: "A1B2C3D4",
				Country:                 "GB",
			},
		},
	}

	data, err := accountClient.CreateAccount(
		context.TODO(),
		data,
	)

	fmt.Println(err)          // nil
	fmt.Println(data.Data.ID) // id.String()

	// Fetch Account by UUID
	result, err := accountClient.FetchAccountByID(
		context.TODO(),
		id.String(),
	)

	fmt.Println(err)            // nil
	fmt.Println(result.Data.ID) // id.String()

	err = accountClient.DeleteAccountByID(
		context.TODO(),
		result.Data.ID,
		*result.Data.Version,
	)

	fmt.Println(err) // nil
}
