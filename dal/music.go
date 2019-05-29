package dal

import (
	"PocketMusic/dal/model"
)

func GetMusicInfoById(Mid uint) (map[string]interface{},error) {
	whereParams := make(map[string]interface{})
	whereParams["mid"] = Mid
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	var musics []*model.Music
	err := db.Where(condition).Find(&musics).Error
	m := make(map[string]interface{})
	m["name"] = musics[0].Mname
	m["id"] = musics[0].Mid
	m["singer"] = musics[0].Singer
	return m,err
}