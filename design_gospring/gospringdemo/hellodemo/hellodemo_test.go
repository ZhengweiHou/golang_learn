package hellodemo

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/util/sysconf"
)

// func init() {
// 	// Register the Service struct as a bean.
// 	gs.Object(&Service{})

// 	// Provide a [*http.ServeMux] as a bean.
// 	gs.Provide(func(s *Service) *http.ServeMux {
// 		http.HandleFunc("/echo", s.Echo)
// 		http.HandleFunc("/refresh", s.Refresh)
// 		return http.DefaultServeMux
// 	})
// }

const timeLayout = "2006-01-02 15:04:05.999 -0700 MST"

type Service struct {
	StartTime   time.Time          `value:"${start-time}"`
	RefreshTime gs.Dync[time.Time] `value:"${refresh-time}"`
}

func (s *Service) Echo(w http.ResponseWriter, r *http.Request) {
	str := fmt.Sprintf("start-time: %s refresh-time: %s",
		s.StartTime.Format(timeLayout),
		s.RefreshTime.Value().Format(timeLayout))
	_, _ = w.Write([]byte(str))
}

func (s *Service) Refresh(w http.ResponseWriter, r *http.Request) {
	sysconf.Set("refresh-time", time.Now().Format(timeLayout))
	_ = gs.RefreshProperties()
	_, _ = w.Write([]byte("OK!"))
}

func TestHelloDemo(t *testing.T) {
	sysconf.Set("start-time", time.Now().Format(timeLayout))
	sysconf.Set("refresh-time", time.Now().Format(timeLayout))

	// Register the Service struct as a bean.
	gs.Object(&Service{})

	// Provide a [*http.ServeMux] as a bean.
	gs.Provide(func(s *Service) *http.ServeMux {
		http.HandleFunc("/echo", s.Echo)
		http.HandleFunc("/refresh", s.Refresh)
		return http.DefaultServeMux
	})

	gs.Run()

	// strBean := "hello"

	// gs.Object(&strBean)

}
