package utils

type NewStoryChecker interface {
	IsNewSotry(id int64) (bool, error)
	AddPostedStory(id int64) error
}
