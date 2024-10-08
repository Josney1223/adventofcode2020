package dayeight

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func InterpretCommandsOne(path string) (int, error) {
	commands, err := OpenCommandsFile(path)
	if err != nil {
		return -1, err
	}

	commandExecutedList := []bool{}
	commandsNew := [][]string{}

	for _, i := range commands {
		splitString := strings.Split(i, " ")
		commandsNew = append(commandsNew, splitString)

		commandExecutedList = append(commandExecutedList, false)
	}
	return ExecuteCommands(commandsNew, commandExecutedList), nil
}

func ExecuteCommands(commands [][]string, commandExecutedList []bool) int {
	acc := 0
	cursor := 0

	for !commandExecutedList[cursor] {
		command := commands[cursor]
		commandExecutedList[cursor] = true

		switch command[0] {
		case "nop":
			cursor++
		case "jmp":
			jmpRunes := []rune(command[1])
			operator := jmpRunes[0]

			num, err := strconv.Atoi(string(jmpRunes[1:]))
			if err != nil {
				return -1
			}

			if operator == []rune("-")[0] {
				cursor -= num
			} else {
				cursor += num
			}

		case "acc":
			num, err := strconv.Atoi(command[1])
			if err != nil {
				return -1
			}
			acc += num
			cursor++
		}

	}
	return acc
}

func OpenCommandsFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arraySlices := []string{}
	for scanner.Scan() {
		newText := scanner.Text()
		arraySlices = append(arraySlices, newText)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return arraySlices, nil
}
