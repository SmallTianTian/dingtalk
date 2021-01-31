package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiV2UserListRequest() *OapiV2UserListRequest {
	return &OapiV2UserListRequest{}
}

type OapiV2UserListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiV2UserListResponse
	ContainAccessLimit bool
	Cursor             int64
	DeptId             int64
	Language           string
	OrderField         string
	Size               int64
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiV2UserListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2UserListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2UserListRequest) SetContainAccessLimit(containAccessLimit2 bool) {
	this.ContainAccessLimit = containAccessLimit2
}
func (this *OapiV2UserListRequest) GetContainAccessLimit() bool {
	return this.ContainAccessLimit
}
func (this *OapiV2UserListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiV2UserListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiV2UserListRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiV2UserListRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiV2UserListRequest) SetLanguage(language2 string) {
	this.Language = language2
}
func (this *OapiV2UserListRequest) GetLanguage() string {
	return this.Language
}
func (this *OapiV2UserListRequest) SetOrderField(orderField2 string) {
	this.OrderField = orderField2
}
func (this *OapiV2UserListRequest) GetOrderField() string {
	return this.OrderField
}
func (this *OapiV2UserListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiV2UserListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiV2UserListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.user.list"
}
func (this *OapiV2UserListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2UserListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2UserListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2UserListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2UserListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["contain_access_limit"] = this.ContainAccessLimit
	txtParams["cursor"] = this.Cursor
	txtParams["dept_id"] = this.DeptId
	txtParams["language"] = this.Language
	txtParams["order_field"] = this.OrderField
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiV2UserListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *OapiV2UserListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2UserListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2UserListResponse struct {
	taobao.TaobaoResponse
	Result PageResult `json:"result,omitempty"`
}
type ListUserResponse struct {
	Active           bool    `json:"active,omitempty"`
	Admin            bool    `json:"admin,omitempty"`
	Avatar           string  `json:"avatar,omitempty"`
	Boss             bool    `json:"boss,omitempty"`
	DeptIdList       []int64 `json:"dept_id_list,omitempty"`
	DeptOrder        int64   `json:"dept_order,omitempty"`
	Email            string  `json:"email,omitempty"`
	ExclusiveAccount bool    `json:"exclusive_account,omitempty"`
	Extension        string  `json:"extension,omitempty"`
	HideMobile       bool    `json:"hide_mobile,omitempty"`
	HiredDate        int64   `json:"hired_date,omitempty"`
	JobNumber        string  `json:"job_number,omitempty"`
	Leader           bool    `json:"leader,omitempty"`
	Mobile           string  `json:"mobile,omitempty"`
	Name             string  `json:"name,omitempty"`
	OrgEmail         string  `json:"org_email,omitempty"`
	Remark           string  `json:"remark,omitempty"`
	StateCode        string  `json:"state_code,omitempty"`
	Telephone        string  `json:"telephone,omitempty"`
	Title            string  `json:"title,omitempty"`
	Unionid          string  `json:"unionid,omitempty"`
	Userid           string  `json:"userid,omitempty"`
	WorkPlace        string  `json:"work_place,omitempty"`
}
