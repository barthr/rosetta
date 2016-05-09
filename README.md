# Rosetta-snippets-cli

[![Go Report Card](https://goreportcard.com/badge/github.com/barthr/rosetta)](https://goreportcard.com/report/github.com/barthr/rosetta)
[![Build Status](https://travis-ci.org/barthr/rosetta.svg?branch=master)](https://travis-ci.org/barthr/rosetta)

Cli application for http://rosettacode.org/wiki/Rosetta_Code, it opens the webpage or returns the url from the search term and also goes to the specific language

```sh
$ go get github.com/barthr/rosetta
```

## Examples


### Settings 

**Setting a default language**
```sh
rosetta language Go
```


**Displaying ur current settings**
```sh
rosetta settings
```

**Reset settings**
```sh
rosetta reset
```


### Flags

```sh
search:
   -r : only return url
   -l : specify language (overrides the default language)
```


### Search

**Search on hello with language java and open Web browser**
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

After entering the search u will be prompted to choose one from the selection

### Extra information
- Rosetta caches the searches!
