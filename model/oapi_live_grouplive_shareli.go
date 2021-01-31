package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiLiveGroupliveSharelistRequest() *OapiLiveGroupliveSharelistRequest {
	return &OapiLiveGroupliveSharelistRequest{}
}

type OapiLiveGroupliveSharelistRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiLiveGroupliveSharelistResponse
	Cid             string
	LiveUuid        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiLiveGroupliveSharelistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiLiveGroupliveSharelistRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiLiveGroupliveSharelistRequest) SetCid(cid2 string) {
	this.Cid = cid2
}
func (this *OapiLiveGroupliveSharelistRequest) GetCid() string {
	return this.Cid
}
func (this *OapiLiveGroupliveSharelistRequest) SetLiveUuid(liveUuid2 string) {
	this.LiveUuid = liveUuid2
}
func (this *OapiLiveGroupliveSharelistRequest) GetLiveUuid() string {
	return this.LiveUuid
}
func (this *OapiLiveGroupliveSharelistRequest) GetApiMethodName() string {
	return "dingtalk.oapi.live.grouplive.sharelist"
}
func (this *OapiLiveGroupliveSharelistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiLiveGroupliveSharelistRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiLiveGroupliveSharelistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiLiveGroupliveSharelistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiLiveGroupliveSharelistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cid"] = this.Cid
	txtParams["live_uuid"] = this.LiveUuid
	return txtParams
}
func (this *OapiLiveGroupliveSharelistRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cid, "cid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiLiveGroupliveSharelistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiLiveGroupliveSharelistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiLiveGroupliveSharelistResponse struct {
	taobao.TaobaoResponse
	Errcode      int64    `json:"errcode,omitempty"`
	Errmsg       string   `json:"errmsg,omitempty"`
	ShareCidList []string `json:"share_cid_list,omitempty"`
}
