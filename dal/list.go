package dal

import (
	"PocketMusic/dal/model"
	"time"
)

func GetLists(uid uint)([]map[string]interface{},error) {
	whereParams := make(map[string]interface{})
	whereParams["uid"] = uid
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	var lists []*model.List
	err := db.Where(condition).Find(&lists).Error
	result := make([]map[string]interface{},0)
	for key := range lists {
		res := make(map[string]interface{})
		res["id"] = lists[key].ID
		res["name"] = lists[key].Name
		res["num"] = lists[key].Num
		res["cover_path"] = lists[key].CoverPath
		result = append(result,res)
	}
	return result,err
}

func AddList(uid uint,name string) (uint,error) {
	list := &model.List{
		RecordMeta :model.RecordMeta{
			CreatedAt: time.Now().Local(),
			UpdatedAt: time.Now().Local(),
		},
		Name: name,
		Status: 0,
		Num: 0,
		UID: uid,
	}
	err := db.Save(list).Error
	return list.ID, err
}

func AddListCoverPath(coverPath string,id uint) error {
	whereParams := make(map[string]interface{})
	whereParams["id"] = id
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	updateParams := make(map[string]interface{})
	updateParams["cover_path"] = coverPath
	return db.Model(&model.List{}).Where(condition).Updates(updateParams).Error
}

func AddListSong(lid,mid uint) error {
	listSong := &model.ListSong{
		RecordMeta: model.RecordMeta{
			CreatedAt: time.Now().Local(),
			UpdatedAt: time.Now().Local(),
		},
		Lid: lid,
		Mid: mid,
		Status: 0,
	}
	err := db.Save(listSong).Error
	return err
}

func DeleteListSong(lid,mid uint) error {
	whereParams := make(map[string]interface{})
	whereParams["lid"] = lid
	whereParams["mid"] = mid
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	updateParams := make(map[string]interface{})
	updateParams["status"] = 1
	updateParams["updated_at"] = time.Now().Local()
	return db.Model(&model.ListSong{}).Where(condition).Updates(updateParams).Error
}