package dal

import (
	"PocketMusic/dal/model"
	"time"
)

func GetAllComment(Mid uint) ([]map[string]interface{}, error){
	whereParams := make(map[string]interface{})
	whereParams["mid"] = Mid
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	var comments []*model.Comment
	err := db.Table("comments").Where(condition).Order("updated_at desc").Error
	result := make([]map[string]interface{}, 0)
	for key := range comments{
		newComment := make(map[string]interface{})
		newComment["mid"] = comments[key].MID
		newComment["uid"] = comments[key].UID
		newComment["content"] = comments[key].Content
		newComment["updated_at"] = comments[key].UpdatedAt
		result = append(result,newComment)
	}
	return result, err
}

func CreateComment(Uid,Mid uint, content string) (uint, error)  {
	comment := &model.Comment{
		RecordMeta: model.RecordMeta{
			CreatedAt: time.Now().Local(),
			UpdatedAt: time.Now().Local(),
		},
		Content:  content,
		UID:   Uid,
		MID:   Mid,
		Status: 0,
	}
	err := db.Save(comment).Error
	return comment.UID, err
}
