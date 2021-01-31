package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"time"
)

func NewOapiSmartdevicePrintdetailGetRequest() *OapiSmartdevicePrintdetailGetRequest {
	return &OapiSmartdevicePrintdetailGetRequest{}
}

type OapiSmartdevicePrintdetailGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdevicePrintdetailGetResponse
	Cursor          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdevicePrintdetailGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdevicePrintdetailGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdevicePrintdetailGetRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiSmartdevicePrintdetailGetRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiSmartdevicePrintdetailGetRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiSmartdevicePrintdetailGetRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiSmartdevicePrintdetailGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.printdetail.get"
}
func (this *OapiSmartdevicePrintdetailGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdevicePrintdetailGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdevicePrintdetailGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdevicePrintdetailGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdevicePrintdetailGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiSmartdevicePrintdetailGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartdevicePrintdetailGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdevicePrintdetailGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdevicePrintdetailGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  PageVO `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
type PrintDetailVO struct {
	DeptExtInfo    string    `json:"dept_ext_info,omitempty"`
	DeptFullName   string    `json:"dept_full_name,omitempty"`
	DeptLevel1Name string    `json:"dept_level1_name,omitempty"`
	DeptLevel2Name string    `json:"dept_level2_name,omitempty"`
	DeptLevel3Name string    `json:"dept_level3_name,omitempty"`
	Origin         string    `json:"origin,omitempty"`
	PageColorType  int64     `json:"page_color_type,omitempty"`
	PageDoubleType int64     `json:"page_double_type,omitempty"`
	PageSizeType   string    `json:"page_size_type,omitempty"`
	Pages          int64     `json:"pages,omitempty"`
	PrintDate      time.Time `json:"print_date,omitempty"`
	PrinterNick    string    `json:"printer_nick,omitempty"`
	UserName       string    `json:"user_name,omitempty"`
}
