package leetcode

import "sort"

const maxNonOverlapPicks = 4

// maximumWeight picks at most 4 pairwise non-overlapping intervals with the
// largest total weight, breaking ties by the lexicographically smallest list
// of original indices.
//
// Interval i starts at li and ends at ri (inclusive); two intervals overlap
// when they share any point, so the next interval may only start strictly
// after the current one ends.
func maximumWeight(intervals [][]int) []int {
	n := len(intervals)

	// Attach the original index to each interval and sort by start.
	arr := make([][4]int, n) // [l, r, weight, originalIndex]
	for i, e := range intervals {
		arr[i] = [4]int{e[0], e[1], e[2], i}
	}
	sort.Slice(arr, func(a, b int) bool {
		if arr[a][0] != arr[b][0] {
			return arr[a][0] < arr[b][0]
		}
		return arr[a][1] < arr[b][1]
	})

	// nxt[i] = first position j > i with arr[j][0] > arr[i][1]
	// (endpoints can not be shared, so the comparison is strict).
	nxt := make([]int, n)
	for i := 0; i < n; i++ {
		lo, hi := i+1, n
		for lo < hi {
			mid := (lo + hi) / 2
			if arr[mid][0] > arr[i][1] {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		nxt[i] = lo
	}

	// f[i][k] = max total weight from interval i onward using at most k picks,
	// g[i][k] = lexicographically smallest index list achieving f[i][k].
	f := make([][maxNonOverlapPicks + 1]int64, n+1)
	g := make([][maxNonOverlapPicks + 1][]int, n+1)
	for k := 0; k <= maxNonOverlapPicks; k++ {
		g[n][k] = []int{}
	}

	for i := n - 1; i >= 0; i-- {
		g[i][0] = []int{}
		for k := 1; k <= maxNonOverlapPicks; k++ {
			skipW := f[i+1][k]
			skipL := g[i+1][k]

			takeW := f[nxt[i]][k-1] + int64(arr[i][2])
			takeL := insertSortedIndex(g[nxt[i]][k-1], arr[i][3])

			if takeW > skipW || (takeW == skipW && lexicographicallyLess(takeL, skipL)) {
				f[i][k] = takeW
				g[i][k] = takeL
			} else {
				f[i][k] = skipW
				g[i][k] = skipL
			}
		}
	}
	return g[0][maxNonOverlapPicks]
}

// insertSortedIndex returns a new sorted slice with x inserted into the sorted
// slice list (all elements are distinct original indices).
func insertSortedIndex(list []int, x int) []int {
	out := make([]int, len(list)+1)
	i := 0
	for i < len(list) && list[i] < x {
		out[i] = list[i]
		i++
	}
	out[i] = x
	for i < len(list) {
		out[i+1] = list[i]
		i++
	}
	return out
}

// lexicographicallyLess reports whether a is lexicographically smaller than b,
// treating a strict prefix as smaller (shorter array wins when equal so far).
func lexicographicallyLess(a, b []int) bool {
	m := len(a)
	if len(b) < m {
		m = len(b)
	}
	for i := 0; i < m; i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}
