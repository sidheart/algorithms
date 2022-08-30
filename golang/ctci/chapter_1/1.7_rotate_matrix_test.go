package chapter_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var red = Pixel{1, 0, 0, 0}
var green = Pixel{0, 1, 0, 0}
var blue = Pixel{0, 0, 1, 0}
var redder = Pixel{2, 0, 0, 0}
var greener = Pixel{0, 2, 0, 0}
var bluer = Pixel{0, 0, 2, 0}
var reddest = Pixel{3, 0, 0, 0}
var greenest = Pixel{0, 3, 0, 0}
var bluest = Pixel{0, 0, 3, 0}

var testImage = image{
{red, green, blue},
{redder, greener, bluer},
{reddest, greenest, bluest},
}

var testImageCW = image{
{reddest, redder, red},
{greenest, greener, green},
{bluest, bluer, blue},
}


func TestEqual(t *testing.T) {
	assert.True(t, testImage.Equal(testImage))
	assert.True(t, testImageCW.Equal(testImageCW))
	assert.False(t, testImage.Equal(testImageCW))
}

func TestRotateCW(t *testing.T) {

	assert.True(t, testImageCW.Equal(rotateCW(testImage)))
}