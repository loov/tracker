package issue

type ID int

type Status string

const (
	Created Status = "Created"
	Closed         = "Closed"
)

type Info struct {
	ID      ID
	Caption string
	Desc    string
	Status  Status
}
