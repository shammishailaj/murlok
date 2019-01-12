package murlok

// type app struct {
// 	server       *http.Server
// 	host         string
// 	defaultURL   string
// 	allowedHosts map[string]struct{}
// 	defaultView  View
// 	backend      Backend
// }

// func (a *app) WithCustomServer(s *http.Server) Application {
// 	a.server = s
// 	return a
// }

// func (a *app) WithPackageConfig(c PackageConfig) Application {
// 	a.pkgconfs = append(a.pkgconfs, c)
// 	return a
// }

// func (a *app) WithCustomView(v View) Application {
// 	a.defaultView = v
// 	return a
// }

// func (a *app) Run(url string, allowedHosts ...string) {
// 	a.defaultURL = url

// 	a.allowedHosts = make(map[string]struct{}, len(allowedHosts)+1)
// 	for _, host := range allowedHosts {
// 		a.allowedHosts[host] = struct{}{}
// 	}

// 	if err := a.runServer(); err != nil {
// 		panic(err)
// 	}

// 	if err := a.run(); err != nil {
// 		panic(err)
// 	}
// }

// func (a *app) runForBuild(path string, c PackageConfig) error {
// 	f, err := os.Create(path)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()

// 	enc := json.NewEncoder(f)
// 	return enc.Encode(c)
// }

// func (a *app) runServer() error {
// 	if a.server == nil {
// 		a.server = &http.Server{Addr: ":0"}
// 	}

// 	http.Handle("/", a)

// 	go func() {
// 		if err := a.server.ListenAndServe(); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	return nil
// }

// func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	res := rpcResponse{
// 		Version: "2.0",
// 	}

// 	defer func() {
// 		if len(res.ID) == 0 {
// 			w.WriteHeader(http.StatusNoContent)
// 			return
// 		}

// 		w.WriteHeader(http.StatusOK)
// 		enc := json.NewEncoder(w)
// 		enc.Encode(res)
// 	}()

// 	if r.Header.Get("Content-Type") != "application/json" {
// 		w.WriteHeader(http.StatusUnsupportedMediaType)
// 		return
// 	}

// 	req := rpcRequest{}
// 	dec := json.NewDecoder(r.Body)
// 	if err := dec.Decode(&req); err != nil {
// 		res.Err = rpcErr{
// 			Code:    -32700,
// 			Message: err.Error(),
// 		}

// 		return
// 	}

// 	if req.Version != "2.0" {
// 		res.Err = rpcErr{
// 			Code:    -32700,
// 			Message: "json-rpc version not supported",
// 		}

// 		return
// 	}

// 	res.ID = req.ID

// 	if err := a.backend.Call(req.Method, &res.Result, req.Params); err != nil {
// 		res.Err = rpcErr{
// 			Code:    -32603,
// 			Message: err.Error(),
// 		}
// 	}
// }

// type rpcRequest struct {
// 	Version string                   `json:"jsonrpc"`
// 	Method  string                   `json:"method"`
// 	Params  []map[string]interface{} `json:"params,omitempty"`
// 	ID      string                   `json:"id,omitempty"`
// }

// type rpcResponse struct {
// 	Version string      `json:"jsonrpc"`
// 	Result  interface{} `json:"result,omitempty"`
// 	Err     rpcErr      `json:"error,omitempty"`
// 	ID      string      `json:"id"`
// }

// type rpcErr struct {
// 	Code    int    `json:"code"`
// 	Message string `json:"message"`
// }
