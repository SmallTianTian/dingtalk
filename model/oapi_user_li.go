package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiUserListRequest() *OapiUserListRequest {
	return &OapiUserListRequest{}
}

type OapiUserListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserListResponse
	DepartmentId    int64
	Lang            string
	Offset          int64
	Order           string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiUserListRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiUserListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserListRequest) SetDepartmentId(departmentId2 int64) {
	this.DepartmentId = departmentId2
}
func (this *OapiUserListRequest) GetDepartmentId() int64 {
	return this.DepartmentId
}
func (this *OapiUserListRequest) SetLang(lang2 string) {
	this.Lang = lang2
}
func (this *OapiUserListRequest) GetLang() string {
	return this.Lang
}
func (this *OapiUserListRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiUserListRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiUserListRequest) SetOrder(order2 string) {
	this.Order = order2
}
func (this *OapiUserListRequest) GetOrder() string {
	return this.Order
}
func (this *OapiUserListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiUserListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiUserListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.list"
}
func (this *OapiUserListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["department_id"] = this.DepartmentId
	txtParams["lang"] = this.Lang
	txtParams["offset"] = this.Offset
	txtParams["order"] = this.Order
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiUserListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiUserListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserListResponse struct {
	taobao.TaobaoResponse

	HasMore  bool       `json:"hasMore,omitempty"`
	Userlist []Userlist `json:"userlist,omitempty"`
}
type Userlist struct {
	Active     bool      `json:"active,omitempty"`
	Avatar     string    `json:"avatar,omitempty"`
	Department string    `json:"department,omitempty"`
	DingId     string    `json:"dingId,omitempty"`
	Email      string    `json:"email,omitempty"`
	Extattr    string    `json:"extattr,omitempty"`
	HiredDate  time.Time `json:"hiredDate,omitempty"`
	IsAdmin    bool      `json:"isAdmin,omitempty"`
	IsBoss     bool      `json:"isBoss,omitempty"`
	IsHide     bool      `json:"isHide,omitempty"`
	IsLeader   bool      `json:"isLeader,omitempty"`
	Jobnumber  string    `json:"jobnumber,omitempty"`
	Mobile     string    `json:"mobile,omitempty"`
	Name       string    `json:"name,omitempty"`
	Order      int64     `json:"order,omitempty"`
	OrgEmail   string    `json:"orgEmail,omitempty"`
	Position   string    `json:"position,omitempty"`
	Remark     string    `json:"remark,omitempty"`
	Tel        string    `json:"tel,omitempty"`
	Unionid    string    `json:"unionid,omitempty"`
	Userid     string    `json:"userid,omitempty"`
	WorkPlace  string    `json:"workPlace,omitempty"`
}
