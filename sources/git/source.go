package git

import "github.com/go-git/go-git/v5"

type gitConfigSource struct {
	name       string
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

func (source *gitConfigSource) Load(application, profile, label string) (map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (source *gitConfigSource) Close() error {
	//TODO implement me
	panic("implement me")
}
