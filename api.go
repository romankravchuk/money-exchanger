package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/romankravchuk/money-converter/types"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

type JSONAPIServer struct {
	listenAddr string
	svc        CurrencyConverter
}

func NewJSONAPIServer(listenAddr string, svc CurrencyConverter) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleConvertCurrency))
	http.ListenAndServe(s.listenAddr, nil)
}

func makeHTTPHandlerFunc(apiFn APIFunc) http.HandlerFunc {
	ctx := context.Background()

	return func(w http.ResponseWriter, r *http.Request) {
		requestID, err := uuid.NewRandom()
		ctx = context.WithValue(ctx, "requestID", requestID.String())

		if err != nil {
			writeJson(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
			return
		}

		if err = apiFn(ctx, w, r); err != nil {
			writeJson(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func (s *JSONAPIServer) handleConvertCurrency(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	amount := r.URL.Query().Get("amount")

	parsedAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return err
	}

	result, err := s.svc.Convert(ctx, from, to, parsedAmount)
	if err != nil {
		return err
	}

	convertResp := types.ConvertResponse{
		Query: types.ConvertQuery{
			From:   from,
			To:     to,
			Amount: parsedAmount,
		},
		Result: result,
	}

	return writeJson(w, http.StatusOK, &convertResp)
}

func writeJson(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
