package f64

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"math"
	"testing"

	"github.com/colinrgodsey/cartesius/f64/filters"
)

var testGridFilters = [...]filters.GridFilter{
	filters.Box, filters.Linear, filters.Hermite, filters.MitchellNetravali,
	filters.CatmullRom, filters.BSpline, filters.Gaussian, filters.Lanczos,
	filters.Hann, filters.Hamming, filters.Blackman, filters.Bartlett,
	filters.Welch, filters.Cosine,
}

func TestGrids(t *testing.T) {
	const (
		maxBadRMS = 0.065
		maxRMS    = 0.060
	)
	for fidx, filter := range testGridFilters {
		max := maxRMS
		if fidx == 0 || fidx == 5 {
			max = maxBadRMS
		}
		str := testInterp(func(samples []Vec3) Function2D {
			x, _ := Grid2D(samples, filter)
			return x
		}, fidx, 1, max)
		if str != "" {
			t.Fatalf(str)
		}
	}
}

func TestGridLinear(t *testing.T) {
	samples := []Vec3{
		{0, 0, 0},
		{0, 1, 0.5},
		{1, 0, 0.5},
		{1, 1, 1},
	}
	interp, _ := Grid2D(samples, filters.Linear)
	expect := func(x, y, z float64) {
		res, err := interp(Vec2{x, y})
		if err != nil || res != z {
			t.Fatalf("Failed. Expecting %v, got %v. error: %v", z, res, err)
		}
	}

	for _, s := range samples {
		expect(s[0], s[1], s[2])
	}
	expect(0.5, 0.5, 0.5)
	expect(0.75, 0.75, 0.75)
}

func TestMicrosphere(t *testing.T) {
	const (
		maxRMS = 0.058
	)
	str := testInterp(MicroSphere2D, "microsphere", 4, maxRMS)
	if str != "" {
		t.Fatalf(str)
	}
}

func testInterp(interpGen func([]Vec3) Function2D, filter interface{}, stride int, maxRMS float64) string {
	const (
		sampleSize = 100.0
		imgFac     = 100.0 / 30.0
	)

	var samples []Vec3
	sVals := getTestGreyImage(testImage30)
	vVals := getTestGreyImage(testImage100)
	for y, ys := range sVals {
		for x, v := range ys {
			s := Vec3{(float64(x) + 0.5) * imgFac, (float64(y) + 0.5) * imgFac, v}
			samples = append(samples, s)
		}
	}

	interp := interpGen(samples)

	sampleChan := func(stride int) <-chan Vec2 {
		positions := make(chan Vec2, 32)
		go func() {
			for x := 0; x < sampleSize; x += stride {
				for y := 0; y < sampleSize; y += stride {
					positions <- Vec2{float64(x) + 0.5, float64(y) + 0.5}
				}
			}
			close(positions)
		}()
		return positions
	}

	var sum, num float64
	for sample := range interp.Multi(sampleChan(stride)) {
		valid := vVals[int(sample[1])][int(sample[0])]
		e := valid - sample[2]
		sum += e * e
		num++
	}
	rms := math.Sqrt(sum / num)
	fmt.Printf("Filter: %v, RMS: %v\n", filter, rms)
	if rms >= maxRMS {
		return fmt.Sprintf("RMS value %v exceeded max %v for filter %v", rms, maxRMS, filter)
	}
	return ""
}

/* base64 jpg */
const testImage100 = "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"
const testImage50 = "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"
const testImage30 = "/9j/4AAQSkZJRgABAQAASABIAAD/4QCMRXhpZgAATU0AKgAAAAgABQESAAMAAAABAAEAAAEaAAUAAAABAAAASgEbAAUAAAABAAAAUgEoAAMAAAABAAIAAIdpAAQAAAABAAAAWgAAAAAAAABIAAAAAQAAAEgAAAABAAOgAQADAAAAAQABAACgAgAEAAAAAQAAAB6gAwAEAAAAAQAAAB4AAAAA/+0AOFBob3Rvc2hvcCAzLjAAOEJJTQQEAAAAAAAAOEJJTQQlAAAAAAAQ1B2M2Y8AsgTpgAmY7PhCfv/CABEIAB4AHgMBIgACEQEDEQH/xAAfAAABBQEBAQEBAQAAAAAAAAADAgQBBQAGBwgJCgv/xADDEAABAwMCBAMEBgQHBgQIBnMBAgADEQQSIQUxEyIQBkFRMhRhcSMHgSCRQhWhUjOxJGIwFsFy0UOSNIII4VNAJWMXNfCTc6JQRLKD8SZUNmSUdMJg0oSjGHDiJ0U3ZbNVdaSVw4Xy00Z2gONHVma0CQoZGigpKjg5OkhJSldYWVpnaGlqd3h5eoaHiImKkJaXmJmaoKWmp6ipqrC1tre4ubrAxMXGx8jJytDU1dbX2Nna4OTl5ufo6erz9PX29/j5+v/EAB8BAAMBAQEBAQEBAQEAAAAAAAECAAMEBQYHCAkKC//EAMMRAAICAQMDAwIDBQIFAgQEhwEAAhEDEBIhBCAxQRMFMCIyURRABjMjYUIVcVI0gVAkkaFDsRYHYjVT8NElYMFE4XLxF4JjNnAmRVSSJ6LSCAkKGBkaKCkqNzg5OkZHSElKVVZXWFlaZGVmZ2hpanN0dXZ3eHl6gIOEhYaHiImKkJOUlZaXmJmaoKOkpaanqKmqsLKztLW2t7i5usDCw8TFxsfIycrQ09TV1tfY2drg4uPk5ebn6Onq8vP09fb3+Pn6/9sAQwACAwMDBAMEBQUEBgYGBgYICAcHCAgNCQoJCgkNEwwODAwODBMRFBEPERQRHhgVFRgeIx0cHSMqJSUqNTI1RUVc/9sAQwECAwMDBAMEBQUEBgYGBgYICAcHCAgNCQoJCgkNEwwODAwODBMRFBEPERQRHhgVFRgeIx0cHSMqJSUqNTI1RUVc/9oADAMBAAIRAxEAAAHoLeicV0/NMPC67Z38xdZXtfntNFf/2gAIAQEAAQUC/Rl0oGxvTCLGdTO2SouI9xty1yKLmnijRdX1xzRPbA82EuQx1kukZCTSO4kKprhOXvKSP//aAAgBAxEBPwHs/9oACAECEQE/Aez/2gAIAQEABj8CpzAQPsYSJUpIdJZhT0DUmKcezXq00fRFWvqp9U6UD+S689az5eTCl1OlBR+2T6h9MlKeTClCpo9BH9pq66vyfCr0Q//EADMQAQADAAICAgICAwEBAAACCwERACExQVFhcYGRobHB8NEQ4fEgMEBQYHCAkKCwwNDg/9oACAEBAAE/IUSeMmaXCedz92NUXIS/TQrMRmpMf1QAicFo/VTJVwf7FfPiYiF+XqtSX4U5oSdE/wBE9UaDFyKDnzXouNvH3xWct3JSaOHzSgQzMn91AVa0eJP6qbBHiC//2gAMAwEAAhEDEQAAEMMOKP/EADMRAQEBAAMAAQIFBQEBAAEBCQEAESExEEFRYSBx8JGBobHRweHxMEBQYHCAkKCwwNDg/9oACAEDEQE/EPwf/9oACAECEQE/EPwf/9oACAEBAAE/ENJwQJxM6PxZBvzLAM9M/FCcger+7CsYEASFyJFHVYcrlv2of5oQqCOQ/b+W/udEUOqO4JaccxRDLhnnaDxIVqB2VII0oCdJD9WYKG4nbVw+qq5JEiR0MQFmmOYx+GyMioIR8olPUUsng4JoYDy4muo5cgD4i//Z"

func getTestGreyImage(b64str string) (out [][]float64) {
	raw, err := base64.StdEncoding.DecodeString(b64str)
	if err != nil {
		panic(err)
	}
	img, err := jpeg.Decode(bytes.NewReader(raw))
	if err != nil {
		panic(err)
	}
	for y := 0; y < img.Bounds().Dy(); y++ {
		var outY []float64
		for x := 0; x < img.Bounds().Dx(); x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			c := float64(r) / 0xffff
			outY = append(outY, c)
		}
		out = append(out, outY)
	}
	return
}
