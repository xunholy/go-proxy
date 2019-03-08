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

When a new release is available you can update your proxy using by running:

```bash
brew upgrade proxy
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

### Start

Usage example:

```
proxy start [command options] [arguments...]
```

```
OPTIONS:
   --port PORT, -p PORT  set custom CNTLM PORT (default: 3128)
   --all, -a             set all CNTLM config
```

### Stop

Usage example:

```bash
proxy stop [command options] [arguments...]
```

```bash
OPTIONS:
   --all, -a  unset all CNTLM config
```

### Set

Usage example:

```bash
proxy set command [command options] [arguments...]
```

Subcommand options:

```bash
COMMANDS:
     npm       set npm proxy config
     git       set git proxy config
     username  proxy set username
     password  proxy set password
     domain    proxy set domain
```

All subcommands have the following options:

```bash
OPTIONS:
   --port PORT, -p PORT  set custom CNTLM PORT (default: 3128)
```

To be able to set the CNTLM authentication password the *password* subcommand can be used. EG:

```bash
proxy set password
```

Once the password is entered and then encrypted, it will be stored in **cntlm.conf**.


### Unset

Usage example:

```bash
proxy unset command [command options] [arguments...]
```

Subcommand options:

```bash
COMMANDS:
     npm  unset npm proxy config
     git  unset git proxy config
```

All subcommands have the following options:

```bash
OPTIONS:
   --port PORT, -p PORT  set custom CNTLM PORT (default: 3128)
```

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing.

## Prerequisites

What things you need to install:

* [Golang](https://golang.org/dl/)

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

## Built With

* [GOlang](https://golang.org/dl/) - Programming Language
* [urfave/cli](https://github.com/urfave/cli) - The CLI framework
* [Travis](https://maven.apache.org/) - Continuous Integration Tool
* [VSCode](https://code.visualstudio.com/) - IDE

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/xUnholy/CONTRIBUTION.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Michael Fornaro** - *Initial work* - [LinkedIn](https://www.linkedin.com/in/michael-fornaro-5b756179/)

See also the list of [contributors](https://github.com/xUnholy/go-proxy/contributors) who participated in this project.

## License

This project is licensed under the GPL License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

Wish to acknowledge the following people for their inspiration, and guidance with this project and/or in general:

* **Prateek Nayak** - [Github](https://github.com/prateeknayak)
* **Jasper Brekelmans** - [Github](https://github.com/jbrekelmans)
