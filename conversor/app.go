package conversor

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func SetAppContext(app *App, ctx context.Context) {
	app.ctx = ctx
}

func (a *App) GetFile() ([]string, error) {
	return readImage(a.ctx)
}

func (a *App) ConvertTo(source, destType string) error {
	image, err := openImage(source)
	if err != nil {
		return err
	}
	savePath, err := getSavePath(a.ctx, getImageName(source), destType)
	if err != nil || savePath == "" {
		return err
	}
	return saveImage(
		&SaveData{Path: savePath, Img: image},
		destType,
	)
}

func (a *App) ConvertMultiple(source []string, destType string) error {
	var images = make([]*SaveData, 0)
	for _, path := range source {
		image, err := openImage(path)
		if err != nil {
			return err
		}
		images = append(images, &SaveData{
			Path: fmt.Sprint(getImageName(path), ".", destType),
			Img:  image,
		})
	}
	savePath, err := getSavePath(a.ctx, "converted-files", "zip")
	if err != nil || savePath == "" {
		return err
	}
	return saveZip(images, savePath, destType)
}

func (a *App) PopMessage(title, message string) {
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   title,
		Message: message,
	})
}

func (a *App) PopError(title, message string) {
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   title,
		Message: message,
	})
}
