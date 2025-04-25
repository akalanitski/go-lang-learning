package main

type User struct {
	ID   int
	Name string
}

func (db Datastore) GetUser(id int) User {
	return db.GetUser(id)
}

func (db Datastore) SetUsers(user User) {
	db.SetUsers(user)
}

type Datastore interface {
	GetUser(id int) User
	SetUser(User)
}

func main() {

}
