package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"github.com/jbirdvegas/prettyJson/prettyJson"
	"io/ioutil"
	"os"
	"strings"
)

const (
	fileArgDescription     = "json file to pretty print"
	collapseArgDescription = "unpretty print json"
	shortHand              = " (shorthand)"
)

var fileFlag = flag.String("file", "", fileArgDescription)
var collapse = flag.Bool("collapse", false, collapseArgDescription)

func init() {
	flag.StringVar(fileFlag, "f", "", fileArgDescription+shortHand)
	flag.BoolVar(collapse, "c", false, collapseArgDescription+shortHand)
}

func main() {
	flag.Parse()
	defer func() {
		if r := recover(); r != nil {
			os.Exit(1)
		}
	}()

	args := os.Args[1:len(os.Args)]

	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if info.Size() == 0 {
		for _, arg := range args {
			if strings.Contains(arg, "collapse") {
				continue
			} else if fileFlag != nil && *fileFlag != "" {
				contents, err := ioutil.ReadFile(*fileFlag)
				prettyJson.CheckError(err)
				prettyJson.PrettyPrintBytes(contents, *collapse)
				break
			} else if _, err := os.Stat(arg); !os.IsNotExist(err) {
				contents, err := ioutil.ReadFile(arg)
				prettyJson.CheckError(err)
				prettyJson.PrettyPrintBytes(contents, *collapse)
			} else {
				var f interface{}
				err := json.Unmarshal([]byte(arg), &f)
				if err != nil {
					panic(err)
				}
				prettyJson.PrettyPrintObject(f, *collapse)
			}
		}
		return
	}

	reader := bufio.NewReader(os.Stdin)
	allData, err := ioutil.ReadAll(reader)
	prettyJson.CheckError(err)
	prettyJson.PrettyPrintBytes(allData, *collapse)
}
