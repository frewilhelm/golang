package menu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Item struct {
	Id   string
	Name string
	Desc string
}

type Menu struct {
	Items []Item `json:"pizza"`
}

func Create(p string) (Menu, error) {
	var m Menu
	mJson, err := os.Open(p)
	if err != nil {
		return m, fmt.Errorf("menu: error opening json-file '%s': %w", p, err)
	}
	defer mJson.Close()

	bJson, err := ioutil.ReadAll(mJson)
	if err != nil {
		return m, fmt.Errorf("menu: error reading json-file '%s': %w", p, err)
	}

	if err := json.Unmarshal([]byte(bJson), &m); err != nil {
		return m, fmt.Errorf("menu: error unmarshalling json-file '%s': %w", p, err)
	}

	return m, nil
}

func (m Menu) String() string {
	var s string
	for _, i := range m.Items {
		s += fmt.Sprintf("\t%s:\t%s (%s)\n", i.Id, i.Name, i.Desc)
	}
	return s
}
