package board

import "fmt"

var board = [8]string{
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
