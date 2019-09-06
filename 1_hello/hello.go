package main

import "fmt"

func main() {
	fmt.Println(Hello("World", ""))
}

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const franchHelloPrefix = "Bonjor, "
const spanish = "Spanish"
const french = "French"

// Hello aaa
func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = franchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
