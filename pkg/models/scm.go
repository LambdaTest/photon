package models

import (
	"github.com/LambdaTest/photon/pkg/errs"
	"github.com/drone/go-scm/scm"
)

// SCMDriver identifies source code management driver.
type SCMDriver string

// SCMDriver values.
const (
	DriverGithub    SCMDriver = "github"
	DriverGitlab    SCMDriver = "gitlab"
	DriverBitbucket SCMDriver = "bitbucket"
	// DriverGithubSelfHosted    SCMDriver = "github_self_hosted"
	// DriverGitlabSelfHosted    SCMDriver = "gitlab_self_hosted"
	// DriverBitbucketSelfHosted SCMDriver = "bitbucket_self_hosted"
)

// VerifyDriver verifies if the SCMDriver is valid.
func (d SCMDriver) VerifyDriver() error {
	switch d {
	case DriverGithub:
	// case DriverGithubSelfHosted:
	case DriverGitlab:
	// case DriverGitlabSelfHosted:
	case DriverBitbucket:
	default:
		return errs.ErrInvalidDriver
	}
	return nil
}

// SCMProvider return a new git scm client based on the driver.
type SCMProvider interface {
	//  GetClient returns a new git scm client
	GetClient(scmClientID SCMDriver) (*SCM, error)
}

// SCM is wrapper around scm.Client
type SCM struct {
	Client     *scm.Client
	SelfHosted bool
}
