package dal

import (
	"PocketMusic/dal/model"
)

func GetLikeNum(Uid string) (int,error) {
	whereParam := make(map[string]interface{})
	whereParam["uid"] = Uid
	whereParam["status"] = 0
	condition := CombineCondition(whereParam)
	count := 0
	err := db.Table("likes").Where(condition).Count(&count).Error
	return count, err
}

func AddLike(Uid string, Mid uint) error {
	var count uint
	err:=db.Table("likes").Where("uid = ? and mid = ?",Uid,Mid).Count(&count).Error
	//println(count)
	if err==nil {
		if count == 0 {
			like := &model.Like{
				Uid:    Uid,
				Mid:    Mid,
				Status: 0,
			}
			err := db.Save(like).Error
			return err
		} else {
			_,err := db.DB().Exec("update likes set status = 0 where uid = ? and mid = ?",Uid,Mid)
			return err
		}
	}
	return err
}

func DeleteLike(Uid string, Mid uint) error {
	whereParams := make(map[string]interface{})
	whereParams["uid"] = Uid
	whereParams["mid"] = Mid
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	updateParams := make(map[string]interface{})
	updateParams["status"] = 1
	return db.Model(&model.Like{}).Where(condition).Updates(updateParams).Error
}

func GetLikeList(Uid string) ([]uint,error) {
	whereParams := make(map[string]interface{})
	whereParams["uid"] = Uid
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	var likes []*model.Like
	err := db.Where(condition).Find(&likes).Error
	var res []uint
	for key := range likes {
		res = append(res,likes[key].Mid)
	}
	return res,err
}