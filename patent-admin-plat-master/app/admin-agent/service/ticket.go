package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin-agent/model"
	"go-admin/app/admin-agent/service/dtos"
	aModels "go-admin/app/admin/models"
	"go-admin/app/user-agent/service/dto"
	cDto "go-admin/common/dto"
	"time"
)

type Ticket struct {
	service.Service
}

func (e *Ticket) GetTicketPages(req *dtos.TicketPagesReq, list *[]model.Ticket, count *int64) error {
	var err error
	var data model.Ticket
	data.Type = req.Type
	data.Status = req.Status
	data.CreateBy = req.UserID
	err = e.Orm.Model(&data).
		Scopes(cDto.Paginate(req.GetPageSize(), req.GetPageIndex())).
		Where(&data).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Query)).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

func (e *Ticket) GetReportTicketPages(req *dtos.TicketPagesReq, reportType string, list *[]model.Ticket, count *int64) error {
	var err error
	var ticketData model.Ticket
	var tickets []model.Ticket
	ticketData.Type = req.Type
	ticketData.Status = req.Status
	ticketData.CreateBy = req.UserID
	err = e.Orm.Model(&ticketData).
		Where(&ticketData).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Query)).
		Find(&tickets).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	ticketIDs := make([]int, 0, len(tickets))
	ticketsMap := make(map[int]model.Ticket)
	for _, t := range tickets {
		ticketIDs = append(ticketIDs, t.ID)
		ticketsMap[t.ID] = t
	}

	tl := model.ReportRelation{}
	tls := make([]model.ReportRelation, 0)
	err = e.Orm.Model(&tl).
		Where("ticket_id IN ?", ticketIDs).
		Find(&tls).Limit(-1).Offset(-1).
		Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	reportIDs := make([]int, 0, len(tickets))
	tlMap := make(map[int]model.ReportRelation)
	for _, rel := range tls {
		reportIDs = append(reportIDs, rel.ReportId)
		tlMap[rel.ReportId] = rel
	}

	reports := make([]model.Report, 0)
	if len(reportIDs) != 0 {
		var data model.Report
		data.Type = reportType
		err = e.Orm.Model(&data).
			Scopes(cDto.Paginate(req.GetPageSize(), req.GetPageIndex())).
			Where(&data).
			Find(&reports, reportIDs).Limit(-1).Offset(-1).
			Count(count).Error
		if err != nil {
			e.Log.Errorf("db error:%s", err)
			return err
		}
	}

	for _, report := range reports {
		tid := tlMap[report.ReportId].TicketId
		if ticket, ok := ticketsMap[tid]; ok {
			resTicket := ticket
			resTicket.RelObj = report
			*list = append(*list, resTicket)
		}
	}
	return nil
}

func (e *Ticket) GetTicketList(req *dtos.TicketListReq, list *[]model.Ticket, count *int64) error {
	var err error
	var data model.Ticket
	data.Type = req.Type
	data.Status = req.Status
	data.CreateBy = req.UserID
	err = e.Orm.Model(&data).
		Where(&data).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

func (e *Ticket) GetTicketListByIds(ids []int) (list []model.Ticket, err error) {
	var data model.Ticket
	err = e.Orm.Model(&data).
		Where(&data).
		Find(&list, ids).Limit(-1).Offset(-1).
		Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return nil, err
	}
	return
}

func (e *Ticket) GetByID(id int) (*model.Ticket, error) {
	var err error
	var data model.Ticket
	err = e.Orm.Model(&data).
		First(&data, id).
		Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return nil, err
	}
	return &data, nil
}

func (e *Ticket) Create(req *dtos.TicketDBReq) (*model.Ticket, error) {
	var user aModels.SysUser
	err := e.Orm.Model(&user).Debug().
		First(&user, req.UserID).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}

	var t model.Ticket
	req.Generate(&t)
	t.CreateBy = req.UserID

	// gen logs
	ols := newOptLogs()
	if len(user.NickName) != 0 {
		ols.addCreateOptLog(user.NickName)
	} else {
		ols.addCreateOptLog(user.Username)
	}
	t.OptLogs = ols.String()

	t.Status = dtos.TicketStatusOpen

	err = e.Orm.Create(&t).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}

	return &t, nil
}

func (e *Ticket) Update(id int, req *dtos.TicketDBReq) error {
	var user aModels.SysUser
	err := e.Orm.Model(&user).Debug().
		First(&user, req.UserID).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	var t model.Ticket
	db := e.Orm.First(&t, id)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateTicket error: %s", err)
		return err
	}

	switch req.FilesOpt {
	case dto.FilesAdd:
		req.GenerateAndAddFiles(&t)
	case dto.FilesDelete:
		req.GenerateAndDeleteFiles(&t)
	default:
		req.Generate(&t)
	}

	t.UpdateBy = req.UserID

	// gen logs
	ols, err := unmarshalOptLogs([]byte(t.OptLogs))
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if len(user.NickName) != 0 {
		ols.addProcessOptLog(user.NickName, req.OptMsg)
	} else {
		ols.addProcessOptLog(user.Username, req.OptMsg)
	}
	t.OptLogs = ols.String()

	update := e.Orm.Model(&t).Where("id = ?", &t.ID).Updates(&t)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	return nil
}

func (e *Ticket) Close(id int, req *dtos.TicketDBReq) error {
	var user aModels.SysUser
	err := e.Orm.Model(&user).Debug().
		First(&user, req.UserID).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	var t model.Ticket
	db := e.Orm.First(&t, id)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service CloseTicket error: %s", err)
		return err
	}

	req.Generate(&t)
	t.UpdateBy = req.UserID

	// update logs
	ols, err := unmarshalOptLogs([]byte(t.OptLogs))
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if len(user.NickName) != 0 {
		if len(req.OptMsg) != 0 {
			ols.addCloseOptLog(user.NickName, req.OptMsg)
		} else {
			ols.addCompleteOptLog(user.NickName)
		}
	} else {
		if len(req.OptMsg) != 0 {
			ols.addCloseOptLog(user.NickName, req.OptMsg)
		} else {
			ols.addCompleteOptLog(user.Username)
		}
	}
	t.OptLogs = ols.String()

	t.Status = dtos.TicketStatusClosed

	update := e.Orm.Model(&t).Where("id = ?", &t.ID).Updates(&t)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	return nil
}

func (e *Ticket) Finish(id int, req *dtos.TicketDBReq) error {
	var user aModels.SysUser
	err := e.Orm.Model(&user).Debug().
		First(&user, req.UserID).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	var t model.Ticket
	db := e.Orm.First(&t, id)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service CloseTicket error: %s", err)
		return err
	}

	req.Generate(&t)
	t.UpdateBy = req.UserID

	// update logs
	ols, err := unmarshalOptLogs([]byte(t.OptLogs))
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if len(user.NickName) != 0 {
		if len(req.OptMsg) != 0 {
			ols.addCloseOptLog(user.NickName, req.OptMsg)
		} else {
			ols.addCompleteOptLog(user.NickName)
		}
	} else {
		if len(req.OptMsg) != 0 {
			ols.addCloseOptLog(user.NickName, req.OptMsg)
		} else {
			ols.addCompleteOptLog(user.Username)
		}
	}
	t.OptLogs = ols.String()

	t.Status = dtos.TicketStatusFinished

	update := e.Orm.Model(&t).Where("id = ?", &t.ID).Updates(&t)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	return nil
}

func (e *Ticket) Remove(id int) error {
	var err error
	var data model.Ticket

	db := e.Orm.Model(&data).
		Delete(&data, id)
	if err = db.Error; err != nil {
		e.Log.Errorf("Error found in  RemoveTicket : %s", err)
		return err
	}

	return nil
}

type OptLogs struct {
	ols []OptLog
}

func newOptLogs() *OptLogs {
	return &OptLogs{ols: make([]OptLog, 0)}
}

func unmarshalOptLogs(data []byte) (*OptLogs, error) {
	var ols []OptLog
	if err := json.Unmarshal(data, &ols); err != nil {
		return nil, err
	}
	return &OptLogs{ols: ols}, nil
}

func (ols *OptLogs) addCreateOptLog(creator string) {
	ols.ols = append(ols.ols, OptLog{
		Operator:  creator,
		Operation: "创建工单",
		Timestamp: time.Now(),
	})
}

func (ols *OptLogs) addProcessOptLog(processor string, optMsg string) {
	ols.ols = append(ols.ols, OptLog{
		Operator:  processor,
		Operation: fmt.Sprintf("处理工单：%s", optMsg),
		Timestamp: time.Now(),
	})
}

func (ols *OptLogs) addCompleteOptLog(closer string) {
	ols.ols = append(ols.ols, OptLog{
		Operator:  closer,
		Operation: "关闭工单：处理完成",
		Timestamp: time.Now(),
	})
}

func (ols *OptLogs) addCloseOptLog(closer string, optMsg string) {
	ols.ols = append(ols.ols, OptLog{
		Operator:  closer,
		Operation: fmt.Sprintf("关闭工单：%s", optMsg),
		Timestamp: time.Now(),
	})
}

func (ols *OptLogs) String() string {
	bs, _ := json.Marshal(ols.ols)
	return string(bs)
}

type OptLog struct {
	Operator  string    `json:"operator"`
	Operation string    `json:"operation"`
	Timestamp time.Time `json:"timestamp"`
}
