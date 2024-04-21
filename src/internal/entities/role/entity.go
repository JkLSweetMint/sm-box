package role

type Role struct {
	ID        int64
	ProjectID int64

	Title        string
	Inheritances []*Inheritance
}

type Inheritance struct {
	*Role
}
