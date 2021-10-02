package kana

import "fmt"

var HiraganaTable [][]rune = [][]rune{
	{'あ', 'い', 'う', 'え', 'お'},
	{'か', 'き', 'く', 'け', 'こ'},
	{'が', 'ぎ', 'ぐ', 'げ', 'ご'},
	{'さ', 'し', 'す', 'せ', 'そ'},
	{'ざ', 'じ', 'ず', 'ぜ', 'ぞ'},
	{'た', 'ち', 'つ', 'て', 'と'},
	{'だ', 'ぢ', 'づ', 'で', 'ど'},
	{'な', 'に', 'ぬ', 'ね', 'の'},
	{'は', 'ひ', 'ふ', 'へ', 'ほ'},
	{'ば', 'び', 'ぶ', 'べ', 'ぼ'},
	{'ぱ', 'ぴ', 'ぷ', 'ぺ', 'ぽ'},
	{'ま', 'み', 'む', 'め', 'も'},
	{'や', ' ', 'ゆ', ' ', 'よ'},
	{'ゃ', ' ', 'ゅ', ' ', 'ょ'},
	{'ら', 'り', 'る', 'れ', 'ろ'},
	{'わ', 'ゐ', ' ', 'ゑ', 'を'},
	{'ん', 'っ'},
}

var HiraganaMap map[rune][]int = map[rune][]int{}

func Col(r rune, col int) (rune, error) {
	pos, ok := HiraganaMap[r]
	if !ok {
		return ' ', fmt.Errorf("invalid rune %v, not in HiraganaMap", r)
	}

	if pos[0] < 0 || pos[0] >= len(HiraganaTable) {
		return ' ', fmt.Errorf("internal: invalid index %v for HiraganaTable", pos)
	}

	if pos[1] < 0 || pos[1] >= len(HiraganaTable[0]) {
		return ' ', fmt.Errorf("internal: invalid index %v for HiraganaTable", pos)
	}

	rt := HiraganaTable[pos[0]][pos[1]]
	if rt == ' ' {
		return ' ', fmt.Errorf("invalid rune lookup: %v in HiraganaTable is empty", pos)
	}

	return rt, nil
}
