package description_merger_test

import (
	"testing"

	"github.com/sirbrokealot/tde/decklist_extractor"
	"github.com/sirbrokealot/tde/description_merger"
)

func TestHandlingOfInvalidBracketInput(t *testing.T) {
	description := "[[Puca's Mischief]] + [[card:Homeward Path]] allows one-sided creature to noncreature card swaps. Also for non existing [[Non Existant Card]]"
	decklist := map[string]decklist_extractor.CardInfo{
		"Puca's Mischief": {
			Count:              1,
			Name:               "Puca's Mischief",
			Foil:               true,
			Edition:            "MYSTOR",
			CustomInformations: []string{""},
			AltArtInformations: "",
			IsCommander:        false,
		},
	}
	expected := "[[card:Puca's Mischief (MYSTOR) *F*]] + [[card:Homeward Path]] allows one-sided creature to noncreature card swaps. Also for non existing Non Existant Card"
	merged, err := description_merger.MergeDescriptionWithDecklist(description, decklist)

	if err != nil {
		t.Errorf("Error merging description with decklist: %v", err)
	}

	if merged != expected {
		t.Errorf("Merged description does not match expected result.\nExpected: %s\nGot:      %s", expected, merged)
	}
}

func TestMergeComboLinkAndCarLinkInSameLine(t *testing.T) {
	description := "[[Illusions of Grandeur + Zedruu the Greathearted]]: Cast [[card:Illusions of Grandeur]], donate it, and gain +20 life. The new controller of [[card:Illusions of Grandeur]]"
	decklist := map[string]decklist_extractor.CardInfo{
		"Illusions of Grandeur": {
			Count:              1,
			Name:               "Illusions of Grandeur",
			Foil:               false,
			Edition:            "ICE",
			CustomInformations: []string{"Bad_Gifts"},
			AltArtInformations: "",
			IsCommander:        false,
		},
	}
	expected := "[[Illusions of Grandeur + Zedruu the Greathearted]]: Cast [[card:Illusions of Grandeur (ICE)]], donate it, and gain +20 life. The new controller of [[card:Illusions of Grandeur (ICE)]]"
	merged, err := description_merger.MergeDescriptionWithDecklist(description, decklist)

	if err != nil {
		t.Errorf("Error merging description with decklist: %v", err)
	}

	if merged != expected {
		t.Errorf("Merged description does not match expected result.\nExpected: %s\nGot:      %s", expected, merged)
	}
}

func TestMergeDescriptionWithDecklist(t *testing.T) {
	description := "This is a sample description with card names like Sol Ring and Mana Crypt."
	decklist := map[string]decklist_extractor.CardInfo{
		"Sol Ring": {
			Name:    "Sol Ring",
			Edition: "C21",
			Foil:    false,
		},
		"Mana Crypt": {
			Name:    "Mana Crypt",
			Edition: "2XM",
			Foil:    true,
		},
	}

	expected := "This is a sample description with card names like [[card:Sol Ring (C21)]] and [[card:Mana Crypt (2XM) *F*]]."
	merged, err := description_merger.MergeDescriptionWithDecklist(description, decklist)

	if err != nil {
		t.Errorf("Error merging description with decklist: %v", err)
	}

	if merged != expected {
		t.Errorf("Merged description does not match expected result.\nExpected: %s\nGot:      %s", expected, merged)
	}
}

func TestMergeDescriptionWithDecklistPlusJoinedCards(t *testing.T) {
	description := "This is a sample description with card names like [[Sol Ring + Mana Crypt]] and Academy Ruins."
	decklist := map[string]decklist_extractor.CardInfo{
		"Sol Ring": {
			Name:    "Sol Ring",
			Edition: "C21",
			Foil:    false,
		},
		"Mana Crypt": {
			Name:    "Mana Crypt",
			Edition: "2XM",
			Foil:    true,
		},
		"Academy Ruins": {
			Name:    "Academy Ruins",
			Edition: "2XM",
			Foil:    false,
		},
	}

	expected := "This is a sample description with card names like [[Sol Ring + Mana Crypt]] and [[card:Academy Ruins (2XM)]]."
	merged, err := description_merger.MergeDescriptionWithDecklist(description, decklist)

	if err != nil {
		t.Errorf("Error merging description with decklist: %v", err)
	}

	if merged != expected {
		t.Errorf("Merged description does not match expected result.\nExpected: %s\nGot:      %s", expected, merged)
	}
}

func TestMergeDescriptionWithCardLinksAndStandaloneNames(t *testing.T) {
	description := "This is a sample description with card names like [[card:Sol Ring]] and Mana Crypt."
	decklist := map[string]decklist_extractor.CardInfo{
		"Sol Ring": {
			Name:    "Sol Ring",
			Edition: "C21",
			Foil:    false,
		},
		"Mana Crypt": {
			Name:    "Mana Crypt",
			Edition: "2XM",
			Foil:    true,
		},
	}

	expected := "This is a sample description with card names like [[card:Sol Ring (C21)]] and [[card:Mana Crypt (2XM) *F*]]."
	merged, err := description_merger.MergeDescriptionWithDecklist(description, decklist)

	if err != nil {
		t.Errorf("Error merging description with decklist: %v", err)
	}

	if merged != expected {
		t.Errorf("Merged description does not match expected result.\nExpected: %s\nGot:      %s", expected, merged)
	}
}

func TestMergeDescriptionWithMultipleCardLinksAndCombinations(t *testing.T) {
	description := "This combo uses [[card:Sol Ring]] and [[card:Mana Crypt]] as ramp, and the main combo is [[Sol Ring + Mana Crypt + Academy Ruins]]."
	decklist := map[string]decklist_extractor.CardInfo{
		"Sol Ring": {
			Name:    "Sol Ring",
			Edition: "C21",
			Foil:    false,
		},
		"Mana Crypt": {
			Name:    "Mana Crypt",
			Edition: "2XM",
			Foil:    true,
		},
		"Academy Ruins": {
			Name:    "Academy Ruins",
			Edition: "2XM",
			Foil:    false,
		},
	}

	expected := "This combo uses [[card:Sol Ring (C21)]] and [[card:Mana Crypt (2XM) *F*]] as ramp, and the main combo is [[Sol Ring + Mana Crypt + Academy Ruins]]."
	merged, err := description_merger.MergeDescriptionWithDecklist(description, decklist)

	if err != nil {
		t.Errorf("Error merging description with decklist: %v", err)
	}

	if merged != expected {
		t.Errorf("Merged description does not match expected result.\nExpected: %s\nGot:      %s", expected, merged)
	}
}
func TestMergeDescriptionWithVariousCardFormats(t *testing.T) {
	description := `This is a combo deck using Whirlwind of Thought and Zedruu the Greathearted.
The main combo is [[Whirlwind of Thought + Zedruu the Greathearted]].
Here are the card links: [[card:Whirlwind of Thought]] and [[card:Zedruu the Greathearted]].
I currently own the version [[card:Whirlwind of Thought (IKO:348) *F*]] and [[card:Zedruu the Greathearted *A:32828*]].`

	decklist := map[string]decklist_extractor.CardInfo{
		"Whirlwind of Thought": {
			Name:    "Whirlwind of Thought",
			Edition: "IKO:348",
			Foil:    true,
		},
		"Zedruu the Greathearted": {
			Name:               "Zedruu the Greathearted",
			AltArtInformations: "*A:32828*",
			IsCommander:        true,
			Foil:               false,
		},
	}

	expected := `This is a combo deck using [[card:Whirlwind of Thought (IKO:348) *F*]] and [[card:Zedruu the Greathearted *A:32828*]].
The main combo is [[Whirlwind of Thought + Zedruu the Greathearted]].
Here are the card links: [[card:Whirlwind of Thought (IKO:348) *F*]] and [[card:Zedruu the Greathearted *A:32828*]].
I currently own the version [[card:Whirlwind of Thought (IKO:348) *F*]] and [[card:Zedruu the Greathearted *A:32828*]].`

	merged, err := description_merger.MergeDescriptionWithDecklist(description, decklist)

	if err != nil {
		t.Errorf("Error merging description with decklist: %v", err)
	}

	if merged != expected {
		t.Errorf("Merged description does not match expected result.\nExpected: %s\nGot:      %s", expected, merged)
	}
}

func TestMergeDescriptionWithBeneficialPairings(t *testing.T) {
	description := `offers two beneficial card pairings: [[Paradox Haze + Zedruu the Greathearted]] and [[Paradox Haze + Puca's Mischief]].`

	decklist := map[string]decklist_extractor.CardInfo{
		"Puca's Mischief": {
			Name:    "Puca's Mischief",
			Edition: "MYSTOR",
			Foil:    true,
		},
		"Paradox Haze": {
			Name:    "Paradox Haze",
			Edition: "MYSTOR",
			Foil:    true,
		},
		"Zedruu the Greathearted": {
			Name:               "Zedruu the Greathearted",
			AltArtInformations: "*A:32828*",
			IsCommander:        true,
			Foil:               false,
		},
	}

	expected := `offers two beneficial card pairings: [[Paradox Haze + Zedruu the Greathearted]] and [[Paradox Haze + Puca's Mischief]].`

	merged, err := description_merger.MergeDescriptionWithDecklist(description, decklist)

	if err != nil {
		t.Errorf("Error merging description with decklist: %v", err)
	}

	if merged != expected {
		t.Errorf("Merged description does not match expected result.\nExpected: %s\nGot:      %s", expected, merged)
	}
}

func TestMergeDescriptionWithComboNotInDecklist(t *testing.T) {
	description := `[[Academy Ruins + Mindslaver]] == gain control of all of target player's turns for [[symbol:11]][[symbol:U]].`

	academyRuins := decklist_extractor.CardInfo{
		Count:              1,
		Name:               "Academy Ruins",
		Foil:               false,
		Edition:            "",
		CustomInformations: []string{"land", "recursion"},
		AltArtInformations: "",
		IsCommander:        false,
	}

	decklist := map[string]decklist_extractor.CardInfo{
		"Academy Ruins": academyRuins,
	}

	expected := `[[Academy Ruins + Mindslaver]] == gain control of all of target player's turns for [[symbol:11]][[symbol:U]].`

	merged, err := description_merger.MergeDescriptionWithDecklist(description, decklist)

	if err != nil {
		t.Errorf("Error merging description with decklist: %v", err)
	}

	if merged != expected {
		t.Errorf("Merged description does not match expected result.\nExpected: %s\nGot:      %s", expected, merged)
	}
}

func TestMergeDescriptionWithSwappedComboNotInDecklist(t *testing.T) {
	description := `[[Mindslaver + Academy Ruins]] == gain control of all of target player's turns for [[symbol:11]][[symbol:U]].`

	academyRuins := decklist_extractor.CardInfo{
		Count:              1,
		Name:               "Academy Ruins",
		Foil:               false,
		Edition:            "",
		CustomInformations: []string{"land", "recursion"},
		AltArtInformations: "",
		IsCommander:        false,
	}

	decklist := map[string]decklist_extractor.CardInfo{
		"Academy Ruins": academyRuins,
	}

	expected := `[[Mindslaver + Academy Ruins]] == gain control of all of target player's turns for [[symbol:11]][[symbol:U]].`

	merged, err := description_merger.MergeDescriptionWithDecklist(description, decklist)

	if err != nil {
		t.Errorf("Error merging description with decklist: %v", err)
	}

	if merged != expected {
		t.Errorf("Merged description does not match expected result.\nExpected: %s\nGot:      %s", expected, merged)
	}
}

func TestMergeDescriptionWithUnknownCard(t *testing.T) {
	description := `[[card:Zedruu the Greathearted *A:32828*]] [[card:Teferi's Protection (2X2) (000) *F*]] [[card:Paradox Haze (TSP) *F*]] [[card:Card Not on The  List (2X2) (000) *F*]]`

	decklist := map[string]decklist_extractor.CardInfo{
		"Zedruu the Greathearted": {
			Name: "Zedruu the Greathearted",
		},
		"Teferi's Protection": {
			Name: "Teferi's Protection",
		},
		"Paradox Haze": {
			Name: "Paradox Haze",
		},
	}

	expected := `[[card:Zedruu the Greathearted]] [[card:Teferi's Protection]] [[card:Paradox Haze]] [[card:Card Not on The  List (2X2) (000) *F*]]`

	merged, err := description_merger.MergeDescriptionWithDecklist(description, decklist)

	if err != nil {
		t.Errorf("Error merging description with decklist: %v", err)
	}

	if merged != expected {
		t.Errorf("Merged description does not match expected result.\nExpected: %s\nGot:      %s", expected, merged)
	}
}

func TestExtractCardNameFromCardLink(t *testing.T) {
	testCases := []struct {
		input      string
		wantOutput string
	}{
		{"[[card:Paradox Haze (TSP) *F*]]", "Paradox Haze"},
		{"[[card:Teferi's Protection (2X2) (000) *F*]]", "Teferi's Protection"},
		{"[[card:Zedruu the Greathearted *A:32828*]]", "Zedruu the Greathearted"},
		{"[[card:Card Not on The List (2X2) (000) *F*]]", "Card Not on The List"},
		{"[[card:Some Cardname]]", "Some Cardname"},
	}

	for _, tc := range testCases {
		gotOutput, err := description_merger.ExtractCardNameFromCardLink(tc.input)
		if err != nil {
			t.Errorf("Error extracting card name: %v", err)
		}
		if gotOutput != tc.wantOutput {
			t.Errorf("ExtractCardNameFromCardLink(%q) = %q; want %q", tc.input, gotOutput, tc.wantOutput)
		}
	}
}
