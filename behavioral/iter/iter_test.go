package iter

import (
	"fmt"
	"testing"
)

func TestIteratorPattern(t *testing.T) {
	users := NewUserCollection([]*User{
		{name: "Jack", age: 11},
		{name: "Lucy", age: 12},
		{name: "Rose", age: 13},
	})

	iter := users.NewIterator()

	for iter.HasNext() {
		u := iter.GetNext()
		fmt.Println(u)
	}
}
