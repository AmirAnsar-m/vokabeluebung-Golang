package main

import (
	"bufio" 
   "fmt"
	"os"
	"strings"
   "unicode"
)

type Vokabel struct {
	Begriff   string
	Bedeutung string
}

func main() {
	var vocabList []Vokabel

	vocabList = loadVokabeln()

	for {
		fmt.Println("Wählen Sie eine Aktion:")
		fmt.Println("1. Vokabel hinzufügen")
		fmt.Println("2. Vokabeln üben")
		fmt.Println("3. Alle Vokabeln löschen")
		fmt.Println("4. Einzelne Vokabel löschen")
		fmt.Println("5. Beenden")

		var auswahl int
		_, err := fmt.Scan(&auswahl)
		if err != nil {
			fmt.Println("Ungültige Eingabe, bitte erneut versuchen.")
			continue
		}
		clearBuffer()

		switch auswahl {
		case 1:
			addVokabeln(&vocabList)
			saveVokabeln(vocabList)
		case 2:
			pruefeVokabeln(vocabList)
		case 3:
			vocabList = []Vokabel{}
			saveVokabeln(vocabList)
			fmt.Println("Alle Vokabeln wurden gelöscht.")
		case 4:
			deleteVokabel(&vocabList)
			saveVokabeln(vocabList)
		case 5:
			fmt.Println("Programm wird beendet.")
			return
		default:
			fmt.Println("Ungültige Auswahl, bitte wählen Sie eine gültige Option.")
		}
	}
}

func addVokabeln(vocabList *[]Vokabel) {
	for {
		fmt.Println("Geben Sie ein neues Wort ein (oder 'fertig' zum Beenden):")
		word := readInput()
		if strings.ToLower(word) == "fertig" {
			break
		}

		if !isValidInput(word) {
			fmt.Println("Ungültige Eingabe! Das Wort darf keine Zahlen oder leere Eingaben enthalten.")
			continue
		}

		fmt.Println("Geben Sie die Bedeutung des Wortes ein:")
		meaning := readInput()

		if !isValidInput(meaning) {
			fmt.Println("Ungültige Eingabe! Die Bedeutung darf keine Zahlen oder leere Eingaben enthalten.")
			continue
		}

		*vocabList = append(*vocabList, Vokabel{Begriff: word, Bedeutung: meaning})
		fmt.Println("Das Wort wurde hinzugefügt!")
	}
}

func pruefeVokabeln(vocabList []Vokabel) {
	score := 0
	for _, v := range vocabList {
		fmt.Printf("Was bedeutet '%s'? ", v.Begriff)
		meaning := readInput()
		if strings.ToLower(meaning) == strings.ToLower(v.Bedeutung) {
			fmt.Println("Richtig!")
			score++
		} else {
			fmt.Printf("Leider falsch. Die richtige Antwort ist: %s\n", v.Bedeutung)
		}
	}
	fmt.Printf("Ihr Punktestand: %d von %d\n", score, len(vocabList))
}

func deleteVokabel(vocabList *[]Vokabel) {
	fmt.Println("Geben Sie das Wort ein, das Sie löschen möchten:")
	wordToDelete := readInput()

	for i, v := range *vocabList {
		if strings.ToLower(v.Begriff) == strings.ToLower(wordToDelete) {
			*vocabList = append((*vocabList)[:i], (*vocabList)[i+1:]...)
			fmt.Println("Das Wort wurde gelöscht.")
			return
		}
	}
	fmt.Println("Das Wort wurde nicht gefunden.")
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func isValidInput(input string) bool {
	for _, r := range input {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return len(input) > 0
}

func saveVokabeln(vocabList []Vokabel) {
	file, err := os.Create("vokabeln.txt")
	if err != nil {
		fmt.Println("Fehler beim Erstellen der Datei:", err)
		return
	}
	defer file.Close()

	for _, v := range vocabList {
		_, err := fmt.Fprintf(file, "%s;%s\n", v.Begriff, v.Bedeutung)
		if err != nil {
			fmt.Println("Fehler beim Speichern der Vokabeln:", err)
			return
		}
	}
}

func loadVokabeln() []Vokabel {
	var vocabList []Vokabel

	file, err := os.Open("vokabeln.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return vocabList
		}
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return vocabList
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		if len(parts) == 2 {
			vocabList = append(vocabList, Vokabel{Begriff: parts[0], Bedeutung: parts[1]})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Fehler beim Lesen der Datei:", err)
	}

	return vocabList
}

func clearBuffer() {
	bufio.NewReader(os.Stdin).ReadString('\n') 
}
