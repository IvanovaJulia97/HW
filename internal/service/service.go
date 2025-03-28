//в пакете реализуем функцию определения данных

package service

import (
	"strings"
	"unicode"
)

const (
	А  = ".-"
	Б  = "-..."
	В  = ".--"
	Г  = "--."
	Д  = "-.."
	Е  = "."
	Ж  = "...-"
	З  = "--.."
	И  = ".."
	Й  = ".---"
	К  = "-.-"
	Л  = ".-.."
	М  = "--"
	Н  = "-."
	О  = "---"
	П  = ".--."
	Р  = ".-."
	С  = "..."
	Т  = "-"
	У  = "..-"
	Ф  = "..-."
	Х  = "...."
	Ц  = "-.-."
	Ч  = "---."
	Ш  = "----"
	Щ  = "--.-"
	ЪЬ = "-..-"
	Ы  = "-.--"
	Э  = "..-.."
	Ю  = "..--"
	Я  = ".-.-"

	One   = ".----"
	Two   = "..---"
	Three = "...--"
	Four  = "....-"
	Five  = "....."
	Six   = "-...."
	Seven = "--..."
	Eight = "---.."
	Nine  = "----."
	Zero  = "-----"

	Period       = "......" //.
	Comma        = ".-.-.-" //,
	Colon        = "---..." //:
	QuestionMark = "..--.." //?
	Apostrophe   = ".----." //'
	Hyphen       = "-....-" //-
	Division     = "-..-."  ///
	LeftBracket  = "-.--."  //(
	RightBracket = "-.--.-" //)
	IvertedComma = ".-..-." //“ ”
	DoubleHyphen = "-...-"  //=
	Cross        = ".-.-."  //+
	CommercialAt = ".--.-." //@

	Space = " "
)

type EncodingMap map[rune]string

var DefaultMorse = EncodingMap{
	'А': А,
	'Б': Б,
	'В': В,
	'Г': Г,
	'Д': Д,
	'Е': Е,
	'Ж': Ж,
	'З': З,
	'И': И,
	'Й': Й,
	'К': К,
	'Л': Л,
	'М': М,
	'Н': Н,
	'О': О,
	'П': П,
	'Р': Р,
	'С': С,
	'Т': Т,
	'У': У,
	'Ф': Ф,
	'Х': Х,
	'Ц': Ц,
	'Ч': Ч,
	'Ш': Ш,
	'Щ': Щ,
	'Ь': ЪЬ,
	'Ы': Ы,
	'Ъ': ЪЬ,
	'Э': Э,
	'Ю': Ю,
	'Я': Я,

	'1': One,
	'2': Two,
	'3': Three,
	'4': Four,
	'5': Five,
	'6': Six,
	'7': Seven,
	'8': Eight,
	'9': Nine,
	'0': Zero,

	'.':  Period,
	',':  Comma,
	':':  Colon,
	'?':  QuestionMark,
	'\'': Apostrophe,
	'-':  Hyphen,
	'/':  Division,
	'(':  LeftBracket,
	')':  RightBracket,
	'"':  IvertedComma,
}

var reverseMorse = make(map[string]rune)

func enicode() {
	for k, v := range DefaultMorse {
		reverseMorse[v] = k
	}
}

// функция для конвертации в код Морзе
func ToMorse(str string) string {
	var resultMorse []string

	for _, v := range str {

		if v == ' ' {
			resultMorse = append(resultMorse, "/")
		} else if val, exists := DefaultMorse[unicode.ToUpper(v)]; exists {
			resultMorse = append(resultMorse, val)
		}

	}
	return strings.Join(resultMorse, " ")

}

//функция для конвертации в Текст

func ToText(code string) string {
	var resultText []rune

	text := strings.Split(code, "   ")
	for _, w := range text {
		word := strings.Split(w, " ")
		for _, l := range word {
			if val, exists := reverseMorse[l]; exists {
				resultText = append(resultText, val)
			} else {
				resultText = append(resultText, '?')
			}
		}
		resultText = append(resultText, ' ')
	}
	return strings.TrimSpace(string(resultText))

}

func AutoDetectAndConvert(input string) string {

	if strings.Contains(input, ".") || strings.Contains(input, "-") {
		return ToText(input)
	}
	return ToMorse(input)
}
