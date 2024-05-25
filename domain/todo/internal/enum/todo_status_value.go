package enum

import "github.com/samber/lo"

// ステータスのEnum (内部用)
type StatusValue string

const (
	StatusBacklog = StatusValue("Backlog")
	StatusTodo    = StatusValue("Todo")
	StatusDoing   = StatusValue("Doing")
	StatusDone    = StatusValue("Done")
)

var StatusSeq = []StatusValue{StatusBacklog, StatusTodo, StatusDoing, StatusDone}

func (s StatusValue) Validate() bool { return lo.Contains(StatusSeq, s) }