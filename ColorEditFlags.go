package imgui

const (
	ColorEditFlagsNoAlpha        = 1 << 1 //              // ColorEdit, ColorPicker, ColorButton: ignore Alpha component (read 3 components from the input pointer).
	ColorEditFlagsNoPicker       = 1 << 2 //              // ColorEdit: disable picker when clicking on colored square.
	ColorEditFlagsNoOptions      = 1 << 3 //              // ColorEdit: disable toggling options menu when right-clicking on inputs/small preview.
	ColorEditFlagsNoSmallPreview = 1 << 4 //              // ColorEdit, ColorPicker: disable colored square preview next to the inputs. (e.g. to show only the inputs)
	ColorEditFlagsNoInputs       = 1 << 5 //              // ColorEdit, ColorPicker: disable inputs sliders/text widgets (e.g. to show only the small preview colored square).
	ColorEditFlagsNoTooltip      = 1 << 6 //              // ColorEdit, ColorPicker, ColorButton: disable tooltip when hovering the preview.
	ColorEditFlagsNoLabel        = 1 << 7 //              // ColorEdit, ColorPicker: disable display of inline text label (the label is still forwarded to the tooltip and picker).
	ColorEditFlagsNoSidePreview  = 1 << 8 //              // ColorPicker: disable bigger color preview on right side of the picker, use small colored square preview instead.

	// User Options (right-click on widget to change some of them). You can set application defaults using SetColorEditOptions(). The idea is that you probably don't want to override them in most of your calls, let the user choose and/or call SetColorEditOptions() during startup.
	ColorEditFlagsAlphaBar         = 1 << 9  //              // ColorEdit, ColorPicker: show vertical alpha bar/gradient in picker.
	ColorEditFlagsAlphaPreview     = 1 << 10 //              // ColorEdit, ColorPicker, ColorButton: display preview as a transparent color over a checkerboard, instead of opaque.
	ColorEditFlagsAlphaPreviewHalf = 1 << 11 //              // ColorEdit, ColorPicker, ColorButton: display half opaque / half checkerboard, instead of opaque.
	ColorEditFlagsHDR              = 1 << 12 //              // (WIP) ColorEdit: Currently only disable 0.0f..1.0f limits in RGBA edition (note: you probably want to use ImGuiColorEditFlags_Float flag as well).
	ColorEditFlagsRGB              = 1 << 13 // [Inputs]     // ColorEdit: choose one among RGB/HSV/HEX. ColorPicker: choose any combination using RGB/HSV/HEX.
	ColorEditFlagsHSV              = 1 << 14 // [Inputs]     // "
	ColorEditFlagsHEX              = 1 << 15 // [Inputs]     // "
	ColorEditFlagsUint8            = 1 << 16 // [DataType]   // ColorEdit, ColorPicker, ColorButton: _display_ values formatted as 0..255.
	ColorEditFlagsFloat            = 1 << 17 // [DataType]   // ColorEdit, ColorPicker, ColorButton: _display_ values formatted as 0.0f..1.0f floats instead of 0..255 integers. No round-trip of value via integers.
	ColorEditFlagsPickerHueBar     = 1 << 18 // [PickerMode] // ColorPicker: bar for Hue, rectangle for Sat/Value.
	ColorEditFlagsPickerHueWheel   = 1 << 19 // [PickerMode] // ColorPicker: wheel for Hue, triangle for Sat/Value.
)
