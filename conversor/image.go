package conversor

import (
	"archive/zip"
	"context"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type SaveData struct {
	Path string
	Img  image.Image
}

// Todo: webp, ico?
var encoders = map[string]func(io.Writer, image.Image) error{
	"png": png.Encode,
	"bmp": bmp.Encode,
	"tiff": func(w io.Writer, i image.Image) error {
		return tiff.Encode(w, i, &tiff.Options{
			Compression: tiff.Deflate,
			Predictor:   true,
		})
	},
	"jpg": func(w io.Writer, i image.Image) error {
		return jpeg.Encode(w, i, &jpeg.Options{
			Quality: 100,
		})
	},
}

// Todo: svg?, ico?
func readImage(ctx context.Context) ([]string, error) {
	return runtime.OpenMultipleFilesDialog(ctx, runtime.OpenDialogOptions{
		Title:                      "Selecione a imagem a ser convertida",
		DefaultDirectory:           "",
		DefaultFilename:            "",
		ShowHiddenFiles:            false,
		CanCreateDirectories:       true,
		ResolvesAliases:            true,
		TreatPackagesAsDirectories: true,
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Jpeg",
				Pattern:     "*.jpg;*.jpeg",
			},
			{
				DisplayName: "Png",
				Pattern:     "*.png",
			},
			{
				DisplayName: "Bitmap",
				Pattern:     "*.bmp",
			},
			{
				DisplayName: "WebP",
				Pattern:     "*.webp",
			},
			{
				DisplayName: "Tiff",
				Pattern:     "*.tif;*.tiff",
			},
			{
				DisplayName: "All supported files",
				Pattern:     "*.jpg;*.jpeg;*.png;*.tif;*.tiff;*.webp;*.bmp",
			},
		},
	})
}

func getImageName(source string) string {
	var defaultName = filepath.Base(source)
	return defaultName[:len(defaultName)-len(filepath.Ext(defaultName))]
}

func openImage(source string) (image.Image, error) {
	file, err := os.Open(source)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	image, _, err := image.Decode(file)
	return image, err
}

func saveImage(saveData *SaveData, destType string) error {
	var file, err = os.Create(saveData.Path)
	if err != nil {
		return err
	}
	defer file.Close()
	return encoders[destType](file, saveData.Img)
}

func saveZip(dataArray []*SaveData, savePath, destType string) error {
	var zipFile, err = os.Create(savePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	var zipWriter = zip.NewWriter(zipFile)
	defer zipWriter.Close()
	for _, saveData := range dataArray {
		file, err := zipWriter.Create(saveData.Path)
		if err != nil {
			return err
		}
		if err := encoders[destType](file, saveData.Img); err != nil {
			return err
		}
	}
	return nil
}

func getSavePath(ctx context.Context, filename, destType string) (string, error) {
	return runtime.SaveFileDialog(ctx, runtime.SaveDialogOptions{
		Title:                      "Escolha onde salvar a imagem",
		CanCreateDirectories:       true,
		TreatPackagesAsDirectories: true,
		DefaultFilename:            filename,
		Filters: []runtime.FileFilter{
			{
				DisplayName: strings.ToTitle(destType),
				Pattern:     "*." + destType,
			},
		},
	})
}
