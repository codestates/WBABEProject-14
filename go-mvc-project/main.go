package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	conf "wba/go-mvc-procjet/config"
	ctl "wba/go-mvc-procjet/controller"
	"wba/go-mvc-procjet/logger"
	"wba/go-mvc-procjet/model"
	rt "wba/go-mvc-procjet/route"
	"wba/go-mvc-procjet/services"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {

	var configFlag = flag.String("config", "./config/config.toml", "toml file to use for configuration")
	flag.Parse()
	cf := conf.NewConfig(*configFlag)

	/* 로그 초기화 */
	if err := logger.InitLogger(cf); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	logger.Debug("ready server....")
	/* Model 선언 */
	if colMenu, colOrder, colReview, err := model.NewModel(cf.DB["account"]["host"].(string)); err != nil {
		panic(err)
		/* Service 선언 */
	} else if OrdererService, err := services.NewOrdererService(colOrder, colReview, context.TODO()); err != nil {
		panic(err)
	} else if TakerService, err := services.NewTakerService(colMenu, context.TODO()); err != nil {
		panic(err)
		/* Controller 선언*/
	} else if oc, err := ctl.NewOrdererController(OrdererService); err != nil {
		panic(err)
	} else if rc, err := ctl.NewTakerController(TakerService); err != nil {
		panic(err)
		/* Router 선언 */
	} else if rt, err := rt.NewRouter(&oc, &rc); err != nil {
		panic(fmt.Errorf("router.NewRouter > %v", err))

	} else {
		/* Server 설정 */
		mapi := &http.Server{
			Addr:           cf.Server.Port,
			Handler:        rt.Idx(),
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		g.Go(func() error {
			return mapi.ListenAndServe()
		})

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		logger.Warn("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			logger.Error("Server Shutdown:", err)
		}

		select {
		case <-ctx.Done():
			logger.Info("timeout of 1 seconds.")
		}

		logger.Info("Server exiting")
	}

	if err := g.Wait(); err != nil {
		logger.Error(err)
	}

}
