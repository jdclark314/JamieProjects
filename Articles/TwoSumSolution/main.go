package main

func twoSumBrute(nums []int, target int) []int {

	for i, firstNum := range nums {
		for j, secNum := range nums {
			if j <= i {
				//if j=i we are at the same element, which is against the question criteria
				//if j < i we have already looked at those values, so we don't need to again?
				// we should analyze how much a difference this makes
				continue
			}
			if firstNum+secNum == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {

}
