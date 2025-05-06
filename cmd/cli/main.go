package main

import (
	"fmt"
	"os"

	"go-kinniku/pkg/fitness/back"
	"go-kinniku/pkg/fitness/chest"
	"go-kinniku/pkg/fitness/legs"
	"go-kinniku/pkg/fitness/shoulders"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "chest":
		trainChest()
	case "back":
		trainBack()
	case "legs":
		trainLegs()
	case "shoulders":
		trainShoulders()
	case "program":
		showFullProgram()
	case "user-chest":
		trainChestWithUserInfo()
	default:
		fmt.Println("無効なコマンドです")
		printUsage()
	}
}

func printUsage() {
	fmt.Println("使用方法: go-kinniku [コマンド]")
	fmt.Println("コマンド:")
	fmt.Println("  chest     - 胸のトレーニングプログラムを実行")
	fmt.Println("  back      - 背中のトレーニングプログラムを実行")
	fmt.Println("  legs      - 脚のトレーニングプログラムを実行")
	fmt.Println("  shoulders - 肩のトレーニングプログラムを実行")
	fmt.Println("  program   - フルトレーニングプログラムを表示")
	fmt.Println("  user-chest - ユーザー情報を使用した胸のトレーニングプログラムを実行")
}

func trainChest() {
	fmt.Println("===== 胸のトレーニングプログラム =====")
	fmt.Println("現在の最大重量: 70kg")
	fmt.Println("目標重量: 100kg")

	// テスト実行
	result := chest.BenchPress(70.0, "100kg")
	fmt.Printf("トレーニング結果: %.1fkg\n", result)

	if chest.CheckFullRange(true) {
		fmt.Println("✓ フルレンジの動きができています")
	}

	if chest.CheckProperForm(true) {
		fmt.Println("✓ 正しいフォームでトレーニングできています")
	}

	if chest.CheckSafety(true) {
		fmt.Println("✓ トレーニングの安全性が確保されています")
	}
}

func trainChestWithUserInfo() {
	fmt.Println("===== ユーザー情報を使用した胸のトレーニングプログラム =====")

	// ユーザー情報の作成
	userInfo := chest.UserInfo{
		CurrentBenchWeight: 70.0,
		TargetBenchWeight:  100.0,
		BodyWeight:         80.0,
		DailyProteinIntake: 180.0,
		CarbIntake:         300.0,
		FatIntake:          60.0,
	}

	fmt.Printf("現在のベンチプレス重量: %.1fkg\n", userInfo.CurrentBenchWeight)
	fmt.Printf("目標ベンチプレス重量: %.1fkg\n", userInfo.TargetBenchWeight)
	fmt.Printf("体重: %.1fkg\n", userInfo.BodyWeight)
	fmt.Printf("1日のプロテイン摂取量: %.1fg\n", userInfo.DailyProteinIntake)

	// トレーニング実行
	result, err := chest.BenchPressWithUserInfo(userInfo)

	if err != nil {
		fmt.Printf("トレーニング結果: %.1fkg\n", result)
		fmt.Printf("エラー: %v\n", err)
	} else {
		fmt.Printf("トレーニング結果: %.1fkg\n", result)
		fmt.Println("✓ 目標重量を達成しました！")

		if chest.CheckFullRange(true) {
			fmt.Println("✓ フルレンジの動きができています")
		}

		if chest.CheckProperForm(true) {
			fmt.Println("✓ 正しいフォームでトレーニングできています")
		}

		if chest.CheckSafety(true) {
			fmt.Println("✓ トレーニングの安全性が確保されています")
		}
	}
}

func trainBack() {
	fmt.Println("===== 背中のトレーニングプログラム =====")
	fmt.Println("現在の最大重量: 80kg")
	fmt.Println("目標重量: 120kg")

	// テスト実行
	result := back.Deadlift(80.0, "120kg")
	fmt.Printf("トレーニング結果: %.1fkg\n", result)

	if back.CheckFullRange(true) {
		fmt.Println("✓ フルレンジの動きができています")
	}

	if back.CheckProperForm(true) {
		fmt.Println("✓ 正しいフォームでトレーニングできています")
	}

	if back.CheckSafety(true) {
		fmt.Println("✓ トレーニングの安全性が確保されています")
	}
}

func trainLegs() {
	fmt.Println("===== 脚のトレーニングプログラム =====")
	fmt.Println("現在の最大重量: 100kg")
	fmt.Println("目標重量: 150kg")

	// テスト実行
	result := legs.Squat(100.0, "150kg")
	fmt.Printf("トレーニング結果: %.1fkg\n", result)

	if legs.CheckFullRange(true) {
		fmt.Println("✓ フルレンジの動きができています")
	}

	if legs.CheckProperForm(true) {
		fmt.Println("✓ 正しいフォームでトレーニングできています")
	}

	if legs.CheckSafety(true) {
		fmt.Println("✓ トレーニングの安全性が確保されています")
	}
}

func trainShoulders() {
	fmt.Println("===== 肩のトレーニングプログラム =====")
	fmt.Println("現在の最大重量: 50kg")
	fmt.Println("目標重量: 80kg")

	// テスト実行
	result := shoulders.OverheadPress(50.0, "80kg")
	fmt.Printf("トレーニング結果: %.1fkg\n", result)

	if shoulders.CheckFullRange(true) {
		fmt.Println("✓ フルレンジの動きができています")
	}

	if shoulders.CheckProperForm(true) {
		fmt.Println("✓ 正しいフォームでトレーニングできています")
	}

	if shoulders.CheckSafety(true) {
		fmt.Println("✓ トレーニングの安全性が確保されています")
	}
}

func showFullProgram() {
	fmt.Println("======= 筋肉TDDトレーニングプログラム =======")
	fmt.Println("1. テスト（Red）： 目標とするパフォーマンスを定義")
	fmt.Println("   - 目標重量を設定")
	fmt.Println("   - 適切なフォームを定義")
	fmt.Println("   - 安全性を確保")

	fmt.Println("\n2. 実装（Green）： 目標達成のための最小限のトレーニング")
	fmt.Println("   - 基本エクササイズの実施")
	fmt.Println("   - 段階的な負荷増加プログラム")
	fmt.Println("   - 適切な栄養摂取と回復戦略")

	fmt.Println("\n3. リファクタリング： トレーニング効率の向上")
	fmt.Println("   - フォームの最適化")
	fmt.Println("   - 補助種目の追加")
	fmt.Println("   - 回復戦略の改善")
	fmt.Println("   - さらなる目標設定")

	fmt.Println("\n週間トレーニングスケジュール:")
	fmt.Println("月曜日: 胸（ベンチプレスと補助種目）")
	fmt.Println("火曜日: 背中（デッドリフトと補助種目）")
	fmt.Println("水曜日: 休息と回復")
	fmt.Println("木曜日: 脚（スクワットと補助種目）")
	fmt.Println("金曜日: 肩（オーバーヘッドプレスと補助種目）")
	fmt.Println("土曜日: 腕と補助トレーニング")
	fmt.Println("日曜日: 完全休息")
}
