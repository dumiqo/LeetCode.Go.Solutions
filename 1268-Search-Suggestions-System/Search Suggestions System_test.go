package leetcode

import (
	"testing"
)

func TestSuggestedProducts(t *testing.T) {
	var tests = []struct {
		name       string
		result     [][]string
		products   []string
		searchWord string
	}{
		{"Search mouse", [][]string{{"mobile", "moneypot", "monitor"}, {"mobile", "moneypot", "monitor"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}}, []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"},
		{"Search havana", [][]string{{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}}, []string{"havana"}, "havana"},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := suggestedProducts(tt.products, tt.searchWord)
			if len(tt.result) != len(tt.result) {
				t.Errorf("got wrong result length %v, want %v", len(ans), len(tt.result))
			}
			for i, l := range ans {
				expected := tt.result[i]
				if len(l) != len(expected) {
					t.Errorf("got wrong result on %v index, length %v, want %v", i, len(l), len(expected))
				}
				for ii, item := range expected {
					res := l[ii]
					if res != item {
						t.Errorf("got wrong result on %v, %v indexs, value %v, want %v", i, ii, res, item)
					}
				}
			}
		})
	}
}
