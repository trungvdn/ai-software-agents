package oauth

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/auth"
)

type CallbackServer struct {
	resultCh chan *auth.AuthorizationResult
	errCh    chan error
	lis      net.Listener
}

func NewCallbackServer() *CallbackServer {
	return &CallbackServer{
		resultCh: make(chan *auth.AuthorizationResult, 1),
		errCh:    make(chan error, 1),
	}
}

func (s *CallbackServer) Start(addr string) error {
	log.Printf("Stgartt==")
	mux := http.NewServeMux()

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
	log.Printf("Listen==")
	go func() {
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			log.Printf("Error net.Listen( : %s", err.Error())
			s.errCh <- err
		}
		if err = http.Serve(ln, mux); err != nil {
			log.Printf("Error  http.Serve(: %s", err.Error())
			s.errCh <- err
		}
		s.lis = ln
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
	return s.lis.Close()
}
