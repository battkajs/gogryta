package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	var ordlista, wordSlice []string
	var klisterOrd, capsFlag string
	var length int
	var exact bool

	flag.IntVar(&length, "len", 15, "lengt of returned password")
	flag.StringVar(&capsFlag, "case", "mix", "lower or upper")
	flag.BoolVar(&exact, "exact", false, "asdas")

	flag.Parse()
	if length < 0 {
		log.Fatalln("'-l'  must  be >= 0")
	}

	stdIn, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}
	if stdIn.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("No Stdin... fetching from default")
		ordlista = fillStdin(ordlista)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			sani := sanitizeWords(scanner.Text())
			ordlista = append(ordlista, sani)
		}
	}
	if exact == false {
		for i := 0; i < length; i = len(klisterOrd) {
			wordSlice = append(wordSlice, wordGenerator(ordlista))
			fmt.Printf("wordslice: %s\n", wordSlice)
			klisterOrd = strings.Join(wordSlice, "")
		}
	} else {
		wordSlice = append(wordSlice, wordGenerator(ordlista))
		if len(wordSlice[0]) < length - 1 {
			klisterOrd = growTrim(wordSlice, length)
		} else {
			wordSlice = append(wordSlice)
			klisterOrd = growTrim(wordSlice, length)
		}
	}

	switch capsFlag {
	case "lower":
		klisterOrd = strings.ToLower(klisterOrd)
	case "upper":
		klisterOrd = strings.ToUpper(klisterOrd)
	case "mix":
		klisterOrd = capsMixer(klisterOrd)
	default:
		log.Fatalln("Error! Usage: '-case lower' or '-case upper'")
	}
	fmt.Println(klisterOrd)
}

func fillStdin(o []string) []string {
	resp, err := http.Get("https://raw.githubusercontent.com/battkajs/ordlista_curated/main/sv_wordlist.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nBody:\n %s\n", resp.StatusCode, scanner)
	}
	for scanner.Scan() {
		o = append(o, scanner.Text())
	}
	return o
}
func sanitizeWords(sani string) string {
	wordSlice := strings.Split(sani, "")

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(wordSlice)/2; i++ {
		randomNumber := rand.Intn(len(wordSlice))
		wordSlice[randomNumber] = strings.ToUpper(wordSlice[randomNumber])
	}

	reg := regexp.MustCompile("^[[:ascii:]]+$")
	if reg.MatchString(sani) == false {
		for index, letter := range wordSlice {
			if reg.MatchString(letter) == false {
				wordSlice[index] = strconv.Itoa(rand.Intn(9))
			}
		}
	}

	processedString := strings.Join(wordSlice, "")
	return processedString
}
func wordGenerator(list []string) string {
	var ordSlice []string
	specChar := []string{"!", "#", "$", "%", "&", "/", "=", "+", "-", "§", "¤"}

	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(list))
	randomSpecIndex := rand.Intn(len(specChar))
	ordSlice = append(ordSlice, list[randomIndex], specChar[randomSpecIndex])

	word := strings.Join(ordSlice, "")
	return word
}
func capsMixer(word string) string {
	var wordByte []string

	rand.Seed(time.Now().UnixNano())

	for _, key := range word {
		wordByte = append(wordByte, string(key))
	}
	for j := 0; j < len(word)/2; j++ {
		rIntIndex := rand.Intn(len(wordByte))
		rUpper := strings.ToUpper(wordByte[rIntIndex])
		wordByte[rIntIndex] = rUpper
		word = strings.Join(wordByte, "")
	}
	return word
}
func growTrim(list []string, length int) string {
	rand.Seed(time.Now().UnixNano())
	if len(list[0]) <= length {
		for i := 0; 
	}
	return word
}
