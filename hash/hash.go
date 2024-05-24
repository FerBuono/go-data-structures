package hash

import (
    "fmt"
    "hash/fnv"
)

const (
    _EMPTY = iota
    _OCCUPIED
    _DELETED

    _INITIAL_CAPACITY = 32
    _RESIZE_FACTOR    = 2
    _MAX_LOAD_FACTOR  = 75
    _MIN_LOAD_FACTOR  = 20
)

type element[K comparable, V any] struct {
    state int
    key   K
    value V
}

type closedHash[K comparable, V any] struct {
    elements  []element[K, V]
    capacity  int
    count     int
    deleted   int
}

type closedHashIterator[K comparable, V any] struct {
    dict      *closedHash[K, V]
    curIndex  int
}

func NewHash[K comparable, V any]() Dictionary[K, V] {
    dict := new(closedHash[K, V])
    dict.capacity = _INITIAL_CAPACITY
    dict.elements = make([]element[K, V], _INITIAL_CAPACITY)
    return dict
}

// Dictionary methods

func (dict *closedHash[K, V]) Save(key K, value V) {
    load := ((dict.count + dict.deleted) * 100) / dict.capacity
    if load > _MAX_LOAD_FACTOR {
        dict.resize(dict.capacity * _RESIZE_FACTOR)
    }
    pos := dict.calculatePos(key)
    if dict.elements[pos].state == _EMPTY {
        dict.count++
    }
    dict.elements[pos].state = _OCCUPIED
    dict.elements[pos].key = key
    dict.elements[pos].value = value
}

func (dict *closedHash[K, V]) Contains(key K) bool {
    pos := dict.calculatePos(key)
    return dict.elements[pos].state == _OCCUPIED
}

func (dict *closedHash[K, V]) Get(key K) V {
    pos := dict.calculatePos(key)
    if dict.elements[pos].state == _OCCUPIED {
        return dict.elements[pos].value
    } else {
        panic("Key does not exist in the dictionary")
    }
}

func (dict *closedHash[K, V]) Delete(key K) V {
    load := (dict.count * 100) / dict.capacity
    if load < _MIN_LOAD_FACTOR && dict.capacity > _INITIAL_CAPACITY {
        dict.resize(dict.capacity / _RESIZE_FACTOR)
    }
    pos := dict.calculatePos(key)
    fmt.Println(pos)
    if dict.elements[pos].state == _OCCUPIED {
        dict.elements[pos].state = _DELETED
        dict.count--
        dict.deleted++
    } else {
        panic("Key does not exist in the dictionary")
    }
    return dict.elements[pos].value
}

func (dict *closedHash[K, V]) Size() int {
    return dict.count
}

func (dict *closedHash[K, V]) Iterate(visitor func(key K, value V) bool) {
    for i := 0; i < dict.capacity; i++ {
        elem := dict.elements[i]
        if elem.state == _OCCUPIED && !visitor(elem.key, elem.value) {
            return
        }
    }
}

func (dict *closedHash[K, V]) Iterator() DictionaryIterator[K, V] {
    iterator := new(closedHashIterator[K, V])
    iterator.dict = dict
    iterator.curIndex = 0
    if dict.elements[0].state != _OCCUPIED {
        iterator.curIndex = iterator.findNext()
    }
    return iterator
}

// DictionaryIterator methods

func (iter *closedHashIterator[K, V]) HasNext() bool {
    return iter.curIndex != iter.dict.capacity
}

func (iter *closedHashIterator[K, V]) Current() (K, V) {
    if !iter.HasNext() {
        panic("Iterator has finished iterating")
    }
    return iter.dict.elements[iter.curIndex].key, iter.dict.elements[iter.curIndex].value
}

func (iter *closedHashIterator[K, V]) Next() K {
    if iter.curIndex == iter.dict.capacity {
        panic("Iterator has finished iterating")
    }
    currentKey := iter.dict.elements[iter.curIndex].key
    iter.curIndex = iter.findNext()
    return currentKey
}

// Auxiliary functions / methods

func convertToBytes[K comparable](key K) []byte {
    return []byte(fmt.Sprintf("%v", key))
}

func hash(key []byte) uint64 {
    h := fnv.New64a()
    h.Write(key)
    return h.Sum64()
}

func (dict *closedHash[K, V]) calculatePos(key K) uint64 {
    pos := hash(convertToBytes(key)) % uint64(dict.capacity)
    for dict.elements[pos].state != _EMPTY {
        if dict.elements[pos].key == key && dict.elements[pos].state == _OCCUPIED {
            return pos
        }
        if pos+1 == uint64(dict.capacity) {
            pos = 0
        } else {
            pos++
        }
    }
    return pos
}

func (dict *closedHash[K, V]) resize(newCapacity int) {
    oldCapacity := dict.capacity
    oldElements := dict.elements
    dict.elements = make([]element[K, V], newCapacity)
    dict.deleted = 0
    dict.capacity = newCapacity
    for i := 0; i < oldCapacity; i++ {
        if oldElements[i].state != _OCCUPIED {
            continue
        }
        pos := dict.calculatePos(oldElements[i].key)
        dict.elements[pos] = oldElements[i]
    }
}

func (iter *closedHashIterator[K, V]) findNext() int {
    for i := iter.curIndex + 1; i < iter.dict.capacity; i++ {
        if iter.dict.elements[i].state == _OCCUPIED {
            return i
        }
    }
    return iter.dict.capacity
}
