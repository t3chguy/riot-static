package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func paginate(x []*Room, page int, size int) []*Room {
	var skip, end int
	length := len(x)

	if skip = (page - 1) * size; skip > length {
		skip = length
	}
	if end = skip + size; end > length {
		end = length
	}

	return x[skip:end]
}

func GetPublicRoomsList(c *gin.Context) {
	data.Once.Do(LoadPublicRooms)

	var page int
	var err error
	if page, err = strconv.Atoi(c.DefaultQuery("page", "1")); err != nil {
		page = 1
	}

	pageSize := 20

	data.RLock()
	c.HTML(http.StatusOK, "rooms.html", gin.H{
		"Rooms":    paginate(data.Ordered, page, pageSize),
		"NumRooms": data.NumRooms,
		"Page":     page,
	})
	data.RUnlock()
}

func GetPublicRoom(c *gin.Context) {
	roomId := c.Param("roomId")

	data.RLock()
	c.HTML(http.StatusOK, "room.html", gin.H{
		"Room": &data.Rooms[roomId],
	})
	data.RUnlock()
}

func GetPublicRoomServers(c *gin.Context) {
	roomId := c.Param("roomId")

	data.RLock()
	c.HTML(http.StatusOK, "room_servers.html", gin.H{
		"Room": &data.Rooms[roomId],
	})
	data.RUnlock()
}

func GetPublicRoomMembers(c *gin.Context) {
	roomId := c.Param("roomId")

	data.RLock()
	c.HTML(http.StatusOK, "room_members.html", gin.H{
		"Room": &data.Rooms[roomId],
	})
	data.RUnlock()
}

var data = struct {
	sync.Once
	sync.RWMutex
	NumRooms int
	Ordered  []*Room
	Rooms    map[string]*Room
}{}

func LoadPublicRooms() {
	fmt.Println("Loading public publicRooms")
	resp, err := cli.PublicRooms(0, "", "")

	if err == nil {
		b := []*Room{}
		c := map[string]*Room{}

		// filter on actually WorldReadable publicRooms
		for _, x := range resp.Chunk {
			if x.WorldReadable {
				var room *Room
				if data.Rooms[x.RoomId] != nil {
					room = data.Rooms[x.RoomId]
				} else {
					room = NewRoom(x)
				}
				b = append(b, room)
				c[x.RoomId] = room
			}
		}

		data.Lock()
		data.Rooms = c
		data.NumRooms = len(b)
		// copy order so we don't encounter slice hell
		data.Ordered = make([]*Room, data.NumRooms)
		copy(data.Ordered, b)

		data.Unlock()
	}

	if err != nil {
		panic(err)
	}
}

func FetchRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		roomId := c.Param("roomId")
		data.Rooms[roomId].Fetch()
	}
}

func FailIfNoRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		roomId := c.Param("roomId")
		if data.Rooms[roomId] == nil {
			c.String(http.StatusNotFound, "Room Not Found")
			c.Abort()
		}
	}
}

func main() {
	setupCli()

	//go LoadPublicRooms()
	// Synchronous cache fill
	data.Once.Do(LoadPublicRooms)

	router := gin.Default()
	router.SetHTMLTemplate(tpl)

	router.GET("/", GetPublicRoomsList)

	roomRouter := router.Group("/room/")
	{
		roomRouter.Use(FailIfNoRoom())
		roomRouter.Use(FetchRoom())

		roomRouter.GET("/:roomId", GetPublicRoom)
		roomRouter.GET("/:roomId/servers", GetPublicRoomServers)
		roomRouter.GET("/:roomId/members", GetPublicRoomMembers)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router.Run(":" + port)
}
