package internal

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const perm os.FileMode = 0777

func writeYAML(filepath string, v interface{}) {

	data, err := yaml.Marshal(&v)
	check(err, "encoding YAML from %+v", v)

	err = ioutil.WriteFile(filepath, data, perm)
	check(err, "writing %s to %s", string(data), filepath)
}

func writeJSON(filepath string, v interface{}) {

	data, err := json.Marshal(&v)
	check(err, "encoding JSON from %+v", v)

	err = ioutil.WriteFile(filepath, data, perm)
	check(err, "writing %s to %s", string(data), filepath)
}
