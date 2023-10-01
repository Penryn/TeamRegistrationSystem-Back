package userService

import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/config/database"
)


func CheckUserExistByPhone(pHone string) error {
	result := database.DB.Where("phone = ?", pHone).First(&models.User{})
	return result.Error
}

func CheckUserinfoExistByPhone(pHone string) error {
	result := database.DB.Where("phone = ?", pHone).First(&models.Userinfo{})
	return result.Error
}

func CheckUserExistByEmail(eMail string) error {
	result := database.DB.Where("email = ?", eMail).First(&models.User{})
	return result.Error
}

func CheckUserinfoExistByEmail(eMail string) error {
	result := database.DB.Where("email = ?", eMail).First(&models.Userinfo{})
	return result.Error
}

func CheckUserExistByName(naMe string) error {
	result := database.DB.Where("name = ?", naMe).First(&models.User{})
	return result.Error
}

func CheckUserinfoExistByName(naMe string) error {
	result := database.DB.Where("name = ?", naMe).First(&models.Userinfo{})
	return result.Error
}

func CheckUserinfoExistByUserid(uid int) (*models.Userinfo, error) {
	var user models.Userinfo
	result := database.DB.Where("user_id = ?", uid).First(&user)
	if result.Error !=nil{
		return nil,result.Error
	}
	return &user,nil
}

func CheckUserExistByAccount(account string) error {
	result := database.DB.Where("phone = ? or email =? or name=?", account,account,account).First(&models.User{})
	return result.Error
}

func GetUserByAccount(aCCount string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("phone = ? or email =? or name=?", aCCount,aCCount,aCCount).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func Register(user models.User) error {
	result := database.DB.Create(&user)
	return result.Error
}

func Updatainfo(info models.Userinfo)error{
	result :=database.DB.Model(&info).Omit("Avatar").Updates(models.Userinfo{Name: info.Name,Phone: info.Phone,Email: info.Email,Birthday:info.Birthday,Address: info.Address,Motto: info.Motto})
	return result.Error
}