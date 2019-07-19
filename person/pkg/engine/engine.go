package engine

import (
	"context"
	"fmt"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"sdkeji/person/pkg/app"
	pb "sdkeji/person/pkg/proto"
	"sdkeji/person/pkg/server"
	"sdkeji/person/pkg/ui/data/swagger"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	std *Engine
)

type Engine struct {
	apiServer *http.Server
	server    *http.Server
	close     chan struct{}
	wg        sync.WaitGroup
}

func Get() *Engine {
	return std
}

func NewStdInstance() *Engine {
	app.Init()
	std = new(Engine)
	std.close = make(chan struct{})
	std.server = &http.Server{Addr: app.Conf.HTTPAddr}
	return std
}

func (e *Engine) Run() {
	go e.registerSignal()

	e.wg.Add(3)
	go e.serveHTTP()
	go e.serveRPC()
	go e.serverApi()
	e.wg.Wait()
}

func (e *Engine) serveRPC() {
	defer e.wg.Done()
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		app.Logger.Info("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterPersonsServer(s, &server.PersonsServer{})
	s.Serve(lis)
}

func (e *Engine) serverApi() {
	defer e.wg.Done()
	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", serveSwaggerFile)
	serveSwaggerUI(mux)
	e.apiServer = &http.Server{
		Addr:    ":5000",
		Handler: mux,
	}
	app.Logger.Info("5000 engine successfully.")
	err := e.apiServer.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			app.Logger.Error("an error was returned while listen and serve engine.", "error", err)
			return
		}
	}
}

func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		app.Logger.Info("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join("pkg/proto", p)
	fmt.Println(p)

	app.Logger.Info("Serving swagger-file: %s", p)

	http.ServeFile(w, r, p)
}

func serveSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/api/"

	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func (e *Engine) serveHTTP() {
	defer e.wg.Done()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterPersonsHandlerFromEndpoint(ctx, mux, "localhost:50052", opts)
	if err != nil {
		return
	}
	e.server.Handler = mux
	err = e.server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			app.Logger.Error("an error was returned while listen and serve engine.", "error", err)
			return
		}
	}
	app.Logger.Info("engine shutdown successfully.")
}

func (e *Engine) shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if err := e.apiServer.Shutdown(ctx); err != nil {
		return err
	}
	return e.server.Shutdown(ctx)
}

func (e *Engine) registerSignal() {
	app.Logger.Info("register signal handler.")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)
	select {
	case sig := <-ch:
		signal.Ignore(syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)
		app.Logger.Info("received signal, try to shutdown engine.", "signal", sig.String())
		close(ch)
		close(e.close)
		err := e.shutdown()
		if err != nil {
			app.Logger.Error("fail to shutdown engine.", "error", err)
		}
	}
}
