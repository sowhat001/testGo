package main

import (
	"encoding/json"
	"fmt"
)

type BaseReq struct {
	PID       string `json:"pid"`
	BID       string `json:"bid"`
	EvType    string `json:"ev_type"`
	Protocol  string `json:"protocol"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

type AjaxTag struct {
	AxPath   string `json:"ax_path"` // TODO 存在疑惑
	AxType   string `json:"ax_type"`
	AxStatus string `json:"ax_status"`
}

type AjaxReq struct {
	BaseReq
	AjaxTag // 这个tag不能写，写了就会嵌套json
}

func TestParseString() {
	m := make(map[int64]string)
	a := "{\"1\":\"\"}"
	// a := "{1:\"\",2:\"\", 14:\"\", 15:\"\"}"
	err := json.Unmarshal([]byte(a), &m)
	fmt.Println(m, err)
	fmt.Println(m[1])
	fmt.Println(len(m))
}

func TestParseMultiLayerStruct() {
	t := &AjaxReq{
		BaseReq: BaseReq{
			PID:       "tt",
			BID:       "vv",
			EvType:    "ee",
			Protocol:  "qweqwe",
			StartTime: 0,
			EndTime:   0,
		},
		AjaxTag: AjaxTag{
			AxPath:   "asd",
			AxType:   "qeg",
			AxStatus: "qweqwr",
		},
	}
	var res []byte
	res, err := json.Marshal(t)
	fmt.Println(string(res), err)
}

func TestParseEmpty() {
	a := make(map[int]int)
	v, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(string(v) == "{}") // true
	}
}
