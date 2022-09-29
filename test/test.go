package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/imroc/biu"
	"io"
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
	fmt.Println(biu.ToBinaryString(a)) //a输出结果:[10100000 00000001]
	fmt.Println(fmt.Sprintf("%x", a))  //a输出结果:[a001]
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
