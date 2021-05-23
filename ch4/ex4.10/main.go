// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

const (
	LTAM = "less than a month"
	LTAY = "less than a year"
	MTAY = "more than a year"
)

//!+
func main() {
	var classIssues = make(map[string][]github.Issue)
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	year, month, _ := time.Now().Date()
	for _, item := range result.Items {
		item := *item
		i_year, i_month, _ := item.CreatedAt.Date()
		switch {
		case (month - i_month) <= 1:
			classIssues[LTAM] = append(classIssues[LTAM], item)

		case (year - i_year) <= 1:
			classIssues[LTAY] = append(classIssues[LTAY], item)

		case (year - i_year) > 1:
			classIssues[MTAY] = append(classIssues[MTAY], item)
		}
	}

	for class, issues := range classIssues {
		fmt.Printf("%s:\n %v\n", class, issues)

	}
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
