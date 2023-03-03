package v1

import (
	"github.com/gin-gonic/gin"
	"mooblog/model"
	"mooblog/utils/errmsg"
	"net/http"
	"strconv"
)

//添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CreateArt(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//TODO 查询分类下的所有文章
func GetArtInfo(c *gin.Context){
//todo 查询单个文章信息
func GetArtInfo(c *gin.Context){

}



//查询文章列表
func GetArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := model.GetArt(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑文章
func EditArt(c *gin.Context) {
	//编辑后的用户名不能和其他人同名,需要
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.EditArt(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除分类
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id")) //因User.go文件内删除传的是int，所以转化下
	code = model.DeleteArt(id)
	c.JSON(http.StatusOK, gin.H{
		"satus":   code,
		"message": errmsg.GetErrMsg(code),
	})
}
