package identity

type FullName struct {
	FirstName string
	LastName  string
}

func (f FullName) String() string {
	return f.FirstName + " " + f.LastName
}

func NewFullName(firstName, lastName string) FullName {
	return FullName {
		FirstName: firstName,
        LastName: lastName,
	}
}