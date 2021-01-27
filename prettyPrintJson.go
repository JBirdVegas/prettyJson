package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			os.Exit(1)
		}
	}()

	fileFlag := flag.String("file", "", "json file to pretty print")
	flag.Parse()

	args := os.Args[1:len(os.Args)]

	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if info.Size() == 0 {
		for _, arg := range args {
			if fileFlag != nil && *fileFlag != "" {
				contents, err := ioutil.ReadFile(*fileFlag)
				checkError(err)
				prettyPrintBytes(contents)
				break
			} else if _, err := os.Stat(arg); !os.IsNotExist(err) {
				contents, err := ioutil.ReadFile(arg)
				checkError(err)
				prettyPrintBytes(contents)
			} else {
				var f interface{}
				err := json.Unmarshal([]byte(arg), &f)
				if err != nil {
					panic(err)
				}
				prettyPrintObject(f)
			}
		}
		return
	}

	reader := bufio.NewReader(os.Stdin)
	allData, err := ioutil.ReadAll(reader)
	checkError(err)
	prettyPrintBytes(allData)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func prettyPrintObject(data interface{}) {
	d, err := json.Marshal(data)
	checkError(err)
	prettyPrintBytes(d)
}

func prettyPrintBytes(data []byte) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, data, "", "    ")
	checkError(err)
	_, err = os.Stdout.WriteString(prettyJSON.String())
	checkError(err)
	_, err = os.Stdout.WriteString(LineBreak)
	checkError(err)
}
