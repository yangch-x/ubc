package projectionPo

import (
	"UBC/api/library/xerr"
	"UBC/models"
	"context"
	"gorm.io/datatypes"
	"reflect"

	"UBC/api/internal/svc"
	"UBC/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFieldsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFieldsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFieldsLogic {
	return &UpdateFieldsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFieldsLogic) UpdateFields(req *types.UpdateProjectionPoFields) error {
	// 字段映射表：前端字段名 -> 数据库字段名
	fieldMapping := map[string]string{
		"arriveDt":       "arrive_dt",
		"customerCode":   "customer_code",
		"customerPo":     "customer_po",
		"styleCode":      "style_code",
		"styleName":      "style_name",
		"color":          "color",
		"fabrication":    "fabrication",
		"poQty":          "po_qty",
		"costPrice":      "cost_price",
		"ttlBuy":         "ttl_buy",
		"salePrice":      "sale_price",
		"ttlSell":        "ttl_sell",
		"exporter":       "exporter",
		"waterResistant": "notes", // 临时映射到notes字段，如果有专门的字段请修改
		"notes":          "notes",
		"country":        "country",
		"poItems":        "po_items",
	}

	// 构建更新字段映射
	updates := make(map[string]interface{})

	reqValue := reflect.ValueOf(req).Elem()
	reqType := reflect.TypeOf(req).Elem()

	for i := 0; i < reqValue.NumField(); i++ {
		field := reqType.Field(i)
		value := reqValue.Field(i)

		// 跳过id字段
		if field.Name == "Id" {
			continue
		}

		// 获取json标签名
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			continue
		}

		// 移除optional标签
		fieldName := jsonTag
		if len(jsonTag) > 0 && jsonTag[len(jsonTag)-9:] == ",optional" {
			fieldName = jsonTag[:len(jsonTag)-9]
		}

		// 检查字段是否有值（非零值）
		if !value.IsZero() {
			if dbField, exists := fieldMapping[fieldName]; exists {
				// 特殊处理poItems字段，转换为JSON类型
				if fieldName == "poItems" {
					updates[dbField] = datatypes.JSON(value.String())
				} else {
					updates[dbField] = value.Interface()
				}
			}
		}
	}

	// 调用模型的UpdateFields方法
	projection := &models.ProjectionPo{}
	err := projection.UpdateFields(req.Id, updates)
	if err != nil {
		l.Errorf("[UpdateFields] err:%v, id:%d, updates:%+v", err, req.Id, updates)
		return xerr.ServerCommonError
	}

	l.Infof("[UpdateFields] success, id:%d, updated fields:%+v", req.Id, updates)
	return nil
}
