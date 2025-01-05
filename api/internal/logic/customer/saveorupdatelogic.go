package customer

import (
	"UBC/api/library/xerr"
	"UBC/models"
	"context"

	"UBC/api/internal/svc"
	"UBC/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveOrUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateLogic {
	return &SaveOrUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveOrUpdateLogic) SaveOrUpdate(req *types.SaveOrUpdateCustomer) error {

	customer := &models.Customer{
		CustomerID:      req.CustomerID,
		CustomerCode:    req.CustomerCode,
		CustomerEmail:   req.CustomerEmail,
		CustomerName:    req.CustomerName,
		BillingContact:  req.BillingContact,
		NotifyContact:   req.NotifyContact,
		PaymentTerm:     req.PaymentTerm,
		ShipTo:          req.ShipTo,
		SalesPerson:     req.SalesPerson,
		UbcMerchandiser: req.UbcMerchandiser,
		Country:         req.Country,
		DischargeLoc:    req.DischargeLoc,
		Status:          req.Status,
		DueDateGap:      req.DueDateGap,
		Code:            req.Code,
	}

	err := customer.SaveOrUpdate()
	if err != nil {
		l.Errorf("[SaveOrUpdate] err:%v", err)
		return xerr.ServerCommonError
	}
	return nil
}
