package toml

import "time"

func Fuzz(data []byte) int {

        type cats struct {
                Plato  string
                Cauchy string
        }
        type simple struct {
                Age     int
                Colors  [][]string
                Pi      float64
                YesOrNo bool
                Now     time.Time
                NowEast time.Time
                NowWest time.Time
                Andrew  string
                Kait    string
                My      map[string]cats
        }

        var val simple

    if _, err := Decode(string(data),&val); err != nil {
      return 0
    }
    return 1
}
