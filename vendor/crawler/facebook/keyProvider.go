package facebook

import (
	"math/rand"
	"os"
	"strings"
)

var (
	tokens []string
)

func getAccessToken() string {

	if tokens == nil {
		initTokens()
	}

	i := rand.Int() % len(tokens)
	token := tokens[i]
	return token
}

func initTokens() {
	tokens = strings.Split(os.Getenv("FACEBOOK_PAGE_ACCESS_TOKENS"), ",")
	for i, token := range tokens {
		tokens[i] = strings.Trim(token, " \t\n")
	}
}
