package shipment

import (
	"UBC/api/internal/svc"
	"UBC/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveLogic {
	return &SaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveLogic) Save(req *types.SaveShipment) error {
	//// shipment
	//st := req.ShipmentInfo
	//// 将数据填充到 Shipment 结构体
	//shipment := &models.Shipment{
	//	MasterBlNum:  st.BillOfLanding,
	//	CustomerCode: st.CustomerCode,
	//	ShipFrom:     st.ShipFrom,
	//	ShipMethod:   st.ShipMethod,
	//	OrigCountry:  st.CountryOfOrigin,
	//	Exporter:     st.Manufacture,
	//	ShipName:     st.VesselFlight,
	//	ShipDt:       time.Unix(st.ETDDt, 0).Format("2006-01-02 15:04:05"),
	//	UbcPi:        st.UBCPI,
	//	Markurl:      st.MarksAndNumbers,
	//}
	//pl := req.Packings
	//
	//plist := make([]models.PackingList, len(req.Packings))
	//
	//for i := range pl {
	//	plist[i] = models.PackingList{
	//		ProjID:      pl[i].ProjID,
	//		CartonCnt:   pl[i].Carton,
	//		ItemCnt:     pl[i].QtyPC,
	//		GrossWeight: pl[i].GWKG,
	//		NetWeight:   pl[i].NWKG,
	//		MeasVol:     pl[i].Meas,
	//	}
	//}
	//ivo := req.Invoice
	//invoice := &models.Invoice{
	//	UbcPi:       st.UBCPI,
	//	InvoiceCode: ivo.InvoiceCode,
	//	InvoiceAmt:  ivo.AdditionalCost,
	//	ReceivedAmt: ivo.DepositAmt,
	//	InvoiceDt:   time.Unix(ivo.InvoiceDt, 0).Format("2006-01-02 15:04:05"),
	//	InvoiceDue:  time.Unix(ivo.InvoiceDue, 0).Format("2006-01-02 15:04:05"),
	//}

	return nil
}
