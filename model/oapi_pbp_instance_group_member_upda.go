package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiPbpInstanceGroupMemberUpdateRequest() *OapiPbpInstanceGroupMemberUpdateRequest {
	return &OapiPbpInstanceGroupMemberUpdateRequest{}
}

type OapiPbpInstanceGroupMemberUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPbpInstanceGroupMemberUpdateResponse
	SyncParam       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiPbpInstanceGroupMemberUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPbpInstanceGroupMemberUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPbpInstanceGroupMemberUpdateRequest) SetSyncParam(syncParam2 string) {
	this.SyncParam = syncParam2
}
func (this *OapiPbpInstanceGroupMemberUpdateRequest) GetSyncParam() string {
	return this.SyncParam
}
func (this *OapiPbpInstanceGroupMemberUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.pbp.instance.group.member.update"
}
func (this *OapiPbpInstanceGroupMemberUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPbpInstanceGroupMemberUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPbpInstanceGroupMemberUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPbpInstanceGroupMemberUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPbpInstanceGroupMemberUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["sync_param"] = this.SyncParam
	return txtParams
}
func (this *OapiPbpInstanceGroupMemberUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiPbpInstanceGroupMemberUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPbpInstanceGroupMemberUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type PunchGroupMemberParam struct {
	MemberId string `json:"member_id,omitempty"`
	Type     int64  `json:"type,omitempty"`
}
type PunchGroupSyncMemberParam struct {
	AddMemberList    []PunchGroupMemberParam `json:"add_member_list,omitempty"`
	BizId            string                  `json:"biz_id,omitempty"`
	BizInstId        string                  `json:"biz_inst_id,omitempty"`
	DeleteMemberList []PunchGroupMemberParam `json:"delete_member_list,omitempty"`
	PunchGroupId     string                  `json:"punch_group_id,omitempty"`
}
type OapiPbpInstanceGroupMemberUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode      int64  `json:"errcode,omitempty"`
	Errmsg       string `json:"errmsg,omitempty"`
	PunchGroupId string `json:"punch_group_id,omitempty"`
}
