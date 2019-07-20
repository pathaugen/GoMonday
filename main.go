package main
import (
  "bufio"
  // "bytes"
  "context"
  "fmt"
  "io/ioutil"
  "net/http"
  "strconv"
  "time"
  "os"
)

// Application Version (Golang App, not HTML/CSS/JS Content):
var version = "0.0.01"

// Console Splash:
var breakline = Blue + "=========================================" + Clr
var appinfo = `
  ` + Blue + breakline + Bold + Cyan + `
    ___     __  __             _
   / __|___|  \/  |___ _ _  __| |__ _ _  _
  | (_ / _ \ |\/| / _ \ ' \/ _`+"`"+` / _`+"`"+` | || |
   \___\___/_|  |_\___/_||_\__,_\__,_|\_, |
                                      |__/` + Clr + `
  ` + Cyan + `Application: ` + White + `GoMonday` + Clr + `
  ` + Cyan + `http://localhost` + Yellow + `` + Cyan + `/` + Clr + `
  ` + Blue + breakline + Clr + `
`

// Main application flow: Respond to CLI and web based API requests
func main() {
  // Read and parse user provided flags
  flagParse()

  // Run a webserver to also respond to JSON API requests
  portNumber, _ := strconv.Atoi( webport )
  webServer( portNumber ) // Typically send port 80

  // Reading Input:
  scanner := bufio.NewScanner(os.Stdin)
  var inputText string

  // Listen for input and break the loop if inputText == "q" for quit:
  for ( inputText != "q" ) {

    // Setup shell, clear screen if in verbose mode
    setShell()

    // Display Application Splash:
    fmt.Print( appSplash() )

    // Toggle debug mode
    if ( inputText == "d" ) {

      // Toggle Debug Mode
      if ( debug ) {
        debug = false
      } else {
        debug = true
      }
      fmt.Print( breakspace + Bold + Cyan + "  Toggle Debug: " + Clr + debugStatus() + breakspace )
    } else if inputText != "" { }


    // Monday.com API v2 GraphQL tests that run on application startup
    fmt.Print( breakspace + Bold + Cyan + "  Monday.com API v2 GraphQL Tests:" + Clr + breakspace )
    if key == "" {
      fmt.Print( breakspace + Bold + Red + "  Missing Monday.com API v2 key!" + Clr + breakspace )
      fmt.Print( "  Launch Application with flag: -key [apikey]" + Clr + breakspace )
    } else {
      // fmt.Print( "  API Key: " + key + Clr + breakspace )
      fmt.Print( "  API Key:  " + "[hidden]" + Clr + breakspace )
    }
    if board == "" {
      fmt.Print( breakspace + Bold + Red + "  Missing Monday.com Board ID!" + Clr + breakspace )
      fmt.Print( "  Launch Application with flag: -board [boardid]" + Clr + breakspace )
    } else {
      fmt.Print( "  Board ID: " + board + Clr + breakspace )
    }
    if key != "" && board != "" {

      queryUrl := ""
      // jsonQuery := []byte(`{boards(ids:181745387){columns{id title}groups(ids: [topics, new_group54, new_group47, new_group11, new_group2]){title items {id name column_values{id value}}}}}`)

      // 1. Non-Auth Query Test
      fmt.Print( breakspace + Bold + Cyan + "  Query 1: Non-Auth Monday.com API v2 Query" + Clr + breakspace )
      queryUrl = "https://api.monday.com/v2"
      fmt.Print( "  " + Blue + queryUrl + Clr + breakspace )
      ctx, cncl := context.WithTimeout(context.Background(), time.Second * 10)
      defer cncl()
      req, _ := http.NewRequest( http.MethodPost, queryUrl, nil ) // bytes.NewBuffer(jsonQuery)
      req.Header.Set( "Content-Type", "application/json" )
      // req.Header.Set( "Authorization", key )
      res, _ := http.DefaultClient.Do(req.WithContext(ctx))
      byteValue, _ := ioutil.ReadAll(res.Body)
      fmt.Print( Bold + Blue + string(byteValue) + Clr + breakspace )
      defer res.Body.Close()

      // 2. Auth Query Test
      fmt.Print( breakspace + Bold + Cyan + "  Query 2: Authorized Monday.com API v2 Query" + Clr + breakspace )
      queryUrl = "https://api.monday.com/v2"
      fmt.Print( "  " + Blue + queryUrl + Clr + breakspace )
      ctx, cncl = context.WithTimeout(context.Background(), time.Second * 10)
      defer cncl()
      req, _ = http.NewRequest( http.MethodPost, queryUrl, nil ) // bytes.NewBuffer(jsonQuery)
      req.Header.Set( "Content-Type", "application/json" )
      req.Header.Set( "Authorization", key )
      res, _ = http.DefaultClient.Do(req.WithContext(ctx))
      byteValue, _ = ioutil.ReadAll(res.Body)
      fmt.Print( Bold + Blue + string(byteValue) + Clr + breakspace )
      defer res.Body.Close()

      // 3. Board Query Test
      fmt.Print( breakspace + Bold + Cyan + "  Query 3: Board Monday.com API v2 Query" + Clr + breakspace )
      queryUrl = "https://api.monday.com/v2?query={boards(ids:181745387){columns{id title}}}"
      fmt.Print( "  " + Blue + queryUrl + Clr + breakspace )
      ctx, cncl = context.WithTimeout(context.Background(), time.Second * 10)
      defer cncl()
      req, _ = http.NewRequest( http.MethodPost, queryUrl, nil ) // bytes.NewBuffer(jsonQuery)
      req.Header.Set( "Content-Type", "application/json" )
      req.Header.Set( "Authorization", key )
      res, _ = http.DefaultClient.Do(req.WithContext(ctx))
      byteValue, _ = ioutil.ReadAll(res.Body)
      fmt.Print( Bold + Blue + string(byteValue) + Clr + breakspace )
      defer res.Body.Close()
    }


    // User's command prompt asking for input
    // TODO: Options can include currentusers, stop/start webserver, reload data
    if verbose {
      fmt.Print( breakspace + Bold + Cyan + "  Specify an option:" + Clr +
        breakspace + "  [" + Yellow + "d" + Clr + "] Toggle Debug              " + Clr + debugStatus() )
      fmt.Print( breakspace +
        "  [" + Yellow + "q" + Clr + "] Quit Application" + Clr +
        breakspace + Yellow + "  > " + Clr )
    }

    // Await user input
    scanner.Scan()
    inputText = scanner.Text()
  }

  // Final application cleanup before exit, displaying completed/remaining statistics
  remainingStats()
}

// App Splash
func appSplash() string {
  if verbose {
    return appinfo
  } else {
    return ""
  }
}

// Remaining Statistics
func remainingStats() {
  if verbose {
    // TODO: Meaningful verbose exit values for user
    // fmt.Print( breakspace + Bold + Cyan + "  Remaining Statistics:" + Clr )
  } else {
    fmt.Print( breakspace + "  [" + Bold + Green + "SUCCESS" + Clr + "] " + Bold + Cyan + "Dashboard successful shutdown and exit on" + Clr + " [" + Yellow + "datetime" + Clr + "]" + breakspace )
  }
}

// Debug Status
func debugStatus() string {
  debugStatus := ``
  if debug {
    debugStatus += Bold + Green + "DEBUG ON" + Clr
  } else {
    debugStatus += Bold + Red + "DEBUG OFF" + Clr
  }
  return "[" + debugStatus + "]"
}
