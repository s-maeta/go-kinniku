package shoulders

import (
	"fmt"
	"strconv"
	"strings"
)

// TrainingPlan は肩トレーニングのプランを表す構造体
type TrainingPlan struct {
	WeeklyIncrease     float64  // 週あたりの重量増加 (kg)
	ProteinIntake      float64  // 1日のタンパク質摂取量 (g)
	RecoveryDays       int      // 回復のための休息日数
	TrainingDuration   int      // トレーニング期間（週）
	AccessoryExercises []string // 補助種目
	MobilityWork       []string // 可動域改善エクササイズ
}

// DefaultTrainingPlan はデフォルトのトレーニングプランを返します
func DefaultTrainingPlan() TrainingPlan {
	return TrainingPlan{
		WeeklyIncrease:   2.0, // 週あたり2.0kgずつ増加（肩は小さな筋肉群）
		ProteinIntake:    2.0, // 体重1kgあたり2.0gのタンパク質
		RecoveryDays:     2,   // 肩トレの間に2日休む
		TrainingDuration: 15,  // 15週間のプログラム
		AccessoryExercises: []string{
			"サイドレイズ",
			"フロントレイズ",
			"フェイスプル",
			"ショルダープレス（ダンベル）",
		},
		MobilityWork: []string{
			"胸椎回旋ストレッチ",
			"肩甲骨円運動",
			"ウォールスライド",
		},
	}
}

// OverheadPress はオーバーヘッドプレスのトレーニングを行い、最終的な最大重量を返します
// 1. 三角筋、僧帽筋、上腕三頭筋を強化する基本エクササイズ
// 2. 段階的な重量増加プログラム
// 3. 適切なタンパク質摂取と回復戦略
// 4. 肩の可動域を改善するモビリティワーク
func OverheadPress(currentMaxWeight float64, targetWeightStr string) float64 {
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

	fmt.Println("肩のケガ防止とパフォーマンス向上のための可動域改善エクササイズ:")
	for _, mobility := range plan.MobilityWork {
		fmt.Printf("- %s\n", mobility)
	}

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
func CheckFullRange(isFullExtension bool) bool {
	// バーベルを頭上まで完全に伸ばせているかを確認
	return isFullExtension
}

// CheckProperForm は適切なフォームでトレーニングできているかチェックする
func CheckProperForm(isCoreStable bool) bool {
	// コアを安定させて脊椎をニュートラルにできているかを確認
	return isCoreStable
}

// CheckSafety はトレーニングの安全性を確認する
func CheckSafety(isNoInjury bool) bool {
	// 身体の怪我や故障が発生していないことを確認
	return isNoInjury
}
