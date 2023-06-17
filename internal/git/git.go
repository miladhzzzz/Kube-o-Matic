package git

import (
    "github.com/go-git/go-git/v5"
    "github.com/go-git/go-git/v5/plumbing/object"
    myhttp "github.com/go-git/go-git/v5/plumbing/transport/http"
    urlHelper "net/url"
    "os"
    "strings"
)

type Git interface {
    Clone(url string) (*object.Commit, string, error)
    Fetch() error
    Checkout(branch string) error
    Status() (string, error)
	Data() (*GitRepository)
}

type GitRepository struct {
    url        string
    directory  string
    private    bool
    token      string
    repository *git.Repository
}

func NewGitRepository(url string, private bool, token string) (*GitRepository, error) {
    directory := "/git/" + ExtractUsernameRepo(url)
    fs, _ := os.Stat(directory)
    var repository *git.Repository
    var err error
    if fs == nil {
        repository, err = git.PlainClone(directory, false, &git.CloneOptions{
            URL:               url,
            RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
        })
    } else {
        if private {
            repository, err = git.PlainOpen(directory)
            if err!= nil {
                return nil, err
            }
            err = repository.Pull(&git.PullOptions{RemoteName: "origin", Auth: &myhttp.BasicAuth{
                Username: "abc123",
                Password: token,
            }})
        } else {
            repository, err = git.PlainOpen(directory)
            if err!= nil {
                return nil, err
            }
            err = repository.Fetch(&git.FetchOptions{})
        }
    }
    if err!= nil {
        return nil, err
    }
    return &GitRepository{
        url:        url,
        directory:  directory,
        private:    private,
        token:      token,
        repository: repository,
    }, nil
}

func (r *GitRepository) Clone() (*object.Commit, string, error) {
    return nil, "", nil
}

func (r *GitRepository) Data() string {
	return r.directory
}

func (r *GitRepository) Fetch() error {
    w, err := r.repository.Worktree()
    if err!= nil {
        return err
    }
    err = w.Pull(&git.PullOptions{RemoteName: "origin"})
    if err!= nil && err.Error() == "already up-to-date" {
        return nil
    }
    return err
}


func (r *GitRepository) Checkout(branch string) error {
    head, err := r.repository.Head()
    if err!= nil {
        return err
    }
    err = r.repository.Checkout(&git.CheckoutOptions{
        Branch: head.Name(),
        Create: false,
    })
    return err
}


func (r *GitRepository) Status() (string, error) {
    status, err := r.repository.Status()
    if err!= nil {
        return "", err
    }
    return status.String(), nil
}

func ExtractUsernameRepo(url string) (usernameRepo string) {
    mix, err := urlHelper.Parse(url)
    if err!= nil {
        return ""
    }
    path := mix.Path
    path = strings.TrimSuffix(path, ".git") // remove.git extension
    path = strings.TrimSuffix(path, "/")    // remove trailing slash
    parts := strings.Split(path, "/")
    if len(parts) < 3 {
        return ""
    }
    usernameRepo = parts[1] + "/" + parts[2]
    return usernameRepo
}