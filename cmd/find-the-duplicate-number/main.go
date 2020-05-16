package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}

// 设直线部分长度为x, 环形部分长度为y, 指针在环形中按顺时针旋转。当慢指针走到交点的时候，快指针距离交点距离为z（再往前走z的距离就能到达交点）。
// 那么再经过z的距离，两个指针相遇，因为快慢指针的速度差为1，此时慢指针走过的距离为x+z，快指针走多的距离为x+ny+y-z+2z=x+(n+1)y+z。
// 因为快指针速度是慢指针速度的2倍，所以有等式：x+(n+1)y+z = 2(x+z) => (n+1)y = x+z (1)。此时将快指针重置到初始位置，并将速度设置为1,
// 经过x的距离，原先的慢指针也走了x的距离，相对于交点，其走过的距离则是x+z，由前面的等式(1)可知，慢指针也恰恰回到了交点。
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
