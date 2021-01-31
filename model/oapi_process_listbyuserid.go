package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessListbyuseridRequest() *OapiProcessListbyuseridRequest {
	return &OapiProcessListbyuseridRequest{}
}

type OapiProcessListbyuseridRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessListbyuseridResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiProcessListbyuseridRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessListbyuseridRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessListbyuseridRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiProcessListbyuseridRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiProcessListbyuseridRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiProcessListbyuseridRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiProcessListbyuseridRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiProcessListbyuseridRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiProcessListbyuseridRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.listbyuserid"
}
func (this *OapiProcessListbyuseridRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessListbyuseridRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessListbyuseridRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessListbyuseridRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessListbyuseridRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiProcessListbyuseridRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Offset, "offset"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessListbyuseridRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessListbyuseridRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessListbyuseridResponse struct {
	taobao.TaobaoResponse
	Result HomePageProcessTemplateVo `json:"result,omitempty"`
}
type HomePageProcessTemplateVo struct {
	NextCursor  int64          `json:"next_cursor,omitempty"`
	ProcessList []ProcessTopVo `json:"process_list,omitempty"`
}
