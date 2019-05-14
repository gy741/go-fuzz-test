package m3u8

import "bytes"

func Fuzz(data []byte) int {

        p := NewMasterPlaylist()
        err := p.DecodeFrom(bytes.NewReader(data), false)
        if err != nil {
                return 0
        }
        return 1
}
