package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"time"

	"github.com/LambdaTest/photon/config"
	"github.com/LambdaTest/photon/pkg/global"
	"github.com/LambdaTest/photon/pkg/opentelemetry"

	// use mysql connector
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

// RootCommand will setup and return the root command
func RootCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "photon",
		Long: `photon is a dispatcher component that is responsible for receiving webhooks from git providers 
		and queueing actions to be taken as a response to events received.`,
		Version:      global.BinaryVersion,
		RunE:         run,
		SilenceUsage: true,
	}

	// define flags used for this command
	AttachCLIFlags(&rootCmd)

	return &rootCmd
}

func run(cmd *cobra.Command, args []string) error {
	// create a context that we can cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// timeout in seconds
	const gracefulTimeout = 5000 * time.Millisecond

	// a WaitGroup for the goroutines to tell us they've stopped
	wg := sync.WaitGroup{}

	cfg, err := config.Load(cmd)
	if err != nil {
		log.Printf("Failed to load config: %v", err)
		return err
	}
	// patch logconfig file location with root level log file location
	if cfg.LogFile != "" {
		cfg.LogConfig.FileLocation = filepath.Join(cfg.LogFile, "photon.log")
	}

	app, err := InitializeApp(cfg)
	if err != nil {
		log.Printf("Could not instantiate application %v\n", err)
		return err
	}

	// initialize tracer
	if cfg.Tracing.OtelEndpoint != "" {
		tracerCleanup := opentelemetry.InitTracer(ctx, cfg, app.logger)
		defer func() {
			if tracerErr := tracerCleanup(context.Background()); tracerErr != nil {
				app.logger.Errorf("Failed to cleanup the tracer %v", tracerErr)
			}
		}()
	}

	errChan := make(chan error, 1)
	wg.Add(1)
	// start http server
	go func() {
		defer wg.Done()
		if err := app.server.ListenAndServe(ctx); err != nil {
			app.logger.Errorf("failed to start http server, error %v", err)
			errChan <- err
		}
	}()
	// listen for C-c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// create channel to mark status of waitgroup
	// this is required to brutally kill application in case of
	// timeout
	done := make(chan struct{})

	// asynchronously wait for all the go routines
	go func() {
		// and wait for all go routines
		wg.Wait()
		app.logger.Debugf("main: all goroutines have finished.")
		close(done)
	}()

	// wait for signal channel
	select {
	case <-c:
		app.logger.Debugf("main: received C-c - attempting graceful shutdown ....")
		// tell the goroutines to stop
		app.logger.Debugf("main: telling goroutines to stop")
		cancel()
		select {
		case <-done:
			app.logger.Debugf("Go routines exited within timeout")
		case <-time.After(gracefulTimeout):
			app.logger.Errorf("Graceful timeout exceeded. Brutally killing the application")
		}
	case err := <-errChan:
		cancel()
		return err
	case <-done:
	}
	return nil
}
