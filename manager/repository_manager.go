/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Sat Mar 04 2023 9:48:51 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

 package manager

 import "github.com/rizkyfazri23/dripay/repository"
 
 type RepoManager interface {
	 GatewayRepo() repository.GatewayRepo
 }
 
 type repositoryManager struct {
	 infraManager InfraManager
 }
 
 func (r *repositoryManager) GatewayRepo() repository.GatewayRepo {
	 return repository.NewGatewayRepository(r.infraManager.DbConn())
 }
 
 func NewRepoManager(manager InfraManager) RepoManager {
	 return &repositoryManager{
		 infraManager: manager,
	 }
 }