package main
import "fmt"

// Complete the solveMeFirst function below.
// Found at https://www.hackerrank.com/challenges/solve-me-first/problem.
func solveMeFirst(a uint32,b uint32) uint32{
	return a + b
}

func main() {
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a,&b)
	res = solveMeFirst(a,b)
	fmt.Println(res)
}
