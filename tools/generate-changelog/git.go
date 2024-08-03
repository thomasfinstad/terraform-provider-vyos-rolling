package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/hashicorp/go-version"
	"github.com/leodido/go-conventionalcommits"
	"github.com/leodido/go-conventionalcommits/parser"
)

func GenerateGitChanges() (previousVersion *version.Version, commitsSinceLastVersion []conventionalcommits.Message) {
	repo, err := git.PlainOpen("../..")
	die(err)

	previousVersion, err = version.NewVersion("0.0.0")
	die(err)

	itr, err := repo.Tags()
	die(err)
	err = itr.ForEach(func(ref *plumbing.Reference) error {
		fmt.Println("Comparing versions: ", previousVersion.Original(), ref.Name().Short())
		v, err := version.NewVersion(ref.Name().Short())
		die(err)
		if v.GreaterThan(previousVersion) {
			previousVersion = v
		}
		return nil
	})
	die(err)

	fmt.Println("Previous version: ", previousVersion.Original())
	hash, err := repo.ResolveRevision(plumbing.Revision(previousVersion.Original()))
	die(err)
	releaseCommit, err := repo.CommitObject(*hash)
	die(err)

	commits, err := repo.Log(&git.LogOptions{
		Since: &releaseCommit.Author.When,
		Order: git.LogOrderCommitterTime,
	})
	die(err)

	chgs := make([]conventionalcommits.Message, 0)
	ccm := parser.NewMachine(parser.WithTypes(conventionalcommits.TypesConventional), parser.WithBestEffort())
	err = commits.ForEach(func(commit *object.Commit) error {
		cc, _ := ccm.Parse([]byte(commit.Message))

		chgs = append(chgs, cc)
		return nil
	})
	die(err)

	return previousVersion, chgs
}
