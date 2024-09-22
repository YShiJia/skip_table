/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-09-22 19:17:08
 */

package skip_table

type node[K ordered, V any] struct {
	next []*node[K, V]
	*compareKV[K, V]
}

type SkipTable[K ordered, V any] struct {
	// head的next数组高度与跳表中第一个元素的高度相同
	head *node[K, V]
	size int64
}

func NewSkipTable[K ordered, V any]() *SkipTable[K, V] {
	//需要将head作为哨兵
	return &SkipTable[K, V]{
		head: &node[K, V]{
			next: make([]*node[K, V], 0),
		},
		size: int64(0),
	}
}

func (s *SkipTable[K, V]) Get(key K) (KV[K, V], bool) {
	if n := s.search(key); n != nil {
		return n, true
	}
	return nil, false
}

func (s *SkipTable[K, V]) search(key K) KV[K, V] {
	//跳表中不存在元素
	cur := s.head
	if len(cur.next) == 0 {
		return nil
	}
	// 当前cur为head节点，并不是实际的跳表元素，
	// 从cur开始遍历，将cur作为逻辑上的跳表最小元素
	// 从最高层开始遍历
	for level := len(cur.next) - 1; level >= 0; level-- {
		for cur.next[level] != nil && cur.next[level].key < key {
			cur = cur.next[level]
		}
		if cur.next[level] != nil && cur.next[level].key == key {
			return cur.next[level]
		}
	}
	return nil
}

func (s *SkipTable[K, V]) Put(key K, value V) bool {
	if n := s.search(key); n != nil {
		n.Set(key, value)
	}

	level := roll()

	//保持头结点处于最高层
	for level > len(s.head.next) {
		s.head.next = append(s.head.next, nil)
	}

	newNode := &node[K, V]{
		next:      make([]*node[K, V], level),
		compareKV: NewCompareKV[K, V](key, value),
	}

	cur := s.head
	for level = len(cur.next) - 1; level >= 0; level-- {
		//从最高层开始，找到比key小的元素，将next指向新节点
		for cur.next[level] != nil && cur.next[level].key < key {
			cur = cur.next[level]
		}

		newNode.next[level] = cur.next[level]
		cur.next[level] = newNode
	}

	return true
}

func (s *SkipTable[K, V]) Del(key K) (KV[K, V], bool) {
	cur := s.head
	var ret KV[K, V]
	for level := len(cur.next) - 1; level >= 0; level-- {
		for cur.next[level] != nil && cur.next[level].key < key {
			cur = cur.next[level]
		}
		if cur.next[level] != nil && cur.next[level].key == key {
			//删除节点
			ret = cur.next[level]
			cur.next[level] = cur.next[level].next[level]
		}
	}
	// TODO 1.如果最高节点只有head，需要降低head高度
	// TODO 2.需要添加一个检测：判断跳表中间某一层是否空了，是则需要将该层移除，让其上层下降
	return ret, true
}

func (s *SkipTable[K, V]) Range(begin K, end K) []KV[K, V] {
	var ret []KV[K, V]
	var cur *node[K, V]
	//找到第一个大于等于begin的元素
	if cur = s.ceiling(begin); cur == nil {
		return nil
	}
	for cur != nil && cur.key <= end {
		ret = append(ret, cur)
		cur = cur.next[0]
	}

	return ret
}

// Ceiling 找到 key 值大于等于 target 且 key 值最接近于 target 的节点
func (s *SkipTable[K, V]) Ceiling(target K) (KV[K, V], bool) {
	if n := s.ceiling(target); n != nil {
		return n, true
	}
	return nil, false
}

func (s *SkipTable[K, V]) ceiling(target K) *node[K, V] {
	cur := s.head
	for level := len(cur.next) - 1; level >= 0; level-- {
		for cur.next[level] != nil && cur.next[level].key < target {
			cur = cur.next[level]
		}
		if cur.next[level] != nil && cur.next[level].key >= target {
			return cur.next[level]
		}
	}
	return nil
}

// Floor 找到 SkipTable 中，key 值大于等于 target 且最接近于 target 的节点
func (s *SkipTable[K, V]) Floor(target K) (KV[K, V], bool) {
	if n := s.floor(target); n != nil {
		return n, true
	}
	return nil, false
}

func (s *SkipTable[K, V]) floor(target K) KV[K, V] {
	cur := s.head
	//获取
	for level := len(s.head.next[0].next) - 1; level >= 0; level-- {
		for cur.next[level] != nil && cur.next[level].key < target {
			cur = cur.next[level]
		}
		if cur.next[level] != nil && cur.next[level].key >= target {
			return cur.next[level]
		}
	}
	return nil
}
