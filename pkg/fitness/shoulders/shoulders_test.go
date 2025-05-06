package shoulders

import (
	"testing"
)

func TestOverheadPress(t *testing.T) {
	t.Run("オーバーヘッドプレス80kgのテスト", func(t *testing.T) {
		// Red: まず失敗するテスト
		want := 80.0
		maxWeight := 50.0 // 現在の最大重量

		// 目標の80kgを持ち上げられるようにトレーニングする
		trained := OverheadPress(maxWeight, "80kg")

		if trained < want {
			t.Errorf("want %.2fkg, got %.2fkg", want, trained)
		}
	})

	t.Run("フルレンジの動きができているか", func(t *testing.T) {
		isFullRange := CheckFullRange(true) // バーベルを頭上まで伸ばす
		if !isFullRange {
			t.Error("バーベルを頭上まで完全に伸ばせていません")
		}
	})

	t.Run("正しいフォームでトレーニングできているか", func(t *testing.T) {
		isProperForm := CheckProperForm(true) // コアを安定させて脊椎をニュートラルに
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
