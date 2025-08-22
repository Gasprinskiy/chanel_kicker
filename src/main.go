package main

import (
	"chanel_kicker/src/config"

	"log"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	// defer cancel()

	config := config.NewConfig()

	// подключение к postgres
	pgdb, err := sqlx.Connect("pgx", config.PostgresURL)
	if err != nil {
		log.Fatalln("не удалось подключиться к базе postgres: ", err)
	}
	defer pgdb.Close()

	if err := pgdb.Ping(); err != nil {
		log.Fatal("ошибка при пинге postgres : ", err)
	}

	// инициализация логгера
	// hook := logger.NewPostgresHook(pgdb)
	// logger, err := logger.InitLogger(hook)
	// if err != nil {
	// 	log.Fatalln("Не удалось инициализировать логгер:", err)
	// }

	// инициализация session manager
	// sessionManager := transaction.NewSQLSessionManager(pgdb)

	// инициализация grpc соеденения
	grpcConn, err := grpc.NewClient(config.GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic("ошибка при подключении к grpc серверу: ", err)
	}
	defer grpcConn.Close()
}
