 package controller

 import (
	 "fmt"
	 "net/http"
	 "strconv"
 
	 "github.com/rizkyfazri23/dripay/model/app_error"
	 "github.com/rizkyfazri23/dripay/model/entity"
	 "github.com/rizkyfazri23/dripay/usecase"
	 "github.com/gin-gonic/gin"
 )
 
 type GatewayController struct {
	 BaseController
	 router  *gin.RouterGroup
	 usecase usecase.GatewayUsecase
 }
 
 func NewGatewayController(r *gin.RouterGroup, u usecase.GatewayUsecase) *GatewayController {
	 controller := GatewayController{
		 router:  r,
		 usecase: u,
	 }
	//  r.Group("/gateway")
	 // define routes
	 r.GET("/gateway", controller.GetAll)
	 r.GET("/gateway/:id", controller.GetOne)
	 r.POST("/gateway", controller.Add)
	 r.PUT("/gateway/:id", controller.Edit)
	 r.DELETE("/gateway/:id", controller.Remove)
 
	 return &controller
 }
 
 func (c *GatewayController) GetAll(ctx *gin.Context) {
	 res, err := c.usecase.GetAll()
	 if err != nil {
		 c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError(""))
		 return
	 }
 
	 c.Success(ctx, http.StatusOK, "", "Successfully retrieved all gateway data", res)
 }
 
 func (c *GatewayController) GetOne(ctx *gin.Context) {
	 id, err := strconv.Atoi(ctx.Param("id"))
	 if err != nil {
		 c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("invalid id"))
		 return
	 }
 
	 res, err := c.usecase.GetOne(id)
	 if err != nil {
		 c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("gateway with id %d not found", id)))
		 return
	 }
 
	 c.Success(ctx, http.StatusOK, "", fmt.Sprintf("Successfully retrieved gateway with Gateway_Id %d", id), res)
 }
 
 func (c *GatewayController) Add(ctx *gin.Context) {
	 var gateway entity.Gateway
 
	 if err := ctx.BindJSON(&gateway); err != nil {
		 c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError(""))
		 return
	 }
 
	 if gateway.Gateway_Name == "" {
		 c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("one or more required fields are missing"))
		 return
	 }
 
	 res, err := c.usecase.Add(&gateway)
	 if err != nil {
		 c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("failed to create gateway"))
		 return
	 }
 
	 c.Success(ctx, http.StatusCreated, "01", "Successfully created new gateway", res)
 }
 
 func (c *GatewayController) Edit(ctx *gin.Context) {
	 var gateway entity.Gateway
 
	 if err := ctx.BindJSON(&gateway); err != nil {
		 c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError("invalid request body"))
		 return
	 }
 
	 res, err := c.usecase.Edit(&gateway)
	 if err != nil {
		 c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("gateway with id %d not found", gateway.Gateway_Id)))
		 return
	 }
 
	 c.Success(ctx, http.StatusOK, "", fmt.Sprintf("Successfully updated gateway with Gateway_Id %d", gateway.Gateway_Id), res)
 }
 
 func (c *GatewayController) Remove(ctx *gin.Context) {
	 id, err := strconv.Atoi(ctx.Param("id"))
	 if err != nil {
		 c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("invalid id"))
		 return
	 }
	 err = c.usecase.Remove(id)
	 if err != nil {
		 c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("gateway with id %d not found", id)))
		 return
	 }
 
	 c.Success(ctx, http.StatusOK, "", fmt.Sprintf("Successfully removed gateway with Gateway_Id %d", id), nil)
 }
 