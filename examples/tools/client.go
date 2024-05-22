package tools

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Address       string `yaml:"address"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	AuthSource    string `yaml:"authSource"`
	AuthMechanism string `yaml:"authMechanism"`
}

// dev 環境 (k8s 集群)
var defaultConfig = Config{
	Address:       "mongodb://mongodb-headless.default.svc.cluster.local:27017/",
	Username:      "root",
	Password:      "7lGl2PVgbS",
	AuthSource:    "admin",
	AuthMechanism: "SCRAM-SHA-256",
}

var cli *mongo.Client
var initOnce sync.Once

func GetClient() *mongo.Client {
	if cli != nil {
		return cli
	}
	initOnce.Do(func() {
		credential := options.Credential{
			Username:      defaultConfig.Username,
			Password:      defaultConfig.Password,
			AuthSource:    defaultConfig.AuthSource,
			AuthMechanism: defaultConfig.AuthMechanism,
		}
		clientOptions := options.Client().ApplyURI(defaultConfig.Address).SetAuth(credential)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			panic(err)
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			panic(err)
		}
		cli = client
	})
	return cli
}

func Close() {
	cli.Disconnect(context.Background())

}
