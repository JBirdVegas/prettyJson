package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			os.Exit(1)
		}
	}()

	fileFlag := flag.String("file", "", "json file to pretty print")
	collapse := flag.Bool("collapse", false, "unpretty print json")
	flag.Parse()

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
				checkError(err)
				prettyPrintBytes(contents, *collapse)
				break
			} else if _, err := os.Stat(arg); !os.IsNotExist(err) {
				contents, err := ioutil.ReadFile(arg)
				checkError(err)
				prettyPrintBytes(contents, *collapse)
			} else {
				var f interface{}
				err := json.Unmarshal([]byte(arg), &f)
				if err != nil {
					panic(err)
				}
				prettyPrintObject(f, *collapse)
			}
		}
		return
	}

	reader := bufio.NewReader(os.Stdin)
	allData, err := ioutil.ReadAll(reader)
	checkError(err)
	prettyPrintBytes(allData, *collapse)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func prettyPrintObject(data interface{}, indent bool) {
	d, err := json.Marshal(data)
	checkError(err)
	prettyPrintBytes(d, indent)
}

func prettyPrintBytes(data []byte, collapse bool) {
	prettyJSON := new(bytes.Buffer)
	if collapse {
		err := json.Compact(prettyJSON, data)
		checkError(err)
	} else {
		err := json.Indent(prettyJSON, data, "", "    ")
		checkError(err)
	}
	_, err := os.Stdout.WriteString(prettyJSON.String())
	checkError(err)
	_, err = os.Stdout.WriteString(LineBreak)
	checkError(err)
}
