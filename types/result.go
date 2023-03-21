package types

type Result uint8

const (
	ResultDraw     Result = iota
	ResultWinWhite Result = iota
	ResultWinBlack Result = iota
	ResultUnknown  Result = iota
)

func ParseResult(result string) Result {
	switch result {
	case "1/2-1/2":
		return ResultDraw
	case "1-0":
		return ResultWinWhite
	case "0-1":
		return ResultWinBlack
	}

	return ResultUnknown
}

func (r Result) String() string {
	switch r {
	case ResultDraw:
		return "draw"
	case ResultWinWhite:
		return "white win"
	case ResultWinBlack:
		return "black win"
	case ResultUnknown:
		return "unknown"
	}

	return "N/A"
}
