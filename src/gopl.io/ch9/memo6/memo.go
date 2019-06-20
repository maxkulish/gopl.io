// Package memo provides a concurrency-unsafe
// memoization of a function of type Func
package memo

import "sync"

// A Memo caches the results of calling a Func
type Memo struct {
	f     Func
	cache sync.Map
}

// Func is the type of the function to memorize
type Func func(key string) (interface{}, error)

func New(f Func) *Memo {
	return &Memo{f: f, cache: sync.Map{}}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	res, ok := memo.cache.Load(key)
	if !ok {
		res, err = memo.f(key)
		if err != nil {

		}
		memo.cache.Store(key, res)
	}

	return res, err
}
