package board

import "fmt"

var board = [7]string{
	`
    +---+
        |
        |
        |
       ===
        `,
	`
    +---+
    o   |
        |
        |
       ===
       `,
	`
    +---+
    o   |
    |   |
        |
       ===
        `,
	`
    +---+
    o   |
   /|   |
        |
       ===
        `,
	`
    +---+
    o   |
   /|\  |
        |
       ===
        `,
	`
    +---+
    o   |
   /|\  |
   /    |
       ===
       `,
	`
    +---+
    o   |
   /|\  |
   / \  |
       ===
    `,
}

func DisplayBoard(pos uint) {
	fmt.Println(board[pos])
}
