// 水王问题，一个数在一个集合内大于一半，请用时间复杂度为o(n),空间复杂度为o(1)找到这个书
package leecode

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	// 解题思路，如果存在一个水王数，那么它在集合内的数量必定超过一半。在遍历集合的时候，从头元素开始，默认hp为1，
	// - 如果当前元素和下一个元素不等，那么当前遍历节点元素hp-1，并跳过这两个元素
	// - 如果当前元素和下一个元素相等，那么当前遍历节点元素hp+1，并跳到下一个元素
	// 例如 1，2，3，3，4

	// 测试数据
	source := []int{1, 2, 3, 3, 3, 4, 2, 3, 2, 3, 3}

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
		fmt.Printf("不存在水王数")
	} else {
		count := 0
		for _, val := range source {
			if val == candidate {
				count++
			}
		}
		if count > (len(source) >> 1) {
			fmt.Printf("存在水王数:%d,个数%d", candidate, count)
		} else {
			fmt.Printf("不存在水王数")
		}
	}
}
