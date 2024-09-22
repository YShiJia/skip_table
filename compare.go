/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-09-22 20:04:40
 */

package skip_table

import "cmp"

type ordered interface {
	cmp.Ordered
}

//type comparable interface {
//	gt(c comparable) bool
//	eq(c comparable) bool
//	lt(c comparable) bool
//}
