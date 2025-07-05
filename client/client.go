package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"math"
	"os"
	"time"

	"github.com/golang/geo/r2"
	"github.com/jakub-galecki/cs_stats/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getMapImg(name string) (image.Image, error) {
	f, err := os.Open("./static/de_ancient.png")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}

func drawDot(img *image.RGBA, x, y float64) {
	red := color.RGBA{255, 0, 0, 255}
	radius := 10
	cx := int(math.Round(x))
	cy := int(math.Round(y))

	for dx := -radius; dx <= radius; dx++ {
		for dy := -radius; dy <= radius; dy++ {
			if dx*dx+dy*dy <= radius*radius {
				px := cx + dx
				py := cy + dy
				if px >= 0 && px < img.Bounds().Dx() && py >= 0 && py < img.Bounds().Dy() {
					img.Set(px, py, red)
				}
			}
		}
	}
}

func createFrame(mapImg image.Image, path string, i int, pos r2.Point) error {
	fr := image.NewRGBA(mapImg.Bounds())
	draw.Draw(fr, fr.Bounds(), mapImg, image.Point{}, draw.Src)
	drawDot(fr, pos.X, pos.Y)

	out, err := os.Create(fmt.Sprintf("%s/frame_%03d.png", path, i))
	if err != nil {
		return err
	}
	defer out.Close()
	err = png.Encode(out, fr)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	conn, err := grpc.NewClient("localhost:9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	f, err := os.OpenFile("positions.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	var (
		w      = bufio.NewWriter(f)
		client = pb.NewStatsClient(conn)

		req = &pb.ProcessDemoReq{
			Path:   "/Users/batman/Downloads/1-0a4e7ada-05a4-4004-91f5-74462dc9e9a0-1-1.dem",
			Player: "jac30b",
		}
	)

	stats, err := client.ProcessDemo(context.Background(), req)
	if err != nil {
		panic(err)
	}
	w.WriteString(stats.String() + "\n")

	log.Printf("processing map: %s\n", stats.Map)

	res, err := client.ReplayDemo(context.Background(), req)
	if err != nil {
		panic(err)
	}

	var (
		start = time.Now()
	)

	m, err := getMapImg("")
	if err != nil {
		panic(err)
	}

	i := 0

	for {
		pos, err := res.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				res.CloseSend()
				break
			}
			panic(err)
		}
		if posRepr := pos.String(); posRepr != "" {
			_, err = w.WriteString(posRepr + "\n")
			if err != nil {
				panic(err)
			}

			err = createFrame(m, "output_frames", i, r2.Point{X: pos.X, Y: pos.Y})
			if err != nil {
				panic(err)
			}

			i++
		}
	}

	if err := w.Flush(); err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	log.Printf("finished in %v", elapsed)
}
