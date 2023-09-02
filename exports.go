package module

const (
	topicCreate = "create"
	topicDelete = "delete"
	topicUpdate = "update"
	topicRead   = "read"
	topicIndex  = "index"
)

var (
	Exports []string = []string{
		topicCreate,
		topicDelete,
		topicUpdate,
		topicRead,
		topicIndex,
	}
)
