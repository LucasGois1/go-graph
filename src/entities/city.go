package entities

import "graph/src/value_object"

type CityID struct {
	Id int
}

func (c CityID) Value() any {
	return c.Id
}

type City struct {
	Identifier CityID
	Name       string
}

func (c City) Id() value_object.Identifier[any] {
	return c.Identifier
}
