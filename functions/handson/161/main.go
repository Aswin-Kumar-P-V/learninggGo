package main

import "fmt"

type User struct {
	Id    int
	first string
}

type mockDB struct {
	Users map[int]User
}

func (md mockDB) getUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("USER WITH ID %v NOT FOUNFD", id)
	}
	return user, nil
}

func (md mockDB) saveUser(u User) error {
	_, ok := md.Users[u.Id]
	if ok {
		return fmt.Errorf("USER WITH ID %v ALREADY EXISTS", u.Id)
	}
	md.Users[u.Id] = u
	return nil
}

type Database interface {
	getUser(int) (User, error)
	saveUser(User) error
}

type service struct {
	ds Database
}

func (s service) getUser(id int) (User, error) {
	return s.getUser(id)
}

func (s service) saveUser(u User) error {
	return s.saveUser(u)
}

func main() {
	db := mockDB{
		make(map[int]User),
	}
	s := service{
		ds: db,
	}

	u := User{
		Id:    1,
		first: "Aswin",
	}
	err := s.ds.saveUser(u)
	if err != nil {
		fmt.Println(err)
	}
	u1, err1 := s.ds.getUser(1)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(u1)
}
