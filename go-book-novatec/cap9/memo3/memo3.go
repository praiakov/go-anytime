package memo

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex // guarda cache
	cache map[string]result
}

type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)

		// Entre as duas seções críticas, várias gorrotinas
		// podem concorrer para processar f(key) e atualizar o map
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	memo.mu.Unlock()
	return res.value, res.err
}
