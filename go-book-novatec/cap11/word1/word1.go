// O pacote word oferece utilitários para jogos de palavra
package word

// IsPalindrome informza se s é lida do mesmo jeito de frente para trás e
// de trás para a frente.
// (Nossa primeira tentativa.)
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
