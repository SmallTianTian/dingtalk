package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiV2UserCreateRequest() *OapiV2UserCreateRequest {
	return &OapiV2UserCreateRequest{}
}

type OapiV2UserCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp             OapiV2UserCreateResponse
	DeptIdList       string
	DeptOrderList    string
	DeptTitleList    string
	Email            string
	ExclusiveAccount bool
	Extension        string
	HideMobile       bool
	HiredDate        int64
	JobNumber        string
	LoginEmail       string
	Mobile           string
	Name             string
	OrgEmail         string
	Remark           string
	SeniorMode       bool
	Telephone        string
	Title            string
	TopHttpMethod    string
	TopResponseType  string
	Userid           string
	WorkPlace        string
}

func (this *OapiV2UserCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2UserCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2UserCreateRequest) SetDeptIdList(deptIdList2 string) {
	this.DeptIdList = deptIdList2
}
func (this *OapiV2UserCreateRequest) GetDeptIdList() string {
	return this.DeptIdList
}
func (this *OapiV2UserCreateRequest) SetDeptOrderList(deptOrderList2 string) {
	this.DeptOrderList = deptOrderList2
}
func (this *OapiV2UserCreateRequest) GetDeptOrderList() string {
	return this.DeptOrderList
}
func (this *OapiV2UserCreateRequest) SetDeptTitleList(deptTitleList2 string) {
	this.DeptTitleList = deptTitleList2
}
func (this *OapiV2UserCreateRequest) GetDeptTitleList() string {
	return this.DeptTitleList
}
func (this *OapiV2UserCreateRequest) SetEmail(email2 string) {
	this.Email = email2
}
func (this *OapiV2UserCreateRequest) GetEmail() string {
	return this.Email
}
func (this *OapiV2UserCreateRequest) SetExclusiveAccount(exclusiveAccount2 bool) {
	this.ExclusiveAccount = exclusiveAccount2
}
func (this *OapiV2UserCreateRequest) GetExclusiveAccount() bool {
	return this.ExclusiveAccount
}
func (this *OapiV2UserCreateRequest) SetExtension(extension2 string) {
	this.Extension = extension2
}
func (this *OapiV2UserCreateRequest) SetExtensionString(extension2 string) {
	this.Extension = extension2
}
func (this *OapiV2UserCreateRequest) GetExtension() string {
	return this.Extension
}
func (this *OapiV2UserCreateRequest) SetHideMobile(hideMobile2 bool) {
	this.HideMobile = hideMobile2
}
func (this *OapiV2UserCreateRequest) GetHideMobile() bool {
	return this.HideMobile
}
func (this *OapiV2UserCreateRequest) SetHiredDate(hiredDate2 int64) {
	this.HiredDate = hiredDate2
}
func (this *OapiV2UserCreateRequest) GetHiredDate() int64 {
	return this.HiredDate
}
func (this *OapiV2UserCreateRequest) SetJobNumber(jobNumber2 string) {
	this.JobNumber = jobNumber2
}
func (this *OapiV2UserCreateRequest) GetJobNumber() string {
	return this.JobNumber
}
func (this *OapiV2UserCreateRequest) SetLoginEmail(loginEmail2 string) {
	this.LoginEmail = loginEmail2
}
func (this *OapiV2UserCreateRequest) GetLoginEmail() string {
	return this.LoginEmail
}
func (this *OapiV2UserCreateRequest) SetMobile(mobile2 string) {
	this.Mobile = mobile2
}
func (this *OapiV2UserCreateRequest) GetMobile() string {
	return this.Mobile
}
func (this *OapiV2UserCreateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiV2UserCreateRequest) GetName() string {
	return this.Name
}
func (this *OapiV2UserCreateRequest) SetOrgEmail(orgEmail2 string) {
	this.OrgEmail = orgEmail2
}
func (this *OapiV2UserCreateRequest) GetOrgEmail() string {
	return this.OrgEmail
}
func (this *OapiV2UserCreateRequest) SetRemark(remark2 string) {
	this.Remark = remark2
}
func (this *OapiV2UserCreateRequest) GetRemark() string {
	return this.Remark
}
func (this *OapiV2UserCreateRequest) SetSeniorMode(seniorMode2 bool) {
	this.SeniorMode = seniorMode2
}
func (this *OapiV2UserCreateRequest) GetSeniorMode() bool {
	return this.SeniorMode
}
func (this *OapiV2UserCreateRequest) SetTelephone(telephone2 string) {
	this.Telephone = telephone2
}
func (this *OapiV2UserCreateRequest) GetTelephone() string {
	return this.Telephone
}
func (this *OapiV2UserCreateRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiV2UserCreateRequest) GetTitle() string {
	return this.Title
}
func (this *OapiV2UserCreateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiV2UserCreateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiV2UserCreateRequest) SetWorkPlace(workPlace2 string) {
	this.WorkPlace = workPlace2
}
func (this *OapiV2UserCreateRequest) GetWorkPlace() string {
	return this.WorkPlace
}
func (this *OapiV2UserCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.user.create"
}
func (this *OapiV2UserCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2UserCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2UserCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2UserCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2UserCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_id_list"] = this.DeptIdList
	txtParams["dept_order_list"] = this.DeptOrderList
	txtParams["dept_title_list"] = this.DeptTitleList
	txtParams["email"] = this.Email
	txtParams["exclusive_account"] = this.ExclusiveAccount
	txtParams["extension"] = this.Extension
	txtParams["hide_mobile"] = this.HideMobile
	txtParams["hired_date"] = this.HiredDate
	txtParams["job_number"] = this.JobNumber
	txtParams["login_email"] = this.LoginEmail
	txtParams["mobile"] = this.Mobile
	txtParams["name"] = this.Name
	txtParams["org_email"] = this.OrgEmail
	txtParams["remark"] = this.Remark
	txtParams["senior_mode"] = this.SeniorMode
	txtParams["telephone"] = this.Telephone
	txtParams["title"] = this.Title
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	txtParams["work_place"] = this.WorkPlace
	return txtParams
}
func (this *OapiV2UserCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.DeptIdList, 100, "deptIdList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiV2UserCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2UserCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DeptOrder struct {
	DeptId int64 `json:"dept_id,omitempty"`
	Order  int64 `json:"order,omitempty"`
}
type DeptTitle struct {
	DeptId int64  `json:"dept_id,omitempty"`
	Title  string `json:"title,omitempty"`
}
type OapiV2UserCreateResponse struct {
	taobao.TaobaoResponse
	Result UserCreateResponse `json:"result,omitempty"`
}
type UserCreateResponse struct {
	Userid string `json:"userid,omitempty"`
}
