# Form3 Exercise


## Installation

To install `f3sdk` package, you need to install Go (version 1.16+ is required) and setup your Go project first.

```zsh
$ mkdir example
$ cd example
$ go mod init example.com
```

Then you can use the below Go command to install the latest version of `f3sdk` package.

```zsh
$ go get -u github.com/clivern/f3sdk
```

Or the following for a specific version `v0.1.0`

```zsh
go get -u github.com/clivern/f3sdk@v0.1.0
```

Import the package in your code.

```zsh
import "github.com/clivern/f3sdk"
```

*Please note that: the sdk is private so the above steps not applicable. You have to clone and define the local path in `go.mod`*


## Quick start

Here is an example showing how to fetch, delete and create accounts using the sdk. Please check `example` directory for the full example.

```golang
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
```


## Testing

To run unit test cases locally without docker

```zsh
$ git clone git@github.com:Clivern/f3sdk.git
$ cd f3sdk
$ make install_revive
$ make ci

# To auto format the code
$ make format

# For all commands
$ make
```

In order to run test cases with docker and docker compose

```zsh
# Install docker and docker compose on Ubuntu for example
$ apt-get update
$ apt-get install docker.io -y
$ systemctl enable docker

$ apt-get install docker-compose -y

# Clone the solution
$ git clone git@github.com:Clivern/f3sdk.git
$ cd f3sdk

# Run dependencies
$ docker-compose up -d accountapi postgresql vault

# Check if they are healthy
$ docker ps

# Then run sdk test cases
$ docker-compose run f3sdk
```


## License

Copyright 2019-2021 Form3 Financial Cloud

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
