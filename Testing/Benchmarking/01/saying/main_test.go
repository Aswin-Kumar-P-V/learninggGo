package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	greeting := Greet("Aswin")

	if greeting != "Hello Aswin nice to meet you" {
		t.Error("Expected:", "Hello Aswin nice to meet you", "got:", greeting)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Aswin"))
	//Output:
	//Hello Aswin nice to meet you
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Aswin")
	}
}
