package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args

	var filepath string

	if len(args) == 1 {
		filepath = "pubspec.lock"
	} else {
		filepath = args[1]
	}

	params := make(map[string]interface{})
	bytes, _ := ioutil.ReadFile(filepath)
	_ = yaml.Unmarshal(bytes, &params)

	packages := params["packages"].(map[interface{}]interface{})

	fmt.Println("dependencies:")
	for k, v := range packages {
		name := k
		var version string

		itemMap := v.(map[interface{}]interface{})
		switch itemMap["source"] {
		case "path":
			version = "\n" +
				"    path: " +
				itemMap["description"].(map[interface{}]interface{})["path"].(string)
			break
		case "hosted":
			version = itemMap["version"].(string)
			break
		case "git":
			gitMap := itemMap["description"].(map[interface{}]interface{})
			version = "\n    git:\n" +
				"      url: " + gitMap["url"].(string)
			if gitMap["ref"] != nil {
				version = version + "\n      ref: " +
					gitMap["resolved-ref"].(string)
			}
		default:
			continue
		}

		fmt.Printf("  %s: %s\n", name, version)
	}
}
