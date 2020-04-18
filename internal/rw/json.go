package rw

import (
	"encoding/json"
	"io/ioutil"

	"github.com/knoebelja/generic-tool/internal/utils"
)

var readJSON read = func(filepath string) (v interface{}) {

	data, err := ioutil.ReadFile(filepath)
	utils.Check(err, "reading JSON from %s", filepath)

	err = json.Unmarshal(data, &v)
	utils.Check(err, "storing JSON in %+v:\n%s", v, string(data))

	return
}

var writeJSON write = func(filepath string, v interface{}) {

	data, err := json.Marshal(&v)
	utils.Check(err, "encoding JSON from %+v", v)

	err = ioutil.WriteFile(filepath, data, perm)
	utils.Check(err, "writing JSON to %s:\n%s", filepath, string(data))
}
