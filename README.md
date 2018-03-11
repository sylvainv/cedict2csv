# cedict2csv
Convert CEDICT dictionary file to csv

## Installation

```
go get github.com/sylvainv/cedict2csv
go install github.com/sylvainv/cedict2csv
```

## Usage

```
cedict2csv:
  -cedict string
    	Cedict file to parse
  -output string
    	Output csv
```

```
cedict2csv -h to display the help
cedict2csv -cedict cedict.u8 -output dictionary.csv
```
