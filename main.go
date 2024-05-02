package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/go-transcoder/transcoder/internal/application/command"
	"github.com/go-transcoder/transcoder/internal/application/services"
	postgres2 "github.com/go-transcoder/transcoder/internal/infrastructure/db/postgres"
	"github.com/go-transcoder/transcoder/internal/infrastructure/events/kafka"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {

	// we load the .env.test only if we are working locally
	PROJECT_ENV := os.Getenv("PROJECT_ENV")

	if PROJECT_ENV != "prod" {
		godotenv.Load(".env.test")
	}

	// DATABASE
	DBHOST := os.Getenv("DBHOST")
	DBPORT := os.Getenv("DBPORT")
	DBNAME := os.Getenv("DBNAME")
	DBUSER := os.Getenv("DBUSER")
	DBPASS := os.Getenv("DBPASS")
	SSLMODE := os.Getenv("SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", DBHOST, DBUSER, DBPASS, DBNAME, DBPORT, SSLMODE)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// KAFKA
	KAFKAHOST := os.Getenv("KAFKAHOST")
	eventService := kafka.NewKafkaEventService(KAFKAHOST)

	// AWS
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// APP
	videoRepository := postgres2.NewTranscodeVideoRepo(gormDB)
	videoService := services.NewTranscodeVideoService(videoRepository, eventService)

	transcodeCommad := command.TranscodeCommand{
		FileName: os.Getenv("OBJECT_NAME"),
		FilePath: fmt.Sprintf("uploads/%s", os.Getenv("OBJECT_NAME")),
		S3Cfg:    &cfg,
		S3Bucket: os.Getenv("INPUT_S3_BUCKET"),
	}

	response, err := videoService.Transcode(&transcodeCommad)

	fmt.Sprintf("Response got from transcoding %v", response)
}
