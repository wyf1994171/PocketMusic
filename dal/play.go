package dal

import (
	"PocketMusic/dal/model"
	"time"
)

func GetLikeStuatus(Uid string,Mid uint)(bool,error)  {
	count:=0
	//println("mid:",Mid)
	err:=db.Table("likes").Where("mid = ? and status = 0 and uid =?",Mid,Uid).Count(&count).Error
	//println("count:",count)
	if count>0{
		return true,err
	}
	return false,err
}

func GetMusic(Mid uint,music *model.MusicInfo)  (error) {
	//println("mid:",Mid)
	count:=0
	err:=db.Table("songs").Where("mid =?",Mid).Count(&count).Error
	if err!=nil{
		return err
	}
	if count>0 {
		row := db.DB().QueryRow("select mid,mname,singer,lrc,url from songs where mid =?", Mid)
		err := row.Scan(&music.Mid, &music.Mname, &music.Singer, &music.Lrc, &music.Source)
		//println("music:", music.Singer)
		if err != nil {
			return err
		}
	}
	return err
}
func ChangeLike(Uid string,Mid uint,Status uint)  error{
	count:=0
	err:=db.Table("likes").Where("uid =? and mid =?",Uid,Mid).Count(&count).Error
	if err!=nil{
		return err
	}
	if count>0{
		_,err:=db.DB().Exec("update likes set status =? where uid = ? and mid = ?",Status,Uid,Mid)
		return err
	} else {
		_,err:=db.DB().Exec("insert into likes(uid,mid,status,create_at) values(?,?,?,?)",Uid,Mid,Status,time.Now().Local())
		return err
	}

}