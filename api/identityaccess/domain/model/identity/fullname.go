package identity

type FullName struct {
	FirstName string
	LastName  string
}

func (f FullName) String() string {
	return f.FirstName + " " + f.LastName
}
