package main

import (
  "github.com/hermanschaaf/cedict"
  "bufio"
  "os"
  "encoding/csv"
  "flag"
  "log"
)

func main () {
  cedictFileName := flag.String("cedict", "", "Cedict file to parse")
  outputFileName := flag.String("output", "", "Output csv")
  flag.Parse()

  if (flag.NFlag() != 2) {
    flag.PrintDefaults()
    return
  }

  cedictFile, cedictErr := os.Open(*cedictFileName)
  checkError("Could not read cedict file", cedictErr)
  defer cedictFile.Close()

  cedictReader := cedict.New(bufio.NewReader(cedictFile))

  csvFile, csvErr := os.Create(*outputFileName)
  checkError("Could not create output file", csvErr)
  defer csvFile.Close()

  writer := csv.NewWriter(csvFile)
  defer writer.Flush()

  var data = []string{"simplified", "traditional", "pinyin", "pinyin_with_tones", "pinyin_no_tones", "definition1", "definition2"}
  err := writer.Write(data)
  checkError("Could not write to file", err)

  for {
    err = cedictReader.NextEntry()
    if err != nil {
        break
    }
    entry := cedictReader.Entry()

    data = append([]string{
        entry.Simplified,
        entry.Traditional,
        entry.Pinyin,
        entry.PinyinWithTones,
        entry.PinyinNoTones,
      },
      entry.Definitions[:]...
    )
    err = writer.Write(data)
    checkError("Could not write to file", err)
  }
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
        panic(err.Error())
    }
}


