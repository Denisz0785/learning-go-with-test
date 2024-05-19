package main

import (
	"log"
	"net/http"
	"test_to_learn_go/injiection"
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
	injiection.Greet(w, "world!")
}

func main() {

	log.Fatal(http.ListenAndServe(":3333", http.HandlerFunc(GreetingsHandler)))

}
