package main

func maximumWealth(accounts [][]int) int {

	var max int
	for _, a := range accounts {
		total := sum(a)
		if total > max {
			max = total
		}
	}
	return max
}

func sum(act []int) (total int) {
	for _, n := range act {
		total += n
	}
	return total
}

func main() {

}
