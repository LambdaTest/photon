package cmd

import (
	"github.com/LambdaTest/photon/pkg/hookparser"
	"github.com/LambdaTest/photon/pkg/lumber"
	"github.com/LambdaTest/photon/pkg/models"
	"github.com/LambdaTest/photon/pkg/scm"
	"github.com/google/wire"
)

var scmClientSet = wire.NewSet(provideSCM, provideParser)

func provideSCM(logger lumber.Logger) models.SCMProvider {
	return scm.New(logger)
}

func provideParser(repoStore models.RepoStore, scmProvider models.SCMProvider, logger lumber.Logger) models.HookParser {
	return hookparser.New(scmProvider, repoStore, logger)
}
