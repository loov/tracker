package issue

import "errors"

var ErrNotExist = errors.New("issue does not exist")

// Tracker tracks issues
type Tracker interface {
	// Create opens a new issue with the appropriate info
	Create(info Info) (ID, error)
	// Load gets information about a specific issue
	Load(id ID) (Info, error)
	// Close closes an issue
	Close(id ID) error
	// List gets the all issues
	List() (issues []Info, err error)
}
