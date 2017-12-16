package calculator

import "fmt"
import "math"

func cal(x int, y int, operator string) (result int){

	switch operator{
	case "mul": result = multiply(x, y)
	case "sub": result = subtract(x, y)
	case "sum": result = addition(x, y)
	case "div": result = divide(x, y)
	case "sqrt": result = sqrt(x, y)
	}

	return
}

func xcal(x int, y int) (add int, mul int, sub int, div int){

	mul = multiply(x, y)
	sub = subtract(x, y)
	add = addition(x, y)
	div = divide(x, y)
	return
}

func subtract(x int, y int) (result int){
result = x - y
return result
}

func multiply(x int, y int) (result int){
result = x * y
return result
}

func addition(x int, y int) (result int){
result = x + y
return result
}

func divide(x int, y int) (result int){
result = x / y
return result
}

func sqrt(x, y int)(result int){
	var f = math.Sqrt(float64(x*x + y*y))
	result = int(f)
	return result
}

func main() {
	resultSum := cal(10, 2, "sum")
	fmt.Println("Computed result for Addition: ", resultSum)

	resultMultpy := cal(10, 2, "mul")
	fmt.Println("Computed result for Multiplication: ", resultMultpy)

	resultSubtract := cal(10, 2, "sub")
	fmt.Println("Computed result for Subtraction: ", resultSubtract)

	resultDivide := cal(10, 2, "div")
	fmt.Println("Computed result Division: ", resultDivide)

	resultSqrt := cal(10, 2, "sqrt")
	fmt.Println("Computed result Sqrt: ", resultSqrt)

	a, b, c, d := xcal(10, 2)
	fmt.Printf("Computed results: \n\tAddition: %d \n\tSubtraction: %d \n\tMultipication: %d \n\tDivision: %d", a, b, c, d)
}
