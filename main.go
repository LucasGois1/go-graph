package main

import (
	"fmt"
	"graph/src/entities"
	"graph/src/graph"
)

func main() {
	ourinhos := entities.City{
		Identifier: entities.CityID{Id: 123},
		Name:       "Ourinhos",
	}
	node1 := graph.NewNode(ourinhos)

	jacarezinho := entities.City{
		Identifier: entities.CityID{Id: 234},
		Name:       "Jacarezinho",
	}
	node2 := graph.NewNode(jacarezinho)

	marilia := entities.City{
		Identifier: entities.CityID{Id: 345},
		Name:       "Marília",
	}
	node3 := graph.NewNode(marilia)

	bauru := entities.City{
		Identifier: entities.CityID{Id: 456},
		Name:       "Bauru",
	}
	node4 := graph.NewNode(bauru)

	saoPaulo := entities.City{
		Identifier: entities.CityID{Id: 567},
		Name:       "São Paulo",
	}
	node5 := graph.NewNode(saoPaulo)

	node1.AddEdge(node2, 100)
	node1.AddEdge(node3, 200)

	node2.AddEdge(node1, 100)
	node2.AddEdge(node3, 200)

	node3.AddEdge(node1, 200)
	node3.AddEdge(node2, 200)

	node4.AddEdge(node1, 300)
	node4.AddEdge(node3, 300)
	node4.AddEdge(node5, 400)

	node5.AddEdge(node1, 400)
	node5.AddEdge(node4, 400)

	g := graph.NewGraph[entities.City]()

	g.AddNode(node1)
	g.AddNode(node2)
	g.AddNode(node3)
	g.AddNode(node4)
	g.AddNode(node5)

	city := g.Search(jacarezinho, func(value entities.City) bool {
		return value.Name == "São Paulo"
	})

	fmt.Println(city)
}
