package main

// SliceRemoveInt remove int values from slice
func SliceRemoveInt(s []int64, key int64) (so []int64) {
	rmc := 0
	for i := 0; i < len(s)-rmc; i++ {
		if s[i] == key {
			rmc = rmc + 1
			if i < len(s)-rmc {
				s[i] = s[len(s)-rmc]
			}
			i = i - 1
		}
	}

	return s[0 : len(s)-rmc]
}

// SliceRemoveIntList remove int slice values from slice
func SliceRemoveIntList(s []int64, keys []int64) (so []int64) {
	rmc := 0
	for i := 0; i < len(s)-rmc; i++ {
		for j := 0; j < len(keys); j++ {
			if s[i] == keys[j] {
				rmc = rmc + 1
				if i < len(s)-rmc {
					s[i] = s[len(s)-rmc]
				}
				i = i - 1
				continue
			}
		}
	}

	return s[0 : len(s)-rmc]
}

// SliceRemoveString remove string values from slice
func SliceRemoveString(s []string, key string) (so []string) {
	rmc := 0
	for i := 0; i < len(s)-rmc; i++ {
		if s[i] == key {
			rmc = rmc + 1
			if i < len(s)-rmc {
				s[i] = s[len(s)-rmc]
			}
			i = i - 1
		}
	}

	return s[0 : len(s)-rmc]
}

// SliceRemoveStringList remove string slice values from slice
func SliceRemoveStringList(s []string, keys []string) (so []string) {
	rmc := 0
	for i := 0; i < len(s)-rmc; i++ {
		for j := 0; j < len(keys); j++ {
			if s[i] == keys[j] {
				rmc = rmc + 1
				if i < len(s)-rmc {
					s[i] = s[len(s)-rmc]
				}
				i = i - 1
				continue
			}
		}
	}

	return s[0 : len(s)-rmc]
}
