package dal

import (
	"PocketMusic/dal/model"
)

func GetLikeNum(Uid uint) (int,error) {
	whereParam := make(map[string]interface{})
	whereParam["uid"] = Uid
	whereParam["status"] = 0
	condition := CombineCondition(whereParam)
	count := 0
	err := db.Table("likes").Where(condition).Count(&count).Error
	return count, err
}

func AddLike(Uid string, Mid uint) error {
	like := &model.Like{
		Uid: Uid,
		Mid: Mid,
		Status: 0,
	}
	err := db.Save(like).Error
	return err
}

func DeleteLike(Uid string, Mid uint) error {
	whereParams := make(map[string]interface{})
	whereParams["uid"] = Uid
	whereParams["mid"] = Mid
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	updateParams := make(map[string]interface{})
	updateParams["status"] = 0
	return db.Model(&model.Like{}).Where(condition).Updates(updateParams).Error
}