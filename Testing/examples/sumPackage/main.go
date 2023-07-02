// Package SumPackage is for people who is too lazy to do sum themselves
package sumPackage

//Sum func adds an unlimited number of values of type int and returns the sum
func Sum(num ...int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}
