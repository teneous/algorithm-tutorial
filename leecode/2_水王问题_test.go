// 水王问题，一个数在一个集合内大于一半，找到这个水王数。请用时间复杂度为o(n),空间复杂度为o(1)找到这个树
package leecode

import (
	"fmt"
	"testing"
)

// 解题思路，如果存在一个水王数，那么它在集合内的数量必定超过一半。在遍历集合的时候，从头元素开始，初始hp为1，
// - 如果当前元素和下一个元素不等，那么当前遍历节点元素hp-1
// - 如果当前元素和下一个元素相等，那么当前遍历节点元素hp+1
// - 如果当前hp等于0，则下一个元素的成为新的候选者，赋值hp=1
func TestRun(t *testing.T) {
	// 测试数据
	source := []int{1, 2, 3, 3, 3, 4, 2, 3, 2, 3, 3}
	// 找到水王货
	catchFakeOne(source)
}

func catchFakeOne(source []int) {
	hp := 0
	candidate := 0

	for _, val := range source {
		if hp == 0 {
			candidate = val
			hp = 1
		} else {
			if candidate == val {
				hp++
			} else {
				hp--
			}
		}
	}

	if hp == 0 {
		fmt.Printf("不存在水王数\n")
	} else {
		count := 0
		for _, val := range source {
			if val == candidate {
				count++
			}
		}
		if count > (len(source) >> 1) {
			fmt.Printf("存在水王数，它是：%d，个数：%d个\n", candidate, count)
		} else {
			fmt.Printf("不存在水王数\n")
		}
	}

}
