package iter

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

func (u *User) String() string {
	return fmt.Sprintf("{ name: %s, age: %d }", u.name, u.age)
}

type UserCollection struct {
	users []*User
}

func NewUserCollection(users []*User) *UserCollection {
	return &UserCollection{users: users}
}

func (c *UserCollection) NewIterator() Iterator[User] {
	return &UserIterator{
		index: 0,
		users: c.users,
	}
}

type UserIterator struct {
	index int
	users []*User
}

func (i *UserIterator) GetNext() *User {
	u := i.users[i.index]
	i.index += 1
	
	return u
}

func (i *UserIterator) HasNext() bool {
	return i.index < len(i.users)
}
