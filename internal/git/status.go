package git

import (
	"sort"

	gogit "github.com/go-git/go-git/v5"
)

type FileStatus struct {
	Path   string
	Staged string
	Work   string
}

func (f FileStatus) ShortStatus() string {
	staged := normalizeStatus(f.Staged)
	work := normalizeStatus(f.Work)
	return staged + work
}

func Status(path string) ([]FileStatus, error) {
	repo, err := gogit.PlainOpenWithOptions(path, &gogit.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		return nil, err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return nil, err
	}

	status, err := worktree.Status()
	if err != nil {
		return nil, err
	}

	files := make([]FileStatus, 0, len(status))
	for path, stat := range status {
		files = append(files, FileStatus{
			Path:   path,
			Staged: string(stat.Staging),
			Work:   string(stat.Worktree),
		})
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].Path < files[j].Path
	})
	return files, nil
}

func normalizeStatus(status string) string {
	if status == "" || status == " " {
		return "."
	}
	return status
}
