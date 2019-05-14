package m3u8

import "bytes"
import "bufio"

func Fuzz(data []byte) int {
        a := bytes.NewReader(data)
        p := NewMasterPlaylist()
        err := p.DecodeFrom(bufio.NewReader(a), false)
        if err != nil {
                return 0
        }
        return 1
}
