package main
import (
  "flag"
  "fmt"
  "os"
  "os/exec"
)

// Variables to load in from flags
var (
  // Debug Switch: Always have a quick switch to display debug data.
  debug           = false
  verbose         = false
  key             = ""
  board           = ""
  webport         = "80"
)

// Parse user provided flags
func flagParse() {

  // Setting all flag values from user provided flags
  // helpPtr     := flag.Bool( "help",     false, "Help documentation." )
  versionPtr  := flag.Bool( "version",  false, "Display version." )

  flag.BoolVar(&debug,            "debug",              false,            "View debug information.")
  flag.BoolVar(&verbose,          "verbose",            false,            "Operate in verbose mode.")

  flag.StringVar(&webport,        "webport",            "80",            "Port the webserver will run on.")

  flag.StringVar(&key,            "key",                "",               "Monday.com API v2 Key for GraphQL.")
  flag.StringVar(&board,          "board",              "",               "Monday.com Board ID for API v2 GraphQL Query.")

  flag.Parse()

  // Setup Shell
  setShell()

  // if ( *helpPtr ) { flagHelp() }
  if ( *versionPtr ) { fmt.Print( flagVersion() ); os.Exit(0) }

  // Testing of user provided values to ensure no format errors (report error and exit)
  // if variableName != "" { if b, s := validateString( variableName ); !b { fmt.Print( appSplash() + s ); os.Exit(0) } }
}

// Set Shell and Clear Screen if Verbose
func setShell() {
  shellCommand := ""
  if verbose {
    // Clear screen in verbose mode
    shellCommand += "cls"
  }
  cmd := exec.Command("cmd", "/c", shellCommand)
  cmd.Stdout = os.Stdout
  cmd.Run()
}

// Return the version number to the user
func flagVersion() string { return appSplash() + breakspace + Bold + Green + "  Application Version: " + Clr + version + breakspace }
