package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/sirbrokealot/tde/decklist_extractor"
	"github.com/sirbrokealot/tde/description_extractor"
	"github.com/sirbrokealot/tde/description_merger"
)

func main() {
	decklistFile := flag.String("l", "sample_decklist.txt", "Path to the decklist file")
	flag.StringVar(decklistFile, "decklist", "sample_decklist.txt", "Path to the decklist file")

	descriptionFile := flag.String("i", "sample_description.txt", "Path to the description input file")
	flag.StringVar(descriptionFile, "input_description", "sample_description.txt", "Path to the description input file")

	outputFile := flag.String("o", "output_description.txt", "Path to the output file")
	flag.StringVar(outputFile, "output_description", "output_description.txt", "Path to the output file")

	flag.Parse()

	// Read the decklist from the specified file path
	file, err := os.Open(*decklistFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Extract the decklist by processing the file line by line
	scanner := bufio.NewScanner(file)
	decklist := make(map[string]decklist_extractor.CardInfo)
	for scanner.Scan() {
		line := scanner.Text()
		cardInfo := decklist_extractor.ExtractDecklist(line)
		decklist[cardInfo.Name] = decklist_extractor.CardInfo{
			Count:              cardInfo.Count,
			Name:               cardInfo.Name,
			Foil:               cardInfo.Foil,
			Edition:            cardInfo.Edition,
			CustomInformations: cardInfo.CustomInformations,
			AltArtInformations: cardInfo.AltArtInformations,
			IsCommander:        cardInfo.IsCommander,
		}
	}

	// Check for errors during the line-by-line reading
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Extract the description using the description_extractor package
	description, err := description_extractor.ExtractDeckDescription(*descriptionFile)
	if err != nil {
		panic(err)
	}

	// Merge the decklist with the description
	mergedDescription, err := description_merger.MergeDescriptionWithDecklist(description, decklist)
	if err != nil {
		panic(err)
	}

	// Write the merged description to the output file
	descOutput, err := os.Create(*outputFile)
	if err != nil {
		panic(err)
	}
	defer descOutput.Close()

	fmt.Fprintf(descOutput, "%s", mergedDescription)
}
