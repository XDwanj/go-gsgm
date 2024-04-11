package img_service

import (
	"bufio"
	"image"
	_ "image/jpeg"
	"image/png"
	"math/big"
	"os"

	"golang.org/x/image/draw"
)

// TODO: 没有判断图片格式错误问题
func ZoomLutrisPicture(srcPath, destPath string, newWidth, newHeight int) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer func() { _ = srcFile.Close() }()

	drawBoard, err := preZoom(srcPath, newWidth, newHeight)
	if err != nil {
		return err
	}
	destImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	draw.NearestNeighbor.Scale(
		destImg,
		image.Rect(0, 0, newWidth, newHeight),
		drawBoard,
		image.Rect(0, 0, drawBoard.Bounds().Dx(), drawBoard.Bounds().Dy()),
		draw.Over,
		nil,
	)

	destFile, err := os.OpenFile(destPath, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer func() { _ = destFile.Close() }()

	writer := bufio.NewWriter(destFile)
	if err = png.Encode(writer, destImg); err != nil {
		return err
	}
	if err = writer.Flush(); err != nil {
		return err
	}

	return nil
}

func preZoom(srcPath string, newWidth, newHeight int) (*image.RGBA, error) {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = srcFile.Close() }()

	srcImg, _, err := image.Decode(srcFile)
	if err != nil {
		return nil, err
	}

	var (
		srcWidth   = srcImg.Bounds().Dx()
		srcHeight  = srcImg.Bounds().Dy()
		destWidth  int
		destHeight int
	)

	var (
		destFactor = big.NewRat(int64(newWidth), int64(newHeight))
		srcFactor  = big.NewRat(int64(srcWidth), int64(srcHeight))
	)

	if srcFactor.Cmp(destFactor) < 0 {
		destWidth = srcHeight * int(destFactor.Num().Int64()) / int(destFactor.Denom().Int64())
		destHeight = srcHeight
	} else if srcFactor.Cmp(destFactor) > 0 {
		destWidth = srcWidth
		destHeight = srcWidth * int(destFactor.Denom().Int64()) / int(destFactor.Num().Int64())
	} else {
		destWidth = srcWidth
		destHeight = srcHeight
	}

	destImg := image.NewRGBA(image.Rect(0, 0, destWidth, destHeight))

	var x, y int
	if srcFactor.Cmp(destFactor) < 0 {
		x, y = (destWidth/2)-(srcWidth/2), 0
	} else if srcFactor.Cmp(destFactor) > 0 {
		x, y = 0, (destHeight/2)-(srcHeight/2)
	} else {
		x, y = 0, 0
	}
	draw.Draw(
		destImg,
		image.Rect(x, y, x+srcWidth, y+srcHeight),
		srcImg,
		image.Pt(0, 0),
		draw.Over,
	)

	return destImg, nil
}
