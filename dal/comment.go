package dal

import (
	"PocketMusic/dal/model"
	"time"
)

func AddComment(uid,mid uint, content string) (uint,error) {
	comment := &model.Comment{
		RecordMeta: model.RecordMeta{
			CreatedAt: time.Now().Local(),
			UpdatedAt: time.Now().Local(),
		},
		Content:  content,
		Uid:   uid,
		Mid:   mid,
		Status: 0,
	}
	err := db.Save(comment).Error
	return comment.ID,err
}

func GetComment(mid uint) ([]map[string]interface{},error) {
	whereParams := make(map[string]interface{})
	whereParams["mid"] = mid
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	var comments []*model.Comment
	err := db.Where(condition).Order("updated_at desc").Find(comments).Error
	result := make([]map[string]interface{},0)
	for key := range comments{
		newComment := make(map[string]interface{})
		newComment["content"] = comments[key].Content
		newComment["id"] = comments[key].ID
		newComment["uid"] = comments[key].Uid
		newComment["updated_at"] = comments[key].UpdatedAt
		result = append(result,newComment)
	}
	return result,err
}

func DeleteComment(id uint) error {
	whereParams := make(map[string]interface{})
	whereParams["id"] = id
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	updateParams := make(map[string]interface{})
	updateParams["status"] = 1
	return db.Model(model.Comment{}).Where(condition).Updates(updateParams).Error
}