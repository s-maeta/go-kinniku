package back

import (
	"fmt"
	"strconv"
	"strings"
)

// TrainingPlan は背中トレーニングのプランを表す構造体
type TrainingPlan struct {
	WeeklyIncrease     float64  // 週あたりの重量増加 (kg)
	ProteinIntake      float64  // 1日のタンパク質摂取量 (g)
	RecoveryDays       int      // 回復のための休息日数
	TrainingDuration   int      // トレーニング期間（週）
	AccessoryExercises []string // 補助種目
}

// DefaultTrainingPlan はデフォルトのトレーニングプランを返します
func DefaultTrainingPlan() TrainingPlan {
	return TrainingPlan{
		WeeklyIncrease:   5.0, // 週あたり5.0kgずつ増加（背中はより大きな筋肉群）
		ProteinIntake:    2.2, // 体重1kgあたり2.2gのタンパク質
		RecoveryDays:     3,   // 背中トレの間に3日休む
		TrainingDuration: 10,  // 10週間のプログラム
		AccessoryExercises: []string{
			"ラットプルダウン",
			"ベントオーバーロウ",
			"プルアップ",
			"フェイスプル",
		},
	}
}

// Deadlift はデッドリフトのトレーニングを行い、最終的な最大重量を返します
// 1. 広背筋、僧帽筋、脊柱起立筋を強化する基本エクササイズ
// 2. 段階的な重量増加プログラム
// 3. 適切なタンパク質摂取と回復戦略
// 4. コア筋群強化によるフォーム安定化
func Deadlift(currentMaxWeight float64, targetWeightStr string) float64 {
	plan := DefaultTrainingPlan()

	// 目標重量を数値に変換
	targetWeightStr = strings.TrimSuffix(targetWeightStr, "kg")
	targetWeight, err := strconv.ParseFloat(targetWeightStr, 64)
	if err != nil {
		fmt.Println("目標重量の形式が正しくありません:", err)
		return currentMaxWeight
	}

	// 必要なトレーニング期間を計算
	requiredWeeks := (targetWeight - currentMaxWeight) / plan.WeeklyIncrease

	// トレーニング期間が計画よりも長ければ、より強度の高いプランが必要
	if requiredWeeks > float64(plan.TrainingDuration) {
		fmt.Printf("目標達成には標準的な計画より長い期間 (%.1f週) が必要です\n", requiredWeeks)
		// 補助種目を追加して効率を上げる
		fmt.Println("効率向上のために以下の補助種目を追加します:")
		for _, exercise := range plan.AccessoryExercises {
			fmt.Printf("- %s\n", exercise)
		}
		return currentMaxWeight + float64(plan.TrainingDuration)*plan.WeeklyIncrease
	}

	// 目標を達成
	return targetWeight
}

// CheckFullRange はフルレンジの動きができているかチェックする
func CheckFullRange(isFullRangeMotion bool) bool {
	// 完全な可動域を確認
	return isFullRangeMotion
}

// CheckProperForm は適切なフォームでトレーニングできているかチェックする
func CheckProperForm(isBackNeutral bool) bool {
	// 背中のニュートラルポジションを確認
	return isBackNeutral
}

// CheckSafety はトレーニングの安全性を確認する
func CheckSafety(isNoInjury bool) bool {
	// 身体の怪我や故障が発生していないことを確認
	return isNoInjury
}
