# Rosetta-snippets-cli

[![Go Report Card](https://goreportcard.com/badge/github.com/barthr/rosetta)](https://goreportcard.com/report/github.com/barthr/rosetta)


Cli application for http://rosettacode.org/wiki/Rosetta_Code, it opens the webpage or returns the url from the search term and also goes to the specific language


```sh
$ go get github.com/barthr/rosetta
```

### Example
```sh
rosetta language java
rosetta search hello
```

### OR
```sh
rosetta search -l java hello
```


After entering the search u will be prompted to choose one from the selection


### Work in progress
- Caching
