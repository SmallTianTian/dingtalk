package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiUserListsimpleRequest() *OapiUserListsimpleRequest {
	return &OapiUserListsimpleRequest{}
}

type OapiUserListsimpleRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiUserListsimpleResponse
	ContainAccessLimit bool
	Cursor             int64
	DeptId             int64
	Language           string
	OrderField         string
	Size               int64
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiUserListsimpleRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiUserListsimpleRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserListsimpleRequest) SetContainAccessLimit(containAccessLimit2 bool) {
	this.ContainAccessLimit = containAccessLimit2
}
func (this *OapiUserListsimpleRequest) GetContainAccessLimit() bool {
	return this.ContainAccessLimit
}
func (this *OapiUserListsimpleRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiUserListsimpleRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiUserListsimpleRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiUserListsimpleRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiUserListsimpleRequest) SetLanguage(language2 string) {
	this.Language = language2
}
func (this *OapiUserListsimpleRequest) GetLanguage() string {
	return this.Language
}
func (this *OapiUserListsimpleRequest) SetOrderField(orderField2 string) {
	this.OrderField = orderField2
}
func (this *OapiUserListsimpleRequest) GetOrderField() string {
	return this.OrderField
}
func (this *OapiUserListsimpleRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiUserListsimpleRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiUserListsimpleRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.listsimple"
}
func (this *OapiUserListsimpleRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserListsimpleRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserListsimpleRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserListsimpleRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserListsimpleRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["contain_access_limit"] = this.ContainAccessLimit
	txtParams["cursor"] = this.Cursor
	txtParams["dept_id"] = this.DeptId
	txtParams["language"] = this.Language
	txtParams["order_field"] = this.OrderField
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiUserListsimpleRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *OapiUserListsimpleRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserListsimpleRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserListsimpleResponse struct {
	taobao.TaobaoResponse
	Errcode int64      `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Result  PageResult `json:"result,omitempty"`
}
type ListUserSimpleResponse struct {
	Name   string `json:"name,omitempty"`
	Userid string `json:"userid,omitempty"`
}
