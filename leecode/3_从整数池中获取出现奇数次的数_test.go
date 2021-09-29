// 从整数池中获取出现奇数次的数
package leecode

import (
	"fmt"
	"testing"
)

// 例如「1,3,4,4,2,3,6」 计算最终得到1和6
func TestGetDiffEvenNumberFromPool(t *testing.T) {
	source := []int{1, 2, 3, 3, 3, 2, 4, 5, 1, 4}
	getEventNum(source)
}

// 获取水王数
func getEventNum(source []int) {

	result := 0
	//a^a=0 a^b^b=a a^a^b=a
	for _, val := range source {
		result ^= val
	}
	//既然存在不同的数，那么a^b出来的数据在二进制位上必存在互斥值，即至少一个二进制位上，一个数是0，一个数是1，那么我们需要找到这个1在什么位置
	bitPosition := result & (^result + 1)

	a := 0
	for _, val := range source {
		if (val>>(bitPosition))&1 == 1 {
			a ^= val
		}
	}

	fmt.Printf("even a is :%d,even b is :%d\n", a, result^a)
}
