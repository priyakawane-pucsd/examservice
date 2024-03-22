package restapp

import "context"

type RestApp struct {
}

func NewRestApp(ctx context.Context, configpath string) *RestApp {

	return &RestApp{}
}

func (app *RestApp) Start(ctx context.Context) error {
	return nil
}

func (app *RestApp) Shutdown(ctx context.Context) error {
	return nil
}
