package crawli

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ColumnItem struct {
	Column   string `yaml:"column"`
	Primary  bool   `yaml:"primary,omitempty"`
	Selector string `yaml:"selector"`
	Type     string `yaml:"type"`
	Attr     string `yaml:"attr,omitempty"`
	Command  string `yaml:"command,omitempty"`
}

type Worksheet struct {
	Task struct {
		Label          string        `yaml:"label"`
		URL            string        `yaml:"url"`
		AllowedDomains []string      `yaml:"allowedDomains"`
		UserAgent      string        `yaml:"userAgent"`
		Delay          int           `yaml:"delay"`
		Async          int           `yaml:"async"`
		Trigger        string        `yaml:"trigger,omitempty"`
		Columns        []*ColumnItem `yaml:"columns"`
	} `yaml:"task"`
}

type WorksheetParser struct{}

func NewWorksheetParser() *WorksheetParser {
	return &WorksheetParser{}
}

func (p *WorksheetParser) Parse(fileName string) (*Worksheet, error) {
	workflow := &Worksheet{}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error when reading file: #%v ", err)
	}

	err = yaml.Unmarshal(file, workflow)
	if err != nil {
		fmt.Errorf("Error when unmarshalng: %v", err)
	}

	return workflow, err
}
