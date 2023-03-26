package description_merger

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sirbrokealot/tde/decklist_extractor"
)

func createCardNameRegex(decklist map[string]decklist_extractor.CardInfo) *regexp.Regexp {
	escapedCardNames := make([]string, 0, len(decklist))
	for cardName := range decklist {
		escapedCardNames = append(escapedCardNames, regexp.QuoteMeta(cardName))
	}
	pattern := fmt.Sprintf(`\b(?:%s)\b`, strings.Join(escapedCardNames, "|"))
	return regexp.MustCompile(pattern)
}

// ExtractCardNameFromCardLink extracts the card name from a card link string in the format [[card:card_name ... ]].
// It returns the card name as a string or an error if the extraction fails.
func ExtractCardNameFromCardLink(match string) (string, error) {
	pattern := `\[\[card:([^\]\(\*]*)`
	re := regexp.MustCompile(pattern)

	result := re.FindStringSubmatch(match)
	if len(result) != 2 {
		return "", fmt.Errorf("Failed to extract card name from match: %s", match)
	}

	return strings.TrimSpace(result[1]), nil
}

func MergeDescriptionWithDecklist(description string, decklist map[string]decklist_extractor.CardInfo) (string, error) {
	// Create a regular expression pattern to match card names and links in the description
	cardNameRegex := createCardNameRegex(decklist)
	pattern := regexp.MustCompile(fmt.Sprintf(`(\[\[card:[^\]]+\]\])|(\[\[[^\]]+\s\+\s[^\]]+\]\])|(\[\[([^\]]+)\]\])|(%s)`, cardNameRegex))

	// Track the number of replacements
	replacements := 0

	mergedDescription := pattern.ReplaceAllStringFunc(description, func(match string) string {
		if strings.HasPrefix(match, "[[card:") && strings.HasSuffix(match, "]]") {
			// First match case [[card:card_name * ]] -> Replace with GetCardLink
			cardName, err := ExtractCardNameFromCardLink(match)
			if err != nil {
				return match
			}
			cardInfo, exists := decklist[cardName]
			if exists {
				cardLink := cardInfo.GetCardLink()
				fmt.Printf("Replacing card link: %s with %s\n", match, cardLink)
				replacements++
				return cardLink
			}
		} else if strings.HasPrefix(match, "[[") && strings.Contains(match, " + ") {
			// Second match case [[card_name + another_card_name ]] -> leave it.
			fmt.Printf("Keeping combo link: %s\n", match)
			return match
		} else if strings.HasPrefix(match, "[[") && strings.HasSuffix(match, "]]") {
			// Fourth case: standalone card name wrapped in square brackets
			cardName := strings.Trim(match, "[]")
			cardInfo, exists := decklist[cardName]
			if exists {
				cardLink := cardInfo.GetCardLink()
				fmt.Printf("Replacing invalid bracketed card name: %s to %s\n", match, cardLink)
				replacements++
				return cardLink
			} else {
				if strings.HasPrefix(match, "[[symbol") {
					return match
				} else {
					fmt.Printf("\n	WARNING invalid or unknow brackets detected! \n check the output for the match: %s\n", match)
					return cardName
				}
			}
		} else {
			// Third case: Replace stand-alone card_name with GetCardLink
			cardInfo, exists := decklist[match]
			if exists {
				cardLink := cardInfo.GetCardLink()
				fmt.Printf("Replacing stand alone: %s to %s\n", match, cardLink)
				replacements++
				return cardLink
			}
		}
		return match
	})

	fmt.Printf("\nTotal Replacements: %d\n", replacements)

	return mergedDescription, nil
}
