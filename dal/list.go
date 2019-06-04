package dal

import (
	"PocketMusic/dal/model"
	"time"
)

func GetLists(uid string)([]map[string]interface{},error) {
	whereParams := make(map[string]interface{})
	whereParams["uid"] = uid
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	var lists []*model.List
	err := db.Debug().Where(condition).Find(&lists).Error
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

func AddList(uid string,name string) (uint,error) {
	list := &model.List{
		RecordMeta : model.RecordMeta{
			CreatedAt: time.Now().Local(),
			UpdatedAt: time.Now().Local(),
		},
		Name: name,
		Status: 0,
		Num: 0,
		UID: uid,
	}
	err := db.Save(list).Error
	return list.RecordMeta.ID, err
}

func AddListCoverPath(coverPath string,id uint) error {
	whereParams := make(map[string]interface{})
	whereParams["id"] = id
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	updateParams := make(map[string]interface{})
	updateParams["cover_path"] = coverPath
	return db.Model(&model.List{}).Where(condition).Updates(updateParams).Error
}

func AddListSong(lid,mid uint) error {
	listSong := &model.ListSong{
		RecordMeta: model.RecordMeta{
			CreatedAt: time.Now().Local(),
			UpdatedAt: time.Now().Local(),
		},
		Lid: lid,
		Mid: mid,
		Status: 0,
	}
	err := db.Save(listSong).Error
	return err
}

func GetListSongIds(lid uint) ([]uint,error) {
	whereParams := make(map[string]interface{})
	whereParams["lid"] = lid
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	var listSongs []*model.ListSong
	err := db.Where(condition).Find(&listSongs).Error
	var result []uint
	for key := range listSongs {
		result =append(result,listSongs[key].Mid)
	}
	return result,err
}

func DeleteListSong(lid,mid uint) error {
	whereParams := make(map[string]interface{})
	whereParams["lid"] = lid
	whereParams["mid"] = mid
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	updateParams := make(map[string]interface{})
	updateParams["status"] = 1
	updateParams["updated_at"] = time.Now().Local()
	return db.Model(&model.ListSong{}).Where(condition).Updates(updateParams).Error
}

func DeleteList(Uid string,Lid uint)(error)  {
	//println("uid: ",Uid)
	//println("lid:",Lid)
	//count:=0
	//db.Table("lists").Where("uid =? and id =?",Uid,Lid).Count(&count)
	//println("count:",count)
	_,err:=db.DB().Exec("update lists set status = 1 where id = ? and uid =?",Lid,Uid)
	return err
}

func ChangeNum(Lid uint)error  {
	count:=0
	err:=db.Table("list_songs").Where("lid=? and status=0",Lid).Count(&count).Error
	if err!=nil{
		return err
	}
	_,err=db.DB().Exec("update lists set num = ? where id = ?",count,Lid)
	return err
}
