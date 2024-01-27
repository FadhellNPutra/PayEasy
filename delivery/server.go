package delivery

import (
	"database/sql"
	"fmt"
	"payeasy/config"
	"payeasy/delivery/controller"
	"payeasy/delivery/middleware"
	"payeasy/shared/service"

	"payeasy/repository"
	"payeasy/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	merchantUC usecase.MerchantUseCase
	historyUC  usecase.HistoryUsecase
	usersUC    usecase.UsersUseCase
	authUsc    usecase.AuthUseCase
	engine     *gin.Engine
	jwtService service.JwtService
	host       string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)

	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	controller.NewMerchantController(s.merchantUC, rg, authMiddleware).Route()
	controller.NewHistoryController(s.historyUC, rg, authMiddleware).Route()
	controller.NewUsersController(s.usersUC, rg, authMiddleware).Route()
	controller.NewAuthController(s.authUsc, rg).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running on host %s, becauce error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic(err.Error())
	}

	// Inject DB ke -> repository
	merchantRepo := repository.NewMerchantRepository(db)
	historyRepo := repository.NewHistoryRepository(db)
	usersRepo := repository.NewUsersRepository(db)

	// Inject REPO ke -> useCase
	merchantUC := usecase.NewMerchantUseCase(merchantRepo)
	historyUC := usecase.NewHistoryUsecase(historyRepo)
	usersUC := usecase.NewUsersUseCase(usersRepo)
	jwtService := service.NewJwtService(cfg.TokenConfig)
	authUc := usecase.NewAuthUseCase(usersUC, jwtService)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		authUsc:    authUc,
		merchantUC: merchantUC,
		historyUC:   historyUC,
		usersUC: usersUC,
		engine:     engine,
		jwtService: jwtService,
		host:       host,
	}
}
