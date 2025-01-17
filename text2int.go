package text2int

import "github.com/r3db34n1an/text2int/pkg/text2int"

func Convert(text string) (int, error) {
	return new(text2int.Text2Int).Convert(text)
}
