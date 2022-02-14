package public_stuff

import "fmt"

type PublicAgent struct {
}

func (PublicAgent) SayIAmPublicFunc() string {
	iWantToSay := "I am public func"
	fmt.Println(iWantToSay)
	return iWantToSay
}

// TODO add more examples of public functions e.g.:
// func (a A) Foo() {}
