# P(Go)N - Simple PGN (Portable Game Notation) Chess file parser

Just a very simple parser for PGN files.

## Example

Use the `pgon.Open` function to open a game from a file.

```golang
package main

import (
	"fmt"
	"github.com/artheus/pgon"
)

// This will simply read and parse a pgn file, then print the result
// to stdout.
func main() {
	gameData, err := pgon.Open("/path/to/file.pgn")
	if err != nil {
		panic(err)
    }

	fmt.Printf("metadata: %+v\n", gameData.Metadata) // metadata that might be present in the PGN file
	fmt.Printf("moves: %+v\n", gameData.Moves) // pgon.Move instances from the parsed result
	fmt.Printf("winner side: %+v\n", gameData.Result) // black, white or draw
}
```

Otherwise, you can use the `pgon.ParseGame` function, to parse PGN from a provided io.Reader.

```golang
package main

import (
	"bytes"
	"fmt"
	"github.com/artheus/pgon"
)

// This will simply read and parse a pgn string, then print the result
// to stdout.
func main() {
	const pgnData = `[Event "F/S Return Match"] 
[Site "Belgrade, Serbia JUG"] 
[Date "1992.11.04"] 
[Round "29"] 
[White "Fischer, Robert J."]
[Black "Spassky, Boris V."] 
[Result "1/2-1/2"] 

1. e4 e5 2. Nf3 Nc6 3. Bb5 a6 4. Ba4 Nf6 5. O-O Be7 6. Re1 b5 7. Bb3 d6 8. c3
O-O 9. h3 Nb8 10. d4 Nbd7 11. c4 c6 12. cxb5 axb5 13. Nc3 Bb7 14. Bg5 b4 15.
Nb1 h6 16. Bh4 c5 17. dxe5 Nxe4 18. Bxe7 Qxe7 19. exd6 Qf6 20. Nbd2 Nxd6 21.
Nc4 Nxc4 22. Bxc4 Nb6 23. Ne5 Rae8 24. Bxf7+ Rxf7 25. Nxf7 Rxe1+ 26. Qxe1 Kxf7
27. Qe3 Qg5 28. Qxg5 hxg5 29. b3 Ke6 30. a3 Kd6 31. axb4 cxb4 32. Ra5 Nd5 33.
f3 Bc8 34. Kf2 Bf5 35. Ra7 g6 36. Ra6+ Kc5 37. Ke1 Nf4 38. g3 Nxh3 39. Kd2 Kb5
40. Rd6 Kc5 41. Ra6 Nf2 42. g4 Bd3 43. Re6 1/2-1/2`

	gameData, err := pgon.ParseGame(bytes.NewReader([]byte(pgnData)))
	if err != nil {
		panic(err)
	}

	fmt.Printf("metadata: %+v\n", gameData.Metadata)  // metadata that might be present in the PGN file
	fmt.Printf("moves: %+v\n", gameData.Moves)        // pgon.Move instances from the parsed result
	fmt.Printf("winner side: %+v\n", gameData.Result) // black, white or draw
}
```