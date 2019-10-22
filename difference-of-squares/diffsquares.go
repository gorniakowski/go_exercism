package diffsquares

//SumOfSquares returns sum of squares of n number ie 1^2 + 2^2 +...+n^2
func SumOfSquares(n int) int {
	return (n * (n + 1) *
		(2*n + 1)) / 6
}

//SquareOfSum returns square of the sum of n numbers (1+2...n)^2
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

//Difference return difference between SquareOfsum and SumOfSquares for n numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
