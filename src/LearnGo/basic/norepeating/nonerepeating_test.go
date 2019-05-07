package main

import "testing"

/**
 * @Author: zhangyan
 * @Date: 2019/5/6 23:16
 * @Version 1.0
 */
func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		//nomal
		{"abcabcb", 3},
		{"pwwkew", 3},
		//edge,
		{"", 0},
		{"b", 1},
		{"abcabcabcd", 4},
		//chinese
		{"这里是慕课网", 6},
	}
	for _, tt := range tests {
		if actual := lengthofNoRepeatingSubStr(tt.s); actual != tt.ans {
			t.Errorf("Got %d for input %s;"+
				"expected %d", actual, tt.s, tt.ans)
		}

	}
}

func BenchmarkSubsrt(b *testing.B) {
	s, ans := "这里是慕课网网啊啊啊的真这", 6
	for i := 0; i < 13; i++ {
		s = s + s

	}
	b.Logf("len(s)=%d", len(s))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actual := lengthofNoRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("Got %d for input %s;"+
				"expected %d", actual, s, ans)
		}
	}

}
