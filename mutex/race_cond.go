package mutex

import "sync"

type Counter struct {
	sync.Mutex //remove this line to impelement race condition.
	Val        int
}

func (v *Counter) Add(val int) {
	v.Lock() //remove this line to impelement race condition.
	v.Val++
	v.Unlock() //remove this line to impelement race condition.
}

func (v *Counter) Value() int {
	return v.Val
}

// Basically will impelement goroutine, waitgroup and mutex.
// This example to implement race conditions.
func RaceConditions() int {
	var wg sync.WaitGroup
	var meter Counter
	for index := 0; index < 1000; index++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				meter.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return meter.Value()
}
