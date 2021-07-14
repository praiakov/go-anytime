// O pacote word oferece utilitários para jogos de palavra
package word

import "unicode"

// IsPalindrome informza se s é lida do mesmo jeito de frente para trás e
// de trás para a frente.
// A diferença entre letras maiúsculas e minúsculas é ignorada, assim como
// os caracteres que não são letras
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
