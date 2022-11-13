package main

type Employee struct {
	name        string
	leavesTaken int
}

func (emp Employee) getLeavesLeft() int {
	var leavesLeft = 20 - emp.leavesTaken
	return leavesLeft
}
