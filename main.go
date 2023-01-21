package main
import "fmt"
func subtractProductAndSum(n int) int {
	product := 1
	sum := 0
	for i := 0; i < 5; i++ {

		num := n % (10 ^ (i))
		n -= num * (10 ^ i)
		product *= (num / (10 ^ i))
		sum += num / (10 ^ i)
		if n/(10^(i+1)) == 0 {
			break
		}
	}
	return (product - sum)
}
func main() {
	var m = 234
	fmt.Println(subtractProductAndSum(m))

}
