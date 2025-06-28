package common

import (
	"UBC/api/internal/svc"
	"UBC/api/internal/types"
	"UBC/api/library/xerr"
	"UBC/models"
	"UBC/models/res_models"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const maxFileSize = 10 << 20 // 10 MB

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Packing struct {
	PO          string `json:"po"`
	Color       string `json:"color"`
	StyleNumber string `json:"style_number"`
	Sizes       []struct {
		Size     string  `json:"size"`
		Quantity float64 `json:"quantity"`
	} `json:"sizes"`
	TotalQuantity float64 `json:"total_quantity"`
	CTNS          float64 `json:"CTNS"`
}

type Projection struct {
	ExFtyInHouse   string  `json:"ex_fty_in_house"`
	Customer       string  `json:"customer"`
	CustomerPO     string  `json:"customer_po"`
	StyleNo        string  `json:"style_no"`
	DescStyleName  string  `json:"desc_style_name"`
	Color          string  `json:"color"`
	Fabrication    string  `json:"fabrication"`
	QtyPC          int     `json:"qty_pc"`
	Buy            float64 `json:"buy"`
	TtlBuy         float64 `json:"ttl_buy"`
	Sell           float64 `json:"sell"`
	TtlSell        float64 `json:"ttl_sell"`
	Vendor         string  `json:"vendor"`
	WaterResistant string  `json:"water_resistant"`
}

type PoInfo struct {
	Po                  string `json:"Po"`
	Data                string `json:"data"`
	Due                 string `json:"due"`
	StyleNum            string `json:"styleNum"`
	StyleName           string `json:"styleName"`
	Color               string `json:"color"`
	Description         string `json:"description"`
	Qty                 string `json:"qty"`
	Amount              string `json:"amount"`
	CustomerName        string `json:"customerName"`
	Vendor              string `json:"vendor"`
	From                string `json:"from"`
	ShipTo              string `json:"shipTo"`
	ShipTerms           string `json:"shipTerms"`
	PaymentTerms        string `json:"paymentTerms"`
	LastRevised         string `json:"lastRevised"`
	Reference           string `json:"reference"`
	PoTotal             string `json:"poTotal"`
	Page                string `json:"page"`
	ShipVia             string `json:"shipVia"`
	SpecialInstructions string `json:"specialInstructions"`
	Country             string `json:"country"`
	Items               []Item `json:"items"`
}
type Item struct {
	PO        string `json:"PO#"`
	Style     string `json:"STYLE"`
	Color     string `json:"COLOR"`
	ColorDesc string `json:"COLOR DESCRIPTION"`
	Dimension string `json:"DIMENSION"`
	Size      string `json:"SIZE"`
	UPC       string `json:"UPC#"`
	Qty       string `json:"QTY"`
	Cost      string `json:"COST"`
	Extended  string `json:"EXTENDED"`
}

type PackingList struct {
	Orders []Packing `json:"orders"`
}

type ProjectionList struct {
	Orders []Projection `json:"orders"`
}

func (l *UploadFileLogic) UploadFile(req *types.UploadFile, r *http.Request) (resp *types.UploadRes, err error) {
	_ = r.ParseMultipartForm(maxFileSize)
	file, _, err := r.FormFile("file")
	if err != nil {
		l.Errorf("[UploadFile] get file err:%v", err)
		return nil, xerr.RequestParamError
	}
	defer file.Close()
	if req.UsedFor == "packing" || req.UsedFor == "po" {
		// 创建临时文件
		tempFile, err := os.CreateTemp("", "upload-*.pdf")
		if err != nil {
			return nil, xerr.RequestParamError

		}
		defer tempFile.Close()

		_, err = io.Copy(tempFile, file)
		if err != nil {
			return nil, xerr.RequestParamError

		}
		output, err := exec.Command(l.svcCtx.Config.PythonPath, "py/pdf2text.py", tempFile.Name()).Output()
		if err != nil {
			return nil, xerr.RequestParamError
		}
		if req.UsedFor == "po" {
			return l.doPoFile(string(output))
		} else {
			return l.doPackingFile(string(output))
		}
	} else if req.UsedFor == "projection" {
		content, err := ReadXLSXFromReader(file)
		if err != nil {
			l.Errorf("[UploadFile] read xlsx from reader :%v", err)
			return nil, xerr.RequestParamError
		}
		return l.duProjectionFile(content)
	}

	l.Errorf("[UploadFile] current use for not found:%s", req.UsedFor)
	return nil, xerr.RequestParamError

}

func (l *UploadFileLogic) doPackingFile(text string) (resp *types.UploadRes, err error) {
	cmd := exec.Command(l.svcCtx.Config.PythonPath, "py/packing.py", text)
	output, err := cmd.CombinedOutput()
	if err != nil {
		l.Errorf("[doPackingFile] exec packing python script err: %s", string(output))
		return nil, xerr.ServerCommonError
	}

	scanner := bufio.NewScanner(bytes.NewReader(output))
	var (
		jsonData                 string
		packingList              PackingList
		newShipmenPackingResList []res_models.NewShipmenPackingRes
	)

	capture := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "START_ORDERS_DATA" {
			capture = true
			continue
		}
		if strings.TrimSpace(line) == "END_ORDERS_DATA" {
			capture = false
			break
		}
		if capture {
			jsonData += line
		}
	}

	if err = scanner.Err(); err != nil {
		l.Errorf("[doPackingFile] Error reading output: %s", err)
		return nil, xerr.RequestParamError
	}

	err = json.Unmarshal([]byte(jsonData), &packingList)
	if err != nil {
		l.Errorf("[doPackingFile] Error parsing JSON: %s", err)
		return nil, xerr.RequestParamError
	}

	for i, order := range packingList.Orders {
		newShipmenPackingRes := res_models.NewShipmenPackingRes{
			Id:            fmt.Sprintf("%d", i),
			CustomerPo:    order.PO,
			StyleCode:     order.StyleNumber,
			Color:         order.Color,
			TotalQuantity: order.TotalQuantity,
			CartonCnt:     order.CTNS,
		}
		newShipmenPackingResList = append(newShipmenPackingResList, newShipmenPackingRes)
	}
	return &types.UploadRes{Res: newShipmenPackingResList}, nil

}

func (l *UploadFileLogic) duProjectionFile(text string) (resp *types.UploadRes, err error) {
	cmd := exec.Command(l.svcCtx.Config.PythonPath, "py/projection_output.py", text)
	output, err := cmd.CombinedOutput()
	if err != nil {
		l.Errorf("[duProjectionFile] exec packing python script err: %s", string(output))
		return nil, xerr.ServerCommonError
	}

	scanner := bufio.NewScanner(bytes.NewReader(output))
	var (
		jsonData       string
		projectionList ProjectionList
	)

	capture := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "START_ORDERS_DATA" {
			capture = true
			continue
		}
		if strings.TrimSpace(line) == "END_ORDERS_DATA" {
			capture = false
			break
		}
		if capture {
			jsonData += line
		}
	}

	if err = scanner.Err(); err != nil {
		l.Errorf("[duProjectionFile] Error reading output: %s", err)
		return nil, xerr.RequestParamError
	}

	err = json.Unmarshal([]byte(jsonData), &projectionList)
	if err != nil {
		l.Errorf("[duProjectionFile] Error parsing JSON: %s", err)
		return nil, xerr.RequestParamError
	}
	ps := make([]models.Projection, len(projectionList.Orders))
	for i := range projectionList.Orders {
		ps[i] = models.Projection{
			ArriveDt:     projectionList.Orders[i].ExFtyInHouse,
			CustomerCode: projectionList.Orders[i].Customer,
			CustomerPo:   projectionList.Orders[i].CustomerPO,
			StyleCode:    projectionList.Orders[i].StyleNo,
			StyleName:    projectionList.Orders[i].DescStyleName,
			Color:        projectionList.Orders[i].Color,
			Fabrication:  projectionList.Orders[i].Fabrication,
			PoQty:        projectionList.Orders[i].QtyPC,
			CostPrice:    projectionList.Orders[i].Buy,
			SalePrice:    projectionList.Orders[i].Sell,
			TtlBuy:       projectionList.Orders[i].TtlBuy,
			TtlSell:      projectionList.Orders[i].TtlSell,
		}
	}
	p := models.Projection{}
	err = p.SaveAll(ps)
	if err != nil {
		l.Errorf("upload projection info save all err:%v", err)
		return nil, xerr.DbError
	}
	return nil, err
}

func (l *UploadFileLogic) doPoFile(text string) (resp *types.UploadRes, err error) {
	cmd := exec.Command(l.svcCtx.Config.PythonPath, "py/po.py", text)
	output, err := cmd.CombinedOutput()
	if err != nil {
		l.Errorf("[doPoFile] exec packing python script err: %s", string(output))
		return nil, xerr.ServerCommonError
	}

	var pos []PoInfo
	err = json.Unmarshal(output, &pos)
	if err != nil {
		l.Errorf("[doPackingFile] Error parsing JSON: %s", err)
		return nil, xerr.RequestParamError
	}

	// 转换为 ProjectionPo 切片
	projections := make([]models.ProjectionPo, len(pos))
	for i, po := range pos {
		qtyInt, _ := strconv.Atoi(po.Qty)
		amount, _ := strconv.ParseFloat(strings.ReplaceAll(po.Amount, ",", ""), 64)
		poTotal, _ := strconv.ParseFloat(strings.ReplaceAll(po.PoTotal, ",", ""), 64)
		items, _ := json.Marshal(po.Items)

		// 从items中提取成本和价格信息
		var costPrice float64
		var totalCost float64
		size := ""

		if len(po.Items) > 0 {
			size = po.Items[0].Size
			// 提取第一个items的成本价格
			if po.Items[0].Cost != "" {
				costPrice, _ = strconv.ParseFloat(strings.ReplaceAll(po.Items[0].Cost, ",", ""), 64)
			}
			// 提取第一个items的扩展总价
			if po.Items[0].Extended != "" {
				totalCost, _ = strconv.ParseFloat(strings.ReplaceAll(po.Items[0].Extended, ",", ""), 64)
			}
		}

		// 如果items中没有价格信息，使用po级别的信息
		if costPrice == 0 && amount > 0 {
			costPrice = amount
		}
		if totalCost == 0 && poTotal > 0 {
			totalCost = poTotal
		}

		projections[i] = models.ProjectionPo{
			// 基本信息
			ArriveDt:     po.Due,          // 到货日期 (due -> arrive_dt)
			PoDate:       po.Data,         // PO日期 (data -> po_date)
			UbcPi:        po.Reference,    // UBC PI (reference -> ubc_pi，如果为空则保持空)
			FobLdp:       po.ShipTerms,    // FOB LDP (使用shipTerms -> fob_ldp)
			CustomerCode: po.CustomerName, // 客户代码 (customerName -> customer_code)
			Country:      po.Country,      // 国家 (country -> country)
			CustomerPo:   po.Po,           // 客户PO号 (Po -> customer_po)
			MasterPo:     "",              // 主PO (PDF中确实未提供)
			StyleCode:    po.StyleNum,     // 款号 (styleNum -> style_code)
			StyleName:    po.StyleName,    // 款名 (styleName -> style_name)
			Fabrication:  po.Description,  // 面料/描述 (description -> fabrication)
			Color:        po.Color,        // 颜色 (color -> color)
			Size:         size,            // 尺码 (从items中提取 -> size)

			// 数量和价格
			PoQty:         qtyInt, // PO数量 (qty -> po_qty)
			ShipQty:       0,      // 发货数量 (PDF中未提供)
			SalePrice:     0,
			SaleCustPrice: 0,         // 客户销售价格 (PDF中未提供)
			SaleCurrency:  "USD",     // 销售货币 (默认USD -> sale_currency)
			InvoiceCode:   "",        // 发票代码 (PDF中未提供)
			Receiving:     "",        // 收货 (PDF中未提供)
			Notes:         "",        // 备注 (PDF中的SPECIAL INSTRUCTIONS为空)
			CostPrice:     costPrice, // 成本价格 (从items中提取的COST -> cost_price)
			CostCurrency:  "USD",     // 成本货币 (默认USD -> cost_currency)
			RmbInv:        "",        // 人民币发票 (PDF中未提供)

			// 供应商和付款信息
			Exporter:        po.Vendor,       // 出口商 (vendor -> exporter)
			UbcPayable:      0,               // UBC应付款 (PDF中未提供)
			PayPeriod:       po.PaymentTerms, // 付款期限 (paymentTerms -> pay_period)
			SalesPerson:     "",              // 销售人员 (PDF中未提供)
			SalesCommission: 0,               // 销售佣金 (PDF中未提供)
			CommPaid:        0,               // 已付佣金 (PDF中未提供)

			// 总计
			TtlSell: 0,
			TtlBuy:  totalCost, // 总采购金额 (从items的EXTENDED中提取 -> ttl_buy)

			// JSON和文本字段
			PoItems:             items,                  // PO条目详情 (items -> po_items)
			ShipTo:              po.ShipTo,              // 发货地址 (shipTo -> ship_to)
			ShipFrom:            po.From,                // 发货方 (from -> ship_from)
			ShipTerms:           po.ShipTerms,           // 运输条件 (shipTerms -> ship_terms)
			PaymentTerms:        po.PaymentTerms,        // 付款条件 (paymentTerms -> payment_terms)
			LastRevised:         po.LastRevised,         // 最后修订 (lastRevised -> last_revised)
			PoTotal:             poTotal,                // PO总额 (poTotal -> po_total)
			PageInfo:            po.Page,                // 页面信息 (page -> page_info)
			ShipVia:             po.ShipVia,             // 运输方式 (shipVia -> ship_via，即使为空也设置)
			SpecialInstructions: po.SpecialInstructions, // 特殊说明 (specialInstructions -> special_instructions，即使为空也设置)
		}
	}

	// 保存到数据库
	p := models.ProjectionPo{}
	err = p.SaveAll(projections)
	if err != nil {
		l.Errorf("upload po info save all err:%v", err)
		return nil, xerr.DbError
	}

	return &types.UploadRes{Res: projections}, nil
}

// ReadXLSXFromReader 读取xlsx文件并将其内容转换为字符串
func ReadXLSXFromReader(file io.Reader) (string, error) {
	// 将 io.Reader 内容读取到一个缓冲区中
	buf, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("无法读取文件: %v", err)
	}

	// 创建一个 bytes.Reader 实现 io.ReaderAt 接口
	reader := bytes.NewReader(buf)

	// 打开xlsx文件
	xlFile, err := xlsx.OpenReaderAt(reader, reader.Size())
	if err != nil {
		return "", fmt.Errorf("无法读取文件: %v", err)
	}

	// 使用strings.Builder来构建字符串
	var sb strings.Builder

	// 遍历文件中的每一个表格
	for _, sheet := range xlFile.Sheets {
		// 遍历表格中的每一行
		for _, row := range sheet.Rows {
			// 遍历行中的每一个单元格
			var rowValues []string
			for _, cell := range row.Cells {
				text, _ := cell.FormattedValue()
				rowValues = append(rowValues, text)
			}
			// 将每一行的字段用逗号连接
			sb.WriteString(strings.Join(rowValues, ","))
			sb.WriteString("\n")
		}
	}

	// 返回构建的字符串
	return sb.String(), nil
}
