package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	RESET = "\u001b[0m"
	BOLD_YELLOW = "\u001b[1;33m"
)

func Run(file, subject string, time int, multiplier float64) error {
	fmt.Printf("file=%s, subject=%s, time=%s%d%s, multiplier=%s%.1f%s\n", BOLD_YELLOW + file + RESET, BOLD_YELLOW + subject + RESET, BOLD_YELLOW, time, RESET, BOLD_YELLOW, multiplier, RESET)
	totalExp := Multiply(float64(time), multiplier)

	if time == 0 {
		return fmt.Errorf("Invalid time amount. Please specify a time greater than 0.")
	}

	if multiplier == 0.0 {
		return fmt.Errorf("Invalid multiplier amount. Please specify a multiplier.")
	}

	ofile, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer ofile.Close()

	data := [][]string{}
	if !fileExists(file) {
		header := []string{"subject", "time spent", "comfortability", "total"}
		data = append(data, header)
	}

	data = append(data, []string{subject, intToStr(time), floatToStr(multiplier), floatToStr(totalExp)})

	err = WriteCSV(ofile, data)
	return nil
}

func floatToStr(num float64) string {
	return strconv.FormatFloat(num, 'f', 1, 64)
}

func intToStr(num int) string {
	return strconv.Itoa(num)
}

func Multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func WriteCSV(file io.Writer, data [][]string) error {
	writer := csv.NewWriter(file)
	for _, values := range data {
		if err := writer.Write(values); err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
