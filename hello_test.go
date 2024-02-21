package main

import "testing"

func TestHello(t *testing.T) {
	if ans := Add(1, 2); ans != 4 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}
}
