package leetcode

/*
решение взял из https://leetcode.com/problems/lexicographically-smallest-permutation-greater-than-target/solutions/8485150/1-by-galac79155-0000
не смог решить задачу за выделенное время.

Вместо того чтобы пытаться собрать ответ, алгоритм ищет самую правую позицию в target, которую можно заменить на больший символ из s, чтобы получить перестановку, строго большую, чем target. Если такая позиция найдена, всё, что справа от неё, заполняется наименьшими возможными символами — это даёт лексикографически наименьшую перестановку.

*/
func lexGreaterPermutation(s string, target string) string {
	cnt := make([]int, 26)

	for i := range s {
		cnt[s[i]-'a']++
	}
	for i := range target {
		cnt[target[i]-'a']--
	}

	bad, mx := 0, -1
	for c := 0; c < 26; c++ {
		if cnt[c] < 0 {
			bad++
		}
		if cnt[c] > 0 {
			mx = c
		}
	}

	for i := len(target) - 1; i >= 0; i-- {
		cur := int(target[i] - 'a')
		cnt[cur]++

		if cnt[cur] == 0 {
			bad--
		} else if cnt[cur] == 1 && cur > mx {
			mx = cur
		}

		if bad > 0 || mx <= cur {
			continue
		}

		next := cur + 1
		for cnt[next] == 0 {
			next++
		}

		cnt[next]--

		ans := []byte(target[:i])
		ans = append(ans, byte('a'+next))

		for c := 0; c < 26; c++ {
			for cnt[c] > 0 {
				ans = append(ans, byte('a'+c))
				cnt[c]--
			}
		}

		return string(ans)
	}

	return ""
}
