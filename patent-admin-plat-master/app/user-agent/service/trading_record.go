package service

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service/dto"
	cDto "go-admin/common/dto"
)

type TradingRecord struct {
	service.Service
}

func (e *TradingRecord) GetPages(req *dto.TradingRecordPagesReq, list *[]models.TradingRecord, count *int64) error {
	var err error
	var data models.TradingRecord
	req.Generate(&data)
	if len(req.PatentIDS) == 0 {
		err = e.Orm.Model(&data).
			Scopes(cDto.Paginate(req.GetPageSize(), req.GetPageIndex())).
			Where("properties LIKE ?", fmt.Sprintf("%%%s%%", req.Query)).
			Find(list).Limit(-1).Offset(-1).
			Count(count).Error
	} else {
		err = e.Orm.Model(&data).
			Scopes(cDto.Paginate(req.GetPageSize(), req.GetPageIndex())).
			Where("patent_id IN ?", req.PatentIDS).
			Where("properties LIKE ?", fmt.Sprintf("%%%s%%", req.Query)).
			Find(list).Limit(-1).Offset(-1).
			Count(count).Error
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

func (e *TradingRecord) Create(req *dto.TradingRecordReq) (*models.TradingRecord, error) {
	var data models.TradingRecord
	req.Generate(&data)
	err := e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}

	return &data, nil
}

func (e *TradingRecord) Update(req *dto.TradingRecordReq) (*models.TradingRecord, error) {
	var data models.TradingRecord
	db := e.Orm.First(&data, req.ID)
	if err := db.Error; err != nil {
		e.Log.Errorf("Service Update TradingRecord error: %s", err)
		return nil, err
	}

	req.Generate(&data)
	err := e.Orm.Updates(data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}

	return &data, nil
}

func (e *TradingRecord) Delete(id int) error {
	var data models.TradingRecord
	err := e.Orm.Model(&data).Delete(&data, id).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}
