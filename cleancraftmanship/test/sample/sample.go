package sample

type StackTest struct {
}

func (s StackTest) CanCreateStack() *StackTest {
	return &StackTest{}
}

func (s StackTest) IsEmpty() bool {
	return true
}