[![GolangCI](https://golangci.com/badges/github.com/xUnholy/go-proxy.svg)](https://golangci.com)
[![Build Status](https://travis-ci.com/xUnholy/go-proxy.svg?branch=develop)](https://travis-ci.com/xUnholy/go-proxy)
[![Coverage Status](https://coveralls.io/repos/github/xUnholy/go-proxy/badge.svg?branch=develop)](https://coveralls.io/github/xUnholy/go-proxy?branch=develop)

# CNTLM Proxy Configuration Tool

This project is a CLI for managing proxy configuring on a local dev machine. To configure CTNLM and other dev tools can be troublesome and time consuming. 

This tool will be able to dynamically set proxy configuration within CNTLM and other supported tools that have tool specific proxy support. 

The following tools are currently supported:
* git
* npm

## Contents
------

* [Getting Started](#Getting-Started)
* [Prerequisites](#Prerequisites)
* [Commands](#Commands)
* [Running Tests](#Running-Tests)
* [Package](#Package) 
* [Built With](#Built-With)
* [Contributing](#Contributing)
* [Versioning](#Versioning)
* [Authors](#Authors)
* [License](#License)
* [Acknowledgments](#Acknowledgments)
   
## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing.

## Prerequisites

What things you need to install.

* [Golang](https://golang.org/dl/)

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

```
proxy stop [command options] [arguments...]
```

```
OPTIONS:
   --all, -a  unset all CNTLM config
```

### Set

Usage example:

```
proxy set command [command options] [arguments...]
```

Subcommand options:

```
COMMANDS:
     npm       set npm proxy config
     git       set git proxy config
     password  proxy set password
```

All subcommands have the following options:

```
OPTIONS:
   --port PORT, -p PORT  set custom CNTLM PORT (default: 3128)
```

To be able to set the CNTLM authentication password the *password* subcommand can be used. EG:

```
proxy set password
```

Once the password is entered and then encrypted, it will be stored in **cntlm.conf**.


### Unset

Usage example:

```
proxy unset command [command options] [arguments...]
```

Subcommand options:

```
COMMANDS:
     npm  unset npm proxy config
     git  unset git proxy config
```

All subcommands have the following options:

```
OPTIONS:
   --port PORT, -p PORT  set custom CNTLM PORT (default: 3128)
```

## Running Tests

Use the following steps to execute the test suite:

```go
go test -v  ./...
```

*Note: The test suite will require you to have the supported tools installed and available locally.*

## Package

To be able to package this tool, use the below command:

```go
go build -o proxy .
```

Alternatively you can build the tool using Docker & running the container.

```bash
docker build -t cntlm-proxy:latest .
```

Executing the proxy tool within a running docker container:

```
docker run --network=host -ti <image-name> [command options] [arguments...] 
```

*Note: Running inside the container functionality is still a WIP.*

## Built With

* [GOlang](https://golang.org/dl/) - Programming Language
* [urfave/cli](https://github.com/urfave/cli) - The CLI framework
* [Travis](https://maven.apache.org/) - Conitnuous Integration Tool
* [VSCode](https://code.visualstudio.com/) - IDE

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/xUnholy/CONTRIBUTION.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Michael Fornaro** - *Initial work* - [Linkedin](https://www.linkedin.com/in/michael-fornaro-5b756179/)

See also the list of [contributors](https://github.com/xUnholy/go-proxy/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

Wish to aknowledge the following groups of people for their inspiration, and guidance with this project:

* **Prateek Nayak** - [Innablr](https://innablr.com.au/)
* **Jasper Brekelmans** - [Cognizant](https://www.cognizant.com/)
