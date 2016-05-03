# Rosetta-snippets-cli

[![Go Report Card](https://goreportcard.com/badge/github.com/barthr/rosetta)](https://goreportcard.com/report/github.com/barthr/rosetta)


Cli application for http://rosettacode.org/wiki/Rosetta_Code, it opens the webpage or returns the url from the search term and also goes to the specific language


```sh
$ go get github.com/barthr/rosetta
```

## Examples

**Search on hello with the language from settings**
```sh
rosetta search hello
```

**Search on hello with language java and open <INSERT WEB BROWSER HERE>**
```sh
rosetta search -l java hello
```

**Search on hello with language java and return the url**
```sh
rosetta search -l java -r hello
```

**Search on hello with default language settings and only return the url**
```sh
rosetta search -r hello
```

**Reset settings**
```sh
rosetta reset
```


After entering the search u will be prompted to choose one from the selection


### Work in progress
- Caching
