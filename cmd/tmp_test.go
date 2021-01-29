package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/SmallTianTian/go-tools/slice"
)

const dir = "/Users/tick/WorkSpace/eclipse-workspace/ding/src/main/java/com/dingtalk/api"
const target = "../model"

func Test_main(t *testing.T) {
	reqFiles := listRequestFiles("request")
	respFiles := listRequestFiles("response")
	for _, v := range reqFiles[1:] {
		reqlines := ReadTxtFileEachLine(filepath.Join(dir, "request", v))
		req := ParseRequest(reqlines)
		resplines := ReadTxtFileEachLine(filepath.Join(dir, "response", strings.ReplaceAll(v, "Request", "Response")))
		resp := ParseResp(resplines)

	}
}

func listRequestFiles(path string) (files []string) {
	err := filepath.Walk(filepath.Join(dir, path), func(path string, info os.FileInfo, err error) error {
		if name := info.Name(); strings.HasSuffix(name, "java") {
			files = append(files, name)
		}
		return nil
	})
	MustNotError(err)
	return
}

func ReadTxtFileEachLine(path string) (lines []string) {
	r := bufio.NewReader(bytes.NewBuffer(ReadFile(path)))
	for {
		line, _, err := r.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		MustNotError(err)
		lines = append(lines, string(line))
	}
	return
}

func ReadFile(path string) []byte {
	bs, err := ioutil.ReadFile(path)
	MustNotError(err)
	return bs
}

func MustNotError(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	ppp = []string{"privat", "protec", "public"}
)

func ParseRequest(lines []string) *Request {
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) < 6 {
			continue
		}
		if slice.StringInSlice(ppp, line[:6]) {

		}
	}

	return nil
}

func ParseResp(lines []string) *Request {
	return nil
}

type Method struct {
	Name   string
	Params []string
	Ret    string
	Body   string
}

type Request struct {
	Field  []string
	Method []*Method
}

func (req *Request) ToStruct(name string) {
	sb := strings.Builder{}
	sb.WriteString("type ")
	sb.WriteString(name)
	sb.WriteString("struct {\n")
	for _, f := range req.Field {
		sb.WriteString(f)
		sb.WriteString("\n")
	}
	sb.WriteString("}\n")
	for _, m := range req.Method {
		sb.WriteString(fmt.Sprintf("func (this *%s) %s(%s)(%s){\n", name, m.Name, strings.Join(m.Params, m.Ret)))
		sb.WriteString(m.Body)
		sb.WriteString("}")
	}
}

type Response struct {
	File   []string
	Json   []string
	Method []*Method
}

func (req *Response) ToStruct(name string) {

}
