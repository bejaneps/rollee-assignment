package rollee

import "sync"

type ID = int

// We suppose L is always valid with len (l.Values) >= 1).
type List struct {
	ID     ID
	Values []int
}

func Fold(initialValue int, f func(int, int) int, l List) map[ID]int {
	for _, val := range l.Values {
		initialValue = f(initialValue, val)
	}

	return map[ID]int{
		l.ID: initialValue,
	}
}

func FoldChan(initialValue int, f func(int, int) int, ch chan List) map[ID]int {
	result := map[ID]int{}

	for c := range ch {
		if _, ok := result[c.ID]; !ok {
			result[c.ID] = initialValue
		}

		for _, val := range c.Values {
			newInitialValue := f(result[c.ID], val)
			result[c.ID] = newInitialValue
		}
	}

	return result
}

func FoldChanX(initialValue int, f func(int, int) int, chs ...chan List) map[ID]int {
	result := map[ID]int{}
	wg := &sync.WaitGroup{}
	receiveMaps := make(chan map[ID]int, len(chs))

	for _, ch := range chs {
		wg.Add(1)

		go func(ch chan List) {
			receiveMaps <- FoldChan(initialValue, f, ch)

			wg.Done()
		}(ch)
	}

	wg.Wait()
	close(receiveMaps)

	for m := range receiveMaps {
		for k, v := range m {
			if _, ok := result[k]; ok {
				result[k] = f(result[k], v)
			} else {
				result[k] = v
			}
		}
	}

	return result
}
