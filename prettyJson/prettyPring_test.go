package prettyJson

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestCollapse(t *testing.T) {
	prettyFileContents, err := ioutil.ReadFile("testPretty.json")
	if err != nil {
		panic(err)
	}
	prettyString := strings.Trim(string(prettyFileContents), " \r\n")

	collapsedFileContents, err := ioutil.ReadFile("test.json")
	if err != nil {
		panic(err)
	}

	collapsedString := strings.Trim(string(collapsedFileContents), " \r\n")
	collapsedJson := strings.Trim(createNewJsonString([]byte(prettyString), true, false).String(), " \r\n")

	if collapsedString != collapsedJson {
		t.Fatal(fmt.Sprintf("No match: %s, %s", string(collapsedFileContents), collapsedJson))
	}
}

func TestPretty(t *testing.T) {
	testPrettyJson, err := ioutil.ReadFile("testPretty.json")
	if err != nil {
		panic(err)
	}
	trimmedPrettyJson := strings.Trim(string(testPrettyJson), " \r\n")

	testJson, err := ioutil.ReadFile("test.json")
	if err != nil {
		panic(err)
	}

	trimmedCollapsedJson := strings.Trim(string(testJson), " \r\n")
	prettyJson := strings.Trim(createNewJsonString([]byte(trimmedCollapsedJson), false, false).String(), " \r\n")

	if trimmedPrettyJson != prettyJson {
		t.Fatal(fmt.Sprintf("No match: %s, %s", trimmedPrettyJson, prettyJson))
	}
}
