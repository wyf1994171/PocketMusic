package dal

func GetLikeNum(Uid uint) (int,error) {
	whereParam := make(map[string]interface{})
	whereParam["uid"] = Uid
	condition := CombineCondition(whereParam)
	count := 0
	err := db.Table("likes").Where(condition).Count(&count).Error
	return count, err
}