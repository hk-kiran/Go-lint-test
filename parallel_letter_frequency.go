//NAME: KIRAN
package letter

import (
	"fmt"
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

func (f FreqMap) String() string {
	var output string = "{ "

	for char, count := range f {
		output += fmt.Sprintf("%v = %c : %d | ", char, char, count)
	}

	output += " }"

	return output
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}

	return m
}

// ConcurrentFrequency counts the frequency of each rune in a given slice of text and returns this
// data as a FreqMap.
func ConcurrentFrequency(s []string) FreqMap {
	var wg sync.WaitGroup
	m := make([]FreqMap, len(s))
	n := FreqMap{}
	for i, str := range s {
		wg.Add(1)
		go func(str string, i int) {
			defer wg.Done()
			m[i] = Frequency(str)

		}(str, i)
	}
	wg.Wait()
	for _, k := range m {
		for index, value := range k {
			n[index] += value
		}
	}
	return n
}
