package config

import "kis-flow/common"

// KisFlowFunctionParam 一个Flow配置中Function的Id及携带固定配置参数
type KisFlowFunctionParam struct {
	FuncName string `yaml:"fname"`  //必须
	Params   FParam `yaml:"params"` //选填,在当前Flow中Function定制固定配置参数
}

// KisFlowConfig 用户贯穿整条流式计算上下文环境的对象
type KisFlowConfig struct {
	KisType  string                 `yaml:"kistype"`
	Status   int                    `yaml:"status"`
	FlowName string                 `yaml:"flow_name"`
	Flows    []KisFlowFunctionParam `yaml:"flows"`
}

// NewFlowConfig 创建一个Flow策略配置对象, 用于描述一个KisFlow信息
func NewFlowConfig(flowName string, enable common.KisOnOff) *KisFlowConfig {
	config := new(KisFlowConfig)
	config.FlowName = flowName
	config.Flows = make([]KisFlowFunctionParam, 0)

	config.Status = int(enable)

	return config
}

// AppendFunctionConfig 添加一个Function Config 到当前Flow中
//有关flow携带的Function配置，这里我们采用通过AppendFunctionConfig动态的去添加，
//目的是为了，今后可能有关kisflow的配置会从数据库/动态远程配置等中提取，那么就需要动态的将配置组合进来。
func (fConfig *KisFlowConfig) AppendFunctionConfig(params KisFlowFunctionParam) {
	fConfig.Flows = append(fConfig.Flows, params)
}
