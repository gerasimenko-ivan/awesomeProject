package test_files

import (
	"awesomeProject/access/public_stuff"
	"testing"
)

// test file must have suffix '_test.go'
// test method must start with 'Test'
// test method can test functions from other packages
func TestPublicFunc(t *testing.T) {
	pa := new(public_stuff.PublicAgent)
	resp := pa.SayIAmPublicFunc()
	if resp != "I am public func" {
		t.Error()
	}
}
