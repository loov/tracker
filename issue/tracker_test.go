package issue

import "testing"

func TrackerSingleIssueTest(tracker Tracker, t *testing.T) {
	id, err := tracker.Create(Info{
		Caption: "Caption",
		Desc:    "Desc",
	})
	if err != nil {
		t.Error(err)
		return
	}

	info, err := tracker.Load(id)
	if err != nil {
		t.Error(err)
		return
	}

	expect := Info{
		ID:      id,
		Caption: "Caption",
		Desc:    "Desc",
		Status:  Open,
	}

	if info != expect {
		t.Errorf("got %#v expected %#v", info, expect)
		return
	}

	list, err := tracker.List()
	if err != nil {
		t.Error(err)
		return
	}
	if len(list) != 1 {
		t.Errorf("got invalid number of issues %v", list)
	}

	err = tracker.Close(id)
	if err != nil {
		t.Error(err)
		return
	}

	info2, err := tracker.Load(info.ID)
	if err != nil {
		t.Error(err)
		return
	}

	expect2 := Info{
		ID:      id,
		Caption: "Caption",
		Desc:    "Desc",
		Status:  Done,
	}
	if info2 != expect2 {
		t.Errorf("got %#v expected %#v", info2, expect2)
		return
	}
}
