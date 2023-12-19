package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ljcbaby/plan/model"
	"github.com/ljcbaby/plan/service"
)

type TagController struct{}

func (c *TagController) CreateTag(ctx *gin.Context) {
	var tag model.Tag
	if err := ctx.ShouldBindJSON(&tag); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  err.Error(),
		})
		return
	}

	if tag.Tag == nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  "Required fields cannot be empty.",
		})
		return
	}

	ts := service.TagService{}

	if err := ts.CreateTag(&tag); err != nil {
		returnMySQLError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, Success(tag.ID))
}

func (c *TagController) GetTagList(ctx *gin.Context) {
	ts := service.TagService{}

	tags, err := ts.GetTagList()
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, Success(tags))
}

func (c *TagController) DeleteTag(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  err.Error(),
		})
		return
	}

	ts := service.TagService{}

	if err := ts.DeleteTag(id); err != nil {
		returnMySQLError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, Success(nil))
}

func (c *TagController) UpdateTag(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  err.Error(),
		})
		return
	}

	var tag model.Tag
	if err := ctx.ShouldBindJSON(&tag); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  err.Error(),
		})
		return
	}

	if tag.Tag == nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code: 1001,
			Msg:  "Required fields cannot be empty.",
		})
		return
	}

	tag.ID = &id

	ts := service.TagService{}

	if err := ts.UpdateTag(&tag); err != nil {
		returnMySQLError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, Success(nil))
}
