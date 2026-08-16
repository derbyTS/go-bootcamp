package main

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 5

type buffer struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

func newBuffer(size int) *buffer {
	b := &buffer{
		items: make([]int, 0, size),
		// cond: sync.NewCond(&b.mu), // We can't do this because we just creating instance of b
	}
	b.cond = sync.NewCond(&b.mu)
	return b
}

// PRODUCER GOROUTINE (b.produce)            CONSUMER GOROUTINE (b.consume)
// ------------------------------            ------------------------------
// 1. Locks b.mu
// 2. Checks len == 5 (TRUE: Buffer Full)
// 3. Calls b.cond.Wait()
//    ├── Unlocks b.mu
//    └── Goes to sleep... (PAUSED)
//                                           4. Locks b.mu
//                                           5. Removes 1 item from b.items (len is now 4)
//                                           6. Calls b.cond.Signal()  ───┐ (Wakes Producer)
//                                           7. Unlocks b.mu             │
// 8. Producer Wakes Up <────────────────────────────────────────────────┘
//    ├── Re-locks b.mu
//    └── Returns from Wait()
// 9. Loop condition checked:
//    len == 5 is now FALSE (len is 4)
// 10. Exits for loop!
// 11. Appends new item (len becomes 5 again)
// 12. Unlocks b.mu

func (b *buffer) produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == bufferSize {
		b.cond.Wait()
	}

	b.items = append(b.items, item)
	fmt.Println("Produce: ", item)
	b.cond.Signal()
}

func (b *buffer) consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == 0 {
		b.cond.Wait()
	}

	item := b.items[0]
	b.items = b.items[1:]

	fmt.Println("Consumed: ", item)
	b.cond.Signal()

	return item
}

func producer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		b.produce(i + 100)
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 10 {
		b.consume()
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	buffer := newBuffer(bufferSize)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(buffer, &wg)
	go consumer(buffer, &wg)

	wg.Wait()
}
