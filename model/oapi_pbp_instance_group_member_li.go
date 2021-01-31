package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiPbpInstanceGroupMemberListRequest() *OapiPbpInstanceGroupMemberListRequest {
	return &OapiPbpInstanceGroupMemberListRequest{}
}

type OapiPbpInstanceGroupMemberListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPbpInstanceGroupMemberListResponse
	BizId           string
	Cursor          int64
	PunchGroupId    string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiPbpInstanceGroupMemberListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPbpInstanceGroupMemberListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPbpInstanceGroupMemberListRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiPbpInstanceGroupMemberListRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiPbpInstanceGroupMemberListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiPbpInstanceGroupMemberListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiPbpInstanceGroupMemberListRequest) SetPunchGroupId(punchGroupId2 string) {
	this.PunchGroupId = punchGroupId2
}
func (this *OapiPbpInstanceGroupMemberListRequest) GetPunchGroupId() string {
	return this.PunchGroupId
}
func (this *OapiPbpInstanceGroupMemberListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiPbpInstanceGroupMemberListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiPbpInstanceGroupMemberListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.pbp.instance.group.member.list"
}
func (this *OapiPbpInstanceGroupMemberListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPbpInstanceGroupMemberListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPbpInstanceGroupMemberListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPbpInstanceGroupMemberListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPbpInstanceGroupMemberListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_id"] = this.BizId
	txtParams["cursor"] = this.Cursor
	txtParams["punch_group_id"] = this.PunchGroupId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiPbpInstanceGroupMemberListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *OapiPbpInstanceGroupMemberListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPbpInstanceGroupMemberListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiPbpInstanceGroupMemberListResponse struct {
	taobao.TaobaoResponse
	Errcode int64      `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Result  PageResult `json:"result,omitempty"`
}
type PositionVo struct {
	MemberId string `json:"member_id,omitempty"`
	Type     int64  `json:"type,omitempty"`
}
