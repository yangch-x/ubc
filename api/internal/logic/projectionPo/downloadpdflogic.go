package projectionPo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"UBC/api/internal/svc"
	"UBC/api/internal/types"
	"UBC/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadPdfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadPdfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadPdfLogic {
	return &DownloadPdfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProjectionPoItem 代表po_items JSON中的单个项目
type ProjectionPoItem struct {
	PO               string `json:"PO#"`
	QTY              string `json:"qTY"`
	UPC              string `json:"UPC#"`
	SIZE             string `json:"sIZE"`
	COLOR            string `json:"cOLOR"`
	STYLE            string `json:"sTYLE"`
	DIMENSION        string `json:"dIMENSION"`
	ColorDescription string `json:"COLOR DESCRIPTION"`
	COST             string `json:"COST"`
	EXTENDED         string `json:"EXTENDED"`
}

// PDFOrderData 用于生成PDF的订单数据结构
type PDFOrderData struct {
	Po                  string    `json:"Po"`
	Data                string    `json:"data"`
	Due                 string    `json:"due"`
	StyleNum            string    `json:"styleNum"`
	StyleName           *string   `json:"styleName"`
	Color               string    `json:"color"`
	Description         string    `json:"description"`
	Qty                 string    `json:"qty"`
	Amount              string    `json:"amount"`
	CustomerName        string    `json:"customerName"`
	Vendor              string    `json:"vendor"`
	From                string    `json:"from"`
	ShipTo              string    `json:"shipTo"`
	ShipTerms           string    `json:"shipTerms"`
	PaymentTerms        string    `json:"paymentTerms"`
	LastRevised         string    `json:"lastRevised"`
	Reference           *string   `json:"reference"`
	PoTotal             string    `json:"poTotal"`
	Page                string    `json:"page"`
	ShipVia             *string   `json:"shipVia"`
	Items               []PDFItem `json:"items"`
	SpecialInstructions *string   `json:"specialInstructions"`
}

// PDFItem 用于PDF中的项目数据
type PDFItem struct {
	PO               string  `json:"PO#"`
	STYLE            string  `json:"STYLE"`
	COLOR            string  `json:"COLOR"`
	ColorDescription string  `json:"COLOR DESCRIPTION"`
	DIMENSION        *string `json:"DIMENSION"`
	SIZE             string  `json:"SIZE"`
	UPC              string  `json:"UPC#"`
	QTY              string  `json:"QTY"`
	COST             string  `json:"COST"`
	EXTENDED         *string `json:"EXTENDED"`
}

func (l *DownloadPdfLogic) DownloadPdf(req *types.IdsReq, w http.ResponseWriter) error {
	// 虽然是ids数组，但根据需求只处理第一个id
	if len(req.Ids) == 0 {
		return fmt.Errorf("no ids provided")
	}

	id := req.Ids[0]

	// 从数据库查询数据
	po := models.ProjectionPo{}
	pos, err := po.SearchByIds([]int{id})
	if err != nil {
		l.Logger.Errorf("Failed to query projection po: %v", err)
		return fmt.Errorf("failed to query projection po: %v", err)
	}

	if len(pos) == 0 {
		return fmt.Errorf("no projection po found with id: %d", id)
	}

	// 取第一条数据
	projectionPo := pos[0]

	// 转换数据格式
	orderData, err := l.convertToOrderData(projectionPo)
	if err != nil {
		l.Logger.Errorf("Failed to convert data: %v", err)
		return fmt.Errorf("failed to convert data: %v", err)
	}

	// 调用Python脚本生成PDF并返回给前端
	err = l.callPythonScriptAndReturnPDF(orderData, w)
	if err != nil {
		l.Logger.Errorf("Failed to call python script: %v", err)
		return fmt.Errorf("failed to call python script: %v", err)
	}

	return nil
}

// convertToOrderData 将数据库查询结果转换为PDF生成所需的格式
func (l *DownloadPdfLogic) convertToOrderData(po *models.ProjectionPo) (*PDFOrderData, error) {
	// 解析po_items JSON
	var poItems []ProjectionPoItem
	if po.PoItems != nil {
		err := json.Unmarshal(po.PoItems, &poItems)
		if err != nil {
			// 如果解析失败，记录错误但继续处理
			l.Error("Failed to unmarshal po_items: %v", err)
			// 创建一个默认的空项目
			poItems = []ProjectionPoItem{}
		}
	}

	// 如果po_items为空，创建一个基于其他字段的默认项目
	if len(poItems) == 0 {
		defaultItem := ProjectionPoItem{
			PO:               po.CustomerPo,
			STYLE:            po.StyleCode,
			COLOR:            po.Color,
			ColorDescription: po.StyleName,
			DIMENSION:        "",
			SIZE:             po.Size,
			UPC:              "",
			QTY:              strconv.Itoa(po.PoQty),
			COST:             fmt.Sprintf("%.2f", po.SalePrice),
			EXTENDED:         "",
		}
		poItems = append(poItems, defaultItem)
	}

	// 转换po_items为PDF items格式
	pdfItems := make([]PDFItem, len(poItems))
	totalQty := 0
	for i, item := range poItems {
		qty, err := strconv.Atoi(item.QTY)
		if err != nil {
			l.Logger.Errorf("Failed to parse QTY '%s': %v", item.QTY, err)
			qty = 0
		}
		totalQty += qty

		// 处理空值
		dimension := ""
		if item.DIMENSION != "" {
			dimension = item.DIMENSION
		}

		// 处理COST字段，如果为空则使用po.SalePrice作为默认值
		itemCost := item.COST
		if itemCost == "" {
			itemCost = fmt.Sprintf("%.2f", po.SalePrice)
		}

		// 处理EXTENDED字段
		extended := item.EXTENDED
		if extended == "" {
			extended = ""
		}

		pdfItems[i] = PDFItem{
			PO:               item.PO,
			STYLE:            item.STYLE,
			COLOR:            item.COLOR,
			ColorDescription: item.ColorDescription,
			DIMENSION:        &dimension,
			SIZE:             item.SIZE,
			UPC:              item.UPC,
			QTY:              item.QTY,
			COST:             itemCost,
			EXTENDED:         &extended,
		}
	}

	// 如果totalQty为0，使用数据库中的PoQty
	if totalQty == 0 {
		totalQty = po.PoQty
	}

	// 构建订单数据 - 根据实际的字段映射关系
	orderData := &PDFOrderData{
		Po:                  po.CustomerPo,                        // customer_po -> Po
		Data:                po.PoDate,                            // po_date -> data
		Due:                 po.ArriveDt,                          // arrive_dt -> due
		StyleNum:            po.StyleCode,                         // style_code -> styleNum
		StyleName:           &po.StyleName,                        // style_name -> styleName
		Color:               po.Color,                             // color -> color
		Description:         po.Fabrication,                       // fabrication -> description
		Qty:                 strconv.Itoa(totalQty),               // 从items计算或使用po_qty -> qty
		Amount:              fmt.Sprintf("%.2f", po.SalePrice),    // sale_price -> amount
		CustomerName:        po.CustomerCode,                      // customer_code -> customerName
		Vendor:              po.Exporter,                          // exporter -> vendor
		From:                po.ShipFrom,                          // ship_from -> from
		ShipTo:              po.ShipTo,                            // ship_to -> shipTo
		ShipTerms:           po.ShipTerms,                         // ship_terms -> shipTerms
		PaymentTerms:        po.PaymentTerms,                      // payment_terms -> paymentTerms
		LastRevised:         po.LastRevised,                       // last_revised -> lastRevised
		Reference:           getStringPtr(po.UbcPi),               // ubc_pi -> reference
		PoTotal:             fmt.Sprintf("%.2f", po.PoTotal),      // po_total -> poTotal
		Page:                po.PageInfo,                          // page_info -> page
		ShipVia:             getStringPtr(po.ShipVia),             // ship_via -> shipVia
		Items:               pdfItems,                             // po_items JSON解析后的数据 -> items
		SpecialInstructions: getStringPtr(po.SpecialInstructions), // special_instructions -> specialInstructions
	}

	return orderData, nil
}

// getStringPtr 辅助函数，将空字符串转换为nil指针
func getStringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// callPythonScriptAndReturnPDF 调用Python脚本生成PDF并直接返回给前端
func (l *DownloadPdfLogic) callPythonScriptAndReturnPDF(orderData *PDFOrderData, w http.ResponseWriter) error {
	// 将orderData转换为JSON数组格式（Python脚本期望的格式）
	jsonArray := []*PDFOrderData{orderData}
	jsonData, err := json.Marshal(jsonArray)
	if err != nil {
		return fmt.Errorf("failed to marshal order data to JSON: %v", err)
	}

	// 获取当前工作目录
	workDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %v", err)
	}

	// Python脚本路径
	pythonScript := filepath.Join(workDir, "py", "pdf_nb.py")

	// 检查Python脚本是否存在
	if _, err := os.Stat(pythonScript); os.IsNotExist(err) {
		return fmt.Errorf("python script not found: %s", pythonScript)
	}

	// 获取配置的Python路径
	pythonPath := l.svcCtx.Config.PythonPath
	if pythonPath == "" {
		// 如果配置为空，尝试默认的Python命令
		pythonCommands := []string{"python", "python3", "py"}
		var lastErr error

		for _, pythonCmd := range pythonCommands {
			// 调用Python脚本，直接传递JSON字符串
			cmd := exec.Command(pythonCmd, pythonScript, string(jsonData))
			cmd.Dir = filepath.Join(workDir, "py")

			// 分别处理stdout和stderr
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err = cmd.Run()

			if err == nil {
				// 输出stderr内容作为调试信息
				if stderr.Len() > 0 {
					l.Infof("Python script stderr output: %s", stderr.String())
				}
				return l.processPythonOutput(stdout.Bytes(), orderData, w, pythonCmd)
			}

			lastErr = err
			l.Errorf("Failed to execute with %s: %v, stderr: %s", pythonCmd, err, stderr.String())
		}

		return fmt.Errorf("python not found in PATH. Please install Python or configure PythonPath in config. Last error: %v", lastErr)
	}

	// 使用配置的Python路径
	cmd := exec.Command(pythonPath, pythonScript, string(jsonData))
	cmd.Dir = filepath.Join(workDir, "py")

	// 分别处理stdout和stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	if err != nil {
		stderrStr := stderr.String()
		l.Errorf("Failed to execute with configured python path %s: %v, stderr: %s", pythonPath, err, stderrStr)
		return fmt.Errorf("failed to execute python script with configured path %s: %v, stderr: %s", pythonPath, err, stderrStr)
	}

	// 输出stderr内容作为调试信息
	if stderr.Len() > 0 {
		l.Infof("Python script stderr output: %s", stderr.String())
	}

	return l.processPythonOutput(stdout.Bytes(), orderData, w, pythonPath)
}

// processPythonOutput 处理Python脚本的输出
func (l *DownloadPdfLogic) processPythonOutput(output []byte, orderData *PDFOrderData, w http.ResponseWriter, pythonCmd string) error {
	// 成功执行，获取PDF文件路径
	outputStr := strings.TrimSpace(string(output))

	if len(outputStr) == 0 {
		l.Errorf("Python script returned empty output")
		return fmt.Errorf("python script returned empty output")
	}

	// 检查是否包含错误信息
	if strings.Contains(outputStr, "错误") || strings.Contains(outputStr, "Error") || strings.Contains(outputStr, "Exception") {
		l.Errorf("Python script output contains error information: %s", outputStr)
		return fmt.Errorf("python script failed with error: %s", outputStr)
	}

	// 检查文件是否存在
	if _, err := os.Stat(outputStr); os.IsNotExist(err) {
		l.Errorf("Generated PDF file does not exist: %s", outputStr)
		return fmt.Errorf("generated PDF file does not exist: %s", outputStr)
	}

	// 读取PDF文件内容
	pdfBytes, err := os.ReadFile(outputStr)
	if err != nil {
		l.Errorf("Failed to read PDF file %s: %v", outputStr, err)
		return fmt.Errorf("failed to read PDF file: %v", err)
	}

	if len(pdfBytes) == 0 {
		return fmt.Errorf("PDF file is empty: %s", outputStr)
	}

	// 验证PDF文件头
	if len(pdfBytes) < 4 || string(pdfBytes[:4]) != "%PDF" {
		return fmt.Errorf("invalid PDF file format")
	}

	// 设置响应头 - 在写入任何内容之前设置
	filename := fmt.Sprintf("Production Document %s-%s TO NBO.pdf", orderData.Po, orderData.StyleNum)
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", filename))

	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition, Content-Length, Content-Type")

	// 直接写入PDF内容
	_, err = w.Write(pdfBytes)
	if err != nil {
		l.Errorf("Failed to write PDF to response: %v", err)
		return fmt.Errorf("failed to write PDF to response: %v", err)
	}

	// 删除临时PDF文件
	err = os.Remove(outputStr)
	if err != nil {
		l.Errorf("Failed to delete temporary PDF file %s: %v", outputStr, err)
		// 不返回错误，因为PDF已经成功返回给前端
	}

	return nil
}
