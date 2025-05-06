package legs

import (
	"fmt"
	"strconv"
	"strings"
)

// TrainingPlan は脚トレーニングのプランを表す構造体
type TrainingPlan struct {
	WeeklyIncrease     float64  // 週あたりの重量増加 (kg)
	ProteinIntake      float64  // 1日のタンパク質摂取量 (g)
	RecoveryDays       int      // 回復のための休息日数
	TrainingDuration   int      // トレーニング期間（週）
	AccessoryExercises []string // 補助種目
	CarboIntake        float64  // 炭水化物摂取量（g/kg体重）
}

// DefaultTrainingPlan はデフォルトのトレーニングプランを返します
func DefaultTrainingPlan() TrainingPlan {
	return TrainingPlan{
		WeeklyIncrease:   7.5, // 週あたり7.5kgずつ増加（脚は最大の筋肉群）
		ProteinIntake:    2.5, // 体重1kgあたり2.5gのタンパク質
		RecoveryDays:     4,   // 脚トレの間に4日休む
		TrainingDuration: 8,   // 8週間のプログラム
		CarboIntake:      4.0, // 体重1kgあたり4.0gの炭水化物
		AccessoryExercises: []string{
			"レッグプレス",
			"レッグエクステンション",
			"レッグカール",
			"カーフレイズ",
		},
	}
}

// Squat はスクワットのトレーニングを行い、最終的な最大重量を返します
// 1. 大腿四頭筋、ハムストリングス、臀筋を強化する基本エクササイズ
// 2. 段階的な重量増加プログラム
// 3. 適切なタンパク質・炭水化物摂取と回復戦略
// 4. コア筋群強化によるフォーム安定化
func Squat(currentMaxWeight float64, targetWeightStr string) float64 {
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

	// 栄養摂取の推奨量を計算（80kg体重を想定）
	bodyWeight := 80.0
	dailyProtein := bodyWeight * plan.ProteinIntake
	dailyCarbs := bodyWeight * plan.CarboIntake

	fmt.Printf("推奨栄養摂取量: タンパク質 %.1fg/日, 炭水化物 %.1fg/日\n", dailyProtein, dailyCarbs)

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
func CheckFullRange(isProperDepth bool) bool {
	// 適切な深さまでスクワットできているかを確認
	return isProperDepth
}

// CheckProperForm は適切なフォームでトレーニングできているかチェックする
func CheckProperForm(isKneeAligned bool) bool {
	// 膝と足先が同じ方向を向いているかを確認
	return isKneeAligned
}

// CheckSafety はトレーニングの安全性を確認する
func CheckSafety(isNoInjury bool) bool {
	// 身体の怪我や故障が発生していないことを確認
	return isNoInjury
}
