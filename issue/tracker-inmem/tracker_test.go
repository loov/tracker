package main

func TestSingleIssue(t *testing.T) {
	tracker := New()
	issue.TrackerSingleIssueTest(tracker, t)
}