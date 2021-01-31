package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewIsvBlazersGeneratecodeRequest() *IsvBlazersGeneratecodeRequest {
	return &IsvBlazersGeneratecodeRequest{}
}

type IsvBlazersGeneratecodeRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            IsvBlazersGeneratecodeResponse
	BizId           string
	Ext             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *IsvBlazersGeneratecodeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *IsvBlazersGeneratecodeRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *IsvBlazersGeneratecodeRequest) GetBizId() string {
	return this.BizId
}
func (this *IsvBlazersGeneratecodeRequest) SetExt(ext2 string) {
	this.Ext = ext2
}
func (this *IsvBlazersGeneratecodeRequest) SetExtString(ext2 string) {
	this.Ext = ext2
}
func (this *IsvBlazersGeneratecodeRequest) GetExt() string {
	return this.Ext
}
func (this *IsvBlazersGeneratecodeRequest) GetApiMethodName() string {
	return "dingtalk.isv.blazers.generatecode"
}
func (this *IsvBlazersGeneratecodeRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *IsvBlazersGeneratecodeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *IsvBlazersGeneratecodeRequest) GetTopApiCallType() string {
	return "top"
}
func (this *IsvBlazersGeneratecodeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *IsvBlazersGeneratecodeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *IsvBlazersGeneratecodeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_id"] = this.BizId
	txtParams["ext"] = this.Ext
	return txtParams
}
func (this *IsvBlazersGeneratecodeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizId, "bizId"); err != nil {
		return err
	}
	return nil
}
func (this *IsvBlazersGeneratecodeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *IsvBlazersGeneratecodeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type IsvBlazersGeneratecodeResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
