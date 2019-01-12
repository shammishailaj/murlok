package murlok

import (
	"encoding/json"
	"net/http"
)

type rpcRequest struct {
	Version string                   `json:"jsonrpc"`
	Method  string                   `json:"method"`
	Params  []map[string]interface{} `json:"params,omitempty"`
	ID      string                   `json:"id,omitempty"`
}

type rpcResponse struct {
	Version string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Err     rpcErr      `json:"error,omitempty"`
	ID      string      `json:"id"`
}

type rpcErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func rpc(w http.ResponseWriter, r *http.Request) {
	res := rpcResponse{
		Version: "2.0",
	}

	defer func() {
		if len(res.ID) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		enc.Encode(res)
	}()

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	req := rpcRequest{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		res.Err = rpcErr{
			Code:    -32700,
			Message: err.Error(),
		}

		return
	}

	if req.Version != "2.0" {
		res.Err = rpcErr{
			Code:    -32700,
			Message: "json-rpc version not supported",
		}

		return
	}

	res.ID = req.ID

	if err := backend.Call(req.Method, &res.Result, req.Params); err != nil {
		res.Err = rpcErr{
			Code:    -32603,
			Message: err.Error(),
		}
	}
}
