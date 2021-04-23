package main

import (
	"encoding/json"
	"fmt"
)

type test struct {
	B int
	C float64
	D string
}

type BaseReq struct {
	PID       string `json:"pid"`
	BID       string `json:"bid"`
	EvType    string `json:"ev_type"`
	Protocol  string `json:"protocol"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

type AjaxTag struct {
	AxPath   string `json:"ax_path"`
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
		fmt.Println(string(v) == "{}") // true
	}

	var arrayTest []*test
	var pointerTest *test
	var structTest test

	fmt.Println("-------marshal Array-------")
	bytes, err := json.Marshal(arrayTest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bytes) // "null"
		fmt.Println(string(bytes))
	}

	fmt.Println("-------unmarshal Array-------")
	err = json.Unmarshal(bytes, &arrayTest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(arrayTest) // []
	}

	fmt.Println("-------marshal pointer-------")
	bytes, err = json.Marshal(pointerTest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bytes) // "null"
		fmt.Println(string(bytes))
	}

	fmt.Println("-------unmarshal pointer-------")
	err = json.Unmarshal(bytes, &pointerTest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pointerTest) // []
	}

	fmt.Println("-------marshal struct-------")
	bytes, err = json.Marshal(structTest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bytes)
		fmt.Println(string(bytes)) // {"B":0,"C":0,"D":""}
	}

	fmt.Println("-------unmarshal struct-------")
	err = json.Unmarshal(bytes, &structTest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(structTest) // {0 0 }
	}

	fmt.Println("-------empty string-------")
	emptyString := ""
	err = json.Unmarshal([]byte(emptyString), &arrayTest) // error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(arrayTest)
	}

	err = json.Unmarshal([]byte(emptyString), &pointerTest) // error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(arrayTest)
	}

	err = json.Unmarshal([]byte(emptyString), &structTest) // error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(arrayTest)
	}
	fmt.Println()
}

func TestSize() {
	s := "[{\"Id\":0,\"Subject\":7,\"Expressions\":[{\"Id\":0,\"Predicate\":{\"Type\":1,\"Negative\":false},\"Objects\":[{\"Id\":0,\"Mode\":1,\"Type\":2,\"Values\":[\"很敏感词\"],\"Ext\":{}},{\"Id\":0,\"Mode\":1,\"Type\":2,\"Values\":[\"很敏感词\"],\"Ext\":{}},{\"Id\":0,\"Mode\":1,\"Type\":2,\"Values\":[\"很敏感词\"],\"Ext\":{}},{\"Id\":0,\"Mode\":1,\"Type\":2,\"Values\":[\"很敏感词\"],\"Ext\":{}}],\"Logic\":1}],\"Logic\":0}]"
	fmt.Println(len(s))
}
