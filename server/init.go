package server

import (
	"fmt"
	"os"
	"sync"

	"github.com/Surya-7890/gateway/server/models"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	POSTGRES_URI string
	REDIS_URI    string
	DB           *gorm.DB
	Redis        *redis.Client
)

func Init() {
	err := godotenv.Load(".env.example")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	POSTGRES_URI = os.Getenv("POSTGRES_URI")
	REDIS_URI = os.Getenv("REDIS_URI")
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go connectPostgres(wg)
	go connectRedis(wg)
	wg.Wait()
}

func connectPostgres(wg *sync.WaitGroup) {
	defer wg.Done()
	db, err := gorm.Open(postgres.Open(POSTGRES_URI))
	if err != nil {
		panic(err)
	}
	DB = db
	migrator := DB.Migrator()
	migrator.CreateTable(&models.User{})
	fmt.Println("connected to postgres")
}

func connectRedis(wg *sync.WaitGroup) {
	defer wg.Done()
	Redis = redis.NewClient(&redis.Options{
		Addr:     REDIS_URI,
		Password: "",
	})
	fmt.Println("connected to redis")
}
