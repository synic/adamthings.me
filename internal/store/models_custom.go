package store

func (c Comment) ApprovalURL() string {
	return "/comment/approve/" + c.ApprovalCode
}
