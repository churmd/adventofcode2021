package day2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	FORWARD Direction = "forward"
	DOWN Direction = "down"
	UP Direction = "up"
)

type Direction = string
type Command struct {
	Dir Direction
	Amount int
}

func (com *Command) UnmarshalText(text []byte) error {
	switch {
	case strings.HasPrefix(string(text), FORWARD): com.Dir = FORWARD
	case strings.HasPrefix(string(text), DOWN): com.Dir = DOWN
	case strings.HasPrefix(string(text), UP): com.Dir = UP
	default: return errors.New("command does not contain a valid direction")
	}

	num := strings.Replace(string(text), com.Dir + " ", "", 1)
	amount, err := strconv.Atoi(num)
	if err != nil {
		return err
	}
	com.Amount = amount
	
	return nil
}

type Position struct {
	Horizontal int
	Depth int
}

func (c Position) String() string {
	return fmt.Sprintf("(Hor: %d, Depth: %d)", c.Horizontal, c.Depth)
}

func (c Position) Follow(com Command) Position {
	switch com.Dir {
	case FORWARD: return c.Add(Position{Horizontal: com.Amount})
	case DOWN: return c.Add(Position{Depth: com.Amount})
	case UP: return c.Add(Position{Depth: -com.Amount})
	}

	panic("followed a command with an unknown direction: " + com.Dir)
}

func (c Position) Add(c2 Position) Position {
	return Position{Horizontal: c.Horizontal + c2.Horizontal, Depth: c.Depth + c2.Depth}
}

func Solution1() {
	commands := getCommands()

	pos := Position{Horizontal:0, Depth:0}
	for _, command := range commands {
		pos = pos.Follow(command)
	}

	fmt.Println("What do you get if you multiply your final horizontal position by your final depth?")
	fmt.Printf("Final position is: %s  answer is: %d\n", pos, pos.Horizontal * pos.Depth)
}

type Position2 struct {
	Horizontal int
	Depth int
	Aim int
}

func (p Position2) String() string {
	return fmt.Sprintf("(Hor: %d, Depth: %d, Aim:%d)", p.Horizontal, p.Depth, p.Aim)
}

func (p *Position2) Follow(com Command) {
	switch com.Dir {
	case FORWARD: {
		p.Horizontal +=com.Amount
		p.Depth += p.Aim * com.Amount
	}
	case DOWN: p.Aim += com.Amount
	case UP: p.Aim -= com.Amount
	default: panic("position 2 followed a command with an unknown direction: " + com.Dir)
	}
}


func Solution2() {
	commands := getCommands()

	pos := Position2{}
	for _, command := range commands {
		pos.Follow(command)
	}

	fmt.Println("What do you get if you multiply your final horizontal position by your final depth?")
	fmt.Printf("Final position is: %s  answer is: %d\n", pos, pos.Horizontal * pos.Depth)
}

func getCommands() []Command {
	lines := strings.Split(input, "\n")

	commands := make([]Command, len(lines))
	for i, line := range lines {
		var com Command
		err := com.UnmarshalText([]byte(line))
		if err != nil {
			panic(fmt.Errorf("could not unmarshal command %s : %w", line, err))
		}
		commands[i] = com
	}

	return commands
}
