package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	cu "github.com/wcademy/simplegoprojectlayout/internal/comandlineutils"
	"github.com/wcademy/simplegoprojectlayout/sumlib"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("can't init logger:", err)
	}

	ints, err := cu.GetArgs(os.Args)
	if err != nil {
		logger.Error("can't get args", zap.Error(err))
		return
	}

	result := sumlib.Sum(ints)

	fmt.Println(result)
}
