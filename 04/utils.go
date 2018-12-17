package main

// max returns the key of the provided map that has the highest value
func max(m map[interface{}]int) (interface{}, int) {
	var maxKey interface{}
	var maxValue int
	for currentKey, currentValue := range m {
		if currentValue > maxValue {
			maxKey = currentKey
			maxValue = currentValue
		}
	}
	return maxKey, maxValue
}
