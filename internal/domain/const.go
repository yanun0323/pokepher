package domain

import "image"

var _defaultWindowsSize = image.Rect(0, 0, 1280, 960)

func DefaultWindowsBounds() (int, int) {
	return _defaultWindowsSize.Dx(), _defaultWindowsSize.Dy()
}

func DefaultWindowsSize() image.Rectangle {
	return _defaultWindowsSize
}

func DefaultWindowsWidth() int {
	return _defaultWindowsSize.Dx()
}

func DefaultWindowsHeight() int {
	return _defaultWindowsSize.Dy()
}
