package main

import (
	"os"
	"todoapp"
	"todoapp/pkg/cache"
	"todoapp/pkg/handler"
	"todoapp/pkg/repository"
	"todoapp/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.Infoln("Запуск сервера")
	if err := godotenv.Load(); err != nil {
		logrus.Println("Ошибка при загрузке переменных окружения .env: \n %s", err.Error())
	}

	if err := initConfig(); err != nil {
		logrus.Fatalf("Ошибка (viper) при инициализации конгфига .yaml: \n %s", err.Error())
	}
	logrus.Infoln("Конфиг YAML инициализирован")

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		// Password: os.Getenv("DB_PASSWORD_GB"),
		Password: os.Getenv("DB_PASSWORD_LC"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("Ошибка при инициализации БД:\n %s ", err.Error())
	}
	logrus.Infoln("База данных Postgers инициализирована")

	client, err := cache.NewRedisClient(cache.Config{
		Addr:     viper.GetString("rdb.addr"),
		Username: viper.GetString("rdb.username"),
		// Password: os.Getenv("RDB_PASSWORD_GB"),
		Password: os.Getenv("rdb.password"),
		DB:       0,
	})
	if err != nil {
		logrus.Fatalf("Ошибка при инициализации Кеша:\n %s ", err.Error())
	}
	logrus.Infoln("Клиент Redis инициализирован")

	repos := repository.NewRepository(db)    // SQL
	caches := cache.NewCache(repos, client)  // Redis
	services := service.NewService(caches)   // Services
	handlers := handler.NewHandler(services) // Handler

	srv := new(todoapp.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoute()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	// viper.SetConfigName("global.config")
	return viper.ReadInConfig()
}
