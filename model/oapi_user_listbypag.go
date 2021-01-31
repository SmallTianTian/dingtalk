package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiUserListbypageRequest() *OapiUserListbypageRequest {
	return &OapiUserListbypageRequest{}
}

type OapiUserListbypageRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserListbypageResponse
	DepartmentId    int64
	Lang            string
	Offset          int64
	Order           string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiUserListbypageRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiUserListbypageRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserListbypageRequest) SetDepartmentId(departmentId2 int64) {
	this.DepartmentId = departmentId2
}
func (this *OapiUserListbypageRequest) GetDepartmentId() int64 {
	return this.DepartmentId
}
func (this *OapiUserListbypageRequest) SetLang(lang2 string) {
	this.Lang = lang2
}
func (this *OapiUserListbypageRequest) GetLang() string {
	return this.Lang
}
func (this *OapiUserListbypageRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiUserListbypageRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiUserListbypageRequest) SetOrder(order2 string) {
	this.Order = order2
}
func (this *OapiUserListbypageRequest) GetOrder() string {
	return this.Order
}
func (this *OapiUserListbypageRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiUserListbypageRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiUserListbypageRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.listbypage"
}
func (this *OapiUserListbypageRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserListbypageRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserListbypageRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserListbypageRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserListbypageRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["department_id"] = this.DepartmentId
	txtParams["lang"] = this.Lang
	txtParams["offset"] = this.Offset
	txtParams["order"] = this.Order
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiUserListbypageRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Offset, "offset"); err != nil {
		return err
	}
	return nil
}
func (this *OapiUserListbypageRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserListbypageRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserListbypageResponse struct {
	taobao.TaobaoResponse
	Errcode  int64      `json:"errcode,omitempty"`
	Errmsg   string     `json:"errmsg,omitempty"`
	HasMore  bool       `json:"hasMore,omitempty"`
	Userlist []Userlist `json:"userlist,omitempty"`
}
