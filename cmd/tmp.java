package com.dingtalk.api.request;

import com.dingtalk.api.response.C0001OapiProcessinstanceExecuteResponse;
import com.taobao.api.ApiRuleException;
import com.taobao.api.BaseTaobaoRequest;
import com.taobao.api.Constants;
import com.taobao.api.internal.util.RequestCheckUtils;
import com.taobao.api.internal.util.TaobaoHashMap;
import java.util.Map;

/* renamed from: com.dingtalk.api.request.OapiProcessinstanceExecuteRequest  reason: case insensitive filesystem */
public class C0000OapiProcessinstanceExecuteRequest extends BaseTaobaoRequest<C0001OapiProcessinstanceExecuteResponse> {
    private String actionerUserid;
    private String processInstanceId;
    private String remark;
    private String result;
    private Long taskId;
    private String topHttpMethod = "POST";
    private String topResponseType = Constants.RESPONSE_TYPE_DINGTALK_OAPI;

    public void setActionerUserid(String actionerUserid2) {
        this.actionerUserid = actionerUserid2;
    }

    public String getActionerUserid() {
        return this.actionerUserid;
    }

    public void setProcessInstanceId(String processInstanceId2) {
        this.processInstanceId = processInstanceId2;
    }

    public String getProcessInstanceId() {
        return this.processInstanceId;
    }

    public void setRemark(String remark2) {
        this.remark = remark2;
    }

    public String getRemark() {
        return this.remark;
    }

    public void setResult(String result2) {
        this.result = result2;
    }

    public String getResult() {
        return this.result;
    }

    public void setTaskId(Long taskId2) {
        this.taskId = taskId2;
    }

    public Long getTaskId() {
        return this.taskId;
    }

    @Override // com.taobao.api.BaseTaobaoRequest, com.taobao.api.TaobaoRequest
    public String getApiMethodName() {
        return "dingtalk.oapi.processinstance.execute";
    }

    @Override // com.taobao.api.BaseTaobaoRequest, com.taobao.api.TaobaoRequest
    public String getTopResponseType() {
        return this.topResponseType;
    }

    @Override // com.taobao.api.BaseTaobaoRequest, com.taobao.api.TaobaoRequest
    public void setTopResponseType(String topResponseType2) {
        this.topResponseType = topResponseType2;
    }

    @Override // com.taobao.api.BaseTaobaoRequest, com.taobao.api.TaobaoRequest
    public String getTopApiCallType() {
        return "oapi";
    }

    @Override // com.taobao.api.BaseTaobaoRequest, com.taobao.api.TaobaoRequest
    public String getTopHttpMethod() {
        return this.topHttpMethod;
    }

    @Override // com.taobao.api.BaseTaobaoRequest, com.taobao.api.TaobaoRequest
    public void setTopHttpMethod(String topHttpMethod2) {
        this.topHttpMethod = topHttpMethod2;
    }

    public void setHttpMethod(String httpMethod) {
        setTopHttpMethod(httpMethod);
    }

    @Override // com.taobao.api.TaobaoRequest
    public Map<String, String> getTextParams() {
        TaobaoHashMap txtParams = new TaobaoHashMap();
        txtParams.put("actioner_userid", this.actionerUserid);
        txtParams.put("process_instance_id", this.processInstanceId);
        txtParams.put("remark", this.remark);
        txtParams.put("result", this.result);
        txtParams.put("task_id", (Object) this.taskId);
        if (this.udfParams != null) {
            txtParams.putAll(this.udfParams);
        }
        return txtParams;
    }

    @Override // com.taobao.api.TaobaoRequest
    public Class<C0001OapiProcessinstanceExecuteResponse> getResponseClass() {
        return C0001OapiProcessinstanceExecuteResponse.class;
    }

    @Override // com.taobao.api.TaobaoRequest
    public void check() throws ApiRuleException {
        RequestCheckUtils.checkNotEmpty(this.actionerUserid, "actionerUserid");
        RequestCheckUtils.checkNotEmpty(this.processInstanceId, "processInstanceId");
        RequestCheckUtils.checkMaxLength(this.remark, 2000, "remark");
        RequestCheckUtils.checkNotEmpty(this.result, "result");
        RequestCheckUtils.checkNotEmpty(this.taskId, "taskId");
    }
}