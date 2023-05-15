package git

import "github.com/go-git/go-git/v5"

type gitConfigSource struct {
	name       string
	priority   int
	repository *git.Repository
}

func (source *gitConfigSource) Name() string {
	return source.name
}

func (source *gitConfigSource) Refresh() error {
	//TODO implement me
	panic("implement me")
}

func (source *gitConfigSource) RefreshRate() int {
	//TODO implement me
	panic("implement me")
}

func (source *gitConfigSource) Priority() int {
	return source.priority
}

func (source *gitConfigSource) Load(application, profile, label string) (map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (source *gitConfigSource) Close() error {
	//TODO implement me
	panic("implement me")
}
