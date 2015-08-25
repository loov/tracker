package tracker

import (
	"sync"

	"github.com/loov/tracker/issue"
)

type Service struct {
	mu     sync.Mutex
	infos  []issue.Info
	lastID ID
}

func New() *Service {
	return &Service{}
}

func (tracker *Service) lock() *Service {
	tracker.mu.Lock()
	return tracker
}
func (tracker *Service) unlock() { tracker.mu.Unlock() }

// Create opens a new issue with the appropriate info
func (tracker *Service) Create(info issue.Info) (issue.ID, error) {
	defer tracker.lock().unlock()

	tracker.lastID++
	info.ID = tracker.lastID
	info.Status = issue.Open
	tracker.infos = append(tracker.infos, info)

	return info.ID, nil
}

// Load gets information about a specific issue
func (tracker *Service) Load(id issue.ID) (issue.Info, error) {
	defer tracker.lock().unlock()

	for _, info := range tracker.infos {
		if info.ID == id {
			return info, nil
		}
	}

	return issue.Info{}, issue.ErrNotExist
}

// Close closes an issue
func (tracker *Service) Close(id issue.ID) error {
	defer tracker.lock().unlock()

	for i := range tracker.infos {
		if tracker.infos[i].ID == id {
			tracker.infos[i].Status = issue.Done
			return nil
		}
	}

	return issue.ErrNotExist
}

// List gets the all issues
func (tracker *Service) List() (issues []issue.Info, err error) {
	defer tracker.lock().unlock()

	issues = make([]issue.Info, len(tracker.infos))
	copy(issues, tracker.infos)

	return issues, nil
}
