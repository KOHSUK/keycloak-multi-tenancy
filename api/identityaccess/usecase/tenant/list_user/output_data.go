package list_user

type User struct {
	Id   string
	Name string
}

type ListUserOutputData struct {
	Users []User
}
