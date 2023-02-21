package record

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Record represents the highest score achieved in the game.
type Record struct {
	Score int
}

// Save saves the current score as the highest record.
func (r *Record) Save(score int) error {
	if score > r.Score {
		r.Score = score

		file, err := os.Create("record.txt")
		if err != nil {
			return err
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		_, err = writer.WriteString(strconv.Itoa(r.Score))
		if err != nil {
			return err
		}

		err = writer.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}

// Load loads the highest record from the file.
func (r *Record) Load() error {
	file, err := os.Open("record.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	text, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	score, err := strconv.Atoi(text[:len(text)-1])
	if err != nil {
		return err
	}

	r.Score = score
	return nil
}

// Display displays the current highest record on the screen.
func (r *Record) Display() {
	fmt.Println("Highest Record:", r.Score)
}
