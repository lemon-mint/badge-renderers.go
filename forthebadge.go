package badgerenderers

import (
	"io"

	"github.com/lemon-mint/badge-renderers.go/forthebadge"
)

func TextWidth(s string) float64 {
	return float64(len(s)) * 9.256
}

func WriteForTheBadge(w io.Writer, label, message, color string) {
	const TEXT_MARGIN = 12
	const FONT_SCALE_UP_FACTOR = 10

	hasLabel := len(label) > 0

	var labelTextMinX, labelRectWidth, messageTextMinX, messageRectWidth float64
	labelTextMinX = TEXT_MARGIN
	if hasLabel {
		labelRectWidth = labelTextMinX + TextWidth(label) + TEXT_MARGIN
		messageTextMinX = labelRectWidth + TEXT_MARGIN
		messageRectWidth = 2*TEXT_MARGIN + TextWidth(message)
	} else {
		messageTextMinX = TEXT_MARGIN
		messageRectWidth = 2*TEXT_MARGIN + TextWidth(message)
	}

	// Draw Label Element

	midx := labelTextMinX + TextWidth(label)/2
	/* x: */ labelX := FONT_SCALE_UP_FACTOR * midx
	/* y: */ _ = 175
	/* textLength: */ labelTextLength := FONT_SCALE_UP_FACTOR * TextWidth(label)

	// Draw Message Element

	midx = messageTextMinX + TextWidth(message)/2
	/* x: */ messageX := FONT_SCALE_UP_FACTOR * midx
	/* y: */ _ = 175
	/* textLength: */ messageTextLength := FONT_SCALE_UP_FACTOR * TextWidth(message)

	var messageRectX float64

	if hasLabel {
		// Draw Label Rectangle

		/* width: */
		_ = labelRectWidth
		/* height: */ _ = 28

		// Draw Message Rectangle

		/* x: */
		messageRectX = labelRectWidth
		/* width: */ _ = messageRectWidth
		/* height: */ _ = 28
	} else {
		// Draw Message Rectangle

		/* width: */
		_ = messageRectWidth
		/* height: */ _ = 28
	}

	svgWidth := labelRectWidth + messageRectWidth

	forthebadge.WriteForTheBadgeMin(w, label, message, hasLabel, svgWidth, labelX, labelTextLength, messageX, messageTextLength, labelRectWidth, messageRectX, messageRectWidth, color)
}