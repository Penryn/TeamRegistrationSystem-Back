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
	result := database.DB.Omit("team_id").Create(&user)
	return result.Error
}

func Updatainfo(info models.Userinfo)error{
	r1 :=database.DB.Model(&info).Updates(models.Userinfo{Name: info.Name,Phone: info.Phone,Email: info.Email,Birthday:info.Birthday,Address: info.Address,Motto: info.Motto})
	r2:=database.DB.Model(&models.User{UserID: info.UserID}).Updates(models.User{Name: info.Name,Phone: info.Phone,Email: info.Email})
	if r1.Error!=nil{
		return r1.Error
	}else if r2.Error!=nil{
		return r2.Error
	}else{
		return nil
	}
}

func UpdataAvatar(info models.Userinfo)error{
	result :=database.DB.Model(&info).Update("avatar",info.Avatar)
	return result.Error
}

func GetInfoList(userID int) ([]models.Userinfo, error) {
	result := database.DB.Where("user_id=?", userID).Find(&models.Userinfo{})
	if result.Error != nil {
		return nil, result.Error
	}
	var infoList []models.Userinfo
	result = database.DB.Where("user_id=?", userID).Find(&infoList)
	if result.Error != nil {
		return nil, result.Error
	}
	return infoList, nil
}

func GetUserByName(nAme string)(*models.User,error){
	var user models.User
	result :=database.DB.Where("name = ?",nAme).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func Compare(t1 string, t2 string) bool {
	return t1 == t2
}

func UpdataPassword(user models.User)error{
	result :=database.DB.Model(&user).Update("password",user.Password)
	return result.Error
}