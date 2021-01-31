package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiV2UserGetRequest() *OapiV2UserGetRequest {
	return &OapiV2UserGetRequest{}
}

type OapiV2UserGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiV2UserGetResponse
	Language        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiV2UserGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2UserGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2UserGetRequest) SetLanguage(language2 string) {
	this.Language = language2
}
func (this *OapiV2UserGetRequest) GetLanguage() string {
	return this.Language
}
func (this *OapiV2UserGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiV2UserGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiV2UserGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.user.get"
}
func (this *OapiV2UserGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2UserGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2UserGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2UserGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2UserGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["language"] = this.Language
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiV2UserGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckMaxLength(this.Language, 6, "language"); err != nil {
		return err
	}
	return nil
}
func (this *OapiV2UserGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2UserGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2UserGetResponse struct {
	taobao.TaobaoResponse
	Result UserGetResponse `json:"result,omitempty"`
}
type DeptLeader struct {
	DeptId int64 `json:"dept_id,omitempty"`
	Leader bool  `json:"leader,omitempty"`
}
type UserRole struct {
	GroupName string `json:"group_name,omitempty"`
	Id        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
}
type UnionEmpMapVo struct {
	CorpId string `json:"corp_id,omitempty"`
	Userid string `json:"userid,omitempty"`
}
type UnionEmpExt struct {
	CorpId          string          `json:"corp_id,omitempty"`
	UnionEmpMapList []UnionEmpMapVo `json:"union_emp_map_list,omitempty"`
	Userid          string          `json:"userid,omitempty"`
}
type UserGetResponse struct {
	Active           bool         `json:"active,omitempty"`
	Admin            bool         `json:"admin,omitempty"`
	Avatar           string       `json:"avatar,omitempty"`
	Boss             bool         `json:"boss,omitempty"`
	DeptIdList       []int64      `json:"dept_id_list,omitempty"`
	DeptOrderList    []DeptOrder  `json:"dept_order_list,omitempty"`
	Email            string       `json:"email,omitempty"`
	ExclusiveAccount bool         `json:"exclusive_account,omitempty"`
	Extension        string       `json:"extension,omitempty"`
	HideMobile       bool         `json:"hide_mobile,omitempty"`
	HiredDate        int64        `json:"hired_date,omitempty"`
	JobNumber        string       `json:"job_number,omitempty"`
	LeaderInDept     []DeptLeader `json:"leader_in_dept,omitempty"`
	Mobile           string       `json:"mobile,omitempty"`
	Name             string       `json:"name,omitempty"`
	OrgEmail         string       `json:"org_email,omitempty"`
	RealAuthed       bool         `json:"real_authed,omitempty"`
	Remark           string       `json:"remark,omitempty"`
	RoleList         []UserRole   `json:"role_list,omitempty"`
	Senior           bool         `json:"senior,omitempty"`
	StateCode        string       `json:"state_code,omitempty"`
	Telephone        string       `json:"telephone,omitempty"`
	Title            string       `json:"title,omitempty"`
	UnionEmpExt      UnionEmpExt  `json:"union_emp_ext,omitempty"`
	Unionid          string       `json:"unionid,omitempty"`
	Userid           string       `json:"userid,omitempty"`
	WorkPlace        string       `json:"work_place,omitempty"`
}
