package common

import (
	"UBC/api/internal/svc"
	"UBC/api/internal/types"
	"UBC/api/library/xerr"
	"UBC/api/utils"
	"UBC/models/res_models"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"os/exec"
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

type Order struct {
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

type OrdersInfo struct {
	Orders []Order `json:"orders"`
}

func (l *UploadFileLogic) UploadFile(req *types.UploadFile, r *http.Request) (resp *types.UploadRes, err error) {
	_ = r.ParseMultipartForm(maxFileSize)
	file, h, err := r.FormFile("file")
	if err != nil {
		l.Errorf("[UploadFile] get file err:%v", err)
		return nil, xerr.RequestParamError
	}
	defer file.Close()

	text, err := utils.ExtractPDFText(file, h.Size)
	if err != nil {
		l.Errorf("[UploadFile] extractPDFText err:%v", err)
		return nil, xerr.RequestParamError

	}
	if req.UsedFor == "packing" {
		return l.doPackingFile(text)
	} else if req.UsedFor == "projection" {
		return l.duProjectionFile()
	}
	l.Errorf("[UploadFile] current use for not found:%s", req.UsedFor)
	return nil, xerr.RequestParamError

}

func (l *UploadFileLogic) doPackingFile(text string) (resp *types.UploadRes, err error) {
	cmd := exec.Command("python", "D:\\PycharmProjects\\order_project\\UBC\\packing.py", text)
	output, err := cmd.CombinedOutput()
	if err != nil {
		l.Errorf("[doPackingFile] exec packing python script err: %s", err)
		return nil, xerr.ServerCommonError
	}

	scanner := bufio.NewScanner(bytes.NewReader(output))
	var (
		jsonData                 string
		ordersData               OrdersInfo
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

	err = json.Unmarshal([]byte(jsonData), &ordersData)
	if err != nil {
		l.Errorf("[doPackingFile] Error parsing JSON: %s", err)
		return nil, xerr.RequestParamError
	}

	for _, order := range ordersData.Orders {
		newShipmenPackingRes := res_models.NewShipmenPackingRes{
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

func (l *UploadFileLogic) duProjectionFile() (resp *types.UploadRes, err error) {
	return nil, err
}
