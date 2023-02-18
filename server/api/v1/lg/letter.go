package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LetterApi struct {
}

var letterService = service.ServiceGroupApp.LgServiceGroup.LetterService

func (letterApi *LetterApi) CreateLetter(c *gin.Context) {
	var letter lg.Letter
	err := c.ShouldBindJSON(&letter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := letterService.CreateLetter(letter); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (letterApi *LetterApi) DeleteLetter(c *gin.Context) {
	var letter lg.Letter
	err := c.ShouldBindJSON(&letter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := letterService.DeleteLetter(letter); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (letterApi *LetterApi) DeleteLetterByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := letterService.DeleteLetterByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (letterApi *LetterApi) UpdateLetter(c *gin.Context) {
	var letter lg.Letter
	err := c.ShouldBindJSON(&letter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := letterService.UpdateLetter(letter); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (letterApi *LetterApi) FindLetter(c *gin.Context) {
	var letter lg.Letter
	err := c.ShouldBindQuery(&letter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reletter, err := letterService.GetLetter(letter.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reletter": reletter}, c)
	}
}

func (letterApi *LetterApi) GetLetterList(c *gin.Context) {
	var pageInfo lgReq.LetterSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := letterService.GetLetterInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
