package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	//fmt.Println(slice2)

	copy(slice1, slice2)
	fmt.Println(slice1)

	ver := time.Now().UnixNano() / 1e6
	fmt.Println(ver)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, ver)
	slotDataDemo := bytebuf.Bytes()
	fmt.Println(slotDataDemo)
	for _, b := range slotDataDemo {
		fmt.Println(int8(b))
	}

	byteBuffer := bytes.NewBuffer(slotDataDemo)
	var da int64
	binary.Read(byteBuffer, binary.BigEndian, &da)
	fmt.Println(da)

	for i := 0; i < 16; i++ {
		d, _ := hex.DecodeString("01000011")
		slotDataDemo = append(slotDataDemo, d...)
	}

	ba := base64.StdEncoding.EncodeToString(slotDataDemo)

	fmt.Println(slotDataDemo)
	fmt.Println(len(slotDataDemo))
	fmt.Println(ba)

	blePhoneKeyData := ISO9797M2Padding("LJ1E6A2U9MG074598", 16)
	fmt.Println(blePhoneKeyData)
	fmt.Println(fmt.Sprintf("%x", 1276309543245865541))
	originStr := "54a3fbb3088e52339c49c22c724a2fdb9491b5b05f7cf2fdafd97b42dafa95aedce11db9470fa459b858e8f741b199959750b55d70efe59f99349f6f9cb49acab68036538a28c04d5365ac19b01c41de95ddf7241c24317fdad8bacb7ebd190f" + strconv.FormatUint(1276309543245865541, 10) + "LJ1E6A2U9MG074598"
	keyId, _ := GetMd5StringHex(originStr)
	fmt.Println("---" + keyId)
	fmt.Println(StringToAsciiHex("JD00M12209130008"))
	de := ISO9797M2Padding("b0362e09cb245ffc11b65cc4749c6245d8e6ba9b2360404bd3c8a1facda5720d014c4a314536413255394d473037343539380000000000017246E8AA0C4A62dac4a5634531570000ffff00000000000000000000", 16)
	fmt.Println(de)
	fmt.Println(strings.Join(strings.Split("b4 9d 27 bd 11 4d b5 eb 82 89 84 5b ad 8d e6 90 95 99 8c 17 d9 46 04 d3 c0 79 ef ca 91 6c cf 71 09 63 34 c3 46 1d 24 a8 65 54 6b ba 64 97 7a e0 8b 33 b6 ce 49 ea 7b 05 3d 7d d3 b3 ab cc 0d be ce 31 5f fc 00 b4 d3 46 f6 e8 1a 37 a5 d1 92 2e 3f 47 c9 f5 48 0d 2b 14 b7 b7 c7 f4 de e1 9a 43", " "), ""))
	d := [16]byte{}
	fmt.Println(d[0:])

	var a uint16 = 0
	a = a | (1 << 15)
	a = a | (1 << 13)
	a = a | (1 << 0)
	fmt.Println(fmt.Sprintf("%x", a)) //a输出结果:[a001]

	var sss uint64
	sss = 18446744073709551615
	fmt.Println(fmt.Sprintf("%x", sss))

	var trialDetail = &TrialAuditDetail{}
	fmt.Println(fmt.Sprintf("%+v", trialDetail))
	trialDetail.ProjectDetail.ProjectName = "123"

	fmt.Println(json.Marshal("{\"project_detail\":{\"business_line_name\":\"测试\",\"project_name\":\"11111\",\"project_type\":\"other\",\"demand_side\":\"quhongli\",\"label_product_operator\":\"quhongli\",\"project_manager\":\"quhongli\",\"task_name\":\"\",\"project_task_type\":\"单轮对话\",\"task_admin_list\":\"quhongli,baihao\",\"target_accuracy\":7,\"target_efficiency\":7,\"worker_comments\":\"111111\",\"assignee_user\":\"zhupeng\"},\"target_detail\":{\"target_accuracy\":7,\"final_accuracy\":8,\"target_efficiency\":7,\"final_efficiency\":9,\"worker_comments\":\"111111\",\"demand_comments\":\"hahaha\"}}"))

	fmt.Printf("%.2f", 0)

	delta := 2
	SubmitTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-11-28 15:01:01", time.Local)
	fmt.Println(fmt.Sprintf("%+v", SubmitTime))
	nowTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-11-27 12:01:01", time.Local)
	fmt.Println(fmt.Sprintf("%+v", nowTime))
	taskBeginAtDay := TruncateToMidnight(nowTime)

	deltaDuration := time.Duration(delta) * 24 * time.Hour
	ds := SubmitTime.Sub(taskBeginAtDay)
	fmt.Println(fmt.Sprintf("%+v", ds))

	packageBeginAtDay := taskBeginAtDay.Add(ds - ds%deltaDuration)
	packageEndAtDay := packageBeginAtDay.Add(deltaDuration)
	fmt.Println(packageBeginAtDay, packageEndAtDay)

	labelTimeSpanMap := make(map[int64]map[string]map[int64][]string)

	if labelOrders, has := labelTimeSpanMap[1]["xxx"][1]; has {
		fmt.Println("xxx")
	} else {
		fmt.Println(labelOrders)
	}
	//fmt.Println(calcSampleSizeByConfidenceUnLimit(float64(0.3), getZValue(99), 99.3))
	//fmt.Println(calcMarginOfErrorByConfidenceUnLimit(getZValue(99), 99.5, 200))
	//fmt.Println(calcMarginOfErrorByConfidenceUnLimit(getZValue(99), 80, 134))

	//fmt.Println(calcSampleSizeByConfidence(float64(0.3), getZValue(99), 99.3, 10000))
	//fmt.Println(calcSampleSizeByConfidence(float64(3), getZValue(95), 99, 10000))
	//fmt.Println(calcSampleSizeByConfidence(float64(3), getZValue(99), 95, 10000))
	//fmt.Println(calcSampleSizeByConfidence(float64(3), getZValue(95), 95, 10000))
	//fmt.Println(calcSampleSizeByConfidence(float64(3), getZValue(99), 90, 10000))
	//fmt.Println(calcSampleSizeByConfidence(float64(3), getZValue(95), 90, 10000))
	//fmt.Println(calcSampleSizeByConfidence(float64(3), getZValue(99), 85, 10000))
	//fmt.Println(calcSampleSizeByConfidence(float64(3), getZValue(95), 85, 10000))
	//fmt.Println(calcSampleSizeByConfidence(float64(3), getZValue(99), 80, 10000))
	fmt.Println(calcSampleSizeByConfidence(float64(2), getZValue(99), 99, 22))

	//fmt.Println(calcMarginOfErrorByConfidence(getZValue(99), 99.5, 500, 10000))
	//fmt.Println(calcMarginOfErrorByConfidence(getZValue(95), 99.5, 500, 10000))

	//fmt.Println(calcMarginOfErrorByConfidence(getZValue(99), 99, 264, 10000))
	//fmt.Println(calcMarginOfErrorByConfidence(getZValue(99), 95, 500, 10000))
	//fmt.Println(calcMarginOfErrorByConfidence(getZValue(95), 95, 500, 10000))
	//fmt.Println(calcMarginOfErrorByConfidence(getZValue(99), 90, 500, 10000))
	//fmt.Println(calcMarginOfErrorByConfidence(getZValue(95), 90, 500, 10000))
	//fmt.Println(calcMarginOfErrorByConfidence(getZValue(99), 85, 500, 10000))
	//fmt.Println(calcMarginOfErrorByConfidence(getZValue(95), 85, 500, 10000))
	//fmt.Println(calcMarginOfErrorByConfidence(getZValue(99), 80, 500, 10000))
	//fmt.Println(calcMarginOfErrorByConfidence(getZValue(95), 99.5, 133, 100000))

	slice := []int64{}
	randShuffle(slice)
	fmt.Println(slice)
}

func randShuffle(slice []int64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

//e 置信区间 z置信度 p 准确率  n 样本量
func calcSampleSizeByConfidenceUnLimit(e, z, p float64) int64 {
	e = e / 100
	p = p / 100
	a := math.Pow(z, 2) * p * (1 - p) / math.Pow(e, 2)
	result := int64(math.Ceil(a))
	return result
}

//e 置信区间 z置信度 p 准确率  n 样本量
func calcMarginOfErrorByConfidenceUnLimit(z, p float64, n int64) float64 {
	if n == 1 {
		return 1
	}

	p = p / 100
	e2 := (math.Pow(z, 2) * p * (1 - p)) / float64(n)
	result := math.Sqrt(e2)
	return result
}

//e 置信区间 z置信度 p 准确率  n 样本量
func calcSampleSizeByConfidence(e, z, p float64, n int64) int64 {
	fn := float64(n)
	e = e / 100
	p = p / 100
	a := math.Pow(z, 2) * p * (1 - p) / math.Pow(e, 2)
	b := 1 + ((math.Pow(z, 2)*p*(1-p))/math.Pow(e, 2)-1)/fn
	result := int64(math.Ceil(a / b))
	return result
}

//e 置信区间 z置信度 p 准确率  n 样本量
func calcMarginOfErrorByConfidence(z, p float64, n int64, N int64) float64 {
	if n == 1 {
		return 1
	}

	fn := float64(n)
	p = p / 100
	a := math.Sqrt(p * (1 - p) / fn)
	b := math.Sqrt(1 - fn/float64(N))

	result := z * a * b
	return result
}

func getZValue(confidenceLevel int) float64 {
	switch confidenceLevel {
	case 80:
		return 1.28
	case 85:
		return 1.44
	case 90:
		return 1.65
	case 95:
		return 1.65
	case 99:
		return 2.58
	}
	return 0
}

func ISO9797M2Padding(origin string, n int) string {
	origin = origin + "80"
	originByteLen := len(origin) / 2

	mod := originByteLen % n
	for i := 0; i < n-mod; i++ {
		origin += "00"
	}
	return origin
}

func GetMd5StringHex(str string) (string, error) {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(m.Sum(nil)), nil
}

func StringToAsciiHex(str string) string {
	result := ""
	for i := 0; i < len(str); i++ {
		intV := int(str[i])
		hexV := fmt.Sprintf("%x", intV)
		result += hexV
	}
	return result
}

type TrialAuditOrder struct {
	ProjectId    int64             `json:"project_id"`
	TaskId       int64             `json:"task_id,omitempty"`
	AssigneeUser string            `json:"assignee_user,omitempty"`
	AuditDetail  *TrialAuditDetail `json:"audit_detail"`
}

type TrialAuditDetail struct {
	ProjectDetail ProjectDetail `json:"project_detail"`
	TrialDetail   *TrialDetail  `json:"trial_detail,omitempty"`
	TargetDetail  TargetDetail  `json:"target_detail"`
}

type ProjectDetail struct {
	BusinessLineName     string `json:"business_line_name"`
	ProjectName          string `json:"project_name"`
	ProjectType          string `json:"project_type"`
	DemandSide           string `json:"demand_side"`
	LabelProductOperator string `json:"label_product_operator"`
	ProjectManager       string `json:"project_manager"`
	TaskName             string `json:"task_name"`
	ProjectTaskType      string `json:"project_task_type"`
	TaskAdminList        string `json:"task_admin_list"`
}

type TrialDetail struct {
	DemandSide *EfficiencyDetail `json:"demand_side"`
	WorkerSide *EfficiencyDetail `json:"worker_side"`
}

type EfficiencyDetail struct {
	LabelUser                string  `json:"label_user"`
	LabelTotalTime           float64 `json:"label_total_time"`
	LabelTotalNum            int64   `json:"label_total_num"`
	LabelAvgEfficiency       float64 `json:"label_avg_efficiency"`
	LabelFirstHourEfficiency float64 `json:"label_first_hour_efficiency"`
	LabelLastHourEfficiency  float64 `json:"label_last_hour_efficiency"`
	LabelAccuracy            float64 `json:"label_accuracy"`
}

type TargetDetail struct {
	TargetAccuracy   float64 `json:"target_accuracy"`
	FinalAccuracy    float64 `json:"final_accuracy,omitempty"`
	TargetEfficiency float64 `json:"target_efficiency"`
	FinalEfficiency  float64 `json:"final_efficiency,omitempty"`
	WorkerComments   string  `json:"worker_comments"`
	DemandComments   string  `json:"demand_comments,omitempty"`
}

func TruncateToMidnight(t time.Time) time.Time {
	// use system location
	location := time.Now().Location()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, location)
}

func must(err any, messageArgs ...interface{}) {
	if err == nil {
		return
	}

	switch e := err.(type) {
	case bool:
		if !e {
			message := messageFromMsgAndArgs(messageArgs...)
			if message == "" {
				message = "not ok"
			}

			panic(message)
		}

	case error:
		message := messageFromMsgAndArgs(messageArgs...)
		if message != "" {
			panic(message + ": " + e.Error())
		} else {
			panic(e.Error())
		}

	default:
		panic("must: invalid err type '" + reflect.TypeOf(err).Name() + "', should either be a bool or an error")
	}
}

// Must is a helper that wraps a call to a function returning a value and an error
// and panics if err is error or false.
// Play: https://go.dev/play/p/TMoWrRp3DyC
func Must[T any](val T, err any, messageArgs ...interface{}) T {
	must(err, messageArgs...)
	return val
}
func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 1 {
		if msgAsStr, ok := msgAndArgs[0].(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msgAndArgs[0])
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}
