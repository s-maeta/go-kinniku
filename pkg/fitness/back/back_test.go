package back

import (
	"testing"
)

func TestDeadlift(t *testing.T) {
	t.Run("デッドリフト120kgのテスト", func(t *testing.T) {
		// Red: まず失敗するテスト
		want := 120.0
		maxWeight := 80.0 // 現在の最大重量

		// 目標の120kgを持ち上げられるようにトレーニングする
		trained := Deadlift(maxWeight, "120kg")

		if trained < want {
			t.Errorf("want %.2fkg, got %.2fkg", want, trained)
		}
	})

	t.Run("フルレンジの動きができているか", func(t *testing.T) {
		isFullRange := CheckFullRange(true) // 完全な可動域
		if !isFullRange {
			t.Error("完全な可動域でトレーニングできていません")
		}
	})

	t.Run("正しいフォームでトレーニングできているか", func(t *testing.T) {
		isProperForm := CheckProperForm(true) // 背中のニュートラルポジション
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
