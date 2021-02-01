package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiExtcontactListlabelgroupsRequest() *OapiExtcontactListlabelgroupsRequest {
	return &OapiExtcontactListlabelgroupsRequest{}
}

type OapiExtcontactListlabelgroupsRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiExtcontactListlabelgroupsResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiExtcontactListlabelgroupsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiExtcontactListlabelgroupsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiExtcontactListlabelgroupsRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiExtcontactListlabelgroupsRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiExtcontactListlabelgroupsRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiExtcontactListlabelgroupsRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiExtcontactListlabelgroupsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.extcontact.listlabelgroups"
}
func (this *OapiExtcontactListlabelgroupsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiExtcontactListlabelgroupsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiExtcontactListlabelgroupsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiExtcontactListlabelgroupsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiExtcontactListlabelgroupsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiExtcontactListlabelgroupsRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckMaxValue(int(this.Size), 100, "size"); err != nil {
		return err
	}
	return nil
}
func (this *OapiExtcontactListlabelgroupsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiExtcontactListlabelgroupsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiExtcontactListlabelgroupsResponse struct {
	taobao.TaobaoResponse
	Results []OpenLabelGroup `json:"results,omitempty"`
}
