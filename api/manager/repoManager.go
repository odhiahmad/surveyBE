package manager

import (
	"survey/api/connect"
	"survey/api/repository"
)

type RepoManager interface {
	//UserAuthRepo() repository.IAuthRepository
	UserRepo() repository.IUserRepository
	RumahRepo() repository.IRumahRepository
}

type repoManager struct {
	connect connect.Connect
}

func (rm *repoManager) UserRepo() repository.IUserRepository {
	return repository.NewUserRepository(rm.connect.SqlDb())
}
func (rm *repoManager) RumahRepo() repository.IRumahRepository {
	return repository.NewRumahRepository(rm.connect.SqlDb())
}

func NewRepoManager(connect connect.Connect) RepoManager {
	return &repoManager{connect}
}
