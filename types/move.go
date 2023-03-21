package types

type Move struct {
	// Number is the move number, like 1 for first, 2 for second etc.
	Number int

	// White is the move white made
	White string

	// Black is the move black made
	Black string

	// Result is the result of the game, should only be
	// present in the last move
	Result string
}
