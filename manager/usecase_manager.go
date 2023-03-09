 package manager

 import "github.com/rizkyfazri23/dripay/usecase"
 
 type UsecaseManager interface {
	 GatewayUsecase() usecase.GatewayUsecase
	 MemberUsecase() usecase.MemberUsecase
 }
 
 type usecaseManager struct {
	 repoManager RepoManager
 }
 
 func (u *usecaseManager) GatewayUsecase() usecase.GatewayUsecase {
	 return usecase.NewGatewayUsecase(u.repoManager.GatewayRepo())
 }

 func (u *usecaseManager) MemberUsecase() usecase.MemberUsecase {
	return usecase.NewMemberUsecase(u.repoManager.MemberRepo())
}
 
 func NewUsecaseManager(rm RepoManager) UsecaseManager {
	 return &usecaseManager{
		 repoManager: rm,
	 }
 }