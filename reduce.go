package reducestring

import (
	"sort"
	"strconv"
)

func Reduce(s string) string {
	counts := map[string]int64{}
	keys := []string{}
	for _, r := range s {
		r := string(r)
		count, ok := counts[r]
		if !ok {
			keys = append(keys, r)
		}
		counts[r] = count + 1
	}

	sort.Strings(keys)

	result := ""
	for _, k := range keys {
		result += k
		result += strconv.FormatInt(counts[k], 10)
	}

	return result
}
