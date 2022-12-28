package jump

import (
	"math"
)

func jump(nums []int) int {
	// Initialize the minDistances to maxInt.
	minDistances := make([]int, len(nums))

	// Start at 1 because distance to 0 should be 0.
	for i := 1; i < len(nums); i++ {
		minDistances[i] = math.MaxInt64
	}

	// Only iterate to the penultimate element, because we can't move forward from the last.
	for i := 0; i < len(nums)-1; i++ {
		steps := nums[i]
		// New index is i + j for all j <= steps.
		for j := 1; j <= steps; j++ {
			newInd := i + j
			if minDistances[i]+1 < minDistances[newInd] {
				minDistances[newInd] = minDistances[i] + 1
			}
		}
	}
	return minDistances[len(nums)-1]
}
