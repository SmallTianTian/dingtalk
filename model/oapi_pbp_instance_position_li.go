package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiPbpInstancePositionListRequest() *OapiPbpInstancePositionListRequest {
	return &OapiPbpInstancePositionListRequest{}
}

type OapiPbpInstancePositionListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPbpInstancePositionListResponse
	BizId           string
	BizInstId       string
	Cursor          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
	Type            int64
}

func (this *OapiPbpInstancePositionListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPbpInstancePositionListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPbpInstancePositionListRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiPbpInstancePositionListRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiPbpInstancePositionListRequest) SetBizInstId(bizInstId2 string) {
	this.BizInstId = bizInstId2
}
func (this *OapiPbpInstancePositionListRequest) GetBizInstId() string {
	return this.BizInstId
}
func (this *OapiPbpInstancePositionListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiPbpInstancePositionListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiPbpInstancePositionListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiPbpInstancePositionListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiPbpInstancePositionListRequest) SetType(type2 int64) {
	this.Type = type2
}
func (this *OapiPbpInstancePositionListRequest) GetType() int64 {
	return this.Type
}
func (this *OapiPbpInstancePositionListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.pbp.instance.position.list"
}
func (this *OapiPbpInstancePositionListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPbpInstancePositionListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPbpInstancePositionListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPbpInstancePositionListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPbpInstancePositionListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_id"] = this.BizId
	txtParams["biz_inst_id"] = this.BizInstId
	txtParams["cursor"] = this.Cursor
	txtParams["size"] = this.Size
	txtParams["type"] = this.Type
	return txtParams
}
func (this *OapiPbpInstancePositionListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizId, "bizId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiPbpInstancePositionListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPbpInstancePositionListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiPbpInstancePositionListResponse struct {
	taobao.TaobaoResponse
	Result PageResult `json:"result,omitempty"`
}
