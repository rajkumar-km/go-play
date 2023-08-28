/*
web_server listens on localhost:8000 and generates the lissajous images based on the parameters
For example:

	http://localhost:8000/lissajous
	http://localhost:8000/lissajous?cycles=2&res=0.001&size=400&nframes=100&delay=10
	http://localhost:8000/counter
*/
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	// Handler for /lissajous
	http.HandleFunc("/lissajous", lissajous)

	// Handler for /counter
	http.HandleFunc("/counter", counter)

	// A common handler for all the other URLs
	http.HandleFunc("/", handler)

	// Register the server to listen and serve from localhost:8000
	fmt.Println("Server listing on localhost:8000")
	http.ListenAndServe("localhost:8000", nil)
}

// Use sync.Mutex to allow concurrent access to request counter
var requestCount uint64
var requestMu sync.Mutex

// Create colors and a color palatte for lissajous image
var red = color.RGBA{0xff, 0x00, 0x00, 0xff}
var green = color.RGBA{0x00, 0xff, 0x00, 0xff}
var blue = color.RGBA{0x00, 0x00, 0xff, 0xff}
var palette = []color.Color{color.Black, red, green, blue}

// lissajous produces the GIF from request params and write to w
func lissajous(w http.ResponseWriter, r *http.Request) {
	requestMu.Lock()
	requestCount++
	requestMu.Unlock()

	cycles, res, size, nframes, delay, err := parseLissajousParams(r)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			fgColorIndex := uint8(rand.Intn(4))
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				fgColorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim) // NOTE: ignoring encoding errors
}

// parseLissajousParams parses the HTTP request parameters for lissajous image
func parseLissajousParams(r *http.Request) (float64, float64, int, int, int, error) {
	var (
		cycles  float64 = 5     // number of complete x oscillator revolutions
		res     float64 = 0.001 // angular resolution
		size    int     = 100   // image canvas covers [-size..+size]
		nframes int     = 64    // number of animation frames
		delay   int     = 8     // delay between frames in 10ms units
	)

	// ParseForm parses the submitted form data from the request and populates r.Form
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	if s := r.Form.Get("cycles"); len(s) > 0 {
		v, err := strconv.ParseFloat(s, 0)
		if err != nil {
			return 0, 0, 0, 0, 0, err
		}
		cycles = v
	}
	if s := r.Form.Get("res"); len(s) > 0 {
		v, err := strconv.ParseFloat(s, 0)
		if err != nil {
			return 0, 0, 0, 0, 0, err
		}
		res = v
	}
	if s := r.Form.Get("size"); len(s) > 0 {
		v, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			return 0, 0, 0, 0, 0, err
		}
		size = int(v)
	}
	if s := r.Form.Get("nframes"); len(s) > 0 {
		v, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			return 0, 0, 0, 0, 0, err
		}
		nframes = int(v)
	}
	if s := r.Form.Get("delay"); len(s) > 0 {
		v, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			return 0, 0, 0, 0, 0, err
		}
		delay = int(v)
	}

	return cycles, res, size, nframes, delay, nil
}

// counter is a HTTP handler to display total lissajous requests
func counter(w http.ResponseWriter, r *http.Request) {
	requestMu.Lock()
	fmt.Fprintf(w, "Total Lissajous Requests: %d", requestCount)
	requestMu.Unlock()
}

// A common handler for all other URLS just prints the request information
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	// ParseForm parses the submitted form data from the request and populates r.Form
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
