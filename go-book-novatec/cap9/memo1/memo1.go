// o pacote memo oferece uma memoização
// segura para concorrência de uma função do tipo Func
package memo

// um memo faz cache dos resultados da chamada a uma Func
type Memo struct {
	f     Func
	cache map[string]result
}

//Func é o tipo da função para memoizar
type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// NOTA: não é seguro para concorrência!
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
