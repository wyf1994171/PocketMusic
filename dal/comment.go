package dal

import (
	"PocketMusic/dal/model"
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/mysql"
	"time"
)

func GetAllComment(Mid string) ([]map[string]interface{}, error){
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

func CreateComment(Uid string,Mid string, content string) error {
	comment := &model.Comment{
		RecordMeta: model.RecordMeta{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UID:   Uid,
		MID:   Mid,
		Content:  content,
		Status: 0,
	}
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	db, err := sql.Open("mysql", "admin:testdb123456@tcp(119.29.111.64:3306)/testdb?parseTime=true")
	checkErr(err)
	var cid int
	rows, err := db.Query("SELECT cid FROM comments")//获得comments表中primarykey：cid的最大值
	for rows.Next() {
		err = rows.Scan(&cid)
		checkErr(err)
	}
	cid = cid + 1
	//插入数据
	stmt, err := db.Prepare("INSERT INTO comments SET cid=?,mid=?,uid=?,content=?,create_at=?,update_at=?")//准备插入
	checkErr(err)
	t := time.Now()
	res, err := stmt.Exec(cid, comment.MID, comment.UID, comment.Content,t,t)//执行插入
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//err := db.Table("comments").Save(comment).Error
	return err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
