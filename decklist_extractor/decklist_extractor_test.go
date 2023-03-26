package decklist_extractor

import (
	"reflect"
	"testing"
)

var (
	unwind = CardInfo{
		Count:              1,
		Name:               "Unwind",
		Foil:               true,
		Edition:            "",
		CustomInformations: []string{"pro"},
		AltArtInformations: "",
		IsCommander:        false,
	}
	vedalkenPlotter = CardInfo{
		Count:              1,
		Name:               "Vedalken Plotter",
		Foil:               true,
		Edition:            "GPT",
		CustomInformations: []string{"Engine_Cards"},
		AltArtInformations: "",
		IsCommander:        false,
	}
	weatheredWayfarer = CardInfo{
		Count:              1,
		Name:               "Weathered Wayfarer",
		Foil:               false,
		Edition:            "ONS",
		CustomInformations: []string{"fix"},
		AltArtInformations: "",
		IsCommander:        false,
	}
	whirlwindOfThought = CardInfo{
		Count:              1,
		Name:               "Whirlwind of Thought",
		Foil:               true,
		Edition:            "IKO:348",
		CustomInformations: []string{"draw"},
		AltArtInformations: "",
		IsCommander:        false,
	}
	zedruuTheGreathearted = CardInfo{
		Count:              1,
		Name:               "Zedruu the Greathearted",
		Foil:               false,
		Edition:            "",
		CustomInformations: []string{},
		AltArtInformations: "*A:32828*",
		IsCommander:        true,
	}
	thievingSkydiver = CardInfo{
		Count:              1,
		Name:               "Thieving Skydiver",
		Foil:               true,
		Edition:            "ZNR:335",
		CustomInformations: []string{"ramp"},
		AltArtInformations: "",
		IsCommander:        false,
	}
	smugglersShare = CardInfo{
		Count:              1,
		Name:               "Smuggler's Share",
		Foil:               false,
		Edition:            "NCC:122",
		CustomInformations: []string{"draw", "ramp"},
		AltArtInformations: "",
		IsCommander:        false,
	}
)

func TestExtractDecklist(t *testing.T) {
	unwind := unwind
	vedalkenPlotter := vedalkenPlotter
	weatheredWayfarer := weatheredWayfarer
	whirlwindOfThought := whirlwindOfThought
	zedruuTheGreathearted := zedruuTheGreathearted
	thievingSkydiver := thievingSkydiver
	smugglersShare := smugglersShare

	tests := []struct {
		input string
		want  CardInfo
	}{
		{"1x Unwind *f* #pro", unwind},
		{"1x Vedalken Plotter (GPT) *f* #Engine_Cards", vedalkenPlotter},
		{"1x Weathered Wayfarer (ONS) #fix", weatheredWayfarer},
		{"1x Whirlwind of Thought (IKO:348) *f* #draw", whirlwindOfThought},
		{"1x Zedruu the Greathearted *A: 32828* *CMDR*", zedruuTheGreathearted},
		{"1x Thieving Skydiver (ZNR:335) *f* #ramp *CMC:4*", thievingSkydiver},
		{"1x Smuggler's Share (NCC:122) #draw #ramp", smugglersShare},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := ExtractDecklist(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ExtractDecklist(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestGetCardLink(t *testing.T) {
	tests := []struct {
		cardInfo CardInfo
		want     string
	}{
		{
			cardInfo: unwind,
			want:     "[[card:Unwind *F*]]",
		},
		{
			cardInfo: vedalkenPlotter,
			want:     "[[card:Vedalken Plotter (GPT) *F*]]",
		},
		{
			cardInfo: weatheredWayfarer,
			want:     "[[card:Weathered Wayfarer (ONS)]]",
		},
		{
			cardInfo: whirlwindOfThought,
			want:     "[[card:Whirlwind of Thought (IKO:348) *F*]]",
		},
		{
			cardInfo: zedruuTheGreathearted,
			want:     "[[card:Zedruu the Greathearted *A:32828*]]",
		},
		{
			cardInfo: thievingSkydiver,
			want:     "[[card:Thieving Skydiver (ZNR:335) *F*]]",
		},
		{
			cardInfo: smugglersShare,
			want:     "[[card:Smuggler's Share (NCC:122)]]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.cardInfo.Name, func(t *testing.T) {
			got := tc.cardInfo.GetCardLink()
			if got != tc.want {
				t.Errorf("GetCardLink() = %s, want %s", got, tc.want)
			}
		})
	}
}
