package norepeating

import "testing"

func TestSubStr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		//Normal
		{"abcabc", 3},
		{"owwkew", 3},
		//edge
		{"bbbbb", 1},
		{"", 0},
		{"abcabcabcd", 4},

		//chinses
		{"一二三三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}
	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr3(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}
func BenchmarkSubStr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s)=%d", len(s))
	b.ResetTimer()
	ans := 8
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr3(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
