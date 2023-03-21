package pgon

import (
	"github.com/artheus/pgon/types"
	"os"
)

func Open(filePath string) (game *types.Game, err error) {
	var file *os.File

	if file, err = os.Open(filePath); err != nil {
		return nil, err
	}
	defer file.Close()

	return ParseGame(file)
}
