Get a package <br />
`go get github.com/gorilla/mux`

List the packages <br />
`go list` <br />
`go list all` <br />
`go list -m all` <br />
`go list -m -versions github.com/gorilla/mux` - to check available version from the remote repo <br />

Verify dependency <br />
`go mod verify` - all modules verified <br />

Get a specific version of a module <br />
`go get github.com/gorilla/mux@v1.6.1` <br />

Getting master will download not released version + pre-release identifier (github.com/gorilla/mux v1.8.1-0.20221209155657-eb99d7a67714) <br />
`go get github.com/gorilla/mux@master` <br />
We can also use "latest" or the closes match to the specified version <br />
`go get 'github.com/gorilla/mux@<v1.7.0'"` <br />

Check if the module is used <br />
`go mod why github.com/gorilla/mux ` <br />

Check the graph of the modules <br />
`go mod graph ` <br />

`go mod edit -module github.com/dmosyan/Learning-Go/module` - modify the module name <br />
`go mod edit go 1.23.1` - modify go version of the app <br /> 
`go mod edit -replace github.com/gorilla/mux=../mux` - replace the module with the local copy <br />

Use `go mod vendor` to create a local `vendor` directory that contains external dependencies. To use that run `go run -mod=vendor .` which will start using the dependencies from the `vendor` directory.











