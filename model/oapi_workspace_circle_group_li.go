package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceCircleGroupListRequest() *OapiWorkspaceCircleGroupListRequest {
	return &OapiWorkspaceCircleGroupListRequest{}
}

type OapiWorkspaceCircleGroupListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceCircleGroupListResponse
	Cursor          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceCircleGroupListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceCircleGroupListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceCircleGroupListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiWorkspaceCircleGroupListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiWorkspaceCircleGroupListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiWorkspaceCircleGroupListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiWorkspaceCircleGroupListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.circle.group.list"
}
func (this *OapiWorkspaceCircleGroupListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceCircleGroupListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceCircleGroupListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceCircleGroupListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceCircleGroupListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiWorkspaceCircleGroupListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Size, "size"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceCircleGroupListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceCircleGroupListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceCircleGroupListResponse struct {
	taobao.TaobaoResponse
	Result OpenQueryGroupResponseDto `json:"result,omitempty"`
}
type OpenGroupDto struct {
	AvatarMediaId  string `json:"avatar_media_id,omitempty"`
	ConversationId string `json:"conversation_id,omitempty"`
	MemberCount    int64  `json:"member_count,omitempty"`
	MemberLimit    int64  `json:"member_limit,omitempty"`
	Name           string `json:"name,omitempty"`
}
type OpenQueryGroupResponseDto struct {
	Groups     []OpenGroupDto `json:"groups,omitempty"`
	HasMore    bool           `json:"has_more,omitempty"`
	NextCursor int64          `json:"next_cursor,omitempty"`
}
