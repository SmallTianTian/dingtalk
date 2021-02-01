package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsJobListRequest() *OapiAtsJobListRequest {
	return &OapiAtsJobListRequest{}
}

type OapiAtsJobListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsJobListResponse
	BizCode         string
	Cursor          string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsJobListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsJobListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsJobListRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsJobListRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsJobListRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiAtsJobListRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiAtsJobListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiAtsJobListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiAtsJobListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.job.list"
}
func (this *OapiAtsJobListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsJobListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsJobListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsJobListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsJobListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["cursor"] = this.Cursor
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiAtsJobListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsJobListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsJobListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAtsJobListResponse struct {
	taobao.TaobaoResponse
	Result AtsPageResult `json:"result,omitempty"`
}
type AtsPageResult struct {
	HasMore    bool          `json:"has_more,omitempty"`
	List       []JobSimpleVo `json:"list,omitempty"`
	NextCursor string        `json:"next_cursor,omitempty"`
}
