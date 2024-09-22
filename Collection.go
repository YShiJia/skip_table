/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-09-22 19:05:57
 */

package skip_table

type Collection[K comparable, V any] interface {
	// Get 根据key获取value，bool表示是否执行成功
	Get(key K) (KV[K, V], bool)
	// Put 插入key-value，返回是否成功
	Put(key K, value V) bool
	// Del 根据key删除元素，返回value，bool表示是否执行成功
	Del(key K) (KV[K, V], bool)
	//Range 返回[begin, end]范围内的切片
	Range(begin K, end K) []KV[K, V]
	// Ceiling 返回>=target的最小元素，bool表示是否执行成功
	Ceiling(target K) (KV[K, V], bool)
	// Floor 返回<=target的最大元素，bool表示是否执行成功
	Floor(target K) (KV[K, V], bool)
}
