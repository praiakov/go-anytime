package memo

import "sync"

type entry struct {
	res   result
	ready chan struct{} //fechado quando res estiver pronto
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guarda cache
	cache map[string]*entry
}

type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// Essa é a primeira requisição para esse chave
		// Essa gorrotina torna-se responsável por calcular
		// o valor e fazer broadcast da condição de pronto
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready) //faz brodcast da condição de pronto
	} else {
		// essa é a uma requisição repetida para essa chave
		memo.mu.Unlock()

		<-e.ready // espera a condição de pronto
	}
	return e.res.value, e.res.err
}
