package issue

// Unique ID for an issue
type ID int64

// Status of an issue
type Status string

const (
	//TODO: should tracker.Create renamed to "Open", because status is Open
	Open = Status("Open") // Open means that the issue needs to be worked on
	//TODO: should "Done" be renamed to "Closed", because tracker has method Close
	Done = Status("Done") // Done means that the issue is completed and delivered
)

// Info contains primary information about an issue
type Info struct {
	// ID is unique identifier to an issue
	ID ID
	// Caption is the clarifying name for the issue
	Caption string
	// Desc is a description for the issue
	Desc string
	// Status is the latest status for the issue
	Status Status
}
