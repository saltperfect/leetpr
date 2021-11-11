package slidingwindow

import "reflect"

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	pmap := make(map[string]int)
	for _, v := range p {
		pmap[string(v)]++
	}

	smap := make(map[string]int)
	for i := 0; i < len(p); i++ {
		smap[string(s[i])]++
	}

	var res []int
	plen := len(p)

	if reflect.DeepEqual(smap, pmap) {
		res = append(res, 0)
	}
	for i := 0; i < len(s)-plen; i++ {

		val := smap[string(s[i])]
		if val == 1 {
			delete(smap, string(s[i]))
		} else {
			smap[string(s[i])]--
		}

		smap[string(s[i+plen])]++

		if reflect.DeepEqual(smap, pmap) {
			res = append(res, i+1)
		}
	}
	return res
}
