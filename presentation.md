---
marp: true
theme: default
paginate: true
backgroundColor: #fff
header: "Go言語で学ぶテスト駆動筋肉開発"
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

# Go 言語で学ぶテスト駆動筋肉開発

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

筋トレのプラクティスをエンジニアに親しみやすく説明するために、TDD（テスト駆動開発）の概念を活用します。

筋トレにおいて自分の体でテストを行いながら重量を増加させていく過程を、システム開発の TDD プロセスに例えて解説します。

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

1. エンジニアに筋トレの基本的なプラクティスを紹介する
2. TDD という親しみやすい概念を通じて筋トレの段階的な進め方を説明する
3. 技術者の思考法を活かした筋トレアプローチを提案する

---

<!-- TDDの流れ概要スライド -->

## 筋肉 TDD の流れ

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
  <div class="arrow">⇅</div>
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
</div>

---

<!-- Red段階の詳細対比スライド -->

## Red: 失敗を明確にする段階

<style scoped>
.container {
  display: flex;
  margin-top: 20px;
  height: 380px;
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
  padding: 8px;
  border-radius: 5px;
  margin-bottom: 10px;
}
.muscle-title { background-color: rgba(255, 99, 71, 0.3); }
.dev-title { background-color: rgba(255, 99, 71, 0.3); }
.point {
  background-color: #f8f8f8;
  padding: 8px;
  margin-bottom: 8px;
  border-left: 4px solid #ff6347;
  flex: 1;
}
.center-text {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  font-size: 22px;
  margin: 0 8px;
}
</style>

<div class="container">
  <div class="column">
    <div class="title muscle-title">筋トレ: Red</div>
    <div class="point">ベンチプレスの目標重量を100kgに設定</div>
    <div class="point">現在の最大重量70kgで試すと失敗する</div>
    <div class="point">フルレンジの動き、正しいフォーム、安全性の基準を設定</div>
    <div class="point">目標に向けて必要な改善点を特定</div>
  </div>
  
  <div class="center-text">
    <div>同じ</div>
    <div>概念</div>
    <div>↔</div>
  </div>
  
  <div class="column">
    <div class="title dev-title">開発: Red</div>
    <div class="point">システムの要求仕様を定義（例: 1000リクエスト/秒）</div>
    <div class="point">期待される動作を検証するテストコードを作成</div>
    <div class="point">実装前のテストを実行すると「失敗」する</div>
    <div class="point">解決すべき技術的課題を明確化</div>
  </div>
</div>

---

<!-- Green段階の詳細対比スライド -->

## Green: 目標達成を実現する段階

<style scoped>
.container {
  display: flex;
  margin-top: 20px;
  height: 380px;
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
  padding: 8px;
  border-radius: 5px;
  margin-bottom: 10px;
}
.muscle-title { background-color: rgba(144, 238, 144, 0.3); }
.dev-title { background-color: rgba(144, 238, 144, 0.3); }
.point {
  background-color: #f8f8f8;
  padding: 8px;
  margin-bottom: 8px;
  border-left: 4px solid #32cd32;
  flex: 1;
}
.center-text {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  font-size: 22px;
  margin: 0 8px;
}
</style>

<div class="container">
  <div class="column">
    <div class="title muscle-title">筋トレ: Green</div>
    <div class="point">効果的なトレーニングプランを作成（週2.5kg増）</div>
    <div class="point">必要な栄養摂取（タンパク質等）を計算</div>
    <div class="point">最小限のアプローチで目標重量100kgの達成を目指す</div>
    <div class="point">「目標を達成できる」という成功体験を得る</div>
  </div>
  
  <div class="center-text">
    <div>同じ</div>
    <div>アプローチ</div>
    <div>↔</div>
  </div>
  
  <div class="column">
    <div class="title dev-title">開発: Green</div>
    <div class="point">テスト通過に必要な最小限のコードを実装</div>
    <div class="point">必要なリソース（メモリ、CPU）を割り当て</div>
    <div class="point">複雑な実装は後回しにして、まず機能させる</div>
    <div class="point">「動くコード」ができるという達成感を得る</div>
  </div>
</div>

---

<!-- Refactor段階の詳細対比スライド -->

## Refactor: 質を高める段階

<style scoped>
.container {
  display: flex;
  margin-top: 20px;
  height: 380px;
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
  padding: 8px;
  border-radius: 5px;
  margin-bottom: 10px;
}
.muscle-title { background-color: rgba(135, 206, 235, 0.3); }
.dev-title { background-color: rgba(135, 206, 235, 0.3); }
.point {
  background-color: #f8f8f8;
  padding: 8px;
  margin-bottom: 8px;
  border-left: 4px solid #4682b4;
  flex: 1;
}
.center-text {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  font-size: 22px;
  margin: 0 8px;
}
</style>

<div class="container">
  <div class="column">
    <div class="title muscle-title">筋トレ: Refactor</div>
    <div class="point">フォームの最適化で力の伝達効率を向上</div>
    <div class="point">補助種目を追加して弱点となる筋肉を強化</div>
    <div class="point">回復戦略を改善（休息、栄養タイミング等）</div>
    <div class="point">次の目標設定（105kgや同重量でのレップ数増加）</div>
  </div>
  
  <div class="center-text">
    <div>同じ</div>
    <div>改善思考</div>
    <div>↔</div>
  </div>
  
  <div class="column">
    <div class="title dev-title">開発: Refactor</div>
    <div class="point">アルゴリズムの最適化でパフォーマンス向上</div>
    <div class="point">コードの整理・モジュール化で保守性向上</div>
    <div class="point">リソース使用効率の改善（キャッシュ戦略等）</div>
    <div class="point">新機能追加やさらなる性能向上への準備</div>
  </div>
</div>

---

<!-- 実際のコード例と筋トレプランの対比 -->

## 具体例: コードと筋トレプランの対比

<style scoped>
.container {
  display: flex;
  margin-top: 20px;
}
.column {
  flex: 1;
  padding: 0 10px;
  font-size: 18px;
}
.title {
  text-align: center;
  font-weight: bold;
  padding: 5px;
  border-radius: 5px;
  margin-bottom: 10px;
}
.muscle-title { background-color: rgba(144, 238, 144, 0.3); }
.dev-title { background-color: rgba(135, 206, 250, 0.3); }
</style>

<div class="container">
  <div class="column">
    <div class="title muscle-title">筋トレプラン</div>
    
```
【Red】目標設定
・ベンチプレス: 100kg
・現状: 70kg → 失敗
・必要増加量: 30kg

【Green】トレーニング計画
・週 2.5kg ずつ増加
・期間: 12 週間
・プロテイン: 体重 ×2g/日
・週 3 回トレーニング

【Refactor】効率化
・補助種目: ダンベルプレス追加
・フォーム改善: 肩甲骨の固定強化
・栄養: プロテイン摂取タイミング最適化

````
  </div>

  <div class="column">
    <div class="title dev-title">Go言語コード</div>

```go
// Red: 失敗するテスト
func TestBenchPress(t *testing.T) {
  want := 100.0
  maxWeight := 70.0
  bodyWeight := 80.0

  // Green: 実装
  trained := BenchPress(maxWeight, "100kg")

  if trained < want {
    t.Errorf("want %.1fkg, got %.1fkg",
             want, trained)
  }

  // Refactor: 品質向上 - プロテイン摂取を強制
  protein := RequireProteinIntake(bodyWeight)
  if protein < bodyWeight*2.0 {
    t.Errorf("必要なプロテイン摂取量: %.1fg",
             bodyWeight*2.0)
  }
}
```

  </div>
</div>

---

<!-- Go言語実装のスライド -->

## Go 言語で実装した筋肉 TDD のコード構造

<style scoped>
.container {
  display: flex;
  justify-content: space-between;
  margin-top: 30px;
}
.column {
  width: 48%;
}
.title {
  font-weight: bold;
  margin-bottom: 10px;
  color: #0070c0;
}
pre {
  font-size: 16px;
  margin: 0;
}
</style>

<div class="container">
  <div class="column">
    <div class="title">トレーニングプラン構造体</div>

```go
type TrainingPlan struct {
  WeeklyIncrease   float64 // 週あたりの重量増加(kg)
  ProteinIntake    float64 // 体重当たりのタンパク質(g)
  RecoveryDays     int     // 回復のための休息日数
  TrainingDuration int     // トレーニング期間（週）
}

func DefaultTrainingPlan() TrainingPlan {
  return TrainingPlan{
    WeeklyIncrease:   2.5,
    ProteinIntake:    2.0, // 体重1kgあたり2g必須
    RecoveryDays:     2,
    TrainingDuration: 12,
  }
}
```
  </div>
  <div class="column">
    <div class="title">リファクタリング後の改善実装</div>

```go
// プロテイン摂取を強制するリファクタリング
func RequireProteinIntake(bodyWeight float64) float64 {
  return bodyWeight * DefaultTrainingPlan().ProteinIntake
}

// BenchPressWithUserInfo改善版
func BenchPressRefactored(info UserInfo) (float64, error) {
  proteinTaken := false

  // deferを使って関数終了時にプロテイン摂取を強制
  defer func() {
    if !proteinTaken {
      log.Printf("プロテイン摂取を強制実行")
      TakeProtein(info.BodyWeight)
    }
  }()

  // プロテイン摂取チェック
  if info.DailyProteinIntake >= RequireProteinIntake(info.BodyWeight) {
    proteinTaken = true
  }

  // トレーニング実行と結果返却
  return calculateTargetWeight(info), nil
}

// 簡略化したプロテイン摂取関数
func TakeProtein(bodyWeight float64) {
  log.Printf("%.1fgのプロテイン摂取完了",
             bodyWeight * DefaultTrainingPlan().ProteinIntake)
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

<!-- 筋肉TDDのメリットスライド -->

## 筋肉 TDD のメリット

<style scoped>
ul {
  list-style-type: none;
  padding: 0;
  margin-top: 40px;
}
li {
  padding: 15px;
  margin-bottom: 10px;
  background-color: #f8f8f8;
  border-left: 5px solid #0070c0;
  font-size: 24px;
}
</style>

- 🎯 **明確な目標と測定基準**を持ち、進捗を客観的に評価できる
- 🔄 **小さな繰り返し**で安全かつ効率的に筋力向上を図れる
- 📈 **段階的な進歩**により挫折リスクを減らし、継続的な取り組みを促進
- 🛠️ **特定の問題点**に対して効果的な対処法を見つけやすくなる
- 🧠 **技術者の思考法**を活かした論理的なトレーニング設計が可能

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
  <div class="label">筋肉開発とシステム開発: 同じサイクル、同じ思考法</div>
  <div class="cycle">
    <div class="box red">目標設定<br>失敗の確認</div>
    <div class="arrow">→</div>
    <div class="box green">最小限の実装<br>目標の達成</div>
    <div class="arrow">→</div>
    <div class="box blue">質の向上<br>継続的改善</div>
    <div class="arrow">→</div>
    <div class="box red">次の目標<br>設定</div>
  </div>

  <p>明確な目標と測定基準を持ち、段階的に進歩することで、<br>どちらの分野でも効率的かつ安全に目標達成を図ることができます！</p>
</div>

---

# ありがとうございました
```
