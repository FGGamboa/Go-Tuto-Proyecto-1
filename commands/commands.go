package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"proyecto-1/expenses"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {
	fmt.Print("->")
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	str = strings.Replace(str, "\r\n", "", 1)

	return str, nil
}

func ShowInConsole(expensesList []float32) {
	fmt.Println(contentString(expensesList))
}

func contentString(expensesList []float32) string {
	builder := strings.Builder{}

	max, min, avg, total := expensesDetails(expensesList)

	fmt.Println("")
	for i, expense := range expensesList {
		builder.WriteString(fmt.Sprintf("Expense: %6.2f\n", expense))

		if i == len(expensesList)-1 {
			builder.WriteString(fmt.Sprintf("==========================\n"))
			builder.WriteString(fmt.Sprintf("Total: %6.2f\n", total))
			builder.WriteString(fmt.Sprintf("Max: %6.2f\n", max))
			builder.WriteString(fmt.Sprintf("Min: %6.2f\n", min))
			builder.WriteString(fmt.Sprintf("Average: %6.2f\n", avg))
			builder.WriteString(fmt.Sprintf("=========================="))
		}
	}

	return builder.String()
}

func expensesDetails(expensesList []float32) (max, min, avg, total float32) {
	if len(expensesList) == 0 {
		return max, min, avg, total
	}
	min = expenses.Min(expensesList...)
	max = expenses.Max(expensesList...)
	total = expenses.Sum(expensesList...)
	avg = expenses.Average(expensesList...)

	return max, min, avg, total
}

func Export(fileName string, list []float32) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	err = newFunction(w, list)
	if err != nil {
		return err
	}
	return w.Flush()
}

func newFunction(w *bufio.Writer, list []float32) error {
	_, err := w.WriteString(contentString(list))
	return err
}
