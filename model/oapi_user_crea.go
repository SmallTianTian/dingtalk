package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiUserCreateRequest() *OapiUserCreateRequest {
	return &OapiUserCreateRequest{}
}

type OapiUserCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserCreateResponse
	Department      string
	Email           string
	Extattr         string
	HiredDate       int64
	IsHide          bool
	IsSenior        bool
	Jobnumber       string
	Mobile          string
	Name            string
	OrderInDepts    string
	OrgEmail        string
	Position        string
	PositionInDepts string
	Remark          string
	Tel             string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
	WorkPlace       string
}

func (this *OapiUserCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiUserCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserCreateRequest) SetDepartment(department2 string) {
	this.Department = department2
}
func (this *OapiUserCreateRequest) GetDepartment() string {
	return this.Department
}
func (this *OapiUserCreateRequest) SetEmail(email2 string) {
	this.Email = email2
}
func (this *OapiUserCreateRequest) GetEmail() string {
	return this.Email
}
func (this *OapiUserCreateRequest) SetExtattr(extattr2 string) {
	this.Extattr = extattr2
}
func (this *OapiUserCreateRequest) SetExtattrString(extattr2 string) {
	this.Extattr = extattr2
}
func (this *OapiUserCreateRequest) GetExtattr() string {
	return this.Extattr
}
func (this *OapiUserCreateRequest) SetHiredDate(hiredDate2 int64) {
	this.HiredDate = hiredDate2
}
func (this *OapiUserCreateRequest) GetHiredDate() int64 {
	return this.HiredDate
}
func (this *OapiUserCreateRequest) SetIsHide(isHide2 bool) {
	this.IsHide = isHide2
}
func (this *OapiUserCreateRequest) GetIsHide() bool {
	return this.IsHide
}
func (this *OapiUserCreateRequest) SetIsSenior(isSenior2 bool) {
	this.IsSenior = isSenior2
}
func (this *OapiUserCreateRequest) GetIsSenior() bool {
	return this.IsSenior
}
func (this *OapiUserCreateRequest) SetJobnumber(jobnumber2 string) {
	this.Jobnumber = jobnumber2
}
func (this *OapiUserCreateRequest) GetJobnumber() string {
	return this.Jobnumber
}
func (this *OapiUserCreateRequest) SetMobile(mobile2 string) {
	this.Mobile = mobile2
}
func (this *OapiUserCreateRequest) GetMobile() string {
	return this.Mobile
}
func (this *OapiUserCreateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiUserCreateRequest) GetName() string {
	return this.Name
}
func (this *OapiUserCreateRequest) SetOrderInDepts(orderInDepts2 string) {
	this.OrderInDepts = orderInDepts2
}
func (this *OapiUserCreateRequest) SetOrderInDeptsString(orderInDepts2 string) {
	this.OrderInDepts = orderInDepts2
}
func (this *OapiUserCreateRequest) GetOrderInDepts() string {
	return this.OrderInDepts
}
func (this *OapiUserCreateRequest) SetOrgEmail(orgEmail2 string) {
	this.OrgEmail = orgEmail2
}
func (this *OapiUserCreateRequest) GetOrgEmail() string {
	return this.OrgEmail
}
func (this *OapiUserCreateRequest) SetPosition(position2 string) {
	this.Position = position2
}
func (this *OapiUserCreateRequest) GetPosition() string {
	return this.Position
}
func (this *OapiUserCreateRequest) SetPositionInDepts(positionInDepts2 string) {
	this.PositionInDepts = positionInDepts2
}
func (this *OapiUserCreateRequest) SetPositionInDeptsString(positionInDepts2 string) {
	this.PositionInDepts = positionInDepts2
}
func (this *OapiUserCreateRequest) GetPositionInDepts() string {
	return this.PositionInDepts
}
func (this *OapiUserCreateRequest) SetRemark(remark2 string) {
	this.Remark = remark2
}
func (this *OapiUserCreateRequest) GetRemark() string {
	return this.Remark
}
func (this *OapiUserCreateRequest) SetTel(tel2 string) {
	this.Tel = tel2
}
func (this *OapiUserCreateRequest) GetTel() string {
	return this.Tel
}
func (this *OapiUserCreateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiUserCreateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiUserCreateRequest) SetWorkPlace(workPlace2 string) {
	this.WorkPlace = workPlace2
}
func (this *OapiUserCreateRequest) GetWorkPlace() string {
	return this.WorkPlace
}
func (this *OapiUserCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.create"
}
func (this *OapiUserCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["department"] = this.Department
	txtParams["email"] = this.Email
	txtParams["extattr"] = this.Extattr
	txtParams["hiredDate"] = this.HiredDate
	txtParams["isHide"] = this.IsHide
	txtParams["isSenior"] = this.IsSenior
	txtParams["jobnumber"] = this.Jobnumber
	txtParams["mobile"] = this.Mobile
	txtParams["name"] = this.Name
	txtParams["orderInDepts"] = this.OrderInDepts
	txtParams["orgEmail"] = this.OrgEmail
	txtParams["position"] = this.Position
	txtParams["positionInDepts"] = this.PositionInDepts
	txtParams["remark"] = this.Remark
	txtParams["tel"] = this.Tel
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	txtParams["workPlace"] = this.WorkPlace
	return txtParams
}
func (this *OapiUserCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiUserCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Userid  string `json:"userid,omitempty"`
}
