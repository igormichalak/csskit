package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/igormichalak/csskit"
	"github.com/igormichalak/csskit/extract"
)

func main() {
	var outFilepath string
	var extractMode bool

	flag.StringVar(&outFilepath, "out", "output.css", "output CSS filepath.")
	flag.BoolVar(&extractMode, "extracted", false, "only prints extracted tokens.")
	flag.Parse()

	sourceFilepaths := flag.Args()

	if len(sourceFilepaths) == 0 {
		fmt.Println("please specify source files.")
		os.Exit(1)
	}

	var validFilepaths []string

	for _, fp := range sourceFilepaths {
		switch ext := filepath.Ext(fp); ext {
		case ".js", ".html", ".gohtml":
			validFilepaths = append(validFilepaths, fp)
		default:
			if len(ext) == 0 {
				fmt.Println("file extensions are required.")
			} else {
				fmt.Printf("unrecognized extension %q.\n", ext)
			}
			os.Exit(1)
		}
	}

	var lexerInput []string

	for _, fp := range validFilepaths {
		file, err := os.Open(fp)
		if err != nil {
			var pathErr *fs.PathError
			if errors.As(err, &pathErr) {
				fmt.Printf("can't open file %q.\n", pathErr.Path)
			} else {
				fmt.Printf("unrecognized error: %s.\n", err)
			}
			os.Exit(1)
		}

		var strs []string

		switch ext := filepath.Ext(fp); ext {
		case ".js":
			strs, err = extract.FromJS(file)
		case ".html", ".gohtml":
			strs, err = extract.FromHTML(file)
		}

		file.Close()

		if err != nil {
			fmt.Printf("unrecognized error: %s.\n", err)
			os.Exit(1)
		}

		lexerInput = append(lexerInput, strs...)
	}

	if extractMode {
		for _, str := range lexerInput {
			fmt.Printf("%q -> ", str)

			lex := csskit.NewLexer(str)
			var tok csskit.Token

			for tok.Type != csskit.TokenEOF {
				tok = lex.NextToken()

				switch tok.Type {
				case csskit.TokenKeyword:
					fmt.Printf("KEYWORD(%s) ", tok.Value)
				case csskit.TokenNumber:
					fmt.Printf("NUMBER(%s) ", tok.Value)
				case csskit.TokenUnit:
					fmt.Printf("UNIT(%s) ", tok.Value)
				case csskit.TokenHyphen:
					fmt.Printf("HYPHEN ")
				case csskit.TokenSpace:
					fmt.Printf("SPACE ")
				case csskit.TokenGarbage:
					fmt.Printf("GARBAGE ")
				case csskit.TokenEOF:
					fmt.Printf("EOF\n")
				}
			}
		}
		os.Exit(0)
	}
}
