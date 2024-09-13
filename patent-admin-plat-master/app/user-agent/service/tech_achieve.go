package service

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service/dto"
	cDto "go-admin/common/dto"
)

type TechAchieve struct {
	service.Service
}

func (e *TechAchieve) GetPages(req *dto.TechAchievePagesReq, list *[]models.TechAchieve, count *int64) error {
	var err error
	var data models.TechAchieve
	req.Generate(&data)
	err = e.Orm.Model(&data).
		Scopes(cDto.Paginate(req.GetPageSize(), req.GetPageIndex())).
		Where("properties LIKE ?", fmt.Sprintf("%%%s%%", req.Query)).
		Where("create_by = ?", req.UserID).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

func (e *TechAchieve) Create(req *dto.TechAchieveReq) (*models.TechAchieve, error) {
	var data models.TechAchieve
	req.Generate(&data)
	err := e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}

	return &data, nil
}

func (e *TechAchieve) Update(req *dto.TechAchieveReq) (*models.TechAchieve, error) {
	var data models.TechAchieve
	db := e.Orm.First(&data, req.ID)
	if err := db.Error; err != nil {
		e.Log.Errorf("Service Update TechAchieve error: %s", err)
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

func (e *TechAchieve) Delete(id int) error {
	var data models.TechAchieve
	err := e.Orm.Model(&data).Delete(&data, id).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}
