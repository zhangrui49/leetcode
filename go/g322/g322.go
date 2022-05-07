package g322

import (
	"math"
	"sort"
)

/**

给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。


*/

//只尝试一遍，不对，需要递归
func coinChangeV2(coins []int, amount int) int {
	sort.Ints(coins)
	r := 0
	for i := len(coins) - 1; i >= 0; i-- {
		coin := coins[i]
		if coin > amount {
			continue
		}
		if coin <= amount {

			mod := amount % coin
			div := amount / coin
			r += div
			amount = mod
		}
		if amount == 0 {
			break
		}
	}
	if amount == 0 {
		return r
	}
	return -1

}

func coinChangeV3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32 //用amount+1
	}

	for i := 0; i <= amount; i++ {

		for _, v := range coins {
			if i-v < 0 {
				continue
			}
			dp[i] = getMin(dp[i], 1+dp[i-v]) //maxInt + 1 会越界变成负数
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	} else {
		return dp[amount]
	}
}

var cache map[int]int

func coinChange(coins []int, amount int) int {
	cache = map[int]int{}
	return canChange(coins, amount)
}

func canChange(coins []int, amount int) int {
	if v, ok := cache[amount]; ok {
		return v
	}

	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	r := math.MaxInt64

	for _, v := range coins {
		sub := canChange(coins, amount-v)
		if sub == -1 {
			continue
		}
		r = getMin(r, sub+1)
	}
	if r == math.MaxInt64 {
		r = -1
	}
	cache[amount] = r
	return r

}

func getMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}
