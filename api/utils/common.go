package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/protobuf/types/known/structpb"
	"gopkg.in/mgo.v2/bson"
	"math"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// 将float64转成精确的int64
func Wrap(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}

// 将int64恢复成正常的float64
func Unwrap(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

// 精准float64
func WrapToFloat64(num float64, retain int) float64 {
	return num * math.Pow10(retain)
}

// 精准int64
func UnwrapToInt64(num int64, retain int) int64 {
	return int64(Unwrap(num, retain))
}

func ParseInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	} else {
		return num
	}
}

func FormatInt(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

func FormatInt64(num int64) string {
	return strconv.FormatInt(num, 10)
}

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func StructToMapStr(obj interface{}) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	data := make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).String()
	}
	return data
}
func InterfaceToMap(data interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// 使用反射获取 interface{} 的值
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		// 如果是指针，则获取指针指向的实际值
		value = value.Elem()
	}

	// 检查类型是否为结构体
	if value.Kind() != reflect.Struct {
		return nil, fmt.Errorf("Input is not a struct")
	}

	// 获取结构体的字段数量
	numFields := value.NumField()

	// 遍历结构体的字段，并将其添加到 map 中
	for i := 0; i < numFields; i++ {
		field := value.Type().Field(i)
		fieldValue := value.Field(i).Interface()
		result[field.Name] = fieldValue
	}

	return result, nil
}

func CdnUrl(cdn, srcUrl string) string {
	if (cdn == "") || (srcUrl == "") {
		return srcUrl
	}
	u, err := url.Parse(srcUrl)
	if err == nil {
		return cdn + u.Path
	}
	return srcUrl
}

func ParseInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	} else {
		return num
	}
}

func WeakDecode(input, output interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		WeaklyTypedInput: true,
		DecodeHook:       mapstructure.ComposeDecodeHookFunc(StringToTimeHookFunc()),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func StringToTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		// Convert it by parsing
		return time.Unix(ParseInt64(data.(string)), 0), nil
	}
}

func DivStr(s string) []int {
	list := make([]int, 0)
	strs := strings.Split(s, ",")
	for _, str := range strs {
		id := ParseInt(str)
		list = append(list, id)
	}
	return list
}

func IsNumber(str string) bool {
	match, _ := regexp.MatchString(`^[\+-]?\d+$`, str)
	return match
}

// BytesToString converts byte slice to string.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes converts string to byte slice.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func GenSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret

	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}

func ToAny(data, s interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, s); err != nil {
		return err
	}
	return nil
}

func ByteToSt(data []byte) *structpb.Struct {
	jsonData := &structpb.Struct{}
	_ = jsonpb.UnmarshalString(string(data), jsonData)
	return jsonData
}

// StructSliceToMap struct切片转为map切片
func StructSliceToMap(source interface{}, filedName string) map[interface{}][]interface{} {
	filedIndex := 0
	v := reflect.ValueOf(source) // 判断，interface转为[]interface{}
	if v.Kind() != reflect.Slice {
		panic("ERROR: Unknown type, slice expected.")
	}
	l := v.Len()
	retList := make([]interface{}, l)
	for i := 0; i < l; i++ {
		retList[i] = v.Index(i).Interface()
	}
	if len(retList) > 0 {
		firstObj := retList[0]
		objT := reflect.TypeOf(firstObj)
		for i := 0; i < objT.NumField(); i++ {
			if objT.Field(i).Name == filedName {
				filedIndex = i
			}
		}
	}

	resMap := make(map[interface{}][]interface{})
	for _, elem := range retList {
		key := reflect.ValueOf(elem).Field(filedIndex).Interface()
		value := make([]interface{}, 0)
		resMap[key] = value
	}

	for _, elem := range retList {
		key := reflect.ValueOf(elem).Field(filedIndex).Interface()
		resMap[key] = append(resMap[key], elem)
	}
	return resMap
}

// StructsToBytes 将 []*structpb.Struct 序列化成字节数组
func StructsToBytes(structs []*structpb.Struct) ([]byte, error) {
	structList := &structpb.ListValue{Values: make([]*structpb.Value, len(structs))}
	for i, s := range structs {
		structList.Values[i] = &structpb.Value{Kind: &structpb.Value_StructValue{StructValue: s}}
	}
	return proto.Marshal(structList)
}

// ConvertBytesToStructs 反序列化
func ConvertBytesToStructs(bytes []byte) []*structpb.Struct {
	structList := &structpb.ListValue{}
	_ = jsonpb.UnmarshalString(string(bytes), structList)

	structs := make([]*structpb.Struct, len(structList.Values))
	for i, v := range structList.Values {
		if structValue, ok := v.GetKind().(*structpb.Value_StructValue); ok {
			structs[i] = structValue.StructValue
		}
	}

	return structs
}

func GetUUID() string {
	return uuid.New().String()
}

func GetMongoUUID() string {
	byteID := bson.NewObjectId()
	return hex.EncodeToString([]byte(byteID))
}

// GetNonEmptyFields 获取不为空的字段
func GetNonEmptyFields(obj interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if value.IsValid() && !isEmpty(value) {
			tag := field.Tag.Get("sqltag")
			m[tag] = value.Interface()
		}
	}
	return m
}

// isEmpty false 表示不为空
func isEmpty(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.String() == ""
	case reflect.Slice:
		return value.IsNil()
	case reflect.Struct:
		return checkEmptySliceFieldsRecursive(value)
	default:
		zero := reflect.Zero(value.Type())
		return value.Interface() == zero.Interface()
	}
}

func checkEmptySliceFieldsRecursive(v reflect.Value) bool {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !isEmpty(field) {
			return false
		}
	}
	return true
}

func ToMapAny(source map[string]interface{}) (map[string]*any.Any, error) {
	rem := make(map[string]*any.Any)
	for k, v := range source {
		jsonParams, err := json.Marshal(v)
		if err != nil {
			return rem, err
		}
		rem[k] = &any.Any{Value: jsonParams}
	}
	return rem, nil
}

func ContainsSlice(s, sub []string) bool {
	sort.Strings(s)
	sort.Strings(sub)
	i, j := 0, 0
	for i < len(s) && j < len(sub) {
		if s[i] < sub[j] {
			i++
		} else if strings.TrimSpace(s[i]) == strings.TrimSpace(sub[j]) {
			i++
			j++
		} else {
			return false
		}
	}
	return j == len(sub)
}

// StructToBytes 将结构体转换为字节切片
func StructToBytes(data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func MapToJSONString(inputMap interface{}) (string, error) {
	jsonString, err := json.Marshal(inputMap)
	if err != nil {
		return "", err
	}
	return string(jsonString), nil
}

func TransformFeedback(inputJSON map[string]interface{}) map[string]interface{} {
	// 只有status和createAt
	if len(inputJSON) == 2 {
		return map[string]interface{}{
			"overallScore": []interface{}{},
			"score":        0,
			"feedbacks":    []interface{}{},
		}
	}

	feedback := make(map[string]interface{})

	for k, v := range inputJSON {
		feedback[strings.ToLower(k)] = v
	}

	scoreMapping := map[string]string{
		"expression":  "📖 Expression",
		"language":    "🧠 Language",
		"personality": "🤔 Personality",
		"reaction":    "❤ Reaction",
	}

	score := make([]map[string]interface{}, 0)
	for k, v := range feedback {
		if title, ok := scoreMapping[k]; ok {
			score = append(score, map[string]interface{}{
				"title": title,
				"score": v,
			})
		}
	}

	feedbacks := make([]map[string]interface{}, 0)
	for title := range scoreMapping {
		advantagesKey := strings.ToLower(title) + " advantages"
		suggestionsKey := strings.ToLower(title) + " suggestions"

		feedbacks = append(feedbacks, map[string]interface{}{
			"title":       title,
			"advantages":  feedback[advantagesKey],
			"suggestions": feedback[suggestionsKey],
		})
	}

	output := map[string]interface{}{
		"overallScore": feedback["overall score"],
		"score":        score,
		"feedbacks":    feedbacks,
	}

	return output
}

func TransformFeedbacks(inputJSON map[string]interface{}) map[string]interface{} {
	// 只有status和createAt
	if len(inputJSON) == 2 {
		return map[string]interface{}{
			"overallScore": []interface{}{},
			"score":        0,
			"feedbacks":    []interface{}{},
		}
	}

	feedback := make(map[string]interface{})

	for k, v := range inputJSON {
		if val, ok := v.(string); ok {
			// 如果是字符串，尝试转换为数字
			if num, err := strconv.Atoi(val); err == nil {
				feedback[strings.ToLower(k)] = num
			} else {
				feedback[strings.ToLower(k)] = val
			}
		} else {
			feedback[strings.ToLower(k)] = v
		}
	}

	scoreMapping := map[string]string{
		"expression":  "📖 Expression",
		"language":    "🧠 Language",
		"personality": "🤔 Personality",
		"reaction":    "❤ Reaction",
	}

	score := make([]map[string]interface{}, 0)
	for k, v := range feedback {
		if title, ok := scoreMapping[k]; ok {
			score = append(score, map[string]interface{}{
				"title": title,
				"score": v,
			})
		}
	}

	feedbacks := make([]map[string]interface{}, 0)
	for title := range scoreMapping {
		advantagesKey := strings.ToLower(title) + " advantages"
		suggestionsKey := strings.ToLower(title) + " suggestions"

		feedbacks = append(feedbacks, map[string]interface{}{
			"title":       title,
			"advantages":  feedback[advantagesKey],
			"suggestions": feedback[suggestionsKey],
		})
	}

	output := map[string]interface{}{
		"overallScore": feedback["overall score"],
		"score":        score,
		"feedbacks":    feedbacks,
	}

	return output
}

func TransformFeedbacks2(inputJSON map[string]interface{}) map[string]interface{} {
	// 只有status和createAt
	if len(inputJSON) == 2 {
		return map[string]interface{}{
			"overallScore": []interface{}{},
			"score":        0,
			"feedbacks":    []interface{}{},
		}
	}

	feedback := make(map[string]interface{})

	for k, v := range inputJSON {
		if val, ok := v.(string); ok {
			// 如果是字符串，尝试转换为数字
			if strings.Contains(val, "/") {
				n := ExtractNumber(val)
				feedback[strings.ToLower(k)] = n
				if n == -1 {
					feedback[strings.ToLower(k)] = val
				}
			} else if num, err := strconv.Atoi(val); err == nil {
				feedback[strings.ToLower(k)] = num
			} else {
				feedback[strings.ToLower(k)] = val
			}
		} else {
			feedback[strings.ToLower(k)] = v
		}
	}

	scoreMapping := map[string]string{
		"expression":  "📖 Expression",
		"language":    "🧠 Language",
		"personality": "🤔 Personality",
		"reaction":    "❤ Reaction",
	}

	score := make([]map[string]interface{}, 0)
	for k, v := range feedback {
		if title, ok := scoreMapping[k]; ok {
			score = append(score, map[string]interface{}{
				"title": title,
				"score": v,
			})
		}
	}

	feedbacks := make([]map[string]interface{}, 0)
	for title := range scoreMapping {
		advantagesKey := strings.ToLower(title) + " advantages"
		suggestionsKey := strings.ToLower(title) + " suggestions"

		feedbacks = append(feedbacks, map[string]interface{}{
			"title":       title,
			"advantages":  feedback[advantagesKey],
			"suggestions": feedback[suggestionsKey],
		})
	}

	output := map[string]interface{}{
		"overallScore": feedback["overall score"],
		"score":        score,
		"feedbacks":    feedbacks,
	}

	return output
}

func ExtractNumber(scoreStr string) int {
	re := regexp.MustCompile(`(\d+)`)
	match := re.FindStringSubmatch(scoreStr)
	if len(match) > 1 {
		num, err := strconv.Atoi(match[1])
		if err == nil {
			return num
		}
	}
	return -1
}
