# csvCol

shows/hides columns of csv files

## Getting Started

Build the executable with `go build`.

`./csvCol --help` will give you this:
```shell script
Usage of csvCol:
  -c value
        list of columns to show
  -f string
        path to input file instead of stdin
  -s string
        separator character (defaults to a comma) (default ",")
  -v    hide (instead of showing) the columns that were selected
```
## Examples
* show the first two columns:
```shell script
./csvCol -c=0 -c=1 -f="./customers.csv"
```

* show all columns *except* the first two:
```shell script
./csvCol -v -c=0 -c=1 -f="./customers.csv"
```

## Built With

* [Golang](https://golang.org/) - The go language
