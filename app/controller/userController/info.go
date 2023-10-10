package userController

import (
	"TeamRegistrationSystem-Back/app/apiExpection"
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/userService"
	"TeamRegistrationSystem-Back/app/utils"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type uinfo struct {
	Name     string `json:"name"  binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Motto    string `json:"motto" binding:"required"`
}

func Updateinfodata(c *gin.Context) {
	//获取用户身份token
	n, er := c.Get("UserID")
	if !er {
		utils.JsonErrorResponse(c, 200, "token获取失败")
		return
	}
	v, _ := n.(int)
	terr :=userService.CheckUserExistByUID(v)
	if terr !=nil{
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	var data uinfo
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	//判断手机号是否符合格式
	name_sample:=regexp.MustCompile(`^[\w\s\p{Han}]{1,25}$`)
	if !name_sample.MatchString(data.Name) {
		utils.JsonErrorResponse(c, 200, "用户名格式错误")
		return
	}
	//判断邮箱是否符合格式
	email_sample:=regexp.MustCompile(`^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$`)
	if !email_sample.MatchString(data.Email) {
		utils.JsonErrorResponse(c, 200, "邮箱格式错误")
		return
	}
	//判断电话是否符合格式
	phone_sample:=regexp.MustCompile(`^1[3456789]\d{9}$`)
	if !phone_sample.MatchString(data.Phone) {
		utils.JsonErrorResponse(c, 200, "电话格式错误")
		return
	}
	//查询手机号是否重复
	err = userService.CheckUserinfoExistByPhone(data.Phone,v)
	if err == nil {
		utils.JsonErrorResponse(c, 200, "手机号已存在")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//查询邮箱是否重复注册
	err = userService.CheckUserinfoExistByEmail(data.Email,v)
	if err == nil {
		utils.JsonErrorResponse(c, 200, "邮箱已存在")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//查询用户名是否存在
	err = userService.CheckUserinfoExistByName(data.Name,v)
	if err == nil {
		utils.JsonErrorResponse(c, 200, "用户名已存在")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	var user *models.Userinfo
	user, err = userService.CheckUserinfoExistByUserid(v)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	err = userService.Updatainfo(models.Userinfo{
		ID:       user.ID,
		UserID:      v,
		Name:     data.Name,
		Phone:    data.Phone,
		Email:    data.Email,
		Birthday: data.Birthday,
		Address:  data.Address,
		Motto:    data.Motto,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func AvatarUpload(c *gin.Context) {
	//获取用户身份token
	n, er := c.Get("UserID")
	if !er {
		utils.JsonErrorResponse(c, 200, "token获取失败")
		return
	}
	v, _ := n.(int)
	terr :=userService.CheckUserExistByUID(v)
	if terr !=nil{
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	//保存图片文件
	file, err := c.FormFile("image")
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	// 创建临时目录
	tempDir, err := os.MkdirTemp("", "tempdir")
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	defer os.RemoveAll(tempDir) // 在处理完之后删除临时目录及其中的文件
	// 在临时目录中创建临时文件
	tempFile := filepath.Join(tempDir, file.Filename)
	f, err := os.Create(tempFile)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	defer f.Close()

	// 将上传的文件保存到临时文件中
	src, err := file.Open()
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	defer src.Close()

	_, err = io.Copy(f, src)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	// 判断文件的MIME类型是否为图片
	mime, err := mimetype.DetectFile(tempFile)
	if err != nil || !strings.HasPrefix(mime.String(), "image/") {
    	utils.JsonErrorResponse(c, 200, "仅允许上传图片文件")
    	return
	}
	//继续保存图片
	filename := uuid.New().String() + filepath.Ext(file.Filename)
	dst := "./uploads/" + filename
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	url := c.Request.Host + "/uploads/" + filename
	var user *models.Userinfo
	user, err = userService.CheckUserinfoExistByUserid(v)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	err = userService.UpdataAvatar(models.Userinfo{
		ID:     user.ID,
		Avatar: url,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"avatar":url,
	})
}


type GetInfoData struct{
	Name string `form:"name" binding:"required"`
}

func GetUserInfo(c *gin.Context) {
	var data GetInfoData
	er :=c.ShouldBindQuery(&data)
	if er != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	var user *models.User
	user,er= userService.GetUserByName(data.Name)
	if er != nil {
		utils.JsonErrorResponse(c,200,"查无此人")
		fmt.Println(data.Name)
		fmt.Println(1)
		return
	}
	var uinfoList []models.Userinfo
	uinfoList, err := userService.GetInfoList(user.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200, "查无此人")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user_info": uinfoList[0],
	})
}
