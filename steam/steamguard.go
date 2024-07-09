package steam

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"time"
)

type Guard struct {
	SharedSecret   string
	IdentitySecret string
}

// GenerateCode generates a Steam Guard code
func (sg *Guard) GenerateCode() string {
	t := time.Now().Unix() / 30
	h := hmac.New(sha1.New, []byte(sg.SharedSecret))
	err := binary.Write(h, binary.BigEndian, t)
	if err != nil {
		return ""
	}
	hash := h.Sum(nil)
	offset := hash[19] & 0x0f
	code := (int(hash[offset])&0x7f)<<24 |
		(int(hash[offset+1])&0xff)<<16 |
		(int(hash[offset+2])&0xff)<<8 |
		(int(hash[offset+3]) & 0xff)
	charset := "23456789BCDFGHJKMNPQRTVWXY"
	result := make([]byte, 5)
	for i := 0; i < 5; i++ {
		result[i] = charset[code%len(charset)]
		code /= len(charset)
	}
	return string(result)
}

// GenerateConfirmationKey generates a confirmation key for a given tag
func (sg *Guard) GenerateConfirmationKey(tag string) string {
	t := time.Now().Unix()
	h := hmac.New(sha1.New, []byte(sg.IdentitySecret))
	h.Write([]byte{0, 0, 0, 0})
	err := binary.Write(h, binary.BigEndian, t)
	if err != nil {
		return ""
	}
	h.Write([]byte(tag))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// TimeAlign returns the current Unix timestamp
func (sg *Guard) TimeAlign() int64 {
	return time.Now().Unix()
}
