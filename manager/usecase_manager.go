 package manager

 import "github.com/rizkyfazri23/dripay/usecase"
 
 type UsecaseManager interface {
	 GatewayUsecase() usecase.GatewayUsecase
 }
 
 type usecaseManager struct {
	 repoManager RepoManager
 }
 
 func (u *usecaseManager) GatewayUsecase() usecase.GatewayUsecase {
	 return usecase.NewGatewayUsecase(u.repoManager.GatewayRepo())
 }
 
 func NewUsecaseManager(rm RepoManager) UsecaseManager {
	 return &usecaseManager{
		 repoManager: rm,
	 }
 }