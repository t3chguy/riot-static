// This file is automatically generated by qtc from "room-common.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:1
package templates

//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:1
import "github.com/t3chguy/riot-static/mxclient"

//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:6
func StreamPrintRoomHeader(qw422016 *qt422016.Writer, roomInfo mxclient.RoomInfo) {
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:6
	qw422016.N().S(`<table id="roomHeader"><tr><td rowspan="2">`)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:10
	if roomInfo.AvatarURL.IsValid() {
		//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:10
		qw422016.N().S(`<img class="roomLogo" src="`)
		//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:11
		qw422016.E().S(roomInfo.AvatarURL.ToThumbURL(48, 48, "crop"))
		//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:11
		qw422016.N().S(`" />`)
		//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:12
	} else {
		//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:12
		qw422016.N().S(`<img class="roomLogo" src="./img/logo_missing.png" />`)
		//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:14
	}
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:14
	qw422016.N().S(`</td><td><h2>`)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:16
	qw422016.E().S(roomInfo.Name)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:16
	qw422016.N().S(`</h2></td><td class="rightAlign"><a href="./room/`)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:18
	qw422016.E().S(roomInfo.RoomID)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:18
	qw422016.N().S(`/members">`)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:18
	qw422016.N().D(roomInfo.NumMembers)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:18
	qw422016.N().S(` `)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:18
	qw422016.N().S(`Members</a></td></tr><tr><td class="maxWidth">`)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:22
	qw422016.E().S(roomInfo.Topic)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:22
	qw422016.N().S(`</td><td class="rightAlign"><a href="./room/`)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:24
	qw422016.E().S(roomInfo.RoomID)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:24
	qw422016.N().S(`/servers">`)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:24
	qw422016.N().D(roomInfo.NumServers)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:24
	qw422016.N().S(` `)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:24
	qw422016.N().S(`Servers</a></td></tr></table>`)
//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:28
}

//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:28
func WritePrintRoomHeader(qq422016 qtio422016.Writer, roomInfo mxclient.RoomInfo) {
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:28
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:28
	StreamPrintRoomHeader(qw422016, roomInfo)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:28
	qt422016.ReleaseWriter(qw422016)
//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:28
}

//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:28
func PrintRoomHeader(roomInfo mxclient.RoomInfo) string {
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:28
	qb422016 := qt422016.AcquireByteBuffer()
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:28
	WritePrintRoomHeader(qb422016, roomInfo)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:28
	qs422016 := string(qb422016.B)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:28
	qt422016.ReleaseByteBuffer(qb422016)
	//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:28
	return qs422016
//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:28
}

//line src\github.com\t3chguy\riot-static\templates\room-common.qtpl:31
func RoomBaseUrl(roomID string) string {
	return "./room/" + roomID
}