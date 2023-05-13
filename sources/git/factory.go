package git

import (
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/pkg/errors"

	"github.com/realbucksavage/teller"
	"github.com/realbucksavage/teller/sources"
)

func init() {
	sources.RegisterFactory("git", teller.NewFactory(newConfigSource, defaultConfiguration))
}

func newConfigSource(name string, cfg teller.Configuration) (teller.ConfigSource, error) {

	oCfg := *(cfg.(*Configuration))

	repo, err := git.Clone(memory.NewStorage(), memfs.New(), &git.CloneOptions{
		URL: oCfg.Repository,
	})
	if err != nil {
		return nil, errors.Wrap(err, "clone error")
	}

	return &gitConfigSource{
		name:       name,
		repository: repo,
	}, nil
}
