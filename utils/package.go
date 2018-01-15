package utils

import (
	"encoding/json"
	"io/ioutil"
	"path"
)

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func ReadPackageInfo(p string) (Package, error) {
	var pak Package
	raw, err := ioutil.ReadFile(path.Join(p, "package.json"))
	if err != nil {
		return pak, err
	}
	json.Unmarshal(raw, &pak)
	return pak, nil
}
