package main

import "fmt"

// type person struct {
// 	first string
// 	last  string
// 	age   int
// }

// type secretagent struct {
// 	person
// 	ltk bool
// }

// type person struct {
// 	first    string
// 	last     string
// 	favGames []string
// }

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	// p1 := secretagent{
	// 	person: person{
	// 		first: "aswin",
	// 		last:  "kumar",
	// 		age:   21,
	// 	},
	// 	ltk: true,
	// }
	// fmt.Println(p1)

	// a1 := struct {
	// 	price    int
	// 	quantity int
	// }{
	// 	price:    120,
	// 	quantity: 2,
	// }
	// fmt.Println(a1)
	// p1 := person{
	// 	first:    "Aswin",
	// 	last:     "Kumar",
	// 	favGames: []string{"bgmi", "gta-5"},
	// }
	// p2 := person{
	// 	first:    "Dakshina",
	// 	last:     "D",
	// 	favGames: []string{"angry birds", "candy crush"},
	// }

	// m := map[string]person{
	// 	p1.first: p1,
	// 	p2.first: p2,
	// }
	// for key, value := range m {
	// 	fmt.Printf("key : %v , value : %v \n", key, value)
	// }
	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "2020",
		model: "tesla",
		doors: 4,
		color: "black",
	}

	v2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "1998",
		model: "mustang",
		doors: 2,
		color: "red",
	}
	fmt.Println(v1.engine.electric)
	fmt.Println(v2)
}
