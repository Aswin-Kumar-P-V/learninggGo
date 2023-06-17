package main

import "fmt"

func main() {
	// a := map[string]int{
	// 	"me":  21,
	// 	"sis": 18,
	// }
	// fmt.Println(a["me"])

	// b := make(map[string]int)

	// b["hello"] = 1
	// b["world"] = 2

	// for key, val := range b {
	// 	fmt.Println(key, val)
	// }
	// delete(b, "hello")
	// if val, ok := b["world"]; !ok {
	// 	fmt.Println("value doesnt exits")
	// } else {
	// 	fmt.Println(val)
	// }
	// m := make(map[string][]string)
	// m["aswin"] = []string{"kumar", "p", "v"}
	// m["Deskina"] = []string{"d"}
	// m["test"] = []string{"asas", "adasfa"}

	// fmt.Println(m)

	// for k, v := range m {
	// 	fmt.Println(k)
	// 	for _, v2 := range v {
	// 		fmt.Println(v2)
	// 	}
	// }
	// delete(m, "test")
	// fmt.Println("----------------")
	// for k, v := range m {
	// 	fmt.Println(k)
	// 	for _, v2 := range v {
	// 		fmt.Println(v2)
	// 	}
	// }

	xs := []string{"apple", "banana", "cherry", "dog", "elephant", "cat", "dog", "frog", "grape", "apple", "bird", "horse", "kiwi", "lemon", "mango", "orange", "pear", "cat", "dog", "rabbit", "strawberry", "watermelon", "apple", "lemon", "avocado", "dog", "banana", "cherry", "cat", "grape", "elephant", "frog", "orange", "kiwi", "lemon", "mango", "apple", "pear", "strawberry", "watermelon", "bird", "horse", "cat", "dog", "rabbit", "grape", "lemon", "kiwi", "orange", "avocado", "pear", "watermelon", "apple", "banana", "cherry", "dog", "frog", "grape", "apple", "lemon", "mango", "orange", "pear", "cat", "dog", "rabbit", "strawberry", "watermelon", "bird", "horse", "kiwi", "lemon", "mango", "apple", "pear", "cat", "dog", "rabbit", "grape", "lemon", "kiwi", "orange", "avocado", "pear", "watermelon"}

	m := make(map[string]int)
	for _, v := range xs {
		m[v] = m[v] + 1
	}
	for key, val := range m {
		fmt.Printf("key : %v. values :%v \n", key, val)
	}
}
