package consistent

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash maps bytes to uint32
type Hash func(data []byte) uint32

// Map constains all hashed keys
type Map struct {
	hash     Hash
	replicas int
	noteRing []int          // Hash ring, virtual note sorted
	hashMap  map[int]string // map of virtual note and real note
}

// New creates a Map instance
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

//添加真实的节点到Map中，生成对应的虚拟节点
func (m *Map) Add(notes ...string) {
	for _, note := range notes {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + note)))
			m.noteRing = append(m.noteRing, hash)
			m.hashMap[hash] = note
		}
	}
	sort.Ints(m.noteRing)
}

//删除真实的节点和对应的虚拟节点
func (m *Map) Remove(notes ...string) {
	for _, note := range notes {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + note)))
			_, ok := m.hashMap[hash]
			if !ok {
				return
			}
			idx := sort.SearchInts(m.noteRing, hash)
			idx = idx % len(m.noteRing)
			m.noteRing = append(m.noteRing[:idx], m.noteRing[idx+1:]...)
			delete(m.hashMap, hash)
		}
	}
}

//选取key对应映射的真实节点名称
//key即我们想使用一致性hash的key
func (m *Map) Get(key string) string {
	if len(m.noteRing) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	// Binary search for appropriate replica.
	//sort.Search本质上使用了二分查找法
	idx := sort.Search(len(m.noteRing), func(i int) bool {
		return m.noteRing[i] >= hash
	})

	return m.hashMap[m.noteRing[idx%len(m.noteRing)]]
}
