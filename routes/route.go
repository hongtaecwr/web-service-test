package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"web-service-test/handler/auth_handler"
	"web-service-test/handler/order_handler"
	"web-service-test/handler/product_handler"
	"web-service-test/handler/user_handler"
	"web-service-test/middlewares"
	"web-service-test/repositories/auth_repository"
	"web-service-test/repositories/order_repository"
	"web-service-test/repositories/product_repository"
	"web-service-test/repositories/users_repository"
	"web-service-test/usecases/auth_usecase"
	"web-service-test/usecases/order_usecase"
	"web-service-test/usecases/product_usecase"
	"web-service-test/usecases/user_usecase"
)

func NewRouter(e *echo.Echo) {
	setLogin := middlewares.SetLogin()
	middlewares.SetCORS(e)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "AccessTime => ${time_rfc3339_nano}\n" +
			"Host => ${host}, RemoteIP => ${remote_ip},\n" +
			"Method => ${method}, URI => ${uri}, Status => ${status},\n" +
			"Error => ${error},\n" +
			"UserAgent => ${user_agent}\n" +
			"--------------\n",
		Output: e.Logger.Output(),
	}))

	route := e.Group("/api/v1")

	authRepo := auth_repository.NewAuthRepository()
	authUseCase := auth_usecase.NewAuthUseCase(authRepo)
	authHandler := auth_handler.NewAuthHandler(authUseCase)

	userRepo := users_repository.NewUserRepository()
	userUseCase := user_usecase.NewUserUseCase(userRepo)
	userHandler := user_handler.NewUserHandler(userUseCase)

	orderRepo := order_repository.NewOrderRepository()
	orderUseCase := order_usecase.NewOrderUseCase(orderRepo)
	orderHandler := order_handler.NewOrderHandler(orderUseCase)

	productRepo := product_repository.NewProductRepository()
	productUseCase := product_usecase.NewProductUseCase(productRepo)
	productHandler := product_handler.NewProductHandler(productUseCase)

	route.POST("/register", authHandler.Register)
	route.POST("/login", authHandler.Login)

	usersRoute := route.Group("/user", setLogin, middlewares.IsUser)
	usersRoute.GET("/detail", userHandler.GetUserDetail)
	usersRoute.GET("/order-histories", userHandler.GetUserOrderHistories)

	productRoute := route.Group("/product")
	productRoute.GET("/list", productHandler.GetAllProducts)
	productRoute.GET("/detail", productHandler.GetProductDetail)

	orderRoute := route.Group("/order", setLogin, middlewares.IsUser)
	orderRoute.POST("/create", orderHandler.CreateOrder)
	orderRoute.POST("/cancel", orderHandler.CancelOrder)
	orderRoute.GET("/detail", orderHandler.OrderDetail)

}
