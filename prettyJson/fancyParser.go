package prettyJson

import (
	"bytes"
	"time"
)

const (
	RedStart    = "\033[1;32m"
	YellowStart = "\033[1;34m"
	ColorEnd    = "\033[0m"
)

func newline(dst *bytes.Buffer, prefix, indent string, depth int, addColor bool) {
	dst.WriteByte('\n')
	dst.WriteString(prefix)
	for i := 0; i < depth; i++ {
		dst.WriteString(indent)
	}
	if addColor {
		dst.WriteString(YellowStart)
	}
}

// Modified from json.Indent
func ColorfulIndent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
	time.Sleep(time.Second * 2)
	origLen := dst.Len()
	scan := newScanner()
	defer freeScanner(scan)
	needIndent := false
	depth := 0
	for _, c := range src {
		scan.bytes++
		v := scan.step(scan, c)
		if v == scanSkipSpace {
			continue
		}
		if v == scanError {
			break
		}
		if needIndent && v != scanEndObject && v != scanEndArray {
			needIndent = false
			depth++
			newline(dst, prefix, indent, depth, true)
		}

		// Emit semantically uninteresting bytes
		// (in particular, punctuation in strings) unmodified.
		if v == scanContinue {
			dst.WriteByte(c)
			continue
		}

		// Add spacing around real punctuation.
		switch c {
		case '{', '[':
			// delay indent so that empty object and array are formatted as {} and [].
			needIndent = true
			dst.WriteByte(c)
		case ',':
			dst.WriteString(ColorEnd)
			dst.WriteByte(c)
			newline(dst, prefix, indent, depth, true)

		case ':':
			dst.WriteString(ColorEnd)
			dst.WriteByte(c)
			dst.WriteByte(' ')
			dst.WriteString(RedStart)

		case '}', ']':
			dst.WriteString(ColorEnd)
			if needIndent {
				// suppress indent in empty object/array
				needIndent = false
			} else {
				depth--
				newline(dst, prefix, indent, depth, false)
			}
			dst.WriteByte(c)

		default:
			dst.WriteByte(c)
		}

	}
	if scan.eof() == scanError {
		dst.Truncate(origLen)
		return scan.err
	}
	return nil
}
