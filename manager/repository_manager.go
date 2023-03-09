
 package manager

 import "github.com/rizkyfazri23/dripay/repository"
 
 type RepoManager interface {
	 GatewayRepo() repository.GatewayRepo
	 MemberRepo() repository.MemberRepo
 }
 
 type repositoryManager struct {
	 infraManager InfraManager
 }
 
 func (r *repositoryManager) GatewayRepo() repository.GatewayRepo {
	 return repository.NewGatewayRepository(r.infraManager.DbConn())
 }

 func (r *repositoryManager) MemberRepo() repository.MemberRepo {
	return repository.NewMemberRepository(r.infraManager.DbConn())
}
 
 func NewRepoManager(manager InfraManager) RepoManager {
	 return &repositoryManager{
		 infraManager: manager,
	 }
 }