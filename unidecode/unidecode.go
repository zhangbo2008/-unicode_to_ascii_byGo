// Unidecode implements a unicode transliterator, which
// replaces non-ASCII characters with their ASCII
// counterparts

package unidecode

import (
	"unicode"
)


/*
把unicode里面的编码,用最近的ascii码来替换.

*/
// Given an unicode encoded string, returns
// another string with non-ASCII characters replaced
// with their closest ASCII counterparts.
// e.g. Unicode("áéíóú") => "aeiou" 
func Unidecode(s string) string {
	str := ""
	for _, c := range s {


		//如果本身字符就是ascii,直接放上就行
		if c <= unicode.MaxASCII {
			str += string(c)
			continue
		}



		// 不是unicode码,超出边界了.那么就忽略这些错误.
		if c > unicode.MaxRune {
			/* Ignore reserved chars */
			continue
		}


		//其他的进行map替换即可.
		if d, ok := transliterations[c]; ok {
			str += d
		}









	}
	return str
}
