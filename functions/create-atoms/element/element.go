package element

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
)

type Element struct {
	Number    int32  `yaml:"number"`
	Symbol    string `yaml:"symbol"`
	Module    string `yaml:"module"`
	Name      string `yaml:"element"`
	Group     int32  `yaml:"group"`
	Period    int32  `yaml:"period"`
	Weight    string `yaml:"weight"`
	Density   string `yaml:"density"`
	Melt      string `yaml:"melt"`
	Boil      string `yaml:"boil"`
	C         string `yaml:"C"`
	X         string `yaml:"X"`
	Abundance string `yaml:"abundance"`
	Property  string `yaml:"property"`
}

type Elements struct {
	Source   string    `yaml:"source"`
	Elements []Element `yaml:"elements"`
}

func ReadElements() Elements {
	content, err := os.ReadFile("/etc/elements.yaml")
	if err != nil {
		content, err = os.ReadFile("elements.yaml")
		if err != nil {
			log.Fatal(err)
		}
	}

	e := Elements{}
	err = yaml.Unmarshal([]byte(content), &e)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return e
}

type Atom struct {
	Name   string
	Symbol string
	Number int32
	Period int32
	Group  int32
}

type Atoms struct {
	ElementByName   map[string]Atom
	ElementByNumber map[int32]Atom
}

func (a *Atoms) read() {
	a.ElementByName = make(map[string]Atom)
	a.ElementByNumber = make(map[int32]Atom)
	elements := ReadElements()

	for _, e := range elements.Elements {
		atom := Atom{
			Symbol: e.Symbol,
			Name:   e.Name,
			Number: e.Number,
			Period: e.Period,
			Group:  e.Group,
		}
		a.ElementByName[strings.ToLower(e.Symbol)] = atom
		a.ElementByNumber[e.Number] = atom
	}
}

func NewAtoms() *Atoms {
	atom := &Atoms{}
	atom.read()
	return atom
}
