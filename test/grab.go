package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/imroc/biu"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func revert(x int) int {
	if math.MinInt32 > x || x > math.MaxInt32 {
		return 0
	}

	newX := 0
	for x != 0 {
		mod := x % 10
		x = x / 10
		newX = newX*10 + mod
	}
	fmt.Println(newX)
	return newX
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])
	maxInt := Max(dp[0], dp[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = Max(nums[i]+dp[i-2], dp[i-1])
		maxInt = Max(maxInt, dp[i])
	}
	return maxInt
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type VehicleUserRelation struct {
	AuthID int64  `json:"authID"`
	UserID int64  `json:"userID"` // 用户ID
	VID    string `json:"vID"`    // 车辆ID
	Type   int64  `json:"type"`   // bang 0:取消，1:绑定，2:更新
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	//fmt.Println(slice2)

	copy(slice1, slice2)
	fmt.Println(slice1)
	test := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(test))

	dev := VehicleUserRelation{
		AuthID: 10,
		UserID: 1229353276036562125,
		VID:    "55c94e6f7e9f325be358c483cc0095c3",
		Type:   1,
	}
	put, _ := json.Marshal(dev)
	fmt.Println(string(put))
	str := "{\"userID\":1229353276036562125,\"vID\":\"55c94e6f7e9f325be358c483cc0095c3\",\"type\":1}"
	des := VehicleUserRelation{}
	json.Unmarshal([]byte(str), &des)
	fmt.Println(fmt.Sprintf("%+v", des))

	type MockKeyData struct {
		BncmSeid      string `json:"bncm_seid"`
		PhoneAuthPara string `json:"phone_auth_para"`
		BleRootKeyId  string `json:"ble_root_key_id"`
		Vin           string `json:"vin"`
		BluetoothMac  string `json:"bluetooth_mac"`
	}

	type data struct {
		Version  int64
		SlotData []byte
	}

	dataDemo := data{Version: time.Now().UnixNano() / 1e6}

	ver := time.Now().UnixNano() / 1e6
	ver = 1661485555111
	fmt.Println(ver)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, ver)

	slotDataDemo := bytebuf.Bytes()

	fmt.Println("11111")
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

	dataDemo.SlotData = slotDataDemo
	f, err := os.Create("slot")
	jsonData, _ := json.Marshal(dataDemo)
	f.Write(jsonData)

	fmt.Println("21111111")
	fmt.Println(dataDemo)

	mock := MockKeyData{
		BncmSeid:      "7a28c4802011954620580c6b83aba912",
		PhoneAuthPara: "aa76c80c",
		BleRootKeyId:  "0f68fd89ebee4b8aa5755bbee5a7d863",
		Vin:           "L6T79T2E1NP007313",
		BluetoothMac:  "43:6e:f3:7c:659c",
	}
	mstr, _ := json.Marshal(mock)
	fmt.Println(string(mstr))

	blePhoneKeyData := ISO9797M2Padding("LJ1E6A2U9MG074598", 16)
	fmt.Println(blePhoneKeyData)

	fmt.Println(fmt.Sprintf("%x", 1276309543245865541))

	originStr := "54a3fbb3088e52339c49c22c724a2fdb9491b5b05f7cf2fdafd97b42dafa95aedce11db9470fa459b858e8f741b199959750b55d70efe59f99349f6f9cb49acab68036538a28c04d5365ac19b01c41de95ddf7241c24317fdad8bacb7ebd190f" + strconv.FormatUint(1276309543245865541, 10) + "LJ1E6A2U9MG074598"
	keyId, _ := GetMd5StringHex(originStr)
	fmt.Println("---" + keyId)
	fmt.Println(StringToAsciiHex("LJ1E6A2U9MG074598"))
	fmt.Println(fmt.Sprintf("%x", 1658504357))
	fmt.Println(fmt.Sprintf("%x", 1665478999))
	de := ISO9797M2Padding("b0362e09cb245ffc11b65cc4749c6245d8e6ba9b2360404bd3c8a1facda5720d014c4a314536413255394d473037343539380000000000017246E8AA0C4A62dac4a5634531570000ffff00000000000000000000", 16)
	fmt.Println(de)
	fmt.Println(strings.Join(strings.Split("b4 9d 27 bd 11 4d b5 eb 82 89 84 5b ad 8d e6 90 95 99 8c 17 d9 46 04 d3 c0 79 ef ca 91 6c cf 71 09 63 34 c3 46 1d 24 a8 65 54 6b ba 64 97 7a e0 8b 33 b6 ce 49 ea 7b 05 3d 7d d3 b3 ab cc 0d be ce 31 5f fc 00 b4 d3 46 f6 e8 1a 37 a5 d1 92 2e 3f 47 c9 f5 48 0d 2b 14 b7 b7 c7 f4 de e1 9a 43", " "), ""))
	d := [16]byte{}
	fmt.Println(d[0:])

	type KeyListItem struct {
		Vid      string `json:"vid"`
		KeyID    string `json:"key_id"`
		PairData string `json:"pair_data"`
	}

	var keyList []KeyListItem
	err = json.Unmarshal([]byte("[{\"vid\":\"LKB3A21K0G5511098\",\"key_id\":\"58355fa0f6e54828b3287a41fbc94df7\"},{\"vid\":\"LKB3A21K0G5511099\",\"key_id\":\"58355fa0f6e54828b3287a41fbc94df8\"}]"), &keyList)
	fmt.Println(err)
	fmt.Println(fmt.Sprintf("%+v", keyList))

	var version int64
	version = 1660726003856
	fmt.Println(version)

	var index int32
	index = 21
	var uIndex uint32
	uIndex = 21

	fmt.Println(index == int32(uIndex))

	var res []*int
	res = nil
	fmt.Println(res)

	type KeyNotify struct {
		UserId string `json:"user_id"`
		VID    string `json:"vid"`
	}
	sendData := KeyNotify{
		UserId: "123",
		VID:    "123",
	}
	sendStr, _ := json.Marshal(sendData)
	fmt.Println(string(sendStr))

	var inputMap map[string]*struct{}

	fmt.Println(inputMap == nil)

	s, _ := hex.DecodeString("0003")
	fmt.Println(s)
	i := 0
	fmt.Println(i << 2)

	var a uint16 = 0

	a = a | (1 << 15)
	a = a | (1 << 13)
	a = a | (1 << 0)

	fmt.Println(biu.ToBinaryString(a)) //a输出结果:[10100000 00000001]
	fmt.Println(fmt.Sprintf("%x", a))  //a输出结果:[a001]

	var asd interface{}
	err = json.Unmarshal([]byte("[{\"vid\":\"LKB3A21K0G5511098\",\"key_id\":\"58355fa0f6e54828b3287a41fbc94df7\"},{\"vid\":\"LKB3A21K0G5511099\",\"key_id\":\"58355fa0f6e54828b3287a41fbc94df8\"}]"), &asd)
	fmt.Println(err)
	fmt.Println(fmt.Sprintf("%+v", asd))
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
