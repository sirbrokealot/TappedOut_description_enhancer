package decklist_extractor

import (
	"fmt"
	"strconv"
	"strings"
)

type CardInfo struct {
	Count              int
	Name               string
	Foil               bool
	Edition            string
	CustomInformations []string
	AltArtInformations string
	IsCommander        bool
}

func ExtractDecklist(line string) CardInfo {
	var result CardInfo

	parts := strings.Split(line, " ")
	if len(parts) < 2 {
		return result
	}

	countStr := strings.Split(parts[0], "x")[0]
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return result
	}

	var edition string
	var nameParts []string
	result.CustomInformations = []string{}

	for i := 1; i < len(parts); i++ {
		if strings.HasPrefix(parts[i], "(") && strings.HasSuffix(parts[i], ")") {
			edition = parts[i][1 : len(parts[i])-1]
		} else if strings.Contains(parts[i], "*") {
			if parts[i] == "*CMDR*" {
				result.IsCommander = true
			} else if parts[i] == "*f*" {
				result.Foil = true
			} else if strings.HasPrefix(parts[i], "*A:") {
				altArt := parts[i]
				for j := i + 1; j < len(parts); j++ {
					if strings.HasSuffix(parts[j], "*") {
						altArt += " " + parts[j]
						break
					}
					altArt += " " + parts[j]
					i++
				}
				result.AltArtInformations = strings.ReplaceAll(strings.TrimSpace(altArt), " ", "")
			} else if strings.HasPrefix(parts[i], "*CMC") {
				continue // Skip prefixes that start with *CMC
			}
		} else if strings.HasPrefix(parts[i], "#") {
			customInfo := strings.TrimPrefix(parts[i], "#")
			result.CustomInformations = append(result.CustomInformations, customInfo)

			// check if there are additional custom information strings
			for j := i + 1; j < len(parts); j++ {
				if strings.HasPrefix(parts[j], "#") {
					customInfo = strings.TrimPrefix(parts[j], "#")
					result.CustomInformations = append(result.CustomInformations, customInfo)
					i++
				} else {
					break
				}
			}
		} else {
			nameParts = append(nameParts, parts[i])
		}
	}

	result.Count = count
	result.Name = strings.Join(nameParts, " ")
	result.Edition = edition

	return result
}
func (ci CardInfo) GetCardLink() string {
	var cardLink string
	if ci.Foil {
		if ci.AltArtInformations != "" {
			cardLink = fmt.Sprintf("[[card:%s %s *F*]]", ci.Name, ci.AltArtInformations)
		} else if ci.Edition != "" {
			cardLink = fmt.Sprintf("[[card:%s (%s) *F*]]", ci.Name, ci.Edition)
		} else {
			cardLink = fmt.Sprintf("[[card:%s *F*]]", ci.Name)
		}
	} else {
		if ci.AltArtInformations != "" {
			cardLink = fmt.Sprintf("[[card:%s %s]]", ci.Name, ci.AltArtInformations)
		} else if ci.Edition != "" {
			cardLink = fmt.Sprintf("[[card:%s (%s)]]", ci.Name, ci.Edition)
		} else {
			cardLink = fmt.Sprintf("[[card:%s]]", ci.Name)
		}
	}
	return cardLink
}
