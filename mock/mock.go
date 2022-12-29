package mock

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func Example() {
	data := make(chan string)
	go Start(time.Second*5, data)

	for {
		v, ok := <-data
		if !ok {
			fmt.Println("The chan has been close!")
		}
		fmt.Println(v)
	}

}

func Start(interval time.Duration, data chan string) {
	for {
		data <- GenerateLineProtocol()
		time.Sleep(interval)
	}

}

func GenerateLineProtocol() string {

	station := []string{"XiaoMaiDao", "LianyunGang"}
	mock := Mock{}
	buf := bytes.Buffer{}

	time := time.Now()

	for _, s := range station {
		wind := Wind{
			Time:      time,
			Station:   s,
			Speed:     mock.Float64(),
			direction: mock.Integer(),
		}
		air := Air{
			Time:        time,
			Station:     s,
			visibility:  mock.Float64(),
			temperature: mock.Float64(),
			pressure:    mock.Float64(),
		}
		sea := Sea{
			Time:        time,
			Station:     s,
			temperature: mock.Float64(),
		}
		buf.Write([]byte(wind.toLineProtocol()))
		buf.Write([]byte("\n"))
		buf.Write([]byte(air.toLineProtocol()))
		buf.Write([]byte("\n"))
		buf.Write([]byte(sea.toLineProtocol()))
		buf.Write([]byte("\n"))
	}
	return buf.String()
}

func GenerateJson(interval time.Duration) {

}

type Mock struct {
}

func (mock Mock) Integer() int64 {
	return rand.Int63()
}

func (mock Mock) String(str string) string {
	return str
}

func (mock Mock) Float64() float64 {
	return rand.Float64()
}

type Wind struct {
	Time      time.Time
	Station   string
	Speed     float64
	direction int64
}

type Air struct {
	Time        time.Time
	Station     string
	visibility  float64
	temperature float64
	pressure    float64
}

type Sea struct {
	Time        time.Time
	Station     string
	temperature float64
}

func (w *Wind) toLineProtocol() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "wind,station=%s speed=%v,direction=%v %v", w.Station, w.Speed, w.direction, w.Time.UnixNano())
	return buf.String()
}

func (a *Air) toLineProtocol() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "air,station=%s visibility=%v,temperature=%v,pressure=%v %v", a.Station, a.visibility, a.temperature, a.pressure, a.Time.UnixNano())
	return buf.String()
}

func (s *Sea) toLineProtocol() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "sea,station=%s temperature=%v %v", s.Station, s.temperature, s.Time.UnixNano())
	return buf.String()
}
