package main

import (
	"log"
	"os"
	"todoapp"
	"todoapp/pkg/cache"
	"todoapp/pkg/handler"
	"todoapp/pkg/repository"
	"todoapp/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка при загрузке переменных окружения .env: \n %s", err.Error())
	}

	if err := initConfig(); err != nil {
		log.Fatalf("Ошибка (viper) при инициализации конгфига .yaml: \n %s", err.Error())
	}
	log.Println("Конфиг YAML инициализирован")

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD_GB"),
		// Password: os.Getenv("DB_PASSWORD_LC"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Ошибка при инициализации БД:\n %s ", err.Error())
	}
	log.Println("База данных Postgers инициализирована")

	client, err := cache.NewRedisClient(cache.Config{
		Addr:     viper.GetString("rdb.addr"),
		Username: viper.GetString("rdb.username"),
		Password: os.Getenv("RDB_PASSWORD_GB"),
		// Password: os.Getenv("RDB_PASSWORD_LC"),
		DB:       0,
	})
	if err != nil {
		log.Fatalf("Ошибка при инициализации Кеша:\n %s ", err.Error())
	}
	log.Println("Клиент Redis инициализирован")

	repos := repository.NewRepository(db)    // SQL
	caches := cache.NewCache(repos, client)  // Redis
	services := service.NewService(caches)   // Services
	handlers := handler.NewHandler(services) // Handler

	srv := new(todoapp.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoute()); err != nil {
		log.Fatalf("error occured while running http server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("global.config")
	return viper.ReadInConfig()
}
