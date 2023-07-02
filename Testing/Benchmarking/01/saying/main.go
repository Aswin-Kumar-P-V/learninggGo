package saying

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello %v nice to meet you", name)
}
