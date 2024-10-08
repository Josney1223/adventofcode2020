package dayeight

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func InterpretCommandsOne(path string) (int, int, error) {
	commands, err := OpenCommandsFile(path)
	if err != nil {
		return -1, -1, err
	}

	commandExecutedList := []bool{}
	commandsNew := [][]string{}

	for _, i := range commands {
		splitString := strings.Split(i, " ")
		commandsNew = append(commandsNew, splitString)

		commandExecutedList = append(commandExecutedList, false)
	}
	acc, t, _ := ExecuteCommands(commandsNew, commandExecutedList)
	return acc, t, nil
}

func FindReplaceCorruptedCommand(path string) (int, int, error) {
	commands, err := OpenCommandsFile(path)
	if err != nil {
		return -1, -1, err
	}

	commandExecutedListDefault := []bool{}
	commandsNew := [][]string{}

	for _, i := range commands {
		splitString := strings.Split(i, " ")
		commandsNew = append(commandsNew, splitString)

		commandExecutedListDefault = append(commandExecutedListDefault, false)
	}

	commandExecutedList := []bool{}
	commandExecutedList = append(commandExecutedList, commandExecutedListDefault...)

	_, _, firstRun := ExecuteCommands(commandsNew, commandExecutedList)

	for _, cmd := range firstRun {
		commandExecutedList = []bool{}
		commandExecutedList = append(commandExecutedList, commandExecutedListDefault...)
		end := -1
		acc := 0

		if commandsNew[cmd][0] == "jmp" {
			commandsNew[cmd][0] = "nop"
			acc, end, _ = ExecuteCommands(commandsNew, commandExecutedList)
			commandsNew[cmd][0] = "jmp"
		} else if commandsNew[cmd][0] == "nop" {
			commandsNew[cmd][0] = "jmp"
			acc, end, _ = ExecuteCommands(commandsNew, commandExecutedList)
			commandsNew[cmd][0] = "nop"
		}

		if end == 1 {
			return acc, 1, nil
		}
	}

	return 0, -1, nil
}

// Notação BigO: Performance (n), Memoria (n)
//
// Execute uma lista de commandos, pare quando chegar no primeiro espaço
// vazio depois do fim da lista ou caso repita algum comando
func ExecuteCommands(commands [][]string, commandExecutedList []bool) (int, int, []int) {
	acc := 0
	cursor := 0
	steps := []int{}

	for !commandExecutedList[cursor] && cursor != len(commands) {

		command := commands[cursor]
		commandExecutedList[cursor] = true
		steps = append(steps, cursor)

		switch command[0] {
		case "nop":
			cursor++

		case "jmp":
			jmpRunes := []rune(command[1])
			operator := jmpRunes[0]

			num, err := strconv.Atoi(string(jmpRunes[1:]))
			if err != nil {
				log.Println(err)
				return -1, -1, steps
			}

			if operator == []rune("-")[0] {
				cursor -= num
			} else {
				cursor += num
			}

		case "acc":
			accRunes := []rune(command[1])
			operator := accRunes[0]

			num, err := strconv.Atoi(string(accRunes[1:]))
			if err != nil {
				log.Println(err)
				return -2, -2, steps
			}

			if operator == []rune("-")[0] {
				acc -= num
			} else {
				acc += num
			}
			cursor++
		}

		if cursor >= len(commands) || cursor < 0 {
			break
		}
	}

	if cursor == len(commands) {
		return acc, 1, steps
	} else {
		return acc, 0, steps
	}
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
