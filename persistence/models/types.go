package models

type TaskPriority string

const (
	P0TaskPriority TaskPriority = "p0"
	P1TaskPriority TaskPriority = "P1"
	P2TaskPriority TaskPriority = "P2"
	P3TaskPriority TaskPriority = "P3"
	P4TaskPriority TaskPriority = "P4"
)

func (t TaskPriority) Valid() bool {
	switch t {
	case P0TaskPriority, P1TaskPriority, P2TaskPriority, P3TaskPriority, P4TaskPriority:
		return true
	}

	return false
}

func ListTaskPriority() []string {
	return []string{
		string(P0TaskPriority),
		string(P1TaskPriority),
		string(P2TaskPriority),
		string(P3TaskPriority),
		string(P4TaskPriority),
	}
}
