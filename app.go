package shift

type App struct {
	Server     Server
	rootDomain *Domain
	Router     *Router
	Name       string
	Config     *AppConfig
}

func New(name string, config *AppConfig) *App {
	if config == nil {
		config = NewAppConfig()
	}

	rootDomain := newDomain("/", nil)
	config.Router.apply(rootDomain.Router)

	return &App{
		Name:       name,
		Server:     &defaultServer{},
		rootDomain: rootDomain,
		Router:     rootDomain.Router,
		Config:     config,
	}
}

func (app *App) Domain(path string, constructor func(d *Domain)) {
	constructor(newDomain(path, app.rootDomain))
}

func (app *App) Run(addr string) error {
	app.Server.Initialize(addr, app.Router)

	return app.Server.ListenAndServe()
}

func (app *App) RunTLS(addr, certFile, keyFile string) error {
	app.Server.Initialize(addr, app.Router)
	return app.Server.ListenAndServeTLS(certFile, keyFile)
}
