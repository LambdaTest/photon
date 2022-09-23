package cmd

import (
	"github.com/LambdaTest/photon/config"
	"github.com/LambdaTest/photon/pkg/db"
	"github.com/LambdaTest/photon/pkg/lumber"
	"github.com/LambdaTest/photon/pkg/models"
	"github.com/LambdaTest/photon/pkg/store/repo"
	"github.com/google/wire"
)

var storeSet = wire.NewSet(
	provideDatabase,
	provideRepoStore,
)

// provideDatabase is a Wire provider function that provides a
// database connection, configured from the environment.
func provideDatabase(cfg *config.Config, logger lumber.Logger) (models.DB, error) {
	return db.Connect(cfg, logger)
}

// provideRepoStore is a Wire provider function that provides a
// repo datastore.
func provideRepoStore(database models.DB, logger lumber.Logger) models.RepoStore {
	return repo.New(database, logger)
}
