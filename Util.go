/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-09-22 19:20:12
 */

package skip_table

import "math/rand"

func roll() int {
	level := 1
	for rand.Int() < 0 {
		level++
	}
	return level
}
