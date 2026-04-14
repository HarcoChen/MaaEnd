package buygoodselastic

import (
	maa "github.com/MaaXYZ/maa-framework-go/v4"
	"github.com/rs/zerolog/log"
)

// Register 注册 buygoodselastic 包提供的自定义动作与识别器。
func Register() {
	if err := InitItemMap("zh_cn"); err != nil {
		log.Warn().
			Err(err).
			Str("component", buyGoodsElasticComponent).
			Msg("failed to init item map during registration, OCR name matching may be disabled")
	}

	maa.AgentServerRegisterCustomAction(buyGoodsElasticSelectItemActionName, &SelectItemAction{})
	maa.AgentServerRegisterCustomAction(buyGoodsElasticReconcileDecisionActionName, &ReconcileDecisionAction{})
	maa.AgentServerRegisterCustomRecognition(buyGoodsElasticRecognitionName, &ItemValueChangeRecognition{})
}
