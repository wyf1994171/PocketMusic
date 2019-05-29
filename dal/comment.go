package dal

import (
	"PocketMusic/dal/model"
	_ "github.com/Go-SQL-Driver/mysql"
	"time"
)

func GetAllComment(Mid uint) ([]map[string]interface{}, error){
	whereParams := make(map[string]interface{})
	whereParams["mid"] = Mid
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	var comments []*model.Comment
	result := make([]map[string]interface{}, 0)
	err := db.Where(condition).Find(&comments).Error
	for key := range comments{
		newComment := make(map[string]interface{})
		newComment["uid"] = comments[key].UID
		newComment["content"] = comments[key].Content
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
		UID:   Uid,
		MID:   Mid,
		Content:  content,
		Status: 0,
	}
	err := db.Save(comment).Error
	return comment.UID, err
}
