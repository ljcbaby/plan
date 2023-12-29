package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ljcbaby/plan/model"
	"github.com/ljcbaby/plan/service"
)

type ProgramController struct{}

func (c *ProgramController) CreateProgram(ctx *gin.Context) {
	var program model.Program
	if err := ctx.ShouldBindJSON(&program); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  err.Error(),
		})
		return
	}

	if program.Name == nil || program.Major == nil || program.Department == nil || program.Grade == nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  "Required fields cannot be empty.",
		})
		return
	}

	ps := service.ProgramService{}

	if program.DependencyID != nil {
		var p model.Program
		if err := ps.GetProgram(*program.DependencyID, &p); err != nil {
			ctx.JSON(http.StatusBadRequest, model.Response{
				Code: 1001,
				Msg:  "Dependency program does not exist.",
			})
			return
		}
	}

	id, err := ps.CreateProgram(&program)
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}

	fmt.Println(id)

	ctx.JSON(http.StatusCreated, Success(program))
}

func (c *ProgramController) DeleteProgram(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  "Invalid ID.",
		})
		return
	}

	ps := &service.ProgramService{}
	err = ps.DeleteProgram(uint(id))
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, Success(nil))
}

func (c *ProgramController) UpdateProgram(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, model.Response{
		Code: -1,
		Msg:  "Under construction.",
	})
}

func (c *ProgramController) GetProgram(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  "Invalid ID.",
		})
		return
	}

	isContentNeeded, err := strconv.Atoi(ctx.DefaultQuery("content", "1"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  "Invalid parameter.",
		})
		return
	}

	var program model.Program
	ps := &service.ProgramService{}
	if isContentNeeded == 0 {
		err = ps.GetProgramWithNoContent(uint(id), &program)
	} else {
		err = ps.GetProgramWithContent(uint(id), &program)
	}
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, model.Response{
				Code: 1001,
				Msg:  "Program not found.",
			})
			return
		}
		returnMySQLError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, Success(program))
}

func (c *ProgramController) CalculateProgram(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, model.Response{
		Code: -1,
		Msg:  "Under construction.",
	})
}

func (c *ProgramController) GetProgramList(ctx *gin.Context) {
	var programs []model.Program
	ps := &service.ProgramService{}
	err := ps.GetProgramList(&programs)
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, Success(programs))
}

func (c *ProgramController) GetProgramFile(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, model.Response{
		Code: -1,
		Msg:  "Under construction.",
	})
}
