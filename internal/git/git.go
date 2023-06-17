package git

import (
    "github.com/go-git/go-git/v5"
    urlHelper "net/url"
    "os"
    "strings"
)

type GitRepository struct {
    Url        string
    Directory  string
    Private    bool
    Token      string
    Repository *git.Repository
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
        Url:        url,
        Directory:  directory,
        Private:    private,
        Token:      token,
        Repository: repository,
    }, nil
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