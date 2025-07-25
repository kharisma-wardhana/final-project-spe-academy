package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"
	"time"

	"github.com/kharisma-wardhana/final-project-spe-academy/config"
	_ "github.com/kharisma-wardhana/final-project-spe-academy/docs"
	"github.com/kharisma-wardhana/final-project-spe-academy/entity"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/http/auth"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/http/handler"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/parser"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/presenter/json"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/mysql"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/redis"
	usecase_account "github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/account"
	usecase_log "github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/log"
	usecase_merchant "github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/merchant"
	usecase_qr "github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/qr"
	usecase_transaction "github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/transaction"

	"github.com/gin-contrib/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"

	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/subosito/gotenv"
)

func init() {
	_ = gotenv.Load()
}

// @title 						Go Skeleton!
// @version 					1.0
// @description 				This is a sample swagger for Go Skeleton
// @termsOfService 				http://swagger.io/terms/
// @contact.name 				API Support
// @contact.email 				rahmat.putra@spesolution.com
// @license.name				Apache 2.0
// @securityDefinitions.apikey 	Bearer
// @in							header
// @name						Authorization
// @license.url 				http://www.apache.org/licenses/LICENSE-2.0.html
// @host 						localhost:7011
// @BasePath /
func main() {
	// Initialize config variable from .env file
	cfg := config.NewConfig()

	app := fiber.New(config.NewFiberConfiguration(cfg))
	app.Get("/apidoc/*", swagger.HandlerDefault)

	// Middleware setup
	setupMiddleware(app, cfg)

	logger, err := config.NewZapLog(cfg.AppEnv)
	if err != nil {
		log.Fatal(err)
	}

	logger = logger.WithOptions(zap.AddCallerSkip(1))

	presenterJson := json.NewJsonPresenter()
	parser := parser.NewParser()

	// RabbitMQ Configuration (if needed)
	queue, err := config.NewRabbitMQInstance(context.Background(), &cfg.RabbitMQOption)
	if err != nil {
		log.Fatal(err)
	}

	// Redis Configuration (if needed)
	redisDB := config.NewRedis(&cfg.RedisOption)

	// MySQL/MariaDB Initialization
	gormLogger := config.NewGormLogMysqlConfig(&cfg.MysqlOption)
	mysqlDB, err := config.NewMysql(cfg.AppEnv, &cfg.MysqlOption, gormLogger)
	if err != nil {
		log.Fatal(err)
	}

	// AUTH : Write authetincation mechanism method (JWT, Basic Auth, etc.)

	// REPOSITORY : Write repository code here (database, cache, etc.)
	accountRepo := mysql.NewAccountRepository(mysqlDB)
	merchantRepo := mysql.NewMerchantRepository(mysqlDB)
	transactionRepo := mysql.NewTransactionRepository(mysqlDB)
	qrRepo := redis.NewQRRepository(redisDB)

	// USECASE : Write bussines logic code here (validation, business logic, etc.)
	logUseCase := usecase_log.NewLogUseCase(queue, logger)
	accountUseCase := usecase_account.NewAccountUseCase(logUseCase, accountRepo)
	merchantUseCase := usecase_merchant.NewMerchantUseCase(logUseCase, merchantRepo)
	transactionUseCase := usecase_transaction.NewTransactionUseCase(logUseCase, transactionRepo, qrRepo)
	qrUseCase := usecase_qr.NewQRUseCase(logUseCase, qrRepo, merchantRepo)

	api := app.Group("/api/v1")

	app.Get("/health-check", healthCheck)
	app.Get("/metrics", monitor.New())

	// HANDLER : Write handler code here (HTTP, gRPC, etc.)
	handler.NewMerchantHandler(parser, presenterJson, merchantUseCase, transactionUseCase, qrUseCase).Register(api)
	handler.NewAccountHandler(parser, presenterJson, accountUseCase).Register(api)

	signature := auth.NewSignature(parser, accountRepo, merchantRepo)
	app.Use(signature.VerifySignature)

	handler.NewTransactionHandler(parser, presenterJson, transactionUseCase).Register(api)

	// Handle Route not found
	app.Use(routeNotFound)

	runServerWithGracefulShutdown(app, cfg.ApiPort, 30)
}

func setupMiddleware(app *fiber.App, cfg *config.Config) {
	// Enable CORS if API shared in public
	if cfg.AppEnv == "production" {
		app.Use(
			cors.New(cors.Config{
				AllowCredentials: true,
				AllowOrigins:     cfg.AllowedCredentialOrigins,
				AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
			}),
		)
	}

	app.Use(
		logger.New(logger.Config{
			Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
			TimeFormat: "02-Jan-2006 15:04:05",
			TimeZone:   "Asia/Jakarta",
		}),
		recover.New(recover.Config{
			StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
				fmt.Println(c.Request().URI())
				stacks := fmt.Sprintf("panic: %v\n%s\n", e, debug.Stack())
				log.Println(stacks)
			},
			EnableStackTrace: true,
		}),
	)
}

func runServerWithGracefulShutdown(app *fiber.App, apiPort string, shutdownTimeout int) {
	var wg sync.WaitGroup
	wg.Add(1)

	// Run server in a goroutine
	go func() {
		defer wg.Done()
		log.Printf("Starting REST server, listening at %s\n", apiPort)
		if err := app.Listen(apiPort); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Capture OS signals for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down REST server...")

	// Timeout context for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(shutdownTimeout)*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("Error during server shutdown: %v", err)
	} else {
		log.Println("REST server shut down gracefully")
	}

	// Wait for goroutines to exit
	wg.Wait()
	log.Println("All tasks completed. Exiting application.")
}

var healthCheck = func(c *fiber.Ctx) error {
	return c.JSON(entity.GeneralResponse{
		Code:    200,
		Message: "OK!",
	})
}

var routeNotFound = func(c *fiber.Ctx) error {
	return c.Status(404).JSON(entity.GeneralResponse{
		Code:    404,
		Message: "Route Not Found!",
	})
}
