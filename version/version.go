package version

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const srvRoot = "/srv/"

var (
	buildID           = "unknown"
	buildDate         = "unknown"
	dependencies      = json.RawMessage("")
	loadedVersionInfo = false
)

// Line returns a single line version
func Line() string {
	loadVersion()
	return fmt.Sprintf("BuildID: %s, BuildDate: %s", buildID, buildDate)
}

// Short returns a short version
func Short(serviceName string) map[string]interface{} {
	loadVersion()
	return map[string]interface{}{
		"name":      serviceName,
		"buildDate": buildDate,
		"version":   buildID,
	}
}

func loadVersion() {
	if loadedVersionInfo {
		return
	}
	loadJSON(srvRoot+"dependencies.json", &dependencies)
	loadedVersionInfo = true
}

func loadJSON(path string, out *json.RawMessage) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		*out = json.RawMessage("missing")
	}
	*out = json.RawMessage(content)
}
