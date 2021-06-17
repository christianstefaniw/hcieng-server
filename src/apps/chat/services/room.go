package services

import (
	"context"
	"fmt"
	accounts "hciengserver/src/apps/account/services"
	"hciengserver/src/database"
	"hciengserver/src/hciengserver"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	Id         primitive.ObjectID `bson:"_id" json:"id"`
	clients    map[*client]bool
	broadcast  chan *message
	register   chan *client
	unregister chan *client
	Name       string     `json:"name"`
	Messages   []*message `json:"messages"`
	ctx        context.Context
	cancel     context.CancelFunc
}

// omit fields that are not loaded when all rooms are loaded
type MinRoomData struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name"`
}

var activeRooms sync.Map

func GetRoom(id string) (*Room, bool) {
	rm, ok := activeRooms.Load(id)
	return rm.(*Room), ok
}

func NewRoomAndStore(name string) (*Room, error) {
	ctx, cancel := context.WithCancel(context.Background())
	rm := &Room{
		Id:         primitive.NewObjectID(),
		Name:       name,
		clients:    make(map[*client]bool),
		broadcast:  make(chan *message),
		register:   make(chan *client),
		unregister: make(chan *client),
		Messages:   make([]*message, 0),
		ctx:        ctx,
		cancel:     cancel,
	}

	err := rm.save()

	return rm, err
}

func (r *Room) save() error {
	r.saveToActiveRooms()
	return r.saveToDb()
}

func (r *Room) saveToActiveRooms() {
	activeRooms.LoadOrStore(r.Id.Hex(), r)
}

func (r *Room) saveToDb() error {
	_, err := database.GetMongoDBConn().
		Client().Database(hciengserver.DB_NAME).
		Collection(hciengserver.ROOMS_COLL).InsertOne(context.Background(), r)
	return err
}

func (r *Room) initPrivFields() {
	ctx, cancel := context.WithCancel(context.Background())
	r.clients = make(map[*client]bool)
	r.broadcast = make(chan *message)
	r.register = make(chan *client)
	r.unregister = make(chan *client)
	r.ctx = ctx
	r.cancel = cancel
}

func (r *Room) InitAndServe() {
	r.initPrivFields()
	r.saveToActiveRooms()
	go r.Serve()
}

func (r *Room) CheckClientInRoom(email string) (*client, bool) {
	for roomClient := range r.clients {
		if roomClient.EmailAddr == email {
			return roomClient, true
		}
	}
	return nil, false
}

func (r *Room) saveMessage(msg *message) error {
	query := bson.M{
		"_id": r.Id,
	}
	update := bson.M{
		"$push": bson.M{"messages": msg},
	}

	_, err := database.GetMongoDBConn().
		Client().Database(hciengserver.DB_NAME).
		Collection(hciengserver.ROOMS_COLL).UpdateOne(context.Background(), query, update)

	return err
}

func (r *Room) Serve() {
	for {
		select {
		case msg := <-r.broadcast:
			fmt.Println(msg)
			if err := r.saveMessage(msg); err == nil {
				for c := range r.clients {
					fmt.Println("ok")
					c.msg <- msg
				}
			} else {
				fmt.Fprintf(os.Stderr, "error sending message `%s`: %s", msg.Content, err)
			}
		case client := <-r.register:
			r.clients[client] = true
		case client := <-r.unregister:
			delete(r.clients, client)
		}
	}
}

// user is empty interface to prevent import cycle
// from middleware package
func JoinRoom(rmId string, user interface{}) {
	update := bson.M{
		"$push": bson.M{"rooms": rmId},
	}
	query := bson.M{
		"_id": user.(*accounts.Account).Id,
	}

	database.GetMongoDBConn().Client().
		Database(hciengserver.DB_NAME).
		Collection(hciengserver.ACCOUNT_COLL).
		UpdateOne(context.Background(), query, update)
}
