package main

import (
    "bufio"
    "log"
    "os"
    "strconv"
)

func main() {
    // Stap 1: Selecteer unieke codes en schrijf naar "unique.txt"
    selectUniqueCodes()

    // Stap 2: Verdeel "unique.txt" over vijf bestanden op basis van het eerste cijfer
    splitUniqueByFirstDigit()
}

// Functie om unieke codes te selecteren en naar "unique.txt" te schrijven
func selectUniqueCodes() {
    // Open "input.txt"
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Map om de eerste vijf cijfers bij te houden en slice voor geselecteerde regels
    seen := make(map[string]bool)
    var selected []string

    // Lees "input.txt" regel voor regel
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) < 12 {
            continue // Sla ongeldige regels over
        }
        code := line[:12]
        prefix := code[:5]
        if !seen[prefix] {
            selected = append(selected, line)
            seen[prefix] = true
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    // Schrijf geselecteerde regels naar "unique.txt"
    output, err := os.Create("unique.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer output.Close()
    for _, line := range selected {
        _, err := output.WriteString(line + "\n")
        if err != nil {
            log.Fatal(err)
        }
    }
}

// Functie om "unique.txt" te verdelen over vijf bestanden op basis van het eerste cijfer
func splitUniqueByFirstDigit() {
    // Open "unique.txt"
    file, err := os.Open("unique.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Maak een map voor de vijf outputbestanden
    files := make(map[string]*os.File)
    for i := 1; i <= 5; i++ {
        filename := strconv.Itoa(i) + "_unique.txt"
        f, err := os.Create(filename)
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()
        files[strconv.Itoa(i)] = f
    }

    // Lees "unique.txt" regel voor regel
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) < 12 {
            continue // Sla ongeldige regels over
        }
        firstDigit := string(line[0])
        if f, ok := files[firstDigit]; ok {
            _, err := f.WriteString(line + "\n")
            if err != nil {
                log.Fatal(err)
            }
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
