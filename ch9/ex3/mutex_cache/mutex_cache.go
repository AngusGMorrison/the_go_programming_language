// Extend the Func type and the (*Memo).Get method so that callers may provide an optional done
// channel through which they can cancel the operation (§8.9). The results of a cancelled Func call
// should not be cached.
package mutexcache

import (
	"fmt"
	"sync"
)

// Func is the type of the function to memoize.
type Func func(key string, done <-chan struct{}) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

// An entry is the representation of a memoized function call stored in the cache.
type entry struct {
	res    result
	ready  chan struct{}   // closed when res is ready
	cancel <-chan struct{} // entry should be recalculated
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
func (memo *Memo) Get(key string, cancel <-chan struct{}) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e != nil {
		select {
		case <-e.cancel:
			// Function was cancelled; recalculate the entry.
		case <-e.ready:
			value, err = e.res.value, e.res.err
			memo.mu.Unlock()
			return
		}
	}

	e = &entry{
		ready:  make(chan struct{}),
		cancel: cancel,
	}
	memo.cache[key] = e
	memo.mu.Unlock()

	e.res.value, e.res.err = memo.f(key, cancel)
	select {
	case <-cancel:
		return nil, fmt.Errorf("getting %s: cancelled", key)
	default:
		close(e.ready)
		return e.res.value, e.res.err
	}
}
