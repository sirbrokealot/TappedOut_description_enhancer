# TappedOut Description Enhancer

This program merges card information from a TappedOut exported decklist with a
description text file.

## Usage

1. Build the program:
```
go build
```

2. Run the program with command-line arguments specifying the input decklist,
   description, and output files:

```
./tde -l sample_decklist.txt -i sample_description.txt -o output_description.txt
```

This will create an output file `output_description.txt` containing the merged
description with card information.

## Sample Files

Sample files are included in the repository:
- `sample_decklist.txt`: A sample decklist file exported from TappedOut.
- `sample_description_input.txt`: A sample description file with different card
  name formats that can be put in a TappedOut description field.

## Input Formats

The program supports the following card name formats:
- Standalone card names: `Card Name`
- TappedOut card links: `[[card:Card Name]]`
- TappedOut card links with additional information (foil, edition, or alternate
  art): `[[card:Card Name *F* #edition]]`
- Invalid TappedOut card links (missing "card:"): `[[Card Name]]`
- Combo links: `[[Card Name + Another Card Name]]`

The program replaces standalone card names and TappedOut card links with the
corresponding card information from the decklist.

The sample files provided in the repository can be used to test the program. You
can and should create your own decklist and description files following the same
format.


## Additional Limitations

- The program might not handle card names with special characters, such as
apostrophes, correctly. 
- The program assumes that card names are unique within the decklist. 
- If a card has multiple versions with different information (e.g., foil and
non-foil), the program may not handle it correctly. 
- If a card name contains a substring that matches another card name (e.g.,
"Jace, the Mind Sculptor" and "Jace"), the program may produce incorrect results
in some cases.
- The program does not handle cases where card names are split across multiple
 lines in the description.
- If the card name is like a normal name like 'Donate' or 'Smoke' you will get a
  card link for that :). 

These limitations do affect the program's accuracy in certain situations.
However, for most use cases, the program should work as expected when handling
typical TappedOut decklists and descriptions.

