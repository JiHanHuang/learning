package gee

import "testing"

func TestTrie(t *testing.T) {
	head := NewNode()
	head.Add("/")
	head.Add("/a")
	head.Add("/a/b/")
	head.Add("/a/b/c/d")
	head.Add("/b/:name")
	//head.Print()

	uri := "/b/hello"
	n, k := head.Get(uri)
	if n == nil {
		t.Fatal("not find:", uri)
	}
	t.Log("find:", uri, k, n)

	uri = "/b/aa/b/c"
	n, k = head.Get(uri)
	if n == nil {
		t.Fatal("not find:", uri)
	}
	t.Log("find:", uri, k)

	uri = "/a/b/c/"
	n, k = head.Get(uri)
	if n != nil {
		t.Fatal("Fatal match:", uri)
	}
	t.Log("not find:", uri, k)
}
