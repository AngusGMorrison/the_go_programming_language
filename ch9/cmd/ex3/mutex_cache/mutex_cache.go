// Extend the Func type and the (*Memo).Get method so that callers may provide an optional done
// channel through which they can cancel the operation (§8.9). The results of a cancelled Func call
// should not be cached.
package mutexCache

import (
	"fmt"
	"sync"
)

// Func is the type of the function to memoize.
type Func func(key string, done <-chan struct{}) (bool, interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

// An entry is the representation of a memoized function call stored in the cache.
type entry struct {
	res    result
	ready  chan struct{} // closed when res is ready
	recalc chan struct{} // entry should be recalculated
}

// A Memo caches the result of calling a Func.
type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// Get returns the cached value of a function call, or calls the function is the value doesn't yet
// exist.
func (memo *Memo) Get(key string, done <-chan struct{}) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key. This goroutine becomes responsible for
		// computing the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		cancelled, value, err := memo.f(key, done)
		if cancelled {
			go func() { e.recalc <- struct{}{} }()
			return nil, fmt.Errorf("operation cancelled for %q", key)
		}
		e.res = result{value, err}

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()
	Loop:
		for {
			select {
			case <-e.ready:
				break Loop
			case <-e.recalc:
				memo.mu.Lock()
				delete(memo.cache, key)
				memo.mu.Unlock()
				return memo.Get(key, done)
			}
		}

	}
	return e.res.value, e.res.err
}
