package model //文章分类
import (
	"github.com/jinzhu/gorm"
	"mooblog/utils/errmsg"
)

type Category struct {
	//gorm.Model
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type varchar(20);not null" json:"name"`
}

//查询分类是否存在

func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name=?", name).First(&cate)
	if cate.ID > 0 { //id从0开始，如果大于0就存在了
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCSE
}

//新增分类

func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE //否则返回200
}

//todo 查询分类下的所有文章

//查询分类列表
func GetCate(pageSize int, pageNum int) []Category {
	var cate []Category
	err = db.Limit(pageSize).Offset((pageSize - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

//编辑分类信息
func EditCate(id int, data *Category) int { //这里用map方式，详细参见gorm更新操作
	var cate []Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = db.Model(&cate).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCSE
}

//删除分类
func DeleteCate(id int) int {
	var cate []Category
	err = db.Where("id=?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
