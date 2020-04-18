package rw

import (
	"io/ioutil"

	"github.com/knoebelja/generic-tool/internal/utils"
	"gopkg.in/yaml.v2"
)

var readYAML read = func(filepath string) (v interface{}) {

	data, err := ioutil.ReadFile(filepath)
	utils.Check(err, "reading JSON from %s", filepath)

	err = yaml.Unmarshal(data, &v)
	utils.Check(err, "storing YAML in interface %+v:\n%s", v, string(data))

	return
}

var writeYAML write = func(filepath string, v interface{}) {

	data, err := yaml.Marshal(&v)
	utils.Check(err, "encoding YAML from %+v", v)

	err = ioutil.WriteFile(filepath, data, perm)
	utils.Check(err, "writing YAML to %s:\n%s", filepath, string(data))
}
