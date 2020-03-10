package historyplugin

import (
	"context"
	"log"
	"time"

	. "github.com/Monibuca/engine"
	. "github.com/Monibuca/engine/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = &struct {
	MonogdbURL string
	Database   string
	Collection string
}{}

type broadcast struct {
	path          string   `bson:"path"`
	subscriber    []string `bson:"subscriber"`
	publisher     int      `bson:"publisher"`
	publisheraddr string   `bson:"publisheraddr"`
	recordType    string   `bson:"recordType"`
}

var (
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
	ctxback    = context.Background()
)

func init() {
	InstallPlugin(&PluginConfig{
		Name:   "History",
		Type:   PLUGIN_HOOK,
		UI:     CurrentDir("dashboard", "ui", "plugin-history.min.js"),
		Config: config,
		Run:    run,
	})
}

func run() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(config.MonogdbURL))
	if err != nil {
		log.Println(err)
	}
	ctx, cancel := context.WithTimeout(ctxback, 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
	}
	db = client.Database(config.Database)
	collection = db.Collection(config.Collection)
	OnPublishHook.AddHook(onPublish)
	OnSubscribeHook.AddHook(onSubscribe)
}
func onPublish(r * Room){

}
func onSubscribe(s *OutputStream){

}