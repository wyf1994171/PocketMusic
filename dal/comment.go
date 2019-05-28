package dal


import "time"

func GetAllComment(Mid uint) (int, error){
	whereParam := make(map[string]interface{})
	whereParam["mid"] = Mid
	condition := CombineCondition(whereParam)
	count := 0
	err := db.Table("comments").Where(condition).Count(&count).Error
	return count, err
}

func CreateComment(Mid uint) (error)  {
	count:=0
	err:=db.Table("comments").Where("mid =?",Mid).Count(&count).Error
	if err!=nil{
		return err
	}else if count>0{
		_,err:=db.DB().Exec("update comments set status = 0 where mid =?",Mid)
		if err!=nil{
			return err
		}
	}else{
		_,err:=db.DB().Exec("insert into comments(uid,mid,status,update_at,create_at) values(?,?,?,?,?)",1,Mid,1,time.Now().Local(),time.Now().Local())
		if err !=nil{
			return err
		}
	}
	return err
}