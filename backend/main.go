package main

import (
	"log"
	"net/http"

	"calculator/backend/server" // 基于模块的绝对路径

	"calculator/backend/gen/calculator/v1/calculatorv1connect"

	connectgrpchealth "connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	// 初始化计算器服务实例
	calculator := &server.CalculatorServer{}

	// 创建HTTP路由
	mux := http.NewServeMux()

	// ========== 注册Connect服务处理器 ==========
	// 关键：必须先注册业务服务处理器
	calculatorPath, calculatorHandler := calculatorv1connect.NewCalculatorServiceHandler(calculator)
	mux.Handle(calculatorPath, calculatorHandler)

	// ========== 注册健康检查服务 ==========
	healthChecker := connectgrpchealth.NewStaticChecker(
		calculatorv1connect.CalculatorServiceName,
	)
	mux.Handle(connectgrpchealth.NewHandler(healthChecker))

	// ========== 注册反射服务 ==========
	reflector := grpcreflect.NewStaticReflector(
		calculatorv1connect.CalculatorServiceName,
	)
	reflPath, reflHandler := grpcreflect.NewHandlerV1(reflector)
	mux.Handle(reflPath, reflHandler)

	// ========== CORS中间件配置 ==========
	corsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 允许所有来源的跨域请求（生产环境应限制）
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Connect-Protocol-Version, Connect-Timeout")

		// 处理预检请求
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// 传递请求给实际处理器
		mux.ServeHTTP(w, r)
	})

	// ========== 配置HTTP服务器 ==========
	server := &http.Server{
		Addr: ":8080",
		// 启用h2c以支持HTTP/2 without TLS
		Handler: h2c.NewHandler(
			corsHandler,
			&http2.Server{
				MaxConcurrentStreams: 250,
			},
		),
	}

	// ========== 启动服务 ==========
	log.Println("Starting ConnectRPC server on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
