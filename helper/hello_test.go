package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloAssert(t *testing.T) {
	result := HelloWorld("Ray")

	assert.Equal(t, "Hello Ray", result, "Result must be hello ray")
	fmt.Println("TestHelloAssert is Done")
}

// assert vs require = assert: fail; require: failnow

 func TestHello(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		// error
		panic("Result is not Hello")
	}
 }

 func TestHelloRay(t *testing.T){
	result := HelloWorld("Rayn")

	if result != "Hello Ray" {
		t.Fail()  //continue then failed;
		/*
		fail.now() failed immediately
		t.Error(...)
		t.fatal(...)
		*/
	}

	fmt.Println("TestHelloRay is Done")
 }


 // untuk menjalankan unit test: go test or go test -v or go test -v -run=TestHello