/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-09-22 19:25:10
 */

package skip_table

type KV[K any, V any] interface {
	Key() K
	Value() V
	Set(key K, value V) bool
}

func NewKV[K any, V any](key K, value V) *kv[K, V] {
	return &kv[K, V]{
		key:   key,
		value: value,
	}
}

type kv[K any, V any] struct {
	key   K
	value V
}

func (k *kv[K, V]) Key() K {
	return k.key
}

func (k *kv[K, V]) Value() V {
	return k.value
}

func (k *kv[K, V]) Set(key K, value V) bool {
	k.key = key
	k.value = value
	return true
}

func NewCompareKV[K ordered, V any](key K, value V) *compareKV[K, V] {
	return &compareKV[K, V]{
		kv: kv[K, V]{
			key:   key,
			value: value,
		},
	}
}

type compareKV[K ordered, V any] struct {
	kv[K, V]
}
