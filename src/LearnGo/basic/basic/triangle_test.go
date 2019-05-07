package main

import "testing"

/**
 * @Author: zhangyan
 * @Date: 2019/5/6 22:56
 * @Version 1.0
 */
func TestTiangle(t *testing.T) {
	test := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}
	for _, tt := range test {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calctriangle (%d %d);"+"got"+
				"%d;expected %d",
				tt.a, tt.b, actual, tt.c)
		}
	}
}
