package bst

type Dictionary[K comparable, V any] interface {

	// Save saves the key-value pair in the Dictionary. If the key already exists, the associated value is updated.
	Save(key K, value V)

	// Contains determines if a key is already in the dictionary.
	Contains(key K) bool

	// Get returns the value associated with a key. If the key does not belong, it should panic with the message
	// 'The key does not belong to the dictionary'.
	Get(key K) V

	// Delete removes the key from the Dictionary, returning the value that was associated with it. If the key does
	// not belong to the dictionary, it should panic with the message 'The key does not belong to the dictionary'.
	Delete(key K) V

	// Size returns the number of elements in the dictionary.
	Size() int

	// Iterate iterates internally through the dictionary, applying the function passed as a parameter to all elements
	// within it.
	Iterate(func(key K, value V) bool)

	// Iterator returns a DictionaryIterator for this Dictionary.
	Iterator() DictionaryIterator[K, V]
}

type DictionaryIterator[K comparable, V any] interface {

	// HasNext returns if there are more elements to see. That is, if there is an element where the iterator is currently positioned.
	HasNext() bool

	// Current returns the key and the value of the current element where the iterator is positioned.
	// If there is no next element, it should panic with the message 'The iterator has finished iterating'.
	Current() (K, V)

	// Next if there is a next element, returns the current key (equivalent to Current, but only the key), and
	// also advances to the next element in the dictionary. If there is no next element, it should panic with the message
	// 'The iterator has finished iterating'.
	Next() K
}

type OrderedDictionary[K comparable, V any] interface {
	Dictionary[K, V]

	// IterateRange iterates only including elements that are within the indicated range,
	// including them if they are within it.
	IterateRange(desde *K, hasta *K, visitar func(key K, value V) bool)

	// RangeIterator creates a DictionaryIterator that only iterates over keys that are within the indicated range.
	RangeIterator(desde *K, hasta *K) DictionaryIterator[K, V]
}
