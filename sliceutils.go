package sliceutils

func Filter[T any](slc []T, predicate func(T) bool) []T {
	var newSlc []T
	for _, element := range slc {
		if predicate(element) {
			newSlc = append(newSlc, element)
		}
	}
	return newSlc
}

func Map[T, S any](slc []T, mapping func(T) S) []S {
	newSlc := make([]S, len(slc))
	for idx, element := range slc {
		newSlc[idx] = mapping(element)
	}
	return newSlc
}

func Fold[T, S any](slc []T, identity S, accumulator func(S, T) S) S {
	result := identity
	for _, element := range slc {
		result = accumulator(result, element)
	}
	return result
}

func ZipAll[T, S, R any](slc1 []T, default1 T, slc2 []S, default2 S, combiner func(T, S) R) []R {
	n, m := len(slc1), len(slc2)
	short, long := min(n, m), max(n, m)
	newSlc := make([]R, long)
	for i := 0; i < short; i++ {
		newSlc[i] = combiner(slc1[i], slc2[i])
	}
	if n < m {
		for i := short; i < long; i++ {
			newSlc[i] = combiner(default1, slc2[i])
		}
	} else {
		for i := short; i < long; i++ {
			newSlc[i] = combiner(slc1[i], default2)
		}
	}
	return newSlc
}

func Zip[T, S, R any](slc1 []T, slc2 []S, combiner func(T, S) R) []R {
	n, m := len(slc1), len(slc2)
	l := min(n, m)
	newSlc := make([]R, l)
	for i := 0; i < l; i++ {
		newSlc[i] = combiner(slc1[i], slc2[i])
	}
	return newSlc
}

func BagEqual[T comparable](slc1, slc2 []T) bool {
	if len(slc1) != len(slc2) {
		return false
	}
	counts := make(map[T]int)
	for _, e := range slc1 {
		_, ok := counts[e]
		if !ok {
			counts[e] = 0
		} else {
			counts[e]++
		}
	}
	for _, e := range slc2 {
		c, ok := counts[e]
		if !ok {
			return false
		} else if c > 0 {
			counts[e]--
		} else if c == 0 {
			delete(counts, e)
		}
	}
	for _, c := range counts {
		if c == 0 {
			panic("Counter with zero should not exist anymore.")
		}
		return false
	}
	return true
}
