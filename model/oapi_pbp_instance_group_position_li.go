package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiPbpInstanceGroupPositionListRequest() *OapiPbpInstanceGroupPositionListRequest {
	return &OapiPbpInstanceGroupPositionListRequest{}
}

type OapiPbpInstanceGroupPositionListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPbpInstanceGroupPositionListResponse
	BizId           string
	Cursor          int64
	PunchGroupId    string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiPbpInstanceGroupPositionListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPbpInstanceGroupPositionListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPbpInstanceGroupPositionListRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiPbpInstanceGroupPositionListRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiPbpInstanceGroupPositionListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiPbpInstanceGroupPositionListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiPbpInstanceGroupPositionListRequest) SetPunchGroupId(punchGroupId2 string) {
	this.PunchGroupId = punchGroupId2
}
func (this *OapiPbpInstanceGroupPositionListRequest) GetPunchGroupId() string {
	return this.PunchGroupId
}
func (this *OapiPbpInstanceGroupPositionListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiPbpInstanceGroupPositionListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiPbpInstanceGroupPositionListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.pbp.instance.group.position.list"
}
func (this *OapiPbpInstanceGroupPositionListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPbpInstanceGroupPositionListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPbpInstanceGroupPositionListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPbpInstanceGroupPositionListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPbpInstanceGroupPositionListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_id"] = this.BizId
	txtParams["cursor"] = this.Cursor
	txtParams["punch_group_id"] = this.PunchGroupId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiPbpInstanceGroupPositionListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.PunchGroupId, "punchGroupId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiPbpInstanceGroupPositionListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPbpInstanceGroupPositionListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiPbpInstanceGroupPositionListResponse struct {
	taobao.TaobaoResponse
	Result PageResult `json:"result,omitempty"`
}
