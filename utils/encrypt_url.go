package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func EncryptUrl(_url string, _resize string, _width int, _height int, _gravity string, _enlarge int, _extension string) string {
	key := os.Getenv("IMGPROXY_KEY")
	salt := os.Getenv("IMGPROXY_SALT")

	var keyBin, saltBin []byte
	var err error

	if keyBin, err = hex.DecodeString(key); err != nil {
		log.Fatal("Key expected to be hex-encoded string")
	}

	if saltBin, err = hex.DecodeString(salt); err != nil {
		log.Fatal("Salt expected to be hex-encoded string")
	}

	encodedURL := base64.RawURLEncoding.EncodeToString([]byte(_url))
	path := fmt.Sprintf("/%s/%d/%d/%s/%d/%s.%s", _resize, _width, _height, _gravity, _enlarge, encodedURL, _extension)

	mac := hmac.New(sha256.New, keyBin)
	mac.Write(saltBin)
	mac.Write([]byte(path))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return fmt.Sprintf("/%s%s", signature, path)
}

func EncryptUrlString(_url string, _parameters string, width int, height int, _extension string) string {
	key := os.Getenv("IMGPROXY_KEY")
	salt := os.Getenv("IMGPROXY_SALT")

	var keyBin, saltBin []byte
	var err error

	if keyBin, err = hex.DecodeString(key); err != nil {
		log.Fatal("Key expected to be hex-encoded string")
	}

	if saltBin, err = hex.DecodeString(salt); err != nil {
		log.Fatal("Salt expected to be hex-encoded string")
	}

	var path string

	encodedURL := base64.RawURLEncoding.EncodeToString([]byte(_url))
	if len(_parameters) <= 1 {
		path = fmt.Sprintf("/resize:fill:%d:%d/quality:100/%s.%s", width, height, encodedURL, _extension)
	} else {
		path = fmt.Sprintf("/%s/resize:fill:%d:%d/quality:100/%s.%s", _parameters, width, height, encodedURL, _extension)
	}

	mac := hmac.New(sha256.New, keyBin)
	mac.Write(saltBin)
	mac.Write([]byte(path))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return fmt.Sprintf("/%s%s", signature, path)
}

func EncryptUrlParams(_url string, _parameters string, _extension string) string {
	key := os.Getenv("IMGPROXY_KEY")
	salt := os.Getenv("IMGPROXY_SALT")

	var keyBin, saltBin []byte
	var err error

	if keyBin, err = hex.DecodeString(key); err != nil {
		log.Fatal("Key expected to be hex-encoded string")
	}

	if saltBin, err = hex.DecodeString(salt); err != nil {
		log.Fatal("Salt expected to be hex-encoded string")
	}

	var path string

	encodedURL := base64.RawURLEncoding.EncodeToString([]byte(_url))
	if len(_parameters) <= 1 {
		path = fmt.Sprintf("/%s.%s", encodedURL, _extension)
	} else {
		path = fmt.Sprintf("/%s/%s.%s", _parameters, encodedURL, _extension)
	}

	mac := hmac.New(sha256.New, keyBin)
	mac.Write(saltBin)
	mac.Write([]byte(path))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return fmt.Sprintf("/%s%s", signature, path)
}
