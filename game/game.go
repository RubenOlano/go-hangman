package game

import (
	"bufio"
	"fmt"
	"hangman/board"
	"log"
	"math/rand"
	"os"
	"time"
	"unicode"
)

type Game struct {
	word        string
	guesses     uint
	underscores []rune
	won         bool
	lose        bool
	prevGuess   []rune
}

func readFile() []string {
	file, err := os.Open("web2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) > 3 {
			lines = append(lines, scanner.Text())
		}
	}
	return lines
}

func initGame() Game {
	rand.Seed(time.Now().UnixNano())
	words := readFile()
	word := words[rand.Intn(len(words))]
	return Game{word: word, guesses: 0, underscores: nil, won: false, lose: false, prevGuess: nil}
}

var game = initGame()

func setUnderscores() {
	word := game.word
	underscores := make([]rune, 0)
	for range word {
		underscores = append(underscores, '_')
	}
	game.underscores = underscores
}

func fetchGuess() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the character you would like to guess: ")
	guess, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You guessed %c\n", guess)
	validateGuess(guess)
}

func validateGuess(guess rune) {
	correct := false
	underScores := make([]rune, 0)
	for i, letter := range game.word {
		if (unicode.ToUpper(guess)) == unicode.ToUpper(letter) {
			underScores = append(underScores, letter)
			correct = true
		} else if game.underscores[i] != '_' {
			underScores = append(underScores, game.underscores[i])
		} else {
			underScores = append(underScores, '_')
		}
	}
	game.underscores = underScores
	handleGuess(correct, guess)
}

func handleGuess(correct bool, guess rune) {
	if !correct {
		game.prevGuess = append(game.prevGuess, guess)
		game.guesses++
	}
	if string(game.underscores) == game.word {
		game.won = true
		fmt.Print("\033[H\033[2J")
		fmt.Printf("You Win! The word was %s\n", game.word)
	}
	if game.guesses > 5 {
		game.lose = true
		fmt.Printf("Too many incorrect guesses. You lose! The word was %s\n", game.word)
	}
}

func Run() {
	setUnderscores()
	for !game.won && !game.lose {
		fmt.Print("\033[H\033[2J")
		board.DisplayBoard(game.guesses)
		fmt.Printf("Word: %s\n", string(game.underscores))
		fmt.Printf("Guesses: %s\n", string(game.prevGuess))
		fetchGuess()
	}
	board.DisplayBoard(game.guesses)
}
