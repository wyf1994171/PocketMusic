package dal

import (
	"PocketMusic/dal/model"
	"time"
)

func AddRecent(Uid string, Mid uint) error {
	recent := &model.Recent{
		RecordMeta: model.RecordMeta{
			CreatedAt: time.Now().Local(),
			UpdatedAt: time.Now().Local(),
		},
		Uid:    Uid,
		Mid:    Mid,
		Status: 0,
	}
	return db.Save(recent).Error
}

func GetRecent(Uid string) ([]uint,error) {
    whereParams := make(map[string]interface{})
    whereParams["uid"] = Uid
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	var recents []*model.Recent
	err := db.Where(condition).Order("order by updated_at desc").Find(recents).Error
	results := make([]uint,0)
	for key := range recents {
		results = append(results,recents[key].Mid)
	}
	return results, err
}
