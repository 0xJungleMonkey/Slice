package main

func average(salary []int) float64 {
	var mini = salary[0]
	var maxi = salary[0]
	var sum = 0

	for i := 0; i < len(salary); i++ {
		sum = sum + salary[i]
		if salary[i] > maxi {
			maxi = salary[i]

		}
		if salary[i] < mini {
			mini = salary[i]
		}

	}
	sum = sum - maxi - mini
	return float64(sum / (len(salary) - 2))

}
func main1() {
	var m = []int{1000, 2000, 3000, 4000}
	average(m)
}
