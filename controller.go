package controller

// https://github.com/48723247842/SpotifyDBUSController/blob/master/python_app/SpotifyDBusController.py
// https://godoc.org/github.com/guelfey/go.dbus#example-Object-Go
// https://github.com/search?p=4&q=dbus+language%3Ago&type=Repositories
// https://github.com/hoffoo/spotify-ctrl/blob/master/spotify.go

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
		Volume string
		Position string
		LoopStatus string
		Playback string
		Metadata struct {
			Artist string
			Title string
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

// func ( spotify *Controller ) Next() {

// }

// func ( spotify *Controller ) Next() {
// 	info := exec_process( "/bin/bash" , "-c" , "xrandr --prop | grep connected" )
// 	lines := strings.Split( info , "\n" )
// 	for _ , line := range lines {
// 		words := strings.Split( line , " " )
// 		if len( words ) < 3 { continue }
// 		switch words[ 2 ] {
// 			case "Primary":
// 				Name := words[ 0 ]
// 				size := words[ 3 ]
// 				size_components := strings.Split( size , "x" )
// 				X , _ := strconv.Atoi( size_components[ 0 ] )
// 				Y , _ := strconv.Atoi( strings.Split( size_components[ 1 ] , "+" )[ 0 ] )
// 				xdo.Monitors.Primary.Name = Name
// 				xdo.Monitors.Primary.X = X
// 				xdo.Monitors.Primary.Y = Y
// 			case "Secondary":
// 				Name := words[ 0 ]
// 				size := words[ 3 ]
// 				size_components := strings.Split( size , "x" )
// 				X , _ := strconv.Atoi( size_components[ 0 ] )
// 				Y , _ := strconv.Atoi( strings.Split( size_components[ 1 ] , "+" )[ 0 ] )
// 				xdo.Monitors.Secondary.Name = Name
// 				xdo.Monitors.Secondary.X = X
// 				xdo.Monitors.Secondary.Y = Y
// 		}
// 	}
// }

// func ( xdo *Wrapper ) Attach( options ...int ) {
// 	number_of_tries := 20
// 	sleep_milliseconds := 1000
// 	if len( options ) > 0 {
// 		number_of_tries = options[0]
// 	}
// 	if len( options ) > 1 {
// 		sleep_milliseconds = options[1]
// 	}
// 	duration , _ := time.ParseDuration( strconv.Itoa( sleep_milliseconds ) + "ms")
// 	for i := 0; i < number_of_tries; i++ {
// 		cmd := fmt.Sprintf( "xdotool search --desktop 0 --name '%s'" , xdo.Window.Name )
// 		info := exec_process( "/bin/bash" , "-c" , cmd )
// 		lines := strings.Split( info , "\n" )
// 		window_id , error := strconv.Atoi( lines[ 0 ] )
// 		if error != nil {
// 			time.Sleep( duration )
// 		} else {
// 			xdo.Window.Id = window_id
// 			return
// 		}
// 	}
// }

// func ( xdo *Wrapper ) Activate() {
// 	exec_process( "/bin/bash" , "-c" , fmt.Sprintf( "xdotool windowactivate %d" , xdo.Window.Id ) )
// }

// func ( xdo *Wrapper ) Focus() {
// 	exec_process( "/bin/bash" , "-c" , fmt.Sprintf( "xdotool windowfocus %d" , xdo.Window.Id ) )
// }

// func ( xdo *Wrapper ) Refocus() {
// 	xdo.Activate()
// 	sleep_duration , _ := time.ParseDuration( "300ms" )
// 	time.Sleep( sleep_duration )
// 	xdo.Focus()
// }

// func ( xdo *Wrapper ) GetGeometry() {
// 	xdo.Refocus()
// 	info := exec_process( "/bin/bash" , "-c" , "xdotool getactivewindow getwindowgeometry" )
// 	lines := strings.Split( info , "\n" )
// 	geometry_components := strings.Split( strings.Split( lines[ 2 ] , "Geometry: " )[ 1 ] , "x" )
// 	xdo.Window.Geometry.X , _ = strconv.Atoi( geometry_components[0] )
// 	xdo.Window.Geometry.Y , _ = strconv.Atoi( geometry_components[1] )
// 	xdo.Window.Geometry.Center.X = ( xdo.Window.Geometry.X / 2 )
// 	xdo.Window.Geometry.Center.Y = ( xdo.Window.Geometry.Y / 2 )
// }

// func ( xdo *Wrapper ) UnMaximize() {
// 	xdo.Refocus()
// 	exec_process( "/bin/bash" , "-c" , fmt.Sprintf( "wmctrl -rf %d -b remove,maximized_ver,maximized_horz" , xdo.Window.Id ) )
// }

// func ( xdo *Wrapper ) Maximize() {
// 	xdo.Refocus()
// 	exec_process( "/bin/bash" , "-c" , "xdotool key F11" )
// }

// func ( xdo *Wrapper ) FullScreen() {
// 	xdo.Maximize()
// }

// func ( xdo *Wrapper ) MoveMouse( X int , Y int ) {
// 	xdo.Refocus()
// 	exec_process( "/bin/bash" , "-c" , fmt.Sprintf( "xdotool mousemove %d %d" , X , Y ) )
// }

// func ( xdo *Wrapper ) LeftClick() {
// 	xdo.Refocus()
// 	exec_process( "/bin/bash" , "-c" , "xdotool click 1" )
// }

// func ( xdo *Wrapper ) RightClick() {
// 	xdo.Refocus()
// 	exec_process( "/bin/bash" , "-c" , "xdotool click 2" )
// }

// func ( xdo *Wrapper ) DoubleClick() {
// 	xdo.Refocus()
// 	exec_process( "/bin/bash" , "-c" , "xdotool click --repeat 2 --delay 200 1" )
// }

// func ( xdo *Wrapper ) CenterMouse() {
// 	xdo.Refocus()
// 	xdo.MoveMouse( xdo.Window.Geometry.Center.X , xdo.Window.Geometry.Center.Y )
// }

// func ( xdo *Wrapper ) PressKey( key string ) {
// 	xdo.Refocus()
// 	exec_process( "/bin/bash" , "-c" , fmt.Sprintf( "xdotool key '%s'" , key )  )
// }
