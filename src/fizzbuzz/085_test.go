package fizzbuzz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestO85(t *testing.T) {

	file, err := os.Open("/Users/jadsadaputkhiew/Workspace/Go/fizzbuzz/src/zbudr085.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for i, eachline := range txtlines {
		s := strings.Split(eachline, "\t")

		if i > 11 && len(s) > 4 {
			fmt.Println(len(s[4]))
			fmt.Println(i, "-1->", s[4], s[5], s[6], s[7], s[8], s[9], s[10])
				
		}
	}

}
