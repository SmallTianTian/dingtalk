package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiUserUpdateRequest() *OapiUserUpdateRequest {
	return &OapiUserUpdateRequest{}
}

type OapiUserUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserUpdateResponse
	Department      []int64
	Email           string
	Extattr         string
	HiredDate       int64
	IsHide          bool
	IsSenior        bool
	Jobnumber       string
	Lang            string
	ManagerUserid   string
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

func (this *OapiUserUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiUserUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserUpdateRequest) SetDepartment(department2 []int64) {
	this.Department = department2
}
func (this *OapiUserUpdateRequest) GetDepartment() []int64 {
	return this.Department
}
func (this *OapiUserUpdateRequest) SetEmail(email2 string) {
	this.Email = email2
}
func (this *OapiUserUpdateRequest) GetEmail() string {
	return this.Email
}
func (this *OapiUserUpdateRequest) SetExtattr(extattr2 string) {
	this.Extattr = extattr2
}
func (this *OapiUserUpdateRequest) GetExtattr() string {
	return this.Extattr
}
func (this *OapiUserUpdateRequest) SetHiredDate(hiredDate2 int64) {
	this.HiredDate = hiredDate2
}
func (this *OapiUserUpdateRequest) GetHiredDate() int64 {
	return this.HiredDate
}
func (this *OapiUserUpdateRequest) SetIsHide(isHide2 bool) {
	this.IsHide = isHide2
}
func (this *OapiUserUpdateRequest) GetIsHide() bool {
	return this.IsHide
}
func (this *OapiUserUpdateRequest) SetIsSenior(isSenior2 bool) {
	this.IsSenior = isSenior2
}
func (this *OapiUserUpdateRequest) GetIsSenior() bool {
	return this.IsSenior
}
func (this *OapiUserUpdateRequest) SetJobnumber(jobnumber2 string) {
	this.Jobnumber = jobnumber2
}
func (this *OapiUserUpdateRequest) GetJobnumber() string {
	return this.Jobnumber
}
func (this *OapiUserUpdateRequest) SetLang(lang2 string) {
	this.Lang = lang2
}
func (this *OapiUserUpdateRequest) GetLang() string {
	return this.Lang
}
func (this *OapiUserUpdateRequest) SetManagerUserid(managerUserid2 string) {
	this.ManagerUserid = managerUserid2
}
func (this *OapiUserUpdateRequest) GetManagerUserid() string {
	return this.ManagerUserid
}
func (this *OapiUserUpdateRequest) SetMobile(mobile2 string) {
	this.Mobile = mobile2
}
func (this *OapiUserUpdateRequest) GetMobile() string {
	return this.Mobile
}
func (this *OapiUserUpdateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiUserUpdateRequest) GetName() string {
	return this.Name
}
func (this *OapiUserUpdateRequest) SetOrderInDepts(orderInDepts2 string) {
	this.OrderInDepts = orderInDepts2
}
func (this *OapiUserUpdateRequest) GetOrderInDepts() string {
	return this.OrderInDepts
}
func (this *OapiUserUpdateRequest) SetOrgEmail(orgEmail2 string) {
	this.OrgEmail = orgEmail2
}
func (this *OapiUserUpdateRequest) GetOrgEmail() string {
	return this.OrgEmail
}
func (this *OapiUserUpdateRequest) SetPosition(position2 string) {
	this.Position = position2
}
func (this *OapiUserUpdateRequest) GetPosition() string {
	return this.Position
}
func (this *OapiUserUpdateRequest) SetPositionInDepts(positionInDepts2 string) {
	this.PositionInDepts = positionInDepts2
}
func (this *OapiUserUpdateRequest) GetPositionInDepts() string {
	return this.PositionInDepts
}
func (this *OapiUserUpdateRequest) SetRemark(remark2 string) {
	this.Remark = remark2
}
func (this *OapiUserUpdateRequest) GetRemark() string {
	return this.Remark
}
func (this *OapiUserUpdateRequest) SetTel(tel2 string) {
	this.Tel = tel2
}
func (this *OapiUserUpdateRequest) GetTel() string {
	return this.Tel
}
func (this *OapiUserUpdateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiUserUpdateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiUserUpdateRequest) SetWorkPlace(workPlace2 string) {
	this.WorkPlace = workPlace2
}
func (this *OapiUserUpdateRequest) GetWorkPlace() string {
	return this.WorkPlace
}
func (this *OapiUserUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.update"
}
func (this *OapiUserUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["department"] = this.Department
	txtParams["email"] = this.Email
	txtParams["extattr"] = this.Extattr
	txtParams["hiredDate"] = this.HiredDate
	txtParams["isHide"] = this.IsHide
	txtParams["isSenior"] = this.IsSenior
	txtParams["jobnumber"] = this.Jobnumber
	txtParams["lang"] = this.Lang
	txtParams["managerUserid"] = this.ManagerUserid
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
func (this *OapiUserUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiUserUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
