package bst

import (
	"github.com/FerBuono/go-data-structures/dynamic-stack"
)

type nodoBST[K comparable, V any] struct {
	left   *nodoBST[K, V]
	right  *nodoBST[K, V]
	key    K
	value  V
}

type bst[K comparable, V any] struct {
	root *nodoBST[K, V]
	size int
	cmp  func(K, K) int
}

type iterBST[K comparable, V any] struct {
	bst   *bst[K, V]
	stack dynamic_stack.Stack[*nodoBST[K, V]]
	from  *K
	to    *K
}

func NewBST[K comparable, V any](cmp func(K, K) int) OrderedDictionary[K, V] {
	t := new(bst[K, V])
	t.cmp = cmp
	return t
}

// Dictionary methods

func (t *bst[K, V]) Save(key K, value V) {
	node := t.findNode(key, &t.root)
	if *node == nil {
		*node = &nodoBST[K, V]{key: key, value: value}
		t.size++
	} else {
		(*node).value = value
	}
}

func (t *bst[K, V]) Contains(key K) bool {
	return *(t.findNode(key, &t.root)) != nil
}

func (t *bst[K, V]) Get(key K) V {
	node := t.findNode(key, &t.root)
	if *node == nil {
		panic("The key does not belong to the dictionary")
	}
	return (*node).value
}

func (t *bst[K, V]) Delete(key K) V {
	node := t.findNode(key, &t.root)
	return t.deleteNode(node)
}

func (t *bst[K, V]) Size() int {
	return t.size
}

func (t *bst[K, V]) Iterate(f func(K, V) bool) {
	t.iterateInRange(t.root, f, nil, nil)
}

func (t *bst[K, V]) Iterator() DictionaryIterator[K, V] {
	return t.RangeIterator(nil, nil)
}

// OrderedDictionary methods

func (t *bst[K, V]) IterateRange(from *K, to *K, visit func(key K, value V) bool) {
	t.iterateInRange(t.root, visit, from, to)
}

func (t *bst[K, V]) RangeIterator(from *K, to *K) DictionaryIterator[K, V] {
	iter := new(iterBST[K, V])
	iter.bst = t
	iter.stack = dynamic_stack.NewDynamicStack[*nodoBST[K, V]]()
	iter.from = from
	iter.to = to
	first := iter.findFirst(t.root)
	if first != nil {
		iter.stack.Push(first)
		iter.pushLeftChildren(first)
	}
	return iter
}

// DictionaryIterator methods

func (iter *iterBST[K, V]) HasNext() bool {
	return !iter.stack.IsEmpty()
}

func (iter *iterBST[K, V]) Current() (K, V) {
	if !iter.HasNext() {
		panic("The iterator has finished iterating")
	}
	return iter.stack.Top().key, iter.stack.Top().value
}

func (iter *iterBST[K, V]) Next() K {
	if !iter.HasNext() {
		panic("The iterator has finished iterating")
	}
	node := iter.stack.Pop()
	iter.pushLeftChildren(node.right)
	return node.key
}

// Helper methods

func (t *bst[K, V]) findNode(key K, node **nodoBST[K, V]) **nodoBST[K, V] {
	if *node == nil {
		return node
	}
	if t.cmp(key, (*node).key) < 0 {
		if (*node).left == nil || t.cmp(key, (*node).left.key) == 0 {
			return &(*node).left
		}
		return t.findNode(key, &(*node).left)
	} else if t.cmp(key, (*node).key) > 0 {
		if (*node).right == nil || t.cmp(key, (*node).right.key) == 0 {
			return &(*node).right
		}
		return t.findNode(key, &(*node).right)
	} else {
		return node
	}
}

func (t *bst[K, V]) deleteNode(node **nodoBST[K, V]) V {
	if *node == nil {
		panic("The key does not belong to the dictionary")
	}
	value := (*node).value
	if t.countChildren(node) == 0 {
		*node = nil
	} else if t.countChildren(node) == 1 {
		child := t.getChild(node)
		*node = *child
	} else {
		replacement := t.findReplacement(&(*node).left)
		newKey, newValue := (*replacement).key, (*replacement).value
		*replacement = (*replacement).left
		(*node).key = newKey
		(*node).value = newValue
	}
	t.size--
	return value
}

func (t *bst[K, V]) findReplacement(node **nodoBST[K, V]) **nodoBST[K, V] {
	if (*node).right == nil {
		return node
	} else {
		return t.findReplacement(&(*node).right)
	}
}

func (t *bst[K, V]) countChildren(node **nodoBST[K, V]) int {
	if (*node).left != nil && (*node).right != nil {
		return 2
	} else if (*node).left == nil && (*node).right == nil {
		return 0
	} else {
		return 1
	}
}

func (t *bst[K, V]) getChild(node **nodoBST[K, V]) **nodoBST[K, V] {
	if (*node).left != nil {
		return &(*node).left
	} else {
		return &(*node).right
	}
}

func (iter *iterBST[K, V]) pushLeftChildren(node *nodoBST[K, V]) {
	if node == nil {
		return
	}
	if (iter.from == nil || iter.bst.cmp(*iter.from, node.key) <= 0) && (iter.to == nil || iter.bst.cmp(*iter.to, node.key) >= 0) {
		iter.stack.Push(node)
	}
	iter.pushLeftChildren(node.left)
}

func (iter *iterBST[K, V]) findFirst(node *nodoBST[K, V]) *nodoBST[K, V] {
	if node == nil {
		return nil
	}

	if iter.from == nil && iter.to == nil {
		return node
	}

	if iter.from != nil && iter.bst.cmp(*iter.from, node.key) > 0 {
		return iter.findFirst(node.right)
	}
	if iter.to != nil && iter.bst.cmp(*iter.to, node.key) < 0 {
		return iter.findFirst(node.left)
	}
	return node
}

func (t *bst[K, V]) iterateInRange(current *nodoBST[K, V], f func(K, V) bool, from *K, to *K) bool {
	if current == nil {
		return true
	}
	proceed := true
	if from == nil || t.cmp(current.key, *from) > 0 {
		proceed = t.iterateInRange(current.left, f, from, to)
	}
	if proceed && (from == nil || t.cmp(current.key, *from) >= 0) && (to == nil || t.cmp(current.key, *to) <= 0) {
		proceed = f(current.key, current.value)
	}
	if proceed && (to == nil || t.cmp(current.key, *to) < 0) {
		return t.iterateInRange(current.right, f, from, to)
	}
	return false
}
