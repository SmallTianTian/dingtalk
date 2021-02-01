package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiConnectorOpenRequest() *OapiConnectorOpenRequest {
	return &OapiConnectorOpenRequest{}
}

type OapiConnectorOpenRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiConnectorOpenResponse
	ConnectorId     string
	CorpId          string
	CreatorUserid   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiConnectorOpenRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiConnectorOpenRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiConnectorOpenRequest) SetConnectorId(connectorId2 string) {
	this.ConnectorId = connectorId2
}
func (this *OapiConnectorOpenRequest) GetConnectorId() string {
	return this.ConnectorId
}
func (this *OapiConnectorOpenRequest) SetCorpId(corpId2 string) {
	this.CorpId = corpId2
}
func (this *OapiConnectorOpenRequest) GetCorpId() string {
	return this.CorpId
}
func (this *OapiConnectorOpenRequest) SetCreatorUserid(creatorUserid2 string) {
	this.CreatorUserid = creatorUserid2
}
func (this *OapiConnectorOpenRequest) GetCreatorUserid() string {
	return this.CreatorUserid
}
func (this *OapiConnectorOpenRequest) GetApiMethodName() string {
	return "dingtalk.oapi.connector.open"
}
func (this *OapiConnectorOpenRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiConnectorOpenRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiConnectorOpenRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiConnectorOpenRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiConnectorOpenRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["connector_id"] = this.ConnectorId
	txtParams["corp_id"] = this.CorpId
	txtParams["creator_userid"] = this.CreatorUserid
	return txtParams
}
func (this *OapiConnectorOpenRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ConnectorId, "connectorId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiConnectorOpenRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiConnectorOpenRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiConnectorOpenResponse struct {
	taobao.TaobaoResponse
}
