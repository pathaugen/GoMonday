package main
import (
  "fmt"
  "net/http"
  "time"
  "os"
  "path/filepath"
)

// Detect folder binary is running in and prepend to asset loading paths
func getBinaryPath() string {
  ex, err := os.Executable()
  if err != nil {
    panic(err)
  }
  return filepath.Dir(ex)
}

func webServer( port int ) {

  // Configure Webserver
  webConfig( "api" )

  // Start Webserver on PORT #
  webStart( port )

}

// Webserver Configuration
func webConfig( apiEndpoint string ) {
  http.Handle( "/static/", http.StripPrefix( "/static/", http.FileServer( http.Dir( getBinaryPath() + "/assets" ) ) ) )

  http.HandleFunc( "/",                 handlerRoot )
  http.HandleFunc( "/api/monday.com/",  handlerApiMonday )
}

// Webserver Start
func webStart( port int ) {
  if verbose {
    // fmt.Print( breakspace + Bold + Cyan + "  Starting Internal Webserver (port #" + Clr + Yellow + webserverPort + Bold + Cyan + "):" + Clr )
    // fmt.Print( breakspace + "  http://localhost:" + Yellow + webserverPort + Clr + "/" + breakspace )
  }
  go func() {

    srv := &http.Server{
      ReadTimeout: 10 * time.Second,
      WriteTimeout: 15 * time.Second,
    }

    // HTTP Server:
    // TODO: Potentially setup logging and -> log.Println(srv.ListenAndServe()) but with custom logger
    srv.ListenAndServe()

    // HTTPS Server:
    // // log.Fatal( http.ListenAndServeTLS( ":" + webserverPort, "server.crt", "server.key", nil ) )
    // http.ListenAndServeTLS( ":443", "server.crt", "server.key", nil )

  }()
  if verbose {
    // webStatus()
  }
}

// Webserver Root Handler: http://localhost:{port}/
func handlerRoot(w http.ResponseWriter, r *http.Request) {
  output := htmlframework
	fmt.Fprintf( w, output )
}

// Monday.com API Handler
func handlerApiMonday(w http.ResponseWriter, r *http.Request) {
  output := htmlframework
	fmt.Fprintf( w, output )
}


var htmlframework = `
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
  </head>
  <body>
    <!-- loadingcontainer is removed after JS load of data from API -->
    <section style="font:2.0em sans-serif;text-align:center;" id="loadingcontainer">
      <h1>GoMonday</h1>
      <section>
        <!-- We could introduce FontAwesome later for extra visual flair, at expense of longer loads -->
        <i class="fas fa-circle-notch fa-spin"></i>
        Loading..
      </section>
    </section>
    <!-- bodycontainer is where HTML data is inserted after API data request comes back -->
    <section id="bodycontainer"></section>
  </body>
</html>
`
