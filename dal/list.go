package dal

import "PocketMusic/dal/model"

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