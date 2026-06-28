package oauth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/auth"
)

type CallbackServer struct {
	resultCh chan *auth.AuthorizationResult
	errCh    chan error
	server   *http.Server
}

func NewCallbackServer() *CallbackServer {
	return &CallbackServer{
		resultCh: make(chan *auth.AuthorizationResult, 1),
		errCh:    make(chan error, 1),
	}
}

func (s *CallbackServer) Start(addr string) error {
	mux := http.NewServeMux()
	s.server = &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		code := query.Get("code")
		state := query.Get("state")
		if code == "" {
			http.Error(w, "missing code", http.StatusBadRequest)
			s.errCh <- fmt.Errorf("authorization callback missing code")
			return
		}

		fmt.Fprintln(w, "Authorization complete. You may close this window.")
		s.resultCh <- &auth.AuthorizationResult{Code: code, State: state}
	})

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.errCh <- err
		}
	}()
	return nil
}

func (s *CallbackServer) Wait(
	ctx context.Context,
) (*auth.AuthorizationResult, error) {
	select {
	case res := <-s.resultCh:
		return res, nil
	case err := <-s.errCh:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *CallbackServer) Stop(
	ctx context.Context,
) error {
	if s == nil || s.server == nil {
		return nil
	}
	return s.server.Shutdown(ctx)
}
