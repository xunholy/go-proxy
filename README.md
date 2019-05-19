[![GolangCI](https://golangci.com/badges/github.com/xUnholy/go-proxy.svg)](https://golangci.com)
[![Build Status](https://travis-ci.com/xUnholy/go-proxy.svg?branch=master)](https://travis-ci.com/xUnholy/go-proxy)
[![Coverage Status](https://coveralls.io/repos/github/xUnholy/go-proxy/badge.svg?branch=master)](https://coveralls.io/github/xUnholy/go-proxy?branch=master)
[![Maintainability](https://api.codeclimate.com/v1/badges/ff937aa58ba1d149d29a/maintainability)](https://codeclimate.com/github/xUnholy/go-proxy/maintainability)
[![License](https://img.shields.io/badge/license-GPL_v3.0-blue.svg)](https://github.com/xUnholy/go-proxy/blob/master/LICENSE.md)

# CNTLM Proxy Configuration Tool

This project is a CLI for managing proxy configuring on a local dev machine. To configure CTNLM and other dev tools can be troublesome and time consuming.

This tool will be able to dynamically set proxy configuration within CNTLM and other supported tools that have tool specific proxy support.

## Contents

* [Installing](#Installing)
* [Commands](#Commands)
* [Getting Started](#Getting-Started)
  * [Prerequisites](#Prerequisites)
  * [Running Tests](#Running-Tests)
  * [Build And Package](#Build-And-Package)
* [Built With](#Built-With)
* [Contributing](#Contributing)
* [Versioning](#Versioning)
* [Authors](#Authors)
* [License](#License)
* [Acknowledgments](#Acknowledgments)

## Installing

Use the following to be able to install on MacOS via Homebrew:

Running the below command will add the homebrew tap to our repository

```bash
brew tap xUnholy/homebrew-proxy
```

Now you've added our custom tap, you can download with the following command:

```bash
brew install proxy
```

To upgrade your proxy to the latest stable release use the following command:

```bash
brew upgrade proxy
```

*Note: Should the upgrade fail you may be required to remove the previous binary, the output will print what file to remove*

## Commands

The following is a list of all available commands:

* [Start](#Start)
* [Stop](#Stop)
* [Set](#Set)
* [Unset](#Unset)
* [Config](#Config)

### Start

Usage example:

```bash
proxy start [flags]
```

```bash
Flags:
  -h, --help        help for start
      --port PORT   set custom CNTLM PORT (default 3128)
```

### Stop

Usage example:

```bash
proxy stop [flags]
```

```bash
Flags:
      --all    unset all proxy configuration (default true)
  -h, --help   help for stop
```

### Set

Usage example:

```bash
proxy set [command]
```

Subcommand options:

```bash
Available Commands:
  domain      This command will update the domain value in your CNTLM.conf file
  git         This command will set the GIT global proxy values. Both http.proxy and https.proxy will be set
  npm         This command will set the NPM proxy values. Both https-proxy and proxy will be set
  password    This command will update the Password value in your CNTLM.conf file
  username    This command will update the Username value in your CNTLM.conf file
```

All subcommands have the following options:

```bash
Flags:
  -h, --help   help for set
```

To be able to set the CNTLM authentication password the *password* subcommand can be used. EG:

```bash
proxy set password
```

Once the password is entered and then encrypted, it will be stored in **cntlm.conf**.

### Unset

Usage example:

```bash
  proxy unset [command]
```

Subcommand options:

```bash
Available Commands:
  git         This command will unset the GIT global proxy values. Both http.proxy and https.proxy will be unset
  npm         This command will unset the NPM proxy values. Both https-proxy and proxy will be unset
```

All subcommands have the following options:

```bash
Flags:
  -h, --help   help for unset
```

#### Config

Usage example:

```bash
  proxy config [command]
```

Subcommand options:

```bash
Available Commands:
  list        This command will print the proxy configuration profile
  load        This command will load a custom proxy profile
  save        This command will save the proxy configuration to a profile
```

All subcommands have the following options:

```bash
Flags:
  -h, --help   help for unset
```

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing.

## Prerequisites

What things you need to install:

* [golang](https://golang.org/dl/)

## Running Tests

Use the following steps to execute the test suite:

```go
go test -v  ./...
```

OR

```bash
make test
```

*Note: The test suite will require you to have the supported tools installed and available locally.*

## Build And Package

To be able to package this tool, use one of the following below command:

```go
go build -o proxy .
```

OR

```bash
make build
```

Once you have packaged the application the compiled binary can be executed with the following:

```bash
proxy [command options] [arguments...]
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTION.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/xUnholy/go-proxy/tags).

## Authors

* **Michael Fornaro** - *Initial work* - [LinkedIn](https://www.linkedin.com/in/michael-fornaro-5b756179/)

See also the list of [contributors](https://github.com/xUnholy/go-proxy/contributors) who participated in this project.

## License

This project is licensed under the GPL License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

Wish to acknowledge the following people for their inspiration, and guidance with this project and/or in general:

* **Prateek Nayak** - [Github](https://github.com/prateeknayak)
* **Jasper Brekelmans** - [Github](https://github.com/jbrekelmans)
* **Mahesh Rayas** - [Github](https://github.com/maheshrayas)
