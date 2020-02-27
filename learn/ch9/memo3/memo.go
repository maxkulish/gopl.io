// Package memo provides a concurrency-unsafe
// memoization of a function of type Func
package memo

import "sync"

// A Memo caches the results of calling a Func
type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]result
}

// Func is the type of the function to memorize
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Get is concurrency-safe
func (memo *Memo) Get(key string) (value interface{}, err error) {
	// First lock - lookup key in cache
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		// Second lock - update if the lookup return nothing
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}