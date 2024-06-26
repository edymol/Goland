package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}
type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreenting(eb)
	printGreenting(sb)
}

func printGreenting(b bot) {
	fmt.Println(b.getGreeting())
}

//	func (eb englishBot) getGreeting() string {
//		return "Hi there"
//
// the eb can be removed if the vabiables are not being used
func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

//func printGreeting(eb englishBot) {
//	fmt.Println(eb.getGreeting())
//}
//
//func printGreeting(sb englishBot) {
//	fmt.Println(sb.getGreeting())
