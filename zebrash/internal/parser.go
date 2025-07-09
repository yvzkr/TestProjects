package internal

import (
	"fmt"
	"strings"

	"github.com/ingridhq/zebrash/elements"
	"github.com/ingridhq/zebrash/parsers"
	"github.com/ingridhq/zebrash/printers"
)

type Parser struct {
	printer        *printers.VirtualPrinter
	commandParsers []*parsers.CommandParser
}

func NewParser() *Parser {
	return &Parser{
		printer: printers.NewVirtualPrinter(),
		commandParsers: []*parsers.CommandParser{
			parsers.NewLabelHomeParser(),
			parsers.NewLabelReversePrintParser(),
			parsers.NewGraphicBoxParser(),
			parsers.NewGraphicCircleParser(),
			parsers.NewGraphicFieldParser(),
			parsers.NewGraphicDiagonalLineParser(),
			parsers.NewGraphicSymbolParser(),
			parsers.NewChangeDefaultFontParser(),
			parsers.NewChangeFontParser(),
			parsers.NewChangeCharsetParser(),
			parsers.NewFieldOriginParser(),
			parsers.NewFieldTypesetParser(),
			parsers.NewFieldBlockParser(),
			parsers.NewFieldSeparatorParser(),
			parsers.NewFieldDataParser(),
			parsers.NewFieldValueParser(),
			parsers.NewFieldOrientationParser(),
			parsers.NewFieldReversePrintParser(),
			parsers.NewHexEscapeParser(),
			parsers.NewMaxicodeParser(),
			parsers.NewBarcode128Parser(),
			parsers.NewBarcode2of5Parser(),
			parsers.NewBarcode39Parser(),
			parsers.NewBarcodePdf417Parser(),
			parsers.NewBarcodeAztecParser(),
			parsers.NewBarcodeDatamatrixParser(),
			parsers.NewBarcodeQrParser(),
			parsers.NewDownloadGraphicsParser(),
			parsers.NewImageLoadParser(),
			parsers.NewRecallGraphicsParser(),
			parsers.NewBarcodeFieldDefaults(),
			parsers.NewPrintWidthParser(),
		},
	}
}

func (p *Parser) Parse(zplData []byte) ([]elements.LabelInfo, error) {
	var results []elements.LabelInfo
	var resultElements []any

	const startCode = "^XA"
	const endCode = "^XZ"

	commands, err := splitZplCommands(zplData)
	if err != nil {
		return nil, fmt.Errorf("failed to split zpl commands: %w", err)
	}

	for _, command := range commands {
		if strings.ToUpper(command) == startCode {
			p.printer.NextDownloadFormatName = ""
			continue
		}

		if strings.ToUpper(command) == endCode {
			if len(resultElements) > 0 {
				results = append(results, elements.LabelInfo{
					DownloadFormatName: p.printer.NextDownloadFormatName,
					PrintWidth:         p.printer.PrintWidth,
					Elements:           resultElements,
				})
			}

			resultElements = nil
			continue
		}

		for _, cp := range p.commandParsers {
			if !cp.CanParse(command) {
				continue
			}

			el, err := cp.Parse(command, p.printer)
			if err != nil {
				return nil, fmt.Errorf("failed to parse zpl command %v: %w", command, err)
			}

			if el != nil {
				resultElements = append(resultElements, el)
			}
		}
	}

	return results, nil
}

func splitZplCommands(zplData []byte) ([]string, error) {
	data := strings.ReplaceAll(string(zplData), "\n", "")
	data = strings.ReplaceAll(data, "\r", "")
	data = strings.ReplaceAll(data, "\t", "")

	caret := byte('^')
	tilde := byte('~')

	var buff strings.Builder
	var results []string

	const changeTildeCode = "CT"
	const changeCaretCode = "CC"

	for i := 0; i < len(data); i++ {
		c := data[i]
		command := buff.String()

		var isCt, isCC bool
		if buff.Len() == 4 {
			isCt = strings.Index(command, changeTildeCode) == 1
			isCC = strings.Index(command, changeCaretCode) == 1
		}

		if c == caret || c == tilde || isCt || isCC {
			buff.Reset()
			command = normalizeCommand(command, tilde, caret)

			switch {
			case isCt:
				tilde = command[3]
			case isCC:
				caret = command[3]
			default:
				results = append(results, command)
			}
		}

		if err := buff.WriteByte(c); err != nil {
			return nil, err
		}
	}

	command := buff.String()

	if len(command) > 0 {
		command = normalizeCommand(command, tilde, caret)
		results = append(results, command)
	}

	return results, nil
}

func normalizeCommand(command string, tilde, caret byte) string {
	if caret != '^' && command[0] == caret {
		command = "^" + command[1:]
	}

	if tilde != '~' && command[0] == tilde {
		command = "~" + command[1:]
	}

	return strings.TrimLeft(command, " ")
}
