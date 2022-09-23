package repo

import (
	"context"

	"github.com/LambdaTest/photon/pkg/lumber"
	"github.com/LambdaTest/photon/pkg/models"
	"github.com/jmoiron/sqlx"
)

type repoStore struct {
	db     models.DB
	logger lumber.Logger
}

// New returns a new repoStore
func New(db models.DB, logger lumber.Logger) models.RepoStore {
	return &repoStore{db: db, logger: logger}
}

func (r *repoStore) FindActiveByName(ctx context.Context,
	repoName,
	orgName string,
	gitProvider models.SCMDriver) (*models.Repository, error) {
	repo := new(models.Repository)
	return repo, r.db.Execute(func(db *sqlx.DB) error {
		rows := db.QueryRowxContext(ctx, selectActiveQuery, repoName, orgName, gitProvider)
		if err := rows.StructScan(repo); err != nil {
			return err
		}
		return nil
	})
}

const selectActiveQuery = `SELECT
r.id,
r.org_id,
r.strict,
r.name,
r.link,
r.tas_file_name,
r.webhook_secret,
r.admin_id,
r.mask
FROM
repositories r
JOIN organizations o ON
o.id = r.org_id
WHERE
r.active = 1
AND r.name =?
AND o.name =?
AND o.git_provider =?`
