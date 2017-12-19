# shakespeare-analyzer

## Set Up
Install go via instructions: https://golang.org/doc/install

```
go get -u github.com/kardianos/govendor
govendor sync
```

## Run Analyzer in Console Mode
`go run main.go -c`

## Run Analyzer in Server Mode
`go run main.go -s`

Navigate to `localhost:8080`

## Test Code
`go test github.com/paddyquinn/shakespeare-analyzer/analyzer`