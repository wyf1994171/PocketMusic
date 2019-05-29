package dal

import (
	"PocketMusic/dal/model"
	"time"
)

func AddUser(Uid,Nickname,AvatarPath string) error {
	user := &model.User{
		RecordMeta: model.RecordMeta{
			CreatedAt: time.Now().Local(),
			UpdatedAt: time.Now().Local(),
		},
		Uid: Uid,
		Nickname: Nickname,
		AvatarPath: AvatarPath,
	}
	return db.Save(user).Error
}

func GetUserInfoByUid(Uid string) (string,error) {
	whereParams := make(map[string]interface{})
	whereParams["uid"] = Uid
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	var users []*model.User
	err := db.Where(condition).Find(&users).Error
	return users[0].Nickname,err
}

func GetIfUser(Uid string) (bool,error) {
	whereParams := make(map[string]interface{})
	whereParams["uid"] = Uid
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	var users []*model.User
	err := db.Where(condition).Find(&users).Error
	return users != nil , err
}