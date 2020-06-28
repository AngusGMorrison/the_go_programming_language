// Extend the Func type and the (*Memo).Get method so that callers may provide an optional done
// channel through which they can cancel the operation (§8.9). The results of a cancelled Func call
// should not be cached.
package cspcache

import "fmt"

// Func is the type of the function to memoize.
type Func func(key string, cancel <-chan struct{}) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

// An entry is an entry in the cache.
type entry struct {
	res    result
	ready  chan struct{}
	cancel <-chan struct{}
}

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
	cancel   <-chan struct{}
}

type Memo struct{ requests chan request }

// New returns a memoization of f. Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, cancel chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, cancel}
	select {
	case res := <-response:
		return res.value, res.err
	case <-cancel:
		return nil, fmt.Errorf("getting %s: cancelled", key)
	}
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		// If the entry exists and has not been cancelled, deliver it to the client.
		if e, ok := cache[req.key]; ok {
			select {
			case <-e.cancel:
				// Function was cancelled; recreate the entry.
			default:
				go e.deliver(req.response)
				continue
			}
		}

		e := &entry{
			ready:  make(chan struct{}),
			cancel: req.cancel,
		}
		cache[req.key] = e
		go e.call(f, req.key, req.cancel) // call f(req.key, req.cancel)
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string, cancel <-chan struct{}) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key, cancel)
	select {
	case <-cancel:
		// Do nothing.
	default:
		close(e.ready) // broadcast the ready condition
	}
}

func (e *entry) deliver(response chan<- result) {
	select {
	case <-e.ready:
		// Send the result to the client.
		response <- e.res
	case <-e.cancel:
		// Unblock the waiting goroutine when a ready condition cannot occur.
	}
}
