# Form3 Exercise


## Usage



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
