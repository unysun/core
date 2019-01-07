package propertyresolver

import (
	"encoding/json"
	"github.com/project-flogo/core/data/property"
	"io/ioutil"
	"os"
	"strings"

	"github.com/project-flogo/core/support/log"
)

var preload = make(map[string]interface{})

//var log = logger.GetLogger("app-props-json-resolver")

// Comma separated list of json files overriding default application property values
// e.g. FLOGO_APP_PROPS_JSON=app1.json,common.json
const EnvAppPropertyFileConfigKey = "FLOGO_APP_PROPS_JSON"

func init() {

	logger := log.RootLogger()

	filePaths := getExternalFiles()
	if filePaths != "" {
		// Register value resolver
		property.RegisterExternalResolver("json", &JSONFileValueResolver{})

		// preload props from files
		files := strings.Split(filePaths, ",")
		if len(files) > 0 {
			for _, filePath := range files {
				props := make(map[string]interface{})

				file, e := ioutil.ReadFile(filePath)
				if e != nil {
					logger.Errorf("Can not read - %s due to error - %v", filePath, e)
					panic("")
				}
				e = json.Unmarshal(file, &props)
				if e != nil {
					logger.Errorf("Can not read - %s due to error - %v", filePath, e)
					panic("")
				}
			}
		}
	}
}

func getExternalFiles() string {
	key := os.Getenv(EnvAppPropertyFileConfigKey)
	if len(key) > 0 {
		return key
	}
	return ""
}

// Resolve property value from external files
type JSONFileValueResolver struct {
}

func (resolver *JSONFileValueResolver) LookupValue(toResolve string) (interface{}, bool) {
	val, found := preload[toResolve]
	return val, found
}
