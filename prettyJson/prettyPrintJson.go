package prettyJson

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

func CheckError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func PrettyPrintObject(data interface{}, indent bool, addColor bool) {
	d, err := json.Marshal(data)
	CheckError(err)
	PrettyPrintBytes(d, indent, addColor)
}

func PrettyPrintBytes(data []byte, collapse bool, addColor bool) {
	s := createNewJsonString(data, collapse, addColor).String()
	_, err := os.Stdout.WriteString(s)
	CheckError(err)

}

func createNewJsonString(data []byte, collapse bool, addColor bool) *bytes.Buffer {
	prettyJSON := new(bytes.Buffer)
	if collapse {
		err := json.Compact(prettyJSON, data)
		CheckError(err)
	} else {
		if addColor {
			err := ColorfulIndent(prettyJSON, data, "", "    ")
			CheckError(err)
		} else {
			err := json.Indent(prettyJSON, data, "", "    ")
			CheckError(err)
		}

	}
	prettyJSON.Write([]byte(LineBreak))
	return prettyJSON
}
