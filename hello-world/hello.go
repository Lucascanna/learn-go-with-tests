package main

import "fmt"

const helloEn = "Hello, "
const helloSp = "Hola, "
const helloFr = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := helloEn
	switch language {
	case "spanish":
		prefix = helloSp
	case "french":
		prefix = helloFr
	}
	return prefix + name
}

func getPrefix(language string) string {
	switch language {
	case "spanish":
		return helloSp
	case "french":
		return helloFr
	}
	return helloEn
}

func main() {
	fmt.Println(Hello("world", "english"))
}
