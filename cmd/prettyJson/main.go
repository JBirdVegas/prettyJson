package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jbirdvegas/prettyJson/prettyJson"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	fileArgDescription     = "json file to pretty print"
	collapseArgDescription = "unpretty print json"
	versionArgDescription  = "prints this program's version"
	shortHand              = " (shorthand)"
)

var AppVersion = "development"
var AppBuildTime = "unknown"

var fileFlag = flag.String("file", "", fileArgDescription)
var collapse = flag.Bool("collapse", false, collapseArgDescription)
var version = flag.Bool("version", false, versionArgDescription)

func init() {
	flag.StringVar(fileFlag, "f", "", fileArgDescription+shortHand)
	flag.BoolVar(collapse, "c", false, collapseArgDescription+shortHand)
	flag.BoolVar(version, "v", false, versionArgDescription+shortHand)
}

func main() {
	flag.Parse()

	if *version {
		filename := filepath.Base(os.Args[0])
		println(fmt.Sprintf("%s: (%s) %s", filename, AppBuildTime, AppVersion))
	}

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
