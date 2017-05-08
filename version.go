package main

import (
	"strings"

	"gopkg.in/libgit2/git2go.v24"
)

func fillUpVersion(repo *git.Repository, refIter *git.ReferenceIterator, versions *[]Version, dateLayout string) {
	for {
		ref, err := refIter.Next()
		if err != nil {
			break
		}
		if !ref.IsTag() {
			continue
		}
		defer ref.Free()
		obj, err := ref.Peel(git.ObjectTag)
		if err != nil {
			Logger.Errorf("[VERSION] Peeling error : %s", err)
			continue
		}
		defer obj.Free()
		tag, err := obj.AsTag()
		if err != nil {
			Logger.Errorf("[VERSION] AsTag error : %s", err)
			continue
		}
		defer tag.Free()
		*versions = append(*versions, Version{
			Number:  tag.Name(),
			Date:    tag.Tagger().When.Format(dateLayout),
			Author:  tag.Tagger().Name,
			Section: []string{T("doc_version_multiple")},
			Log:     strings.Split(tag.Message(), "\n")[0]})
	}
}

func Versions(dateLayout string) []Version {
	var versions []Version
	repo, err := git.OpenRepository(".")
	if err != nil {
		Logger.Warnf("[VERSION] Can't open repo : %s", err)
		return versions
	}
	defer repo.Free()
	refIter, err := repo.NewReferenceIterator()
	if err != nil {
		Logger.Errorf("[VERSION] Can't get new ref iterator : %s", err)
		return versions
	}
	defer refIter.Free()
	fillUpVersion(repo, refIter, &versions, dateLayout)
	return versions
}
