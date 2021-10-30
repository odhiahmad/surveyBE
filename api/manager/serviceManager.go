package manager

import (
	"survey/api/connect"
	"survey/api/usecase"
)

type ServiceManager interface {
	UserUseCase() usecase.IUserUseCase
	RumahUseCase() usecase.IRumahUseCase
}

type serviceManager struct {
	repo RepoManager
}

func (sm *serviceManager) UserUseCase() usecase.IUserUseCase {
	return usecase.NewUserUseCase(sm.repo.UserRepo())
}

func (sm *serviceManager) RumahUseCase() usecase.IRumahUseCase {
	return usecase.NewRumahUseCase(sm.repo.RumahRepo())
}

func NewServiceManager(connect connect.Connect) ServiceManager {
	return &serviceManager{repo: NewRepoManager(connect)}
}
