package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gen2brain/beeep"
)

var defaultOpeningMoviePath = "Library/Application Support/Steam/steamapps/common/Sid Meier's Civilization V/Civilization V.app/Contents/Assets/Civ5_Opening_Movie_en_US.mov"

func main() {
	err := _main()
	if err != nil {
		beeep.Notify("There was an error", err.Error(), "")
		os.Exit(1)
	}
	beeep.Notify("It worked!", "Randy loves you ^__^", "")
}

func _main() error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	movieSrcPath := filepath.Join(homePath, defaultOpeningMoviePath)
	movieDestPath := fmt.Sprintf("%s.bak", movieSrcPath)

	_, err = os.Stat(movieDestPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	} else if err == nil {
		return nil
	}

	return os.Rename(movieSrcPath, movieDestPath)
}
