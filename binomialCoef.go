package main
import "fmt"
func binomCoeff(a int, b int) int {
	if (b == 0 || b == a) {
		return 1
	}
	return binomCoeff(a-1, b-1) + binomCoeff(a-1, b)
}
func main() {
	fmt.Print("Value of nCr is: " , binomCoeff(5, 2))
}