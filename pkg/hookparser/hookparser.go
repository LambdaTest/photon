package hookparser

import (
	"context"
	"errors"
	"net/http"

	"github.com/LambdaTest/photon/pkg/errs"
	"github.com/LambdaTest/photon/pkg/lumber"
	"github.com/LambdaTest/photon/pkg/models"
	"github.com/drone/go-scm/scm"
	"github.com/segmentio/kafka-go"
)

const maxHeaders = 3

type parser struct {
	scmProvider models.SCMProvider
	logger      lumber.Logger
	repoStore   models.RepoStore
}

// New returns a new HookParser.
func New(
	scmProvider models.SCMProvider,
	repoStore models.RepoStore,
	logger lumber.Logger,
) models.HookParser {
	return &parser{
		scmProvider: scmProvider,
		repoStore:   repoStore,
		logger:      logger}
}

func (p *parser) Parse(req *http.Request, driver models.SCMDriver) (scm.Webhook, []kafka.Header, error) {
	scmz, err := p.scmProvider.GetClient(driver)
	if err != nil {
		p.logger.Errorf("failed to get scm client for driver %s, error: %v", driver, err)
		return nil, nil, err
	}
	var repoID, orgID string
	headers := make([]kafka.Header, 0, maxHeaders)

	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()
	// callback function provides the webhook parser with
	// a per-repository secret key used to verify the webhook
	// payload signature for authenticity.
	fn := func(webhook scm.Webhook) (string, error) {
		if webhook == nil {
			// if the incoming webhook is nil
			// we assume it is an unknown event or action.
			return "", errs.ErrUnknownEvent
		}
		repoDetails := webhook.Repository()
		repo, repoErr := p.repoStore.FindActiveByName(ctx, repoDetails.Name, repoDetails.Namespace, driver)
		if repoErr != nil {
			p.logger.Errorf("failed to get repository details for %s/%s with git provider %s, error %v",
				repoDetails.Namespace,
				repoDetails.Name,
				driver,
				repoErr)
			return "", repoErr
		}
		repoID = repo.ID
		orgID = repo.OrgID
		return repo.Secret, nil
	}

	payload, err := scmz.Client.Webhooks.Parse(req, fn)
	if err != nil {
		p.logger.Errorf("failed to parse webhook, error %v, repoID %s, orgID %s", err, repoID, orgID)
		if errors.Is(err, scm.ErrUnknownEvent) {
			err = errs.ErrUnknownEvent
		}
		if errors.Is(err, scm.ErrSignatureInvalid) {
			err = errs.ErrSignatureInvalid
		}
		if errors.Is(err, errs.ErrRowsNotFound) {
			err = errs.ErrRepoNotFound
		}
		return nil, nil, err
	}
	headers = append(headers, kafka.Header{Key: models.RepoIDHeader, Value: []byte(repoID)},
		kafka.Header{Key: models.GitSCMHeader, Value: []byte(driver)})

	switch v := payload.(type) {
	case *scm.PushHook:
		// github sends push hooks when tags and branches are
		// deleted. These hooks should be ignored.
		if v.Commit.Sha == scm.EmptyCommit {
			return nil, nil, errs.ErrEmptyCommit
		}
		headers = append(headers, kafka.Header{Key: models.EventHeader, Value: []byte(models.EventPush)})
	case *scm.PullRequestHook:
		// support for only PR open and Sync events
		if v.Action != scm.ActionOpen && v.Action != scm.ActionSync {
			p.logger.Debugf("pull request action %d not supported for repo: %s, org: %s, repoID %s, orgID %s",
				v.Action, v.Repo.Name, v.Repo.Namespace, repoID, orgID)
			return nil, nil, errs.ErrNotSupported
		}
		headers = append(headers, kafka.Header{Key: models.EventHeader, Value: []byte(models.EventPullRequest)})
	case *scm.PingHook:
		return nil, nil, errs.ErrPingEvent
	default:
		return nil, nil, errs.ErrNotSupported
	}
	return payload, headers, nil
}
