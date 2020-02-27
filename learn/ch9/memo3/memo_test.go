package memo_test

import (
	"testing"

	memo "gopl.io/ch9/memo3"
	"gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func TestMemo(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe. Test fail
func TestMemoConc(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
