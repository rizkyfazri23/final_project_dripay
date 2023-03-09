package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/dripay/middlewares"
	"github.com/rizkyfazri23/dripay/model/app_error"
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/usecase"
)

type MemberController struct {
	BaseController
	router  *gin.RouterGroup
	usecase usecase.MemberUsecase
}

func NewMemberController(r *gin.RouterGroup, u usecase.MemberUsecase) *MemberController {
	controller := MemberController{
		router:  r,
		usecase: u,
	}
	mmGroup := r.Group("/member")
	mmGroup.Use(middlewares.JwtAuthMiddleware())
	mmGroup.GET("/", controller.GetAll)
	mmGroup.GET("/:id", controller.GetOne)
	mmGroup.PUT("/:id", controller.Edit)
	mmGroup.DELETE("/:id", controller.Remove)

	r.POST("/register", controller.Add)
	r.POST("/login", controller.Login)

	return &controller
}

func (c *MemberController) GetAll(ctx *gin.Context) {
	res, err := c.usecase.GetAll()
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError(""))
		return
	}

	c.Success(ctx, http.StatusOK, "", "Successfully retrieved all member data", res)
}

func (c *MemberController) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("invalid id"))
		return
	}

	res, err := c.usecase.GetOne(id)
	if err != nil {
		c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("member with id %d not found", id)))
		return
	}

	c.Success(ctx, http.StatusOK, "", fmt.Sprintf("Successfully retrieved member with member_id %d", id), res)
}

func (c *MemberController) Add(ctx *gin.Context) {
	var member entity.Member

	if err := ctx.BindJSON(&member); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError(""))
		return
	}

	if member.Username == "" || member.Password == "" || member.Email_Address == "" || member.Contact_Number == "" {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("one or more required fields are missing"))
		return
	}

	res, err := c.usecase.Add(&member)
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("failed to create member"))
		return
	}

	c.Success(ctx, http.StatusCreated, "01", "Successfully created new member", res)
}

func (c *MemberController) Edit(ctx *gin.Context) {
	var member entity.Member

	if err := ctx.BindJSON(&member); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError("invalid request body"))
		return
	}

	res, err := c.usecase.Edit(&member)
	if err != nil {
		c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("member with id %d not found", member.Member_Id)))
		return
	}

	c.Success(ctx, http.StatusOK, "", fmt.Sprintf("Successfully updated member with Member_Id %d", member.Member_Id), res)
}

func (c *MemberController) Remove(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("invalid id"))
		return
	}
	err = c.usecase.Remove(id)
	if err != nil {
		c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("member with id %d not found", id)))
		return
	}

	c.Success(ctx, http.StatusOK, "", fmt.Sprintf("Successfully removed member with Member_Id %d", id), nil)
}

func (c *MemberController) Login(ctx *gin.Context) {
	var input entity.MemberLogin

	if err := ctx.BindJSON(&input); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError("invalid request body"))
		return
	}

	token, err := c.usecase.LoginCheck(input.Username, input.Password)

	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError("Username or Password is Incorrect"))
		return
	}

	c.Success(ctx, http.StatusOK, "", token, nil)
}
