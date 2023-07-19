package hashtable

import "github.com/jkittell/array"

// HashTable stores key/value pairs where the number of hashtable buckets is n number of buckets.
//
// A smaller number of buckets means less buckets will be created,
// but each key stored in the HashTable has a higher likelihood of having to share a bucket with other keys,
// thus slowing down lookups.
//
// A larger number of buckets means more buckets will be created, so each key stored in the Table has a lower
// likelihood of having to share a bucket with other keys, thus speeding up lookups.
type HashTable[K comparable, V any] struct {
	// hash is a function which can hash a key of type K and return the bucket containing the key/value.
	hash            func(K, int) int
	numberOfBuckets int
	table           [][]kv[K, V]
}

// A kv stores generic key/value data in a HashTable.
type kv[K comparable, V any] struct {
	Key   K
	Value V
}

type hash[K comparable] func(K, int) int

// New creates a table with n number of internal buckets which uses the specified hash
// function for an input type K.
func New[K comparable, V any](numberOfBuckets int, hash hash[K]) *HashTable[K, V] {
	return &HashTable[K, V]{
		hash:            hash,
		numberOfBuckets: numberOfBuckets,
		table:           make([][]kv[K, V], numberOfBuckets),
	}
}

// Insert a new key/value pair.
func (ht *HashTable[K, V]) Insert(key K, value V) {
	bucket := ht.hash(key, ht.numberOfBuckets)

	for n, data := range ht.table[bucket] {
		if key == data.Key {
			// overwrite previous value for the same key
			ht.table[bucket][n].Value = value
			return
		}
	}

	// add a new value to the table
	ht.table[bucket] = append(ht.table[bucket], kv[K, V]{
		Key:   key,
		Value: value,
	})
}

func (ht *HashTable[K, V]) Delete(key K) {
	var value V
	bucket := ht.hash(key, ht.numberOfBuckets)
	for n, data := range ht.table[bucket] {
		if key == data.Key {
			// overwrite previous value for the same key
			ht.table[bucket][n].Value = value
			return
		}
	}
}

func (ht *HashTable[K, V]) Search(key K) (V, bool) {
	bucket := ht.hash(key, ht.numberOfBuckets)

	for n, data := range ht.table[bucket] {
		if key == data.Key {
			// match found
			return ht.table[bucket][n].Value, true
		}
	}

	// no match
	var value V
	return value, false
}

func (ht *HashTable[K, V]) Keys() array.Array[K] {
	var keys array.Array[K]
	// Loop through the hash table buckets
	for _, bucket := range ht.table {
		// If the bucket has any keys
		if len(bucket) > 0 {
			for i := 0; i < len(bucket); i++ {
				// put the key in the keys array
				keys.Push(bucket[i].Key)
			}
		}
	}
	return keys
}
