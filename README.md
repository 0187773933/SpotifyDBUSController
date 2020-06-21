# XDoTool Wrapper

```
package main

import (
    "fmt"
    spotify_dbus "github.com/0187773933/SpotifyDBusController"
)

func main() {
    spotify := spotify_dbus.Controller{}
    spotify.Connect()
    fmt.Println( spotify.Metadata() )
}
```