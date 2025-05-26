package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// RSES‑J (Self‑Esteem)
var rsesQuestions = []string{
	"1. 私は、自分自身にだいたい満足している。",
	"2. 時々、自分はまったくダメだと思うことがある。", // 逆転
	"3. 私には、けっこう長所があると感じている。",
	"4. 私は、他の大半の人と同じくらいに物事がこなせる。",
	"5. 私には誇れるものが大してないと感じる。",      // 逆転
	"6. 時々、自分は役に立たないと強く感じることがある。", // 逆転
	"7. 自分は少なくとも他の人と同じくらい価値のある人間だと感じる。",
	"8. 自分のことをもう少し尊敬できたらいいと思う。", // 逆転
	"9. よく、私は落ちこぼれだと思ってしまう。",    // 逆転
	"10. 私は、自分のことを前向きに考えている。",
}
var rsesReverse = map[int]bool{2: true, 5: true, 6: true, 8: true, 9: true}

// GSES‑J (Self‑Efficacy)
var gsesQuestions = []string{
	"1. 何か仕事をする時は、自信を持ってやるほうである。",
	"2. 過去に犯した失敗や嫌な経験を思い出して、暗い気持ちになることがよくある。", // 逆転
	"3. 友人より優れた能力がある。",
	"4. 仕事を終えた後、失敗したと感じることのほうが多い。", // 逆転
	"5. 人と比べて心配性なほうである。",           // 逆転
	"6. 何かを決める時、迷わずに決定するほうである。",
	"7. 何かをする時、うまくゆかないのではないかと不安になることが多い。", // 逆転
	"8. 引っ込み思案なほうだと思う。",                   // 逆転
	"9. 人より記憶力が良いほうである。",
	"10. 結果の見通しがつかない仕事でも、積極的に取り組んでゆくほうだと思う。",
	"11. どうやったら良いか決心がつかずに仕事に取りかかれないことがよくある。", // 逆転
	"12. 友人よりも特に優れた知識を持っている分野がある。",
	"13. どんなことでも積極的にこなすほうである。",
	"14. 小さな失敗でも人よりずっと気にするほうである。", // 逆転
	"15. 積極的に活動するのは苦手なほうである。",     // 逆転
	"16. 世の中に貢献できる力があると思う。",
}
var gsesReverse = map[int]bool{2: true, 4: true, 5: true, 7: true, 8: true, 11: true, 14: true, 15: true}

// main
func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Println("自己肯定感・自己効力感 自動採点ツール")

	// RSES‑J -------------------------------------------------------
	fmt.Println("■ Rosenberg Self‑Esteem Scale (RSES‑J) 10項目")
	rsesScore := 0
	for i, q := range rsesQuestions {
		ans := promptLikert(in, q)
		if rsesReverse[i+1] {
			ans = 5 - ans // 1↔4, 2↔3
		}
		rsesScore += ans
	}
	fmt.Printf("\n▶ RSES‑J 合計点: %d / 40\n", rsesScore)
	interpretRSES(rsesScore)

	// GSES‑J -------------------------------------------------------
	fmt.Println("\n■ General Self‑Efficacy Scale (GSES‑J) 16項目 (はい=1 / いいえ=0)")
	gsesScore := 0
	for i, q := range gsesQuestions {
		ans := promptYesNo(in, q)
		if gsesReverse[i+1] {
			ans = 1 - ans // reverse scoring
		}
		gsesScore += ans
	}
	fmt.Printf("\n▶ GSES‑J 合計点: %d / 16\n", gsesScore)
	interpretGSES(gsesScore)

	fmt.Println("\n―― 完了 ――")
}

// ヘルパ
func promptLikert(in *bufio.Reader, question string) int {
	for {
		fmt.Printf("%s\n1: 強くそう思う / 2: そう思う / 3: そう思わない / 4: 強くそう思わない → ", question)
		line, _ := in.ReadString('\n')
		line = strings.TrimSpace(line)
		v, err := strconv.Atoi(line)
		if err == nil && v >= 1 && v <= 4 {
			return v
		}
		fmt.Println("  ※ 1〜4 の数字で入力してください。")
	}
}

func promptYesNo(in *bufio.Reader, question string) int {
	for {
		fmt.Printf("%s (はい/いいえ) → ", question)
		line, _ := in.ReadString('\n')
		line = strings.TrimSpace(strings.ToLower(line))
		switch line {
		case "はい", "y", "yes", "1":
			return 1
		case "いいえ", "n", "no", "0":
			return 0
		default:
			fmt.Println("  ※ 'はい' または 'いいえ' で入力してください。")
		}
	}
}

func interpretRSES(score int) {
	switch {
	case score <= 20:
		fmt.Println("  → 自己肯定感はやや低めと考えられます。")
	case score <= 29:
		fmt.Println("  → 標準的な範囲です。")
	default:
		fmt.Println("  → 高めの自己肯定感を示しています。")
	}
}

func interpretGSES(score int) {
	switch {
	case score <= 8:
		fmt.Println("  → 自己効力感はやや低めと考えられます。")
	case score <= 11:
		fmt.Println("  → 標準的な範囲です。")
	default:
		fmt.Println("  → 高めの自己効力感を示しています。")
	}
}
