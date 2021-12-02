# Advent of Code 2021
Learning Go while completing AOC 2021

## Install Go
https://golang.org/doc/install

## Run
```bash
# In directory
go run .
```

## Init Module
```bash
go mod init aoc/fileloading
```

## Update Modules
```bash
go mod tidy
```

### Add Local Module
```bash
go mod edit -replace aoc/fileloading=../fileloading
go mod tidy
```
