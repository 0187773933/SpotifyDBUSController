package main

import (
	spotify_dbus "github.com/0187773933/SpotifyDBusController/controller"
	"time"
)

func main() {
	spotify := spotify_dbus.Controller{}
	spotify.Connect()
	spotify.OpenURI( "spotify:playlist:7BmaNDSoPDLHpxHzXaiJAN" )
	time.Sleep( 2 * time.Second )
	spotify.Next()
	time.Sleep( 2 * time.Second )
	spotify.Pause()
	time.Sleep( 2 * time.Second )
	spotify.Play()
	time.Sleep( 2 * time.Second )
	spotify.Previous()
}
