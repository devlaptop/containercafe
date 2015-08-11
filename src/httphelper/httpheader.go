package httphelper

import (
	"log"
	"net/http"
	"strings"
)

func CopyHeader (dst http.Header, src http.Header) {
	for k, v := range src {
		for _, vv := range v {
			dst.Add(k,vv)
		}
    }
}

func DumpHeader (src http.Header) {
	for k, v := range src {
		log.Printf("%s: ",k)
		for _, vv := range v {
			log.Printf("%s ", vv)
		}
		log.Printf("\n")
    }
}

func IsUpgradeHeader (h http.Header) bool{
	for k, _ := range h {
		if strings.ToUpper(k)=="UPGRADE" {
				return true
		}
    }
	return false
}

// This is equivalent to Header.Get(key)
func GetHeader (h http.Header, key string) string{
	for k, v := range h {
		if strings.ToUpper(k)==strings.ToUpper(key) {
			val := ""
			for _, vv := range v {
				val += vv
			}
			return val
		}
    }
	return ""
}

func IsStreamHeader(h http.Header) bool{
	val := GetHeader(h, "Content-Type")
	if val == "application/octet-stream" {
		return true
	}
	return false
}

func IsDockerHeader(h http.Header) bool{
	val := GetHeader(h, "Content-Type")
	if val == "application/vnd.docker.raw-stream" {
		return true
	}
	return false
}
