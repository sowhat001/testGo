package main

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
	_ "net/http/pprof"

	"github.com/sirupsen/logrus"
)

func TestHttpPProf() {
	http.HandleFunc("/test", handler)
	err := http.ListenAndServe(":4396", nil)
	if err != nil {
		logrus.Fatal(err)
	}
}

func handler(writer http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		_, _ = writer.Write([]byte(err.Error()))
		return
	}
	doSomeThingOne(10000)
	buff := genSomeBytes()
	b, err := ioutil.ReadAll(buff)
	if nil != err {
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Write(b)
}

func doSomeThingOne(times int) {
	for i := 0; i < times; i++ {
		for j := 0; j < times; j++ {

		}
	}
}

func genSomeBytes() *bytes.Buffer {
	var buff bytes.Buffer
	for i := 1; i < 20000; i++ {
		buff.Write([]byte{'0' + byte(rand.Intn(10))})
		buff.Write([]byte{' '})
	}
	return &buff
}
