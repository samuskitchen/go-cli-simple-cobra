# Go CLI Simple With Cobra

This is the basic addition operation cli built in golang using cobra library based on the article of a [blog](https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177).

## Getting started

To install the library, run

```sh
go get -u github.com/samuskitchen/go-cli-simple-cobra
```

## Use the cli

```sh
go-cli-simple-cobra --help

Simple example of a command line with the Cobra library, which uses simple mathematical operations

Usage:
  go-cli-simple-cobra [command]

Available Commands:
  add         Add number
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.go-cli-simple-cobra.yaml)
  -h, --help            help for go-cli-simple-cobra
  -t, --toggle          Help message for toggle

Use "go-cli-simple-cobra [command] --help" for more information about a command.
```

##### Use command add

```sh
go-cli-simple-cobra add --help
You can add whole numbers and decimal numbers through the '-f' or '--float' tag

Usage:
  go-cli-simple-cobra add [flags]
  go-cli-simple-cobra add [command]

Available Commands:
  even        Add even numbers
  odd         Add odd numbers

Flags:
  -f, --float   Add Floating Numbers
  -h, --help    help for add

Global Flags:
      --config string   config file (default is $HOME/.go-cli-simple-cobra.yaml)

Use "go-cli-simple-cobra add [command] --help" for more information about a command.
```