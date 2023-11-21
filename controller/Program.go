package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ljcbaby/plan/model"
)

type ProgramController struct{}

func (c *ProgramController) CreateProgram(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, model.Response{
		Code: -1,
		Msg:  "Under construction.",
	})
}

func (c *ProgramController) DeleteProgram(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, model.Response{
		Code: -1,
		Msg:  "Under construction.",
	})
}

func (c *ProgramController) UpdateProgram(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, model.Response{
		Code: -1,
		Msg:  "Under construction.",
	})
}

func (c *ProgramController) GetProgram(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, model.Response{
		Code: -1,
		Msg:  "Under construction.",
	})
}

func (c *ProgramController) CalculateProgram(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, model.Response{
		Code: -1,
		Msg:  "Under construction.",
	})
}

func (c *ProgramController) GetProgramList(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, model.Response{
		Code: -1,
		Msg:  "Under construction.",
	})
}

func (c *ProgramController) GetProgramFile(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, model.Response{
		Code: -1,
		Msg:  "Under construction.",
	})
}
