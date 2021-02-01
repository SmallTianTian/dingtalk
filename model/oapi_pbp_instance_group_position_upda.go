package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiPbpInstanceGroupPositionUpdateRequest() *OapiPbpInstanceGroupPositionUpdateRequest {
	return &OapiPbpInstanceGroupPositionUpdateRequest{}
}

type OapiPbpInstanceGroupPositionUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPbpInstanceGroupPositionUpdateResponse
	SyncParam       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiPbpInstanceGroupPositionUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPbpInstanceGroupPositionUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPbpInstanceGroupPositionUpdateRequest) SetSyncParam(syncParam2 string) {
	this.SyncParam = syncParam2
}
func (this *OapiPbpInstanceGroupPositionUpdateRequest) GetSyncParam() string {
	return this.SyncParam
}
func (this *OapiPbpInstanceGroupPositionUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.pbp.instance.group.position.update"
}
func (this *OapiPbpInstanceGroupPositionUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPbpInstanceGroupPositionUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPbpInstanceGroupPositionUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPbpInstanceGroupPositionUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPbpInstanceGroupPositionUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["sync_param"] = this.SyncParam
	return txtParams
}
func (this *OapiPbpInstanceGroupPositionUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiPbpInstanceGroupPositionUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPbpInstanceGroupPositionUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type PunchGroupPositionParam struct {
	PositionId   string `json:"position_id,omitempty"`
	PositionName string `json:"position_name,omitempty"`
	PositionType int64  `json:"position_type,omitempty"`
}
type PunchGroupSyncPositionParam struct {
	AddPositionList    []PunchGroupPositionParam `json:"add_position_list,omitempty"`
	BizId              string                    `json:"biz_id,omitempty"`
	BizInstId          string                    `json:"biz_inst_id,omitempty"`
	DeletePositionList []PunchGroupPositionParam `json:"delete_position_list,omitempty"`
	PunchGroupId       string                    `json:"punch_group_id,omitempty"`
}
type OapiPbpInstanceGroupPositionUpdateResponse struct {
	taobao.TaobaoResponse

	PunchGroupId string `json:"punch_group_id,omitempty"`
}
