package fizzbuzz

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
	_ "testing"
	"time"

	"github.com/secsy/goftp"
)

func TestInput1ShouldBeDisplay1(t *testing.T) {
	v := Fizzbuzz(1)
	if "1" != v {
		t.Error("fizzbuzz of 1 should be '1' but have", v)
	}
}

func TestInput2ShouldBeDisplay2(t *testing.T) {
	v := Fizzbuzz(2)
	if "2" != v {
		t.Error("fizzbuzz of 2 should be '2' but have", v)
	}
}

func TestInput3ShouldBeDisplayfizz(t *testing.T) {
	v := Fizzbuzz(3)
	if "fizz" != v {
		t.Error("fizzbuzz of 2 should be 'fizz' but have", v)
	}
}

func TestInput4ShouldBeDisplay4(t *testing.T) {
	v := Fizzbuzz(4)
	if "4" != v {
		t.Error("fizzbuzz of 4 should be '4' but have", v)
	}
}

func TestInput5ShouldBeDisplaybuzz(t *testing.T) {
	v := Fizzbuzz(5)
	if "buzz" != v {
		t.Error("fizzbuzz of 5 should be 'buzz' but have", v)
	}
}

func TestInput6ShouldBeDisplayfizz(t *testing.T) {
	v := Fizzbuzz(6)
	if "fizz" != v {
		t.Error("fizzbuzz of 6 should be 'fizz' but have", v)
	}
}

func TestInput7ShouldBeDisplayfizz(t *testing.T) {
	v := Fizzbuzz(7)
	if "7" != v {
		t.Error("fizzbuzz of 7 should be '7' but have", v)
	}
}

func TestInput15ShouldBeDisplayfizzbuzz(t *testing.T) {
	v := Fizzbuzz(15)
	if "Fizzbuzz" != v {
		t.Error("fizzbuzz of 15 should be 'Fizzbuzz' but have", v)
	}

}

func TestReadFile(t *testing.T) {

	content, err := ioutil.ReadFile("/Users/jadsadaputkhiew/Workspace/Go/fizzbuzz/src/thermopylae.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}

func TestNewScanner(t *testing.T) {
	f, err := os.Open("/Users/jadsadaputkhiew/Workspace/Go/fizzbuzz/src/ZBUDE005.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		fmt.Printf("line: %s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func TestMain2(t *testing.T) {

	file, err := os.Open("/Users/jadsadaputkhiew/Workspace/Go/fizzbuzz/src/ZBUDE005_2.txt")

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

	for _, eachline := range txtlines {
		s := strings.Split(eachline, "\t")
		fmt.Println(" -wbs->", s[2])
		fmt.Println(" -2->", s[16])

	}

}

func TestFtp(t *testing.T) {
	config := goftp.Config{
		User:               "test",
		Password:           "test",
		ConnectionsPerHost: 10,
		Timeout:            10 * time.Second,
		Logger:             os.Stderr,
	}
	client, err := goftp.DialConfig(config, "localhost:2121")
	if err != nil {
		panic(err)
	}

	// Download a file to disk
	readme, err := os.Create("ZBUDE005_3.txt")
	if err != nil {
		panic(err)
	}

	err = client.Retrieve("ZBUDE005_2.txt", readme)
	if err != nil {
		panic(err)
	}

	// Upload a file from disk
	// bigFile, err := os.Open("big_file")
	// if err != nil {
	// 	panic(err)
	// }

	// err = client.Store("big_file", bigFile)
	// if err != nil {
	// 	panic(err)
	// }
}
