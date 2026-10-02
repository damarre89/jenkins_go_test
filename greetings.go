package main

const Greetings = "Hello, world!"

type Greeter interface {
	Greet() string
	AddGreetingsMessage(msg string)
}

type People struct {
	// only one param
	sGreet string
}

func (p *People) AddGreetingsMessage(msg string) error {
	p.sGreet = msg
	return nil
}

func (p *People) Greet() (string, error) {
	return p.sGreet, nil
}
