package dal

import (
	"PocketMusic/dal/model"
	"database/sql"
	_ "github.com/Go-SQL-Driver/mysql"
	"time"
)

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}

func GetAllComment(Mid uint) ([]map[string]interface{}, error){
	db, err := sql.Open("mysql", "admin:testdb123456@tcp(119.29.111.64:3306)/testdb?parseTime=true")
	checkErr(err)
	whereParams := make(map[string]interface{})
	whereParams["mid"] = Mid
	whereParams["status"] = 0
	var comments []*model.Comment
	err, _ = db.Query("select uid, content from comments where mid = Mid and status = 0")
	checkErr(err)
	result := make([]map[string]interface{}, 0)
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
