package test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(name string) {
	p.speak()
	fmt.Println(name)
}

type Dog struct {
	Pet
}

func (d *Dog) speak() {
	fmt.Println("dog")
}

func TestDog(t *testing.T) {
	var dog = new(Dog)
	dog.SpeakTo("xxx")
}
