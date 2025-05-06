package legs

import (
	"testing"
)

func TestSquat(t *testing.T) {
	t.Run("スクワット150kgのテスト", func(t *testing.T) {
		// Red: まず失敗するテスト
		want := 150.0
		maxWeight := 100.0 // 現在の最大重量

		// 目標の150kgを持ち上げられるようにトレーニングする
		trained := Squat(maxWeight, "150kg")

		if trained < want {
			t.Errorf("want %.2fkg, got %.2fkg", want, trained)
		}
	})

	t.Run("フルレンジの動きができているか", func(t *testing.T) {
		isFullRange := CheckFullRange(true) // 適切な深さまで下げる
		if !isFullRange {
			t.Error("適切な深さまでスクワットできていません")
		}
	})

	t.Run("正しいフォームでトレーニングできているか", func(t *testing.T) {
		isProperForm := CheckProperForm(true) // 膝と足先が同じ方向を向いている
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
