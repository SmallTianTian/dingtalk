package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"time"
)

func NewOapiSmartdeviceFocusdetailGetRequest() *OapiSmartdeviceFocusdetailGetRequest {
	return &OapiSmartdeviceFocusdetailGetRequest{}
}

type OapiSmartdeviceFocusdetailGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceFocusdetailGetResponse
	Cursor          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceFocusdetailGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceFocusdetailGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceFocusdetailGetRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiSmartdeviceFocusdetailGetRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiSmartdeviceFocusdetailGetRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiSmartdeviceFocusdetailGetRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiSmartdeviceFocusdetailGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.focusdetail.get"
}
func (this *OapiSmartdeviceFocusdetailGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceFocusdetailGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceFocusdetailGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceFocusdetailGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceFocusdetailGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiSmartdeviceFocusdetailGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartdeviceFocusdetailGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceFocusdetailGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceFocusdetailGetResponse struct {
	taobao.TaobaoResponse
	Result  PageVO `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
type FocusDetailVO struct {
	CreateDate   time.Time `json:"create_date,omitempty"`
	DeptFullName string    `json:"dept_full_name,omitempty"`
	NickName     string    `json:"nick_name,omitempty"`
	Room         string    `json:"room,omitempty"`
	UserName     string    `json:"user_name,omitempty"`
}
type PageVO struct {
	HasMore          bool            `json:"has_more,omitempty"`
	List             []FocusDetailVO `json:"list,omitempty"`
	NextCursor       int64           `json:"next_cursor,omitempty"`
	NextCursorString string          `json:"next_cursor_string,omitempty"`
}
