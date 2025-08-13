package main

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/Knetic/govaluate"
	"github.com/golang-module/carbon/v2"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	//fmt.Println(slice2)

	copy(slice1, slice2)
	fmt.Println(slice1)

	var a uint16 = 0
	a = a | (1 << 15)
	a = a | (1 << 13)
	a = a | (1 << 0)
	fmt.Println(fmt.Sprintf("%x", a)) //a输出结果:[a001]

	demo := &RegisterTimeChecker{}
	fmt.Println(demo.GetKey())

	var q interface{}
	q = 2
	s, ok := q.(bool)
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println("not bool")
	}

	fmt.Println(demo)
	fmt.Println(demo.Id)

	type B struct {
		Id string `json:"id"`
	}
	type C struct {
		Id string `json:"id"`
	}

	type A struct {
		B B
		C C
	}
	s1 := A{B{"b"}, C{"c"}}
	v, _ := json.Marshal(s1)
	fmt.Println(string(v))

	fmt.Println(carbon.CreateFromStdTime(time.Unix(1740758400, 0)).AddMonths(1).StdTime())

	type InviteRule struct {
		AwardRule []*C `json:"award_rule,omitempty"`
	}
	s2, _ := json.Marshal(&InviteRule{AwardRule: []*C{}})
	fmt.Println(string(s2))
	rule := &InviteRule{}
	json.Unmarshal([]byte(s2), rule)
	fmt.Println(fmt.Sprintf("%+v", rule))

	type LotteryChanceExecutorParams struct {
		Count int `json:"count"`
	}

	pa := LotteryChanceExecutorParams{Count: 3}
	s2, _ = json.Marshal(pa)
	fmt.Println(string(s2))

	params := map[string]interface{}{
		"xxx": "xxx",
	}
	_, ok = params["xxxxxx"].(string)
	fmt.Println(ok)

	var headers map[string]string
	for k, _ := range headers {
		fmt.Println(k)
	}
	credits := map[string]int{"xxx": 1}
	fmt.Println(credits["xxx"])
	fmt.Println(credits["xxxx"])

	fmt.Println(math.Ceil(12.131222323232323232*100) / 100)

	type RedPacketListResponse struct {
		Count       int     `json:"count"`
		TotalAmount float64 `json:"total_amount"`
	}
	s3, _ := json.Marshal(&RedPacketListResponse{TotalAmount: 123})
	fmt.Println(string(s3))
	type PacketListResponse struct {
		Count       int `json:"count"`
		TotalAmount int `json:"total_amount"`
	}
	s4 := &PacketListResponse{}
	json.Unmarshal([]byte(s3), s4)
	fmt.Println(fmt.Sprintf("%+v", s4))

	minUpdatedTime := carbon.Now().SetTimeMilli(0, 0, 0, 0).SubDay()
	maxUpdatedTime := carbon.Now()
	fmt.Println(fmt.Sprintf("%+v", maxUpdatedTime))
	fmt.Println(fmt.Sprintf("%+v", minUpdatedTime))
	fmt.Println("----")

	startTime := carbon.Now().SetMinute(0).SetSecond(0).SubHour()
	fmt.Println(fmt.Sprintf("%+v", startTime))
	startTime = startTime.SetHour(startTime.Hour() - 11)
	fmt.Println(fmt.Sprintf("%+v", startTime))
	endTime := startTime.AddHour()
	fmt.Println(fmt.Sprintf("%+v", endTime))
	fmt.Println(fmt.Sprintf("%+v", endTime.Hour()))

	fmt.Println(carbon.CreateFromTimestamp(0))

	fmt.Println(carbon.CreateFromTimestamp(time.Now().Unix()).SubDay().SetTimeMilli(0, 0, 0, 0))

	fmt.Println(carbon.CreateFromStdTime(carbon.CreateFromDateTime(2025, 7, 31, 0, 0, 0).StdTime()).AddMonthsNoOverflow(1).StdTime())

	prizeKey := "member_3d_hidden"
	memberPrefix := "member_"
	if !strings.HasPrefix(prizeKey, memberPrefix) {
		fmt.Println("请联系客服领取奖品")
	}
	parts := strings.SplitN(strings.TrimPrefix(prizeKey, "member_"), "_", 2)
	fmt.Println(parts)
	keyTmp := parts[0]
	fmt.Println(keyTmp)
	_, got := Find(GetCachedTodayUserTask(), func(t int) bool {
		return true
	})
	fmt.Println(got)

	parameters := make(map[string]interface{})

	parameters["a"] = true
	parameters["b"] = false
	parameters["c"] = true
	expr, err := govaluate.NewEvaluableExpression("(!a || !b) && c")

	result, err := expr.Evaluate(parameters)
	if err != nil {
		fmt.Println(fmt.Sprintf("expr evaluate failed: %+v: %+v", "a && !b && c", err))
	}
	fmt.Println(fmt.Sprintf("result:%+v", result))

	fmt.Println(time.Unix(0, 0))

	t := time.Time{}
	ts := &t
	fmt.Println(ts.IsZero())

	visitMap := make(map[string]bool)
	fmt.Println(visitMap["xxx"])

	sortDemo := []struct {
		GroupId string
	}{
		{GroupId: "k12_invite_task"},
		{GroupId: "k12_photo_question"},
	}

	sort.Slice(sortDemo, func(i, j int) bool {
		return sortDemo[i].GroupId > sortDemo[j].GroupId
	})
	fmt.Println(sortDemo)
}

func GetCachedTodayUserTask() []int {
	return nil
}

func Find[S ~[]T, T any](arr S, fn func(T) bool) (T, bool) {
	if idx := slices.IndexFunc(arr, fn); idx > -1 {
		return arr[idx], true
	}
	return *new(T), false
}

type RegisterTimeChecker struct {
	Id int `json:"id"`
}

func (c *RegisterTimeChecker) GetKey() string {
	return reflect.TypeOf(*c).Name()
}
