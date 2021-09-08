/*
已有方法 rand7 可生成 1 到 7 范围内的均匀随机整数，试写一个方法 rand10 生成 1 到 10 范围内的均匀随机整数。
不要使用系统的 Math.random() 方法
*/

/*
已知 rand_N() 可以等概率的生成[1, N]范围的随机数
那么：
(rand_X() - 1) × Y + rand_Y() ==> 可以等概率的生成[1, X * Y]范围的随机数
即实现了 rand_XY()
*/

package g470

//用 Rand7() 实现 Rand10()

func rand10() int {

	for {
		result := (rand7()-1)*7 + rand7()
		if result <= 40 {
			return result%10 + 1
		}
	}

}
