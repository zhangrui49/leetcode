package g2225

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=2225 lang=golang
 *
 * [2225] 找出输掉零场或一场比赛的玩家
 */

// @lc code=start
func findWinners(matches [][]int) [][]int {
	loseCounts := make(map[int]int)
	undefeated := []int{}
	beatedOnce := []int{}

	for _, match := range matches {
		winner := match[0]
		loser := match[1]
		loseCounts[winner] += 0
		loseCounts[loser] += 1
	}

	for player, loseCount := range loseCounts {

		if loseCount == 0 {
			undefeated = append(undefeated, player)
		} else if loseCount == 1 {
			beatedOnce = append(beatedOnce, player)
		}
	}
	sort.Slice(undefeated, func(i, j int) bool {
		return undefeated[i] < undefeated[j]
	})
	sort.Slice(beatedOnce, func(i, j int) bool {
		return beatedOnce[i] < beatedOnce[j]
	})
	return [][]int{undefeated, beatedOnce}
}

// @lc code=end
