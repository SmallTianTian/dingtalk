package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiV2DepartmentGetRequest() *OapiV2DepartmentGetRequest {
	return &OapiV2DepartmentGetRequest{}
}

type OapiV2DepartmentGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiV2DepartmentGetResponse
	DeptId          int64
	Language        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiV2DepartmentGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2DepartmentGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2DepartmentGetRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiV2DepartmentGetRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiV2DepartmentGetRequest) SetLanguage(language2 string) {
	this.Language = language2
}
func (this *OapiV2DepartmentGetRequest) GetLanguage() string {
	return this.Language
}
func (this *OapiV2DepartmentGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.department.get"
}
func (this *OapiV2DepartmentGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2DepartmentGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2DepartmentGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2DepartmentGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2DepartmentGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_id"] = this.DeptId
	txtParams["language"] = this.Language
	return txtParams
}
func (this *OapiV2DepartmentGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeptId, "deptId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiV2DepartmentGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2DepartmentGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2DepartmentGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64           `json:"errcode,omitempty"`
	Errmsg  string          `json:"errmsg,omitempty"`
	Result  DeptGetResponse `json:"result,omitempty"`
}
type DeptGetResponse struct {
	AutoAddUser           bool     `json:"auto_add_user,omitempty"`
	CreateDeptGroup       bool     `json:"create_dept_group,omitempty"`
	DeptGroupChatId       string   `json:"dept_group_chat_id,omitempty"`
	DeptId                int64    `json:"dept_id,omitempty"`
	DeptManagerUseridList []string `json:"dept_manager_userid_list,omitempty"`
	DeptPermits           []int64  `json:"dept_permits,omitempty"`
	Extention             string   `json:"extention,omitempty"`
	FromUnionOrg          bool     `json:"from_union_org,omitempty"`
	GroupContainSubDept   bool     `json:"group_contain_sub_dept,omitempty"`
	HideDept              bool     `json:"hide_dept,omitempty"`
	Name                  string   `json:"name,omitempty"`
	Order                 int64    `json:"order,omitempty"`
	OrgDeptOwner          string   `json:"org_dept_owner,omitempty"`
	OuterDept             bool     `json:"outer_dept,omitempty"`
	OuterPermitDepts      []int64  `json:"outer_permit_depts,omitempty"`
	OuterPermitUsers      []string `json:"outer_permit_users,omitempty"`
	ParentId              int64    `json:"parent_id,omitempty"`
	SourceIdentifier      string   `json:"source_identifier,omitempty"`
	Tags                  string   `json:"tags,omitempty"`
	UserPermits           []string `json:"user_permits,omitempty"`
}
