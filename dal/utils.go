package dal

import (
	"fmt"
	"time"
)

const (
	PAGE_NUM        = "pn"
	PAGE_RECORD_NUM = "rn"
)

type Parameters map[string]interface{}

func CombineCondition(pars Parameters) string {
	condition := ""
	if pars == nil {
		return condition
	}

	first := true
	for k, v := range pars {
		if k == PAGE_NUM || k == PAGE_RECORD_NUM {
			continue
		}

		if first == false {
			condition += " and "
		}
		first = false
		condition += fmt.Sprintf("`%v`=\"%v\"", k, v)
	}
	return condition
}

func SimpleQuery(status bool, pars Parameters, model interface{}) error {
	condition := CombineCondition(pars)
	query := db


	if len(condition) > 0 {
		query = query.Where(condition).Order("updated_at desc")
	} else {
		query = query.Order("updated_at desc")
	}

	if status {
		query = query.Where("status = 0")
	}

	if pnI, exists := pars[PAGE_NUM]; exists {
		rn := pars[PAGE_RECORD_NUM].(int64)
		query = query.Offset(rn * (pnI.(int64) - 1)).Limit(rn)
	}
	return query.Find(model).Error
}

func GetObjById(id uint, model interface{}) error {
	params := make(map[string]interface{})
	params["ID"] = id
	err := SimpleQuery(false, params, model)
	if err != nil {
		return err
	}
	return nil
}

func SoftDelete(WhereParams Parameters, updater string, model interface{}) error {
	WhereParams["status"] = 0
	condition := CombineCondition(WhereParams)
	updateParams := make(map[string]interface{})
	updateParams["updated_at"] = time.Now().Local()
	updateParams["Updater"] = updater
	updateParams["status"] = 1
	return db.Model(&model).Where(condition).Update(updateParams).Error
}

func SoftDeleteById(id uint, updater string, model interface{}) error {
	Params := make(Parameters)
	Params["ID"] = id
	return SoftDelete(Params, updater, model)
}
