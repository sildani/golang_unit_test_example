# golang_unit_test_example

This a quick example of how to unit test in Go lang.

## Dependencies ##

Golang (see https://golang.org/doc/install to install)

## How to run the unit tests ##

1. After cloning the project, set the $GOPATH variable to the root of the clone.
1. Run `cd $GOPATH/src/silvanolte.com/mystring`
1. While in the mystring directory, run `go install` to build and install the mystring library. This will create a compiled module to test. The location will be in the `$GOPATH/pkg` directory.
1. Run `cd $GOPATH/src/silvanolte.com/mystring_test`
1. Run `go test` to run the unit tests.

## How to run a program that uses the library ##

1. Run `cd $GOPATH/src/silvanolte.com/hello`
1. While in the hello directory, run `go install` to build and install the hello program. This will create a compiled executable binary. The location will be in the `$GOPATH/bin` directory.
1. Run `$GOPATH/bin/hello` to execute the program.

## Learn more ##

- https://www.golang-book.com/books/intro/12
- https://www.youtube.com/watch?v=ndmB0bj7eyw
