package chest

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// TrainingPlan は胸部トレーニングのプランを表す構造体
type TrainingPlan struct {
	WeeklyIncrease   float64 // 週あたりの重量増加 (kg)
	ProteinIntake    float64 // 1日のタンパク質摂取量 (g)
	RecoveryDays     int     // 回復のための休息日数
	TrainingDuration int     // トレーニング期間（週）
}

// UserInfo はユーザーの情報を表す構造体
type UserInfo struct {
	CurrentBenchWeight float64 // 現在のベンチプレス重量 (kg)
	TargetBenchWeight  float64 // 目標のベンチプレス重量 (kg)
	BodyWeight         float64 // 現在の体重 (kg)
	DailyProteinIntake float64 // 1日のプロテイン摂取量 (g)
	CarbIntake         float64 // 炭水化物摂取量 (g)
	FatIntake          float64 // 脂質摂取量 (g)
}

// DefaultTrainingPlan はデフォルトのトレーニングプランを返します
func DefaultTrainingPlan() TrainingPlan {
	return TrainingPlan{
		WeeklyIncrease:   2.5, // 週あたり2.5kgずつ増加
		ProteinIntake:    2.0, // 体重1kgあたり2gのタンパク質
		RecoveryDays:     2,   // 胸トレの間に2日休む
		TrainingDuration: 12,  // 12週間のプログラム
	}
}

// BenchPress はベンチプレスのトレーニングを行い、最終的な最大重量を返します
// 1. 胸筋、三角筋、上腕三頭筋を強化する基本エクササイズ
// 2. 段階的な重量増加プログラム
// 3. 適切なタンパク質摂取と回復戦略
// 4. コア筋群強化によるフォーム安定化
func BenchPress(currentMaxWeight float64, targetWeightStr string) float64 {
	// Refactor: より実践的な実装
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
		return currentMaxWeight + float64(plan.TrainingDuration)*plan.WeeklyIncrease
	}

	// 目標を達成
	return targetWeight
}

// BenchPressWithUserInfo はユーザー情報を使用してベンチプレスのトレーニングを行います
// エラー条件：
// 1. ベンチプレスの重量が目標重量に達しなかった
// 2. 必要なプロテイン摂取量を満たしていない
func BenchPressWithUserInfo(userInfo UserInfo) (float64, error) {
	plan := DefaultTrainingPlan()

	// 必要なプロテイン摂取量を計算 (体重あたり2g)
	requiredProtein := userInfo.BodyWeight * 2.0
	totalProteinIntake := userInfo.DailyProteinIntake

	// プロテイン摂取量のチェック
	if totalProteinIntake < requiredProtein {
		return userInfo.CurrentBenchWeight, errors.New(fmt.Sprintf(
			"必要なプロテイン摂取量が不足しています: 必要量 %.1fg, 現在の摂取量 %.1fg",
			requiredProtein, totalProteinIntake))
	}

	// 必要なトレーニング期間を計算
	requiredWeeks := (userInfo.TargetBenchWeight - userInfo.CurrentBenchWeight) / plan.WeeklyIncrease

	// トレーニング結果を計算
	var result float64
	// トレーニング期間が計画よりも長ければ、より強度の高いプランが必要
	if requiredWeeks > float64(plan.TrainingDuration) {
		fmt.Printf("目標達成には標準的な計画より長い期間 (%.1f週) が必要です\n", requiredWeeks)
		result = userInfo.CurrentBenchWeight + float64(plan.TrainingDuration)*plan.WeeklyIncrease
	} else {
		// 目標達成が可能な場合
		result = userInfo.TargetBenchWeight
	}

	// PFC（タンパク質、脂質、炭水化物）バランスの表示
	fmt.Printf("PFCバランス - プロテイン: %.1fg, 脂質: %.1fg, 炭水化物: %.1fg\n",
		totalProteinIntake, userInfo.FatIntake, userInfo.CarbIntake)

	// 結果が目標重量より低ければエラー
	if result < userInfo.TargetBenchWeight {
		return result, errors.New(fmt.Sprintf(
			"目標重量 %.1fkg に達成できませんでした。達成重量: %.1fkg",
			userInfo.TargetBenchWeight, result))
	}

	return result, nil
}

func calcProteinIntake(currentMaxWeight float64) float64 {
	plan := DefaultTrainingPlan()
	return currentMaxWeight * plan.ProteinIntake
}

// CheckFullRange はフルレンジの動きができているかチェックする
func CheckFullRange(isBarTouchingChest bool) bool {
	// バーベルが胸につくことを確認
	return isBarTouchingChest
}

// CheckProperForm は適切なフォームでトレーニングできているかチェックする
func CheckProperForm(isScapulaRetracted bool) bool {
	// 肩甲骨の下方回旋が適切に行われていることを確認
	return isScapulaRetracted
}

// CheckSafety はトレーニングの安全性を確認する
func CheckSafety(isNoInjury bool) bool {
	// 身体の怪我や故障が発生していないことを確認
	return isNoInjury
}
