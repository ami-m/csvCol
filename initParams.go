package main

import (
	"flag"
	"strconv"
	"strings"
)

type runParams struct {
	file      string
	separator rune
	invert    bool
	cols      []int
	colNames  []string
}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return strings.Join(*i, ", ")
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func initParams() runParams {
	var res runParams
	var file string
	var separator string
	var invert bool
	var cols arrayFlags
	var colNames arrayFlags

	flag.StringVar(&file, "f", "", "path to input file instead of stdin")
	flag.StringVar(&separator, "s", ",", "separator character (defaults to a comma)")
	flag.BoolVar(&invert, "v", false, "hide (instead of showing) the columns that were selected")
	flag.Var(&cols, "c", "list of columns to show by index (zero based)")
	flag.Var(&colNames, "C", "list of columns to show by name")

	flag.Parse()

	res.file = file
	res.invert = invert

	var actualCols []int
	for _, v := range cols {
		if index, err := strconv.Atoi(v); err == nil {
			actualCols = append(actualCols, index)
		}
	}
	res.cols = actualCols

	var actualColNames []string
	for _, v := range colNames {
		actualColNames = append(actualColNames, v)
	}
	res.colNames = actualColNames

	separatorRunes := []rune(separator)
	res.separator = separatorRunes[0]

	return res
}
