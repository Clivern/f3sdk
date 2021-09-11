// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package f3sdk

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/franela/goblin"
)

// TestUnitAccountClientMethods test cases
func TestUnitAccountClientMethods(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#Create", func() {
		g.It("status error", func() {
			srv := ServerMock(
				"/v1/organisation/accounts",
				`{"error": "Not Found"}`,
				http.StatusInternalServerError,
				false,
			)

			defer srv.Close()

			accountClient := NewAccountClient(30, srv.URL)

			_, err := accountClient.Create(
				context.TODO(),
				&Data{},
			)

			g.Assert(strings.Contains(err.Error(), "Invalid response")).Equal(true)
		})

		g.It("status created", func() {
			srv := ServerMock(
				"/v1/organisation/accounts",
				`{"data":{"attributes":{"account_number":"41426819","alternative_names":["Sam Holder"],"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","iban":"GB11NWBK40030041426819","name":["Samantha Holder"],"secondary_identification":"A1B2C3D4"},"id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts"}}`,
				http.StatusCreated,
				false,
			)

			defer srv.Close()

			accountClient := NewAccountClient(30, srv.URL)

			data := &Data{
				Data: &AccountData{
					OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
					Type:           "accounts",
					Attributes: &AccountAttributes{
						AccountNumber:           "41426819",
						AlternativeNames:        []string{"Sam Holder"},
						BankID:                  "400300",
						BankIDCode:              "GBDSC",
						BaseCurrency:            "GBP",
						Bic:                     "NWBKGB22",
						Iban:                    "GB11NWBK40030041426819",
						Name:                    []string{"Samantha Holder"},
						SecondaryIdentification: "A1B2C3D4",
					},
				},
			}

			data, err := accountClient.Create(
				context.TODO(),
				data,
			)

			g.Assert(err).Equal(nil)
			g.Assert(data.Data.ID).Equal("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
			g.Assert(data.Data.OrganisationID).Equal("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
			g.Assert(data.Data.Type).Equal("accounts")
			g.Assert(data.Data.Attributes.SecondaryIdentification).Equal("A1B2C3D4")
		})
	})

	g.Describe("#FetchOneByAccountID", func() {
		g.It("status not found", func() {
			srv := ServerMock(
				"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				`{"error": "Not Found"}`,
				http.StatusNotFound,
				false,
			)

			defer srv.Close()

			accountClient := NewAccountClient(30, srv.URL)

			_, err := accountClient.FetchOneByAccountID(
				context.TODO(),
				"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
			)

			g.Assert(strings.Contains(err.Error(), "Invalid response")).Equal(true)
		})

		g.It("status found", func() {
			srv := ServerMock(
				"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				`{"data":{"attributes":{"account_number":"41426819","alternative_names":["Sam Holder"],"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","iban":"GB11NWBK40030041426819","name":["Samantha Holder"],"secondary_identification":"A1B2C3D4"},"id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts"}}`,
				http.StatusOK,
				false,
			)

			defer srv.Close()

			accountClient := NewAccountClient(30, srv.URL)

			data, err := accountClient.FetchOneByAccountID(
				context.TODO(),
				"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
			)

			g.Assert(err).Equal(nil)
			g.Assert(data.Data.ID).Equal("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
		})
	})

	g.Describe("#DeleteOneByAccountID", func() {
		g.It("status no content", func() {
			srv := ServerMock(
				"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				`{}`,
				http.StatusNoContent,
				false,
			)

			defer srv.Close()

			accountClient := NewAccountClient(30, srv.URL)

			err := accountClient.DeleteOneByAccountID(
				context.TODO(),
				"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				1,
			)
			g.Assert(err).Equal(nil)
		})

		g.It("status not found", func() {
			srv := ServerMock(
				"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				`{}`,
				http.StatusNotFound,
				false,
			)

			defer srv.Close()

			accountClient := NewAccountClient(30, srv.URL)

			err := accountClient.DeleteOneByAccountID(
				context.TODO(),
				"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				1,
			)
			g.Assert(err).Equal(ErrResourceNotFound)
		})

		g.It("status conflict", func() {
			srv := ServerMock(
				"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				`{}`,
				http.StatusConflict,
				false,
			)

			defer srv.Close()

			accountClient := NewAccountClient(30, srv.URL)

			err := accountClient.DeleteOneByAccountID(
				context.TODO(),
				"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				1,
			)
			g.Assert(err).Equal(ErrVersionConflict)
		})

		g.It("300 status code", func() {
			srv := ServerMock(
				"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				`{"error": "Invalid Request"}`,
				300,
				false,
			)

			defer srv.Close()

			accountClient := NewAccountClient(30, srv.URL)

			err := accountClient.DeleteOneByAccountID(
				context.TODO(),
				"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				1,
			)
			g.Assert(strings.Contains(err.Error(), "300")).Equal(true)
		})
	})
}

// TestIntegrationAccountClientMethods test cases
func TestIntegrationAccountClientMethods(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#DataType", func() {
		g.It("It should satisfy all provided test cases", func() {

		})
	})
}
