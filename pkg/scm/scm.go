package scm

import (
	"net/http"

	"github.com/LambdaTest/photon/pkg/errs"
	"github.com/LambdaTest/photon/pkg/lumber"
	"github.com/LambdaTest/photon/pkg/models"
	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/bitbucket"
	"github.com/drone/go-scm/scm/driver/github"
	"github.com/drone/go-scm/scm/driver/gitlab"
	"github.com/drone/go-scm/scm/transport/oauth2"
)

// gitClientProvider provides the git scm client
type gitClientProvider struct {
	logger          lumber.Logger
	gitHubClient    *models.SCM
	gitLabClient    *models.SCM
	bitbucketClient *models.SCM
}

// New initializes GitClientProvider
func New(logger lumber.Logger) models.SCMProvider {
	return &gitClientProvider{
		logger:          logger,
		gitHubClient:    &models.SCM{Client: provideGithubClient()},
		gitLabClient:    &models.SCM{Client: provideGitlabClient()},
		bitbucketClient: &models.SCM{Client: provideBitbucketClient()},
	}
}

func (g *gitClientProvider) GetClient(scmClientID models.SCMDriver) (*models.SCM, error) {
	switch scmClientID {
	case models.DriverGithub:
		return g.gitHubClient, nil
	case models.DriverGitlab:
		return g.gitLabClient, nil

	case models.DriverBitbucket:
		return g.bitbucketClient, nil
	default:
		return nil, errs.ErrInvalidDriver
	}
}

func provideGithubClient() *scm.Client {
	client := github.NewDefault()
	client.Client = &http.Client{
		Transport: &oauth2.Transport{
			Source: oauth2.ContextTokenSource(),
			Base:   http.DefaultTransport,
		},
	}
	return client
}

func provideGitlabClient() *scm.Client {
	client := gitlab.NewDefault()
	client.Client = &http.Client{
		Transport: &oauth2.Transport{
			Source: oauth2.ContextTokenSource(),
			Base:   http.DefaultTransport,
		},
	}
	return client
}

func provideBitbucketClient() *scm.Client {
	client := bitbucket.NewDefault()
	client.Client = &http.Client{
		Transport: &oauth2.Transport{
			Source: oauth2.ContextTokenSource(),
			Base:   http.DefaultTransport,
		},
	}
	return client
}
