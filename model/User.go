package model

import (
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
	"mooblog/utils/errmsg"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(200);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"` //角色0管理员1一般阅读者
}

//对数据库操作

//查询用户是否存在

func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username=?", name).First(&users)
	if users.ID > 0 { //id从0开始，如果大于0就存在了
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

//新增用户

func CreateUser(data *User) int {
	//data.Password = ScryptPw(data.Password) //金数据库直接加密，也可用gorm钩子函数函数名必须BeforeSave如下函数
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE //否则返回200
}

func (u *User) BeforeSave() { //钩子函数，命名不能改。和上面函数data.password效果一样
	u.Password = ScryptPw(u.Password)
}

//查询 用户列表,涉及到分页问题
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageSize - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

//删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//数据库密码加密

func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
