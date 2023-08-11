package consistent

import (
	"fmt"
	"strconv"
	"testing"
)

func runCase(t *testing.T, hash *Map, testCases map[string]string) {
	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}
}

func TestHashing(t *testing.T) {
	hash := New(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	// Given the above hash function, this will give replicas with "hashes":
	// 2, 4, 6, 12, 14, 16, 22, 24, 26
	hash.Add("6", "4", "2")

	testCases := map[string]string{
		"1":  "2",
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}
	runCase(t, hash, testCases)

	// Adds 8, 18, 28
	hash.Add("8")
	// 27 should now map to 8.
	testCases["27"] = "8"
	runCase(t, hash, testCases)

	hash.Remove("8")
	testCases["27"] = "2"
	runCase(t, hash, testCases)

	hash.Remove("5")
	fmt.Println(hash.hashMap)
	fmt.Println(hash.noteRing)
}
