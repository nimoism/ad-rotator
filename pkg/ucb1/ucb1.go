package ucb1

import (
	"math"
)

type Arm struct {
	count    int
	winCount int
}

type Arms []Arm

func NewArm(count, winCount int) Arm {
	return Arm{
		count:    count,
		winCount: winCount,
	}
}

func NextArm(arms Arms) (index int, weight float64) {
	totalCount := 0
	for _, option := range arms {
		totalCount += option.count
	}
	for i, option := range arms {
		if option.count == 0 {
			return i, 0
		}
		if w := armWeight(totalCount, option); w > weight {
			index, weight = i, w
		}
	}
	return
}

func armWeight(totalCount int, arm Arm) float64 {
	avg := float64(arm.winCount) / float64(arm.count)
	return avg + math.Sqrt(2*math.Log(float64(totalCount))/float64(arm.count))
}
