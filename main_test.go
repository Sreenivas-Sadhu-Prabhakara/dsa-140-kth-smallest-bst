package main

import "testing"

func TestKthSmallest(t *testing.T) {
	root := tn(5, tn(3, tn(2, tn(1, nil, nil), nil), tn(4, nil, nil)), tn(6, nil, nil))
	wants := []int{1, 2, 3, 4, 5, 6}
	for k := 1; k <= 6; k++ {
		got, ok := KthSmallest(root, k)
		if !ok || got != wants[k-1] {
			t.Errorf("KthSmallest(%d) = (%d,%v), want %d", k, got, ok, wants[k-1])
		}
	}
	if _, ok := KthSmallest(root, 7); ok {
		t.Error("k beyond size should be not-ok")
	}
	if _, ok := KthSmallest(root, 0); ok {
		t.Error("k=0 should be not-ok")
	}
	if _, ok := KthSmallest(nil, 1); ok {
		t.Error("empty tree should be not-ok")
	}
}
