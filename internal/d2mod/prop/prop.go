package prop

import "strconv"

// Prop Property, which on an items is called Affix.  Loaded from UniqueItems, Sets, etc
type Prop struct {
	Name string
	Par  string
	Min  string
	Max  string
	Lvl  int
	Val  Val
}

// Val Atoi conversions of fields in Prop
type Val struct {
	Par int
	Min int
	Max int
}

// Props is a slice of Prop
type Props = []Prop

// NewProp Create a new Prop
func NewProp(name string, par string, min string, max string) Prop {
	/*
		if name == "sock" {
			fmt.Printf("NewProp: [%s][%s][%s][%s]", name, par, min, max)
		}
	*/
	prop := Prop{
		Name: name,
		Par:  par,
		Min:  min,
		Max:  max,
	}
	prop.Val.Par, _ = strconv.Atoi(prop.Par)
	prop.Val.Min, _ = strconv.Atoi(prop.Min)
	prop.Val.Max, _ = strconv.Atoi(prop.Max)
	return prop
}

// GetID Returns a Prop Name used to determine uniqueness
// This code is broken because of props that contain parameters, such as poison damage or blah/lvl
func (p *Prop) GetID() string {
	if p.Name == "aura" {
		// Two auras do not work even if they are different types
		return p.Name
	}
	// Otherwise include both the prop type and sub-type
	return p.Name + p.Par
}
