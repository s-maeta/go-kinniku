---
marp: true
theme: default
paginate: true
backgroundColor: #fff
header: "Go言語で学ぶテスト駆動(筋肉)開発"
style: |
  section {
    font-size: 28px;
  }
  h1 {
    color: #0070c0;
    font-size: 40px;
  }
  h2 {
    color: #0070c0;
    border-bottom: 1px solid #0070c0;
    padding-bottom: 5px;
  }
  table {
    font-size: 24px;
  }
  code {
    background-color: #f8f8f8;
    color: #d73a49;
  }
---

<!-- タイトルスライド -->

# Go 言語で学ぶテスト駆動(筋肉)開発

<style scoped>
h1 {
  font-size: 60px;
  text-align: center;
  margin-top: 160px;
}
p {
  text-align: center;
  font-size: 24px;
  margin-top: 80px;
}
</style>

前田祥貴

---

<!-- テーマスライド -->

## 主なテーマ

<style scoped>
p {
  font-size: 28px;
  line-height: 1.8;
  margin-top: 40px;
}
</style>

システム開発 ≒ 筋肉開発

システム開発の TDD プロセスと比較しながら、筋トレにおいて自分の体でテストを行い重量を増加させていく過程を解説します。

---

<!-- 目的スライド -->

## 目的

<style scoped>
ul {
  font-size: 30px;
  line-height: 2.0;
  margin-top: 60px;
}
</style>

1. TDD の概念を通じて開発（筋肉）の段階的な進め方を知っていただく
2. （おこがましいですが）前田（まえた）の認知を広げる

---

<!-- TDDの流れ概要スライド -->

## システム開発と筋肉 TDD の流れ

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 450px;
}
.row {
  display: flex;
  width: 100%;
  justify-content: center;
  margin-bottom: 20px;
}
.box {
  width: 250px;
  height: 110px;
  margin: 0 20px;
  padding: 15px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border-radius: 10px;
  font-weight: bold;
  text-align: center;
}
.arrow {
  font-size: 24px;
  margin-top: 5px;
}
.red { background-color: rgba(255, 99, 71, 0.7); }
.green { background-color: rgba(144, 238, 144, 0.7); }
.blue { background-color: rgba(135, 206, 235, 0.7); }
.label {
  font-weight: bold;
  font-size: 20px;
  margin-bottom: 5px;
  color: #333;
}
</style>

<div class="container">
  <div class="label">システム開発</div>
  <div class="row">
    <div class="box red">
      Red<br>
      失敗するテストコードを書く
    </div>
    <div class="box green">
      Green<br>
      テスト通過の実装をする
    </div>
    <div class="box blue">
      Refactor<br>
      コードの品質を改善する
    </div>
  </div>
  <div class="arrow">⇅</div>
  <div class="label">筋肉開発</div>
  <div class="row">
    <div class="box red">
      Red<br>
      目標重量を設定し、失敗を確認
    </div>
    <div class="box green">
      Green<br>
      トレーニングで目標達成
    </div>
    <div class="box blue">
      Refactor<br>
      フォームと効率の向上
    </div>
  </div>
</div>

---

<!-- Red段階の詳細対比スライド -->

## Red: 失敗を明確にする段階

<style scoped>
.container {
  display: flex;
  margin-top: 20px;
  height: 400px;
}
.column {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: 0 10px;
}
.title {
  text-align: center;
  font-weight: bold;
  padding: 10px;
  border-radius: 5px;
  margin-bottom: 15px;
}
.muscle-title { background-color: rgba(255, 99, 71, 0.3); }
.dev-title { background-color: rgba(255, 99, 71, 0.3); }
.point {
  background-color: #f8f8f8;
  padding: 10px;
  margin-bottom: 10px;
  border-left: 4px solid #ff6347;
  flex: 1;
}
.center-text {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  font-size: 24px;
  margin: 0 10px;
}
</style>

<div class="container">
  <div class="column">
    <div class="title dev-title">開発: Red</div>
    <div class="point">達成すべき要件を明確に定義（例: ユーザー認証機能）</div>
    <div class="point">要件を検証する具体的なテストコードを先に作成</div>
    <div class="point">まだ実装していないので「失敗」するテストを確認</div>
    <div class="point">テスト失敗の内容から必要な実装内容を明確化</div>
  </div>
  
  <div class="center-text">
    <div>同じ</div>
    <div>概念</div>
    <div>↔</div>
  </div>
  
  <div class="column">
    <div class="title muscle-title">筋トレ: Red</div>
    <div class="point">達成すべき目標を明確に定義（例: ベンチプレス100kg）</div>
    <div class="point">目標達成の基準を設定（正しいフォーム、完全な可動域）</div>
    <div class="point">現状の能力で目標に「挑戦して失敗」を確認</div>
    <div class="point">失敗の内容から必要なトレーニング内容を明確化</div>
  </div>
</div>

---

<!-- Green段階の詳細対比スライド -->

## Green: 目標達成を実現する段階

<style scoped>
.container {
  display: flex;
  margin-top: 20px;
  height: 400px;
}
.column {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: 0 10px;
}
.title {
  text-align: center;
  font-weight: bold;
  padding: 10px;
  border-radius: 5px;
  margin-bottom: 15px;
}
.muscle-title { background-color: rgba(144, 238, 144, 0.3); }
.dev-title { background-color: rgba(144, 238, 144, 0.3); }
.point {
  background-color: #f8f8f8;
  padding: 10px;
  margin-bottom: 10px;
  border-left: 4px solid #32cd32;
  flex: 1;
}
.center-text {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  font-size: 24px;
  margin: 0 10px;
}
</style>

<div class="container">
  <div class="column">
    <div class="title dev-title">開発: Green</div>
    <div class="point">テストが通過する最小限のコードを実装</div>
    <div class="point">コードの効率性より「動作すること」を優先</div>
    <div class="point">設計の美しさよりも機能の実現に集中</div>
    <div class="point">テストが通ることで機能の実現を確認</div>
  </div>
  
  <div class="center-text">
    <div>同じ</div>
    <div>アプローチ</div>
    <div>↔</div>
  </div>
  
  <div class="column">
    <div class="title muscle-title">筋トレ: Green</div>
    <div class="point">目標達成のための最短経路を計画</div>
    <div class="point">漸進的負荷（週2.5kg増）と必要栄養を確保</div>
    <div class="point">フォームの完璧さより「重量が上がること」を優先</div>
    <div class="point">目標達成のマイルストーンを定めて進捗を確認</div>
  </div>
</div>

---

<!-- Refactor段階の詳細対比スライド -->

## Refactor: 質を高める段階

<style scoped>
.container {
  display: flex;
  margin-top: 20px;
  height: 400px;
}
.column {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: 0 10px;
}
.title {
  text-align: center;
  font-weight: bold;
  padding: 10px;
  border-radius: 5px;
  margin-bottom: 15px;
}
.muscle-title { background-color: rgba(135, 206, 235, 0.3); }
.dev-title { background-color: rgba(135, 206, 235, 0.3); }
.point {
  background-color: #f8f8f8;
  padding: 10px;
  margin-bottom: 10px;
  border-left: 4px solid #4682b4;
  flex: 1;
}
.center-text {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  font-size: 24px;
  margin: 0 10px;
}
</style>

<div class="container">
  <div class="column">
    <div class="title dev-title">開発: Refactor</div>
    <div class="point">テスト成功を維持しながらコードを改善</div>
    <div class="point">重複除去、モジュール化でコードを整理</div>
    <div class="point">パフォーマンス向上、メモリ効率化</div>
    <div class="point">将来の拡張性を考慮した設計に洗練</div>
  </div>
  
  <div class="center-text">
    <div>同じ</div>
    <div>改善思考</div>
    <div>↔</div>
  </div>
  
  <div class="column">
    <div class="title muscle-title">筋トレ: Refactor</div>
    <div class="point">目標達成能力を維持しながら質を向上</div>
    <div class="point">フォーム改善で関節への負担を軽減</div>
    <div class="point">補助種目で弱点部位を強化して均整を整える</div>
    <div class="point">回復の質を高め、長期的な発達を促進</div>
  </div>
</div>

---

## 問題: プロテイン摂取不足の解消

<style scoped>
.container {
  display: flex;
  justify-content: space-between;
  margin-top: 15px;
}
.column {
  width: 48%;
}
.title {
  font-weight: bold;
  margin-bottom: 5px;
  color: #0070c0;
  text-align: center;
}
pre {
  font-size: 13px;
  margin: 0;
  line-height: 1.15;
}
.question {
  font-size: 19px;
  text-align: center;
  font-weight: bold;
  margin-top: 10px;
  color: #d73a49;
}
</style>

<div class="question">
  問題: トレーニング時に必要なタンパク質摂取を忘れてしまう！<br>
  どうすれば確実に摂取できるようにできるでしょうか？
</div>

<div class="container">
  <div class="column">
    <div class="title">プロテイン不足を検出するテスト</div>

```go
// ユーザー情報
type User struct {
  Weight  float64 // 体重(kg)
  MaxKg   float64 // ベンチプレス重量(kg)
  Protein float64 // タンパク質摂取量(g)
}

// Red: 栄養不足を検出するテスト
func TestProtein(t *testing.T) {
  u := User{
    Weight: 80.0,  // 体重80kg
    MaxKg: 70.0,   // 現在の最大重量
    Protein: 0.0,  // 開始時は未摂取
  }

  // 必要量: 体重×2g = 160g
  need := u.Weight * 2.0

  // トレーニング実行 (※60%の確率で摂取を忘れる)
  goal := 100.0
  result := Training(u, goal)

  // トレーニング後のプロテイン摂取量をチェック
  if u.Protein < need {
    t.Errorf("プロテイン摂取が不足: %.1fg必要なのに%.1fg",
           need, u.Protein)
  }
}
```

  </div>

  <div class="column">
    <div class="title">修正前: プロテイン摂取を忘れる問題</div>

```go
// トレーニング関数
func Training(u User, goal float64) float64 {
  // たまにプロテイン摂取を忘れることがある
  // プロテインを忘れる確率
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  randomValue := r.Intn(100)

  // 60%の確率でプロテイン摂取を忘れる
  if randomValue < 60 {
    log.Printf("プロテイン: %.1fg", u.Weight*2.0)
    u.Protein = u.Weight * 2.0 // プロテイン摂取
  }

  // トレーニング実行
  gap := goal - u.MaxKg
  weeks := int(math.Ceil(gap / 2.5))
  for i := 0; i < weeks; i++ {
    u.MaxKg += 2.5
  }
  return u.MaxKg

  // プロテイン摂取を60%の確率で忘れる問題あり
}
```

  </div>
</div>

---

## 解決策: defer を使ったプロテイン摂取の保証

<style scoped>
.container {
  display: flex;
  justify-content: center;
  margin-top: 15px;
}
.column {
  width: 90%;
}
.title {
  font-weight: bold;
  margin-bottom: 5px;
  color: #0070c0;
  text-align: center;
}
pre {
  font-size: 13px;
  margin: 0;
  line-height: 1.15;
}
.solution {
  font-size: 19px;
  text-align: center;
  font-weight: bold;
  margin-top: 10px;
  color: #32cd32;
}
</style>

<div class="container">
  <div class="column">
    <div class="solution">
  解決: defer で処理の最後に必ず実行されることで<br>
  プロテイン摂取の仕組みを導入し、忘れを防止！<br>
</div>

```go
// トレーニング関数 (Refactored)
func Training(u User, goal float64) float64 {
  // 関数終了時に自動実行される処理を defer で登録
  //　deferにすることで全ての筋トレが終了したゴールデンタイムに摂取することが可能
  taken := false
  defer func() {
    if !taken {
      // 未摂取なら強制摂取
      log.Printf("プロテイン: %.1fg", u.Weight * 2.0)
      u.Protein = u.Weight * 2.0 // TakeProtein簡略化
    }
  }()

  // 事前チェック
  req := u.Weight * 2.0
  if u.Protein >= req {
    taken = true
  }

  // トレーニング実行
  gap := goal - u.MaxKg
  weeks := int(math.Ceil(gap / 2.5))
  // 週ごとに最大重量を増加
  for i := 0; i < weeks; i++ {
    u.MaxKg += 2.5
  }
  return u.MaxKg
}
```

  </div>
</div>

---

<!-- TDDの応用ポイントスライド -->

## TDD の応用ポイント: 共通する概念

<style scoped>
table {
  width: 100%;
  margin-top: 30px;
  border-collapse: collapse;
}
th {
  background-color: #0070c0;
  color: white;
  padding: 10px;
}
td {
  padding: 10px;
  border: 1px solid #ddd;
}
tr:nth-child(even) {
  background-color: #f2f2f2;
}
</style>

| 筋トレでの適用                                                   | システム開発での適用                                                           |
| ---------------------------------------------------------------- | ------------------------------------------------------------------------------ |
| 小さなサイクルで繰り返す<br>（週単位での目標設定と評価）         | 小さなイテレーションで繰り返す<br>（スプリントごとの目標設定とレビュー）       |
| 測定可能な目標設定<br>（重量、回数、フォームの質など）           | 測定可能な目標設定<br>（パフォーマンス指標、カバレッジ率、品質メトリクスなど） |
| 結果に基づく調整<br>（進捗が見られない場合はプログラムを見直す） | 結果に基づく調整<br>（バグや性能問題が見られる場合は設計を見直す）             |
| 継続的な改善<br>（常により良いフォームと効率を追求する）         | 継続的な改善<br>（常により良いコード品質と効率を追求する）                     |

---

<!-- まとめスライド -->

## まとめ: システム開発と筋肉開発のサイクル

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 30px;
}
.cycle {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 30px;
}
.box {
  width: 180px;
  height: 100px;
  margin: 0 10px;
  padding: 15px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border-radius: 10px;
  font-weight: bold;
  text-align: center;
  font-size: 20px;
}
.arrow {
  font-size: 30px;
  margin: 0 5px;
}
.red { background-color: rgba(255, 99, 71, 0.7); }
.green { background-color: rgba(144, 238, 144, 0.7); }
.blue { background-color: rgba(135, 206, 235, 0.7); }
.label {
  font-weight: bold;
  font-size: 24px;
  margin-bottom: 10px;
}
p {
  text-align: center;
  font-size: 24px;
  margin-top: 20px;
}
</style>

<div class="container">
  <div class="label">システム開発と筋肉開発: 同じサイクル、同じ思考法</div>
  <div class="cycle">
    <div class="box red">目標設定<br>失敗の確認</div>
    <div class="arrow">→</div>
    <div class="box green">最小限の実装<br>目標の達成</div>
    <div class="arrow">→</div>
    <div class="box blue">質の向上<br>継続的改善</div>
    <div class="arrow">→</div>
    <div class="box red">次の目標<br>設定</div>
  </div>

  <p>明確な目標と測定基準を持ち、段階的に進歩することで、<br>どちらの分野でも効率的かつ安全（最小限の怪我で）に目標達成を図ることができます！</p>
</div>

---

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 400px;
}
h1 {
  font-size: 40px;
  color: #0070c0;
  margin-bottom: 30px;
  text-align: center;
}
.message {
  font-size: 32px;
  font-weight: bold;
  margin-bottom: 20px;
  text-align: center;
  color: #333;
}
.subtitle {
  font-size: 28px;
  font-style: italic;
  color: #ff6347;
  text-align: center;
  margin-bottom: 40px;
}
.emoji {
  font-size: 50px;
  margin: 10px 0;
}
.muscle {
  color: #32cd32;
  font-weight: bold;
}
.system {
  color: #0070c0;
  font-weight: bold;
}
</style>

<div class="container">
  <div class="message">TDDで<span class="muscle">筋肉</span>も<span class="system">システム</span>も鍛えよう！</div>
  <div class="subtitle">筋肉開発 ≒ システム開発</div>
  <h1>ありがとうございました</h1>
</div>
