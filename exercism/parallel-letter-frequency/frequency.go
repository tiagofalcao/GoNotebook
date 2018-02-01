// Package letter implements the frequency count of runes in strings.
package letter

// FreqMap stores the frequency of runes.
type FreqMap map[rune]int

// Frequency returns the FreqMap for one string.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency returns FreqMap for one slice of strings.
func ConcurrentFrequency(i []string) FreqMap {
	c := make(chan FreqMap, 10)
	for _, s := range i {
		go func(c chan FreqMap, s string) {
			c <- Frequency(s)
		}(c, s)
	}
	f := <-c
	for x := 1; x < len(i); x++ {
		f.add(<-c)
	}
	return f
}

// add joins FreqMaps.
func (f FreqMap) add(a FreqMap) {
	for r, c := range a {
		f[r] += c
	}
}
