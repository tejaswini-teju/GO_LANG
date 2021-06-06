package main

import "testing"

func TestAdd(t *testing.T) {
	ans := Add(5, 5)
	if ans != 10 {
		t.Errorf("Add was incorrect,got %d, expected:%d", ans, 10)
	}

}
