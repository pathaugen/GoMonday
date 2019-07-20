package main
import ()

// Color Variables
var (

  // Internal variables (not available to other packages)
  breakspace      = "\n"

  // Clear all color settings in a string
  Clr             = "\u001b[0m"

  // Brighten color values
  Bold            = "\u001b[1m"
  Bright          = Bold

  // String colors
  Black           = "\u001b[30m"
  Red             = "\u001b[31m"
  Green           = "\u001b[32m"
  Yellow          = "\u001b[33m"
  Blue            = "\u001b[34m"
  Magenta         = "\u001b[35m"
  Cyan            = "\u001b[36m"
  White           = "\u001b[37m"

  // Brightened color volues
  BrightBlack     = Bright + "\u001b[30m"
  BrightRed       = Bright + "\u001b[31m"
  BrightGreen     = Bright + "\u001b[32m"
  BrightYellow    = Bright + "\u001b[33m"
  BrightBlue      = Bright + "\u001b[34m"
  BrightMagenta   = Bright + "\u001b[35m"
  BrightCyan      = Bright + "\u001b[36m"
  BrightWhite     = Bright + "\u001b[37m"

  // String background colors
  BGBlack         = "\u001b[40m"
  BGRed           = "\u001b[41m"
  BGGreen         = "\u001b[42m"
  BGYellow        = "\u001b[43m"
  BGBlue          = "\u001b[44m"
  BGMagenta       = "\u001b[45m"
  BGCyan          = "\u001b[46m"
  BGWhite         = "\u001b[47m"
  
)
