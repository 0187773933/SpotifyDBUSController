package controller

import (
	"fmt"
	godbus "github.com/godbus/dbus"
	"time"
)

// https://github.com/48723247842/SpotifyDBUSController/blob/master/python_app/SpotifyDBusController.py
// https://godoc.org/github.com/guelfey/go.dbus#example-Object-Go
// https://github.com/search?p=4&q=dbus+language%3Ago&type=Repositories
// https://github.com/hoffoo/spotify-ctrl/blob/master/spotify.go
// https://github.com/sticreations/spotigopher/blob/master/spotigopher/spotigopher.go

type Controller struct {
	dbus godbus.BusObject
	Abilities struct {
		CanControl string
		CanGoNext string
		CanGoPrevious string
		CanPause string
		CanPlay string
		CanSeek string
	}
	Status struct {
		Shuffle string
		MaximumRate string
		MinimumRate string
		Rate string
		Volume float64
		Position int64
		LoopStatus string
		Playback string
		Metadata struct {
			TrackID string
			Artist []string
			Title string
			Album string
			TrackNumber int32
			Rating int
			Status string
			Url string
			ArtUrl string
			ArtFile string
		}
	}
}

func ( spotify *Controller ) Connect() {
	if spotify.dbus == nil {
		connection , err := godbus.SessionBus()
		if err != nil { panic( err ) }
		spotify.dbus = connection.Object( "org.mpris.MediaPlayer2.spotify" , "/org/mpris/MediaPlayer2" )
	}
}

func ( spotify *Controller ) Metadata() {
	metadata , err := spotify.dbus.GetProperty( "org.mpris.MediaPlayer2.Player.Metadata" )
	if err != nil { panic( err ) }
	fmt.Println( metadata )
}

func ( spotify *Controller ) PlaybackStatus() ( status string ) {
	status = "failed"
	playback_status , playback_status_error := spotify.dbus.GetProperty( "org.mpris.MediaPlayer2.Player.PlaybackStatus" )
	if playback_status_error != nil { fmt.Println( "Could not get Playback Status" ); panic( playback_status_error ) }
	status = playback_status.String()
	status = status[ 1 : ( len( status )-1 ) ] // Removes Quotes
	return
}

func ( spotify *Controller ) UpdateStatus() {
	// Metadata
	metadata , metadata_error := spotify.dbus.GetProperty( "org.mpris.MediaPlayer2.Player.Metadata" )
	if metadata_error != nil { fmt.Println( "Could not get Metadata" ); panic( metadata_error ) }
	metadata_map := metadata.Value().( map[ string ]godbus.Variant )
	spotify.Status.Metadata.TrackID = metadata_map["mpris:trackid"].Value().( string )
	spotify.Status.Metadata.Artist = metadata_map["xesam:artist"].Value().( []string )
	spotify.Status.Metadata.Title = metadata_map["xesam:title"].Value().( string )
	spotify.Status.Metadata.Album = metadata_map["xesam:album"].Value().( string )
	spotify.Status.Metadata.TrackNumber = metadata_map["xesam:trackNumber"].Value().( int32 )
	spotify.Status.Metadata.Url = metadata_map["xesam:url"].Value().( string )
	spotify.Status.Metadata.ArtUrl = metadata_map["mpris:artUrl"].Value().( string )

	// Playback Status
	playback_status , playback_status_error := spotify.dbus.GetProperty( "org.mpris.MediaPlayer2.Player.PlaybackStatus" )
	if playback_status_error != nil { fmt.Println( "Could not get Playback Status" ); panic( playback_status_error ) }
	spotify.Status.Playback = playback_status.Value().( string )

	// Volume
	volume_status , volume_status_error := spotify.dbus.GetProperty( "org.mpris.MediaPlayer2.Player.Volume" )
	if volume_status_error != nil { fmt.Println( "Could not get Volume Status" ); panic( volume_status_error ) }
	spotify.Status.Volume = volume_status.Value().( float64 )

	// Position
	position_status , position_status_error := spotify.dbus.GetProperty( "org.mpris.MediaPlayer2.Player.Position" )
	if position_status_error != nil { fmt.Println( "Could not get Position Status" ); panic( position_status_error ) }
	spotify.Status.Position = position_status.Value().( int64 )
}

func ( spotify *Controller ) Next() {
	result := spotify.dbus.Call( "org.mpris.MediaPlayer2.Player.Next" , 0 )
	if result.Err != nil { panic( result.Err ) }
	time.Sleep( 1 * time.Second )
	spotify.UpdateStatus()
}

func ( spotify *Controller ) Previous() {
	result := spotify.dbus.Call( "org.mpris.MediaPlayer2.Player.Previous" , 0 )
	if result.Err != nil { panic( result.Err ) }
	time.Sleep( 1 * time.Second )
	spotify.UpdateStatus()
}

func ( spotify *Controller ) Pause() {
	result := spotify.dbus.Call( "org.mpris.MediaPlayer2.Player.Pause" , 0 )
	if result.Err != nil { panic( result.Err ) }
	time.Sleep( 1 * time.Second )
	spotify.UpdateStatus()
}

func ( spotify *Controller ) Play() {
	result := spotify.dbus.Call( "org.mpris.MediaPlayer2.Player.Play" , 0 )
	if result.Err != nil { panic( result.Err ) }
	time.Sleep( 1 * time.Second )
	spotify.UpdateStatus()
}

func ( spotify *Controller ) PlayPause() {
	result := spotify.dbus.Call( "org.mpris.MediaPlayer2.Player.PlayPause" , 0 )
	if result.Err != nil { panic( result.Err ) }
	time.Sleep( 1 * time.Second )
	spotify.UpdateStatus()
}

func ( spotify *Controller ) Stop() {
	result := spotify.dbus.Call( "org.mpris.MediaPlayer2.Player.Stop" , 0 )
	if result.Err != nil { panic( result.Err ) }
	time.Sleep( 1 * time.Second )
	spotify.UpdateStatus()
}

func ( spotify *Controller ) OpenURI( uri string ) {
	result := spotify.dbus.Call( "org.mpris.MediaPlayer2.Player.OpenUri" , 0 ,  uri )
	if result.Err != nil { panic( result.Err ) }
	time.Sleep( 1 * time.Second )
	spotify.UpdateStatus()
}

// func ( spotify *Controller ) Seek( seconds string ) {
// 	result := spotify.dbus.Call( "org.mpris.MediaPlayer2.Player.Seek" , 0 ,  seconds )
// 	if result.Err != nil { panic( result.Err ) }
// 	sleep_duration , _ := time.ParseDuration( "600ms" )
// 	time.Sleep( sleep_duration )
// 	metadata , err := spotify.dbus.GetProperty( "org.mpris.MediaPlayer2.Player.Metadata" )
// 	if err != nil { panic( err ) }
// 	fmt.Println( metadata )
// }

// func ( spotify *Controller ) SetPosition( track_id string , position int ) {
// 	result := spotify.dbus.Call( "org.mpris.MediaPlayer2.Player.SetPosition" , 0 , track_id , position )
// 	if result.Err != nil { panic( result.Err ) }
// 	sleep_duration , _ := time.ParseDuration( "600ms" )
// 	time.Sleep( sleep_duration )
// 	metadata , err := spotify.dbus.GetProperty( "org.mpris.MediaPlayer2.Player.Metadata" )
// 	if err != nil { panic( err ) }
// 	fmt.Println( metadata )
// }