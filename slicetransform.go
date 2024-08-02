package slicetransform

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

func Reduce[T, S any](slc []T, identity S, accumulator func(S, T) S) S {
	result := identity
	for _, element := range slc {
		result = accumulator(result, element)
	}
	return result
}
