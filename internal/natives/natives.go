package natives

import (
	"encoding/json"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	langTypes = map[string]string{
		"bool":    "boolean",
		"integer": "number",
		"number":  "number",
		"float":   "number",
		"string":  "string",
		"vector3": "{ x: number, y: number, z: number }",
		"entity":  "number",
		"player":  "number",
		"ped":     "number",
	}
)

type List map[string]Namespace

func (l List) Namespaces() []string {
	var namespaces []string
	for k := range l {
		namespaces = append(namespaces, k)
	}

	return namespaces
}

func (l List) Natives() []*Native {
	var n []*Native
	for _, v := range l {
		for _, v2 := range v {
			if v2.Name == "" {
				continue
			}

			n = append(n, v2)
		}
	}

	return n
}

type Namespace map[string]*Native

type Native struct {
	Name        string   `json:"name"`
	Params      Params   `json:"params"`
	Results     string   `json:"results"`
	Description string   `json:"description"`
	Hash        string   `json:"hash"`
	Namespace   string   `json:"ns"`
	Aliases     []string `json:"aliases"`
	JHash       string   `json:"jhash"`
	ManualHash  bool     `json:"manualHash"`
}

func (p *Native) Function() string {
	name := strings.ReplaceAll(p.Name, "_", " ")
	name = cases.Title(language.English).String(name)
	name = strings.ReplaceAll(name, " ", "")

	return name
}

func (p *Native) Comment() string {
	comment := strings.ReplaceAll(p.Description, "\n", "\n--- ")
	comment = strings.ReplaceAll(comment, "`", "")

	return comment
}

func (p *Native) Args() Params {
	var args Params

	for _, v := range p.Params {
		if v.IsReturn() {
			continue
		}

		args = append(args, v)
	}

	return args
}

func (p *Native) Returns() Params {
	var returns Params

	for _, v := range p.Params {
		if v.IsReturn() {
			returns = append(returns, v)
		}
	}

	if len(returns) == 0 {
		returns = append(returns, &Param{
			Name: "result",
			Type: p.Results,
		})
	}

	return returns
}

type Params []*Param

func (p Params) LastIndex() int {
	return len(p) - 1
}

type Param struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (p *Param) IsReturn() bool {
	return p.IsPointer()
}

func (p *Param) IsPointer() bool {
	return strings.HasPrefix(p.Type, "*")
}

func (p *Param) LangType() string {
	if v, ok := langTypes[strings.ToLower(p.Type)]; ok {
		return v
	}

	return "any"
}

func FromJSON(b []byte) (List, error) {
	var list List
	err := json.Unmarshal(b, &list)
	return list, err
}
