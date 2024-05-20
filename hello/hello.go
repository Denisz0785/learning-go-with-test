package main

import (
	"net/http"
	"os"
	"test_to_learn_go/injiection"
	"test_to_learn_go/mocking"
	"time"
)

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "

	spanish = "Spanish"
	french  = "French"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingsPrefix(language) + name

}

func greetingsPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
func GreetingsHandler(w http.ResponseWriter, r *http.Request) {
	injiection.Greet(w, "my World!")
}

func main() {
	sleeper := mocking.ConfigurableSleeper{1 * time.Second, time.Sleep}

	mocking.CountDown(os.Stdout, &sleeper)

}
