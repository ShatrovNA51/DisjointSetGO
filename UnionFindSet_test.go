package unionfindset

import (
	"fmt"
	"testing"
)

func TestUnionFind_Connected(t *testing.T) {
	uf, _ := New(10)
	unionItems(uf, []testUnionData{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
	})

	assertConnected(t, uf, testConnectedResult{0, 7, false})
	assertConnected(t, uf, testConnectedResult{8, 9, true})

	unionItems(uf, []testUnionData{
		{5, 0},
		{7, 2},
		{6, 1},
		{1, 0},
		{0, 7},
	})
	assertConnected(t, uf, testConnectedResult{0, 7, true})
}

func TestUnionFind_Root(t *testing.T) {
	uf, _ := New(5)

	unionItems(uf, []testUnionData{
		{4, 3}, // [0 1 2 4 4]
		{1, 2}, // [0 1 1 4 4]
	})
	assertRoot(t, uf, testRootResult{3, 4})
	assertRoot(t, uf, testRootResult{2, 1})
}

func unionItems(uf *UnionFindSet, items []testUnionData) {
	for _, item := range items {
		err := uf.Union(item.p, item.q)
		if err != nil {
			panic(err)
		}
	}
}

func assetTrue(t *testing.T, condition bool, message string) {
	if !condition {
		t.Error(message)
	}
}

type testUnionData struct {
	p, q int
}

type testConnectedResult struct {
	p, q     int
	expected bool
}

type testRootResult struct {
	i        int
	expected int
}

func assertConnected(t *testing.T, uf *UnionFindSet, r testConnectedResult) {
	connected, err := uf.Connected(r.p, r.q)
	if err != nil {
		panic(err)
	}
	var condition bool
	if r.expected {
		condition = connected
	} else {
		condition = !connected
	}
	assetTrue(t, condition, fmt.Sprintf("uf.Connected(%d, %d) == %t, want %t", r.p, r.q, !r.expected, r.expected))
}

func assertRoot(t *testing.T, uf *UnionFindSet, r testRootResult) {
	root, err := uf.Find(r.i)
	if err != nil {
		panic(err)
	}
	assetTrue(t, r.expected == root, fmt.Sprintf("uf.Root(%d) == %d, want %d", r.i, root, r.expected))
}
