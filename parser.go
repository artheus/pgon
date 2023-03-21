package pgon

import (
	"github.com/artheus/pgon/types"
	"io"
	"regexp"
	"strings"
)

var metadataPattern = regexp.MustCompile("^\\[([a-zA-Z]+)\\s+\"(.*)\"]$")
var movePattern = regexp.MustCompile("[0-9]+\\. ([a-zA-Z]+[0-9]\\+?|O-O(-O)?|1-0|0-1|1/2-1/2|\\*) ([a-zA-Z]+[0-9]\\+?|O-O(-O)?|1-0|0-1|1/2-1/2|\\*)")

const (
	resultDraw     = "1/2-1/2"
	resultWinWhite = "1-0"
	resultWinBlack = "0-1"
	resultUnknown  = "*"
)

func ParseGame(r io.Reader) (game *types.Game, err error) {
	var pgnFileData []byte

	if pgnFileData, err = io.ReadAll(r); err != nil {
		return nil, err
	}

	var gameData []string
	var metaDataLines []string

	// NOTE: Read the file, line by line, and sort data into separate slices for
	//       parsing of the separate data types.
	for _, line := range strings.Split(string(pgnFileData), "\n") {
		if line == "" {
			continue
		}

		// NOTE: Trim whitespace from sides of line
		line = strings.Trim(line, "\n \t")

		// NOTE: Sort metadata and game data into separate slices
		if metadataPattern.MatchString(line) {
			metaDataLines = append(metaDataLines, line)
		} else {
			gameData = append(gameData, line)
		}
	}

	var metadata = map[string]string{}

	// NOTE: Parse metadata using regexp
	for _, md := range metaDataLines {
		submatches := metadataPattern.FindAllStringSubmatch(md, -1)
		metadata[submatches[0][1]] = submatches[0][2]
	}

	// NOTE: parse game data
	var moves []types.Move
	var moveStr = strings.Join(gameData, " ")

	var rMoves = movePattern.FindAllStringSubmatch(moveStr, -1)
	var result string

	for i, move := range rMoves {
		var whiteMove, blackMove string

		whiteMove = move[1]

		if len(move) > 3 {
			blackMove = move[3]
		}

		if isResult(whiteMove) {
			result = whiteMove
			whiteMove = ""
		}

		if isResult(blackMove) {
			result = blackMove
			blackMove = ""
		}

		moves = append(moves, types.Move{
			Number: i + 1,
			White:  whiteMove,
			Black:  blackMove,
			Result: result,
		})
	}

	if result == "" {
		if r, ok := metadata["Result"]; ok {
			result = r
		}
	}

	game = &types.Game{
		Metadata: metadata,
		Moves:    moves,
		Result:   types.ParseResult(result),
	}

	return game, nil
}

func isResult(s string) bool {
	return s == resultDraw || s == resultWinWhite || s == resultWinBlack || s == resultUnknown
}
