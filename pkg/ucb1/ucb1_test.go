package ucb1

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestArmsStart(t *testing.T) {
	tests := []struct {
		name      string
		arms      []Arm
		wantIndex int
	}{
		{"all empties", []Arm{
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
		}, 0},
		{"second empty", []Arm{
			{1, 1},
			{0, 0},
			{0, 0},
			{0, 0},
		}, 1},
		{"last empty", []Arm{
			{1, 1},
			{1, 1},
			{1, 1},
			{0, 0},
		}, 3},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, _ := NextArm(tt.arms)
			require.Equal(t, tt.wantIndex, gotIndex)
		})
	}
}

func TestArmsWinners(t *testing.T) {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	arms := Arms{{count: 1}, {count: 1}, {count: 1}}
	lucks := []float32{0.0, 0.8, 0.4}
	tops := []int{1, 2}
	shows := make(map[int]int)
	for i := 0; i < 10000; i++ {
		armIndex, _ := NextArm(arms)
		shows[armIndex]++
		arms[armIndex].count++
		if rand.Float32() < lucks[armIndex] { //nolint:gosec
			arms[armIndex].winCount++
		}
	}
	totalCount := 0
	for _, count := range shows {
		totalCount += count
	}
	for i := range tops[:len(tops)-1] {
		require.Less(t, shows[tops[i+1]], shows[tops[i]], "Wrong top: %v > %v, seed: %v", tops[i], tops[i+1], seed)
	}
}

func TestAtLeastOneArmsCounted(t *testing.T) {
	arms := Arms{{count: 1}, {count: 1}, {count: 1}}
	lucks := []float32{0.0, 0.8, 0.4, 0.0}
	shows := make(map[int]int)
	for i := 0; i < 10000; i++ {
		armIndex, _ := NextArm(arms)
		shows[armIndex]++
		arms[armIndex].count++
		if rand.Float32() < lucks[armIndex] { //nolint:gosec
			arms[armIndex].winCount++
		}
	}
	totalCount := 0
	for _, count := range shows {
		totalCount += count
	}
	for i := range arms {
		require.Greater(t, arms[i].count, 0)
	}
}
