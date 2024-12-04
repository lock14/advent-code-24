package util

func MapSlice[T, U any](s []T, f func(T) U) []U {
	s2 := make([]U, 0, len(s))
	for _, v := range s {
		s2 = append(s2, f(v))
	}
	return s2
}

func MapSliceErr[T, U any](s []T, f func(T) (U, error)) ([]U, error) {
	s2 := make([]U, 0, len(s))
	for _, v := range s {
		u, err := f(v)
		if err != nil {
			return nil, err
		}
		s2 = append(s2, u)
	}
	return s2, nil
}
