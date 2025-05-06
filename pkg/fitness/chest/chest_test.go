package chest

import (
	"testing"
)

func TestBenchPress(t *testing.T) {
	t.Run("ベンチプレス100kgのテスト", func(t *testing.T) {
		// Red: まず失敗するテスト
		want := 100.0
		maxWeight := 70.0 // 現在の最大重量

		// 目標の100kgを持ち上げられるようにトレーニングする
		trained := BenchPress(maxWeight, "100kg")

		if trained < want {
			t.Errorf("want %.2fkg, got %.2fkg", want, trained)
		}
	})

	t.Run("フルレンジの動きができているか", func(t *testing.T) {
		isFullRange := CheckFullRange(true) // バーベルが胸につくこと
		if !isFullRange {
			t.Error("バーベルが胸につく完全な動きができていません")
		}
	})

	t.Run("正しいフォームでトレーニングできているか", func(t *testing.T) {
		isProperForm := CheckProperForm(true) // 肩甲骨の下方回旋が適切
		if !isProperForm {
			t.Error("正しいフォームでトレーニングできていません")
		}
	})

	t.Run("安全性が確保されているか", func(t *testing.T) {
		isSafe := CheckSafety(true) // 怪我や故障が発生しない
		if !isSafe {
			t.Error("トレーニングの安全性が確保されていません")
		}
	})
}

func TestBenchPressWithUserInfo(t *testing.T) {
	t.Run("目標重量達成とプロテイン摂取量が十分な場合", func(t *testing.T) {
		userInfo := UserInfo{
			CurrentBenchWeight: 70.0,
			TargetBenchWeight:  100.0,
			BodyWeight:         80.0,
			DailyProteinIntake: 180.0, // 80kg * 2g以上 = 160g以上
			CarbIntake:         300.0,
			FatIntake:          60.0,
		}

		result, err := BenchPressWithUserInfo(userInfo)
		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}

		if result < userInfo.TargetBenchWeight {
			t.Errorf("目標重量に達していません: want %.2fkg, got %.2fkg", userInfo.TargetBenchWeight, result)
		}
	})

	t.Run("プロテイン摂取量不足の場合", func(t *testing.T) {
		userInfo := UserInfo{
			CurrentBenchWeight: 70.0,
			TargetBenchWeight:  100.0,
			BodyWeight:         80.0,
			DailyProteinIntake: 120.0, // 80kg * 2g = 160gに満たない
			CarbIntake:         250.0,
			FatIntake:          50.0,
		}

		_, err := BenchPressWithUserInfo(userInfo)
		if err == nil {
			t.Error("プロテイン摂取量不足エラーが検出されませんでした")
		}
	})

	t.Run("目標重量が達成できない長期間が必要な場合", func(t *testing.T) {
		userInfo := UserInfo{
			CurrentBenchWeight: 50.0,
			TargetBenchWeight:  120.0, // 70kgの増加が必要 → 28週間（12週間のプログラムを超える）
			BodyWeight:         75.0,
			DailyProteinIntake: 160.0, // 75kg * 2g以上 = 150g以上
			CarbIntake:         280.0,
			FatIntake:          55.0,
		}

		_, err := BenchPressWithUserInfo(userInfo)
		if err == nil {
			t.Error("目標重量達成不可エラーが検出されませんでした")
		}
	})
}
