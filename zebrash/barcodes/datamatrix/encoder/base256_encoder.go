package encoder

import (
	"fmt"
)

type Base256Encoder struct{}

func NewBase256Encoder() Encoder {
	return Base256Encoder{}
}

func (enc Base256Encoder) getEncodingMode() int {
	return HighLevelEncoder_BASE256_ENCODATION
}

func (enc Base256Encoder) encode(context *EncoderContext) error {
	buffer := make([]byte, 0)
	buffer = append(buffer, []byte{0, 0}...) //Initialize length field
	for context.HasMoreCharacters() {
		c := context.GetCurrentChar()
		buffer = append(buffer, c)

		context.pos++

		newMode := HighLevelEncoder_lookAheadTest(context.GetMessage(), context.pos, enc.getEncodingMode())
		if newMode != enc.getEncodingMode() {
			// Return to ASCII encodation, which will actually handle latch to new mode
			context.SignalEncoderChange(HighLevelEncoder_ASCII_ENCODATION)
			break
		}
	}
	dataCount := len(buffer) - 2
	lengthFieldSize := 1
	currentSize := context.GetCodewordCount() + dataCount + lengthFieldSize
	e := context.UpdateSymbolInfoByLength(currentSize)
	if e != nil {
		return fmt.Errorf("failed to update symbol info by length: %w", e)
	}

	mustPad := (context.GetSymbolInfo().GetDataCapacity() - currentSize) > 0
	if context.HasMoreCharacters() || mustPad {
		if dataCount <= 249 {
			buffer = buffer[1:]
			buffer[0] = byte(dataCount)
		} else if dataCount <= 1555 {
			buffer[0] = byte((dataCount / 250) + 249)
			buffer[1] = byte(dataCount % 250)
		} else {
			return fmt.Errorf("message length not in valid ranges: %v", dataCount)
		}
	}
	for i, c := 0, len(buffer); i < c; i++ {
		context.WriteCodeword(base256Randomize255State(
			buffer[i], context.GetCodewordCount()+1))
	}
	return nil
}

func base256Randomize255State(ch byte, codewordPosition int) byte {
	pseudoRandom := ((149 * codewordPosition) % 255) + 1
	tempVariable := int(ch) + pseudoRandom
	if tempVariable <= 255 {
		return byte(tempVariable)
	}
	return byte(tempVariable - 256)
}
