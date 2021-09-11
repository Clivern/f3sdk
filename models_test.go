// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package f3sdk

import (
	"strings"
	"testing"

	"github.com/franela/goblin"
)

// TestUnitModels test cases
func TestUnitModels(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#DataType", func() {
		g.It("It should satisfy all provided test cases", func() {
			data := &Data{
				Data: &AccountData{
					ID:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
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

			result, err := data.ConvertToJSON()

			g.Assert(err).Equal(nil)
			g.Assert(strings.Contains(result, "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")).Equal(true)
			g.Assert(strings.Contains(result, "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")).Equal(true)
			g.Assert(strings.Contains(result, "accounts")).Equal(true)
			g.Assert(strings.Contains(result, "41426819")).Equal(true)
			g.Assert(strings.Contains(result, "Sam Holder")).Equal(true)
			g.Assert(strings.Contains(result, "400300")).Equal(true)
			g.Assert(strings.Contains(result, "GBDSC")).Equal(true)
			g.Assert(strings.Contains(result, "GBP")).Equal(true)
			g.Assert(strings.Contains(result, "NWBKGB22")).Equal(true)
			g.Assert(strings.Contains(result, "GB11NWBK40030041426819")).Equal(true)
			g.Assert(strings.Contains(result, "Samantha Holder")).Equal(true)
			g.Assert(strings.Contains(result, "A1B2C3D4")).Equal(true)

			err = data.LoadFromJSON([]byte(`{"data":{"attributes":{"account_number":"41426819","alternative_names":["Sam Holder"],"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","iban":"GB11NWBK40030041426819","name":["Samantha Holder"],"secondary_identification":"A1B2C3D4"},"id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts"}}`))

			g.Assert(err).Equal(nil)

			g.Assert(data.Data.ID).Equal("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
			g.Assert(data.Data.OrganisationID).Equal("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
			g.Assert(data.Data.Type).Equal("accounts")
			g.Assert(data.Data.Attributes.AccountNumber).Equal("41426819")
			g.Assert(data.Data.Attributes.AlternativeNames[0]).Equal("Sam Holder")
			g.Assert(data.Data.Attributes.BankID).Equal("400300")
			g.Assert(data.Data.Attributes.BankIDCode).Equal("GBDSC")
			g.Assert(data.Data.Attributes.BaseCurrency).Equal("GBP")
			g.Assert(data.Data.Attributes.Bic).Equal("NWBKGB22")
			g.Assert(data.Data.Attributes.Iban).Equal("GB11NWBK40030041426819")
			g.Assert(data.Data.Attributes.Name[0]).Equal("Samantha Holder")
			g.Assert(data.Data.Attributes.SecondaryIdentification).Equal("A1B2C3D4")
		})
	})
}
