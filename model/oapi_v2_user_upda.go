package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiV2UserUpdateRequest() *OapiV2UserUpdateRequest {
	return &OapiV2UserUpdateRequest{}
}

type OapiV2UserUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiV2UserUpdateResponse
	DeptIdList      string
	DeptOrderList   string
	DeptTitleList   string
	Email           string
	Extension       string
	HideMobile      bool
	HiredDate       int64
	JobNumber       string
	Language        string
	Mobile          string
	Name            string
	OrgEmail        string
	Remark          string
	SeniorMode      bool
	Telephone       string
	Title           string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
	WorkPlace       string
}

func (this *OapiV2UserUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2UserUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2UserUpdateRequest) SetDeptIdList(deptIdList2 string) {
	this.DeptIdList = deptIdList2
}
func (this *OapiV2UserUpdateRequest) GetDeptIdList() string {
	return this.DeptIdList
}
func (this *OapiV2UserUpdateRequest) SetDeptOrderList(deptOrderList2 string) {
	this.DeptOrderList = deptOrderList2
}
func (this *OapiV2UserUpdateRequest) GetDeptOrderList() string {
	return this.DeptOrderList
}
func (this *OapiV2UserUpdateRequest) SetDeptTitleList(deptTitleList2 string) {
	this.DeptTitleList = deptTitleList2
}
func (this *OapiV2UserUpdateRequest) GetDeptTitleList() string {
	return this.DeptTitleList
}
func (this *OapiV2UserUpdateRequest) SetEmail(email2 string) {
	this.Email = email2
}
func (this *OapiV2UserUpdateRequest) GetEmail() string {
	return this.Email
}
func (this *OapiV2UserUpdateRequest) SetExtension(extension2 string) {
	this.Extension = extension2
}
func (this *OapiV2UserUpdateRequest) SetExtensionString(extension2 string) {
	this.Extension = extension2
}
func (this *OapiV2UserUpdateRequest) GetExtension() string {
	return this.Extension
}
func (this *OapiV2UserUpdateRequest) SetHideMobile(hideMobile2 bool) {
	this.HideMobile = hideMobile2
}
func (this *OapiV2UserUpdateRequest) GetHideMobile() bool {
	return this.HideMobile
}
func (this *OapiV2UserUpdateRequest) SetHiredDate(hiredDate2 int64) {
	this.HiredDate = hiredDate2
}
func (this *OapiV2UserUpdateRequest) GetHiredDate() int64 {
	return this.HiredDate
}
func (this *OapiV2UserUpdateRequest) SetJobNumber(jobNumber2 string) {
	this.JobNumber = jobNumber2
}
func (this *OapiV2UserUpdateRequest) GetJobNumber() string {
	return this.JobNumber
}
func (this *OapiV2UserUpdateRequest) SetLanguage(language2 string) {
	this.Language = language2
}
func (this *OapiV2UserUpdateRequest) GetLanguage() string {
	return this.Language
}
func (this *OapiV2UserUpdateRequest) SetMobile(mobile2 string) {
	this.Mobile = mobile2
}
func (this *OapiV2UserUpdateRequest) GetMobile() string {
	return this.Mobile
}
func (this *OapiV2UserUpdateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiV2UserUpdateRequest) GetName() string {
	return this.Name
}
func (this *OapiV2UserUpdateRequest) SetOrgEmail(orgEmail2 string) {
	this.OrgEmail = orgEmail2
}
func (this *OapiV2UserUpdateRequest) GetOrgEmail() string {
	return this.OrgEmail
}
func (this *OapiV2UserUpdateRequest) SetRemark(remark2 string) {
	this.Remark = remark2
}
func (this *OapiV2UserUpdateRequest) GetRemark() string {
	return this.Remark
}
func (this *OapiV2UserUpdateRequest) SetSeniorMode(seniorMode2 bool) {
	this.SeniorMode = seniorMode2
}
func (this *OapiV2UserUpdateRequest) GetSeniorMode() bool {
	return this.SeniorMode
}
func (this *OapiV2UserUpdateRequest) SetTelephone(telephone2 string) {
	this.Telephone = telephone2
}
func (this *OapiV2UserUpdateRequest) GetTelephone() string {
	return this.Telephone
}
func (this *OapiV2UserUpdateRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiV2UserUpdateRequest) GetTitle() string {
	return this.Title
}
func (this *OapiV2UserUpdateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiV2UserUpdateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiV2UserUpdateRequest) SetWorkPlace(workPlace2 string) {
	this.WorkPlace = workPlace2
}
func (this *OapiV2UserUpdateRequest) GetWorkPlace() string {
	return this.WorkPlace
}
func (this *OapiV2UserUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.user.update"
}
func (this *OapiV2UserUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2UserUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2UserUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2UserUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2UserUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_id_list"] = this.DeptIdList
	txtParams["dept_order_list"] = this.DeptOrderList
	txtParams["dept_title_list"] = this.DeptTitleList
	txtParams["email"] = this.Email
	txtParams["extension"] = this.Extension
	txtParams["hide_mobile"] = this.HideMobile
	txtParams["hired_date"] = this.HiredDate
	txtParams["job_number"] = this.JobNumber
	txtParams["language"] = this.Language
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
func (this *OapiV2UserUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.DeptIdList, 100, "deptIdList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiV2UserUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2UserUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2UserUpdateResponse struct {
	taobao.TaobaoResponse
}
