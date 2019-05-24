package dal

import (
	"time"
	"PocketMusic/dal/model"
)

func GetLikeStuatus(Mid uint)(bool,error)  {
	count:=0
	//println("mid:",Mid)
	err:=db.Table("likes").Where("mid = ? and status = 1",Mid).Count(&count).Error
	//println("count:",count)
	if count>0{
		return true,err
	}
	return false,err
}

func AddLike(Mid uint)(bool,error)  {
	count:=0
	err:=db.Table("likes").Where("mid =?",Mid).Count(&count).Error
	if err!=nil{
		return false,err
	}else if count>0{
		_,err:=db.DB().Exec("update likes set status = 1 where mid =?",Mid)
		if err!=nil{
			return false,err
		}
	}else{
		_,err:=db.DB().Exec("insert into likes(uid,mid,status,update_at,create_at) values(?,?,?,?,?)",1,Mid,1,time.Now().Local(),time.Now().Local())
		if err !=nil{
			return false,err
		}
	}
	return true,err
}
func DeleteLike(Mid uint)(bool,error) {
	_,err:=db.DB().Exec("update likes set status =0 where mid = ?",Mid)
	if err!=nil {
		return false, err
	}
	return true,err
}

func GetMusic(Mid uint,music *model.MusicInfo)  (error){
	//println("mid:",Mid)
	row:=db.DB().QueryRow("select mid,mname,singer,lrc,url from songs where mid =?",Mid)
	err:=row.Scan(&music.Mid,&music.Mname,&music.Singer,&music.Lrc,&music.Source)
	println("music:",music.Singer)
	if err!=nil{
		return err
	}
	return err
}