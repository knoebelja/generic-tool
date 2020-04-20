package internal

import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func readYAML(filepath string, v interface{}) {

	data, err := ioutil.ReadFile(filepath)
	check(err, "reading JSON from %s", filepath)

	err = yaml.Unmarshal(data, v)
	check(err, "storing %s in %+v", string(data), v)
}

func readJSON(filepath string, v interface{}) {

	data, err := ioutil.ReadFile(filepath)
	check(err, "reading JSON from %s", filepath)

	err = json.Unmarshal(data, &v)
	check(err, "storing %s in %+v", string(data), v)
}
