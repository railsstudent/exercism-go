package letter

type FreqMap map[rune]int

// Frequency counts frequency of letters in a string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency find the frequency of characters in string array in parallel
// func ConcurrentFrequency(s []string) FreqMap {
// 	var wg sync.WaitGroup
// 	wg.Add(len(s))
//
// 	m := FreqMap{}
// 	for i := range s {
// 		go func(i int) {
// 			defer wg.Done()
// 			subMap := Frequency(s[i])
// 			mutex.Lock()
// 			for k, v := range subMap {
// 				m[k] += v
// 			}
// 			mutex.Unlock()
// 		}(i)
// 	}
// 	wg.Wait()
// 	return m
// }

// ConcurrentFrequency find the frequency of characters in string array in parallel
func ConcurrentFrequency(s []string) FreqMap {
	c := make(chan FreqMap)
	m := FreqMap{}

	// Use this example: https://gobyexample.com/closing-channels
	go func() {
		for _, str := range s {
			m := Frequency(str)
			c <- m
		}
		close(c)
	}()

	for subMap := range c {
		for k, v := range subMap {
			m[k] += v
		}
	}
	return m
}
