package discovery

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/netx"
	"os"
	"strings"

	"google.golang.org/grpc/resolver"
)

const (
	allEths  = "0.0.0.0"
	envPodIp = "POD_IP"
)

type Server struct {
	Name    string `json:"name"`
	Addr    string `json:"addr"`    // 地址
	Version string `json:"version"` // 版本
	Weight  int64  `json:"weight"`  // 权重
}

func BuildPrefix(server Server) string {
	if server.Version == "" {
		return fmt.Sprintf("/%s/", server.Name)
	}

	return fmt.Sprintf("/%s/%s/", server.Name, server.Version)
}

func BuildRegisterPath(server Server) string {
	return server.Name
	//return fmt.Sprintf("%s%s", BuildPrefix(server), server.Addr)
}

func figureOutListenOn(listenOn string) string {
	fields := strings.Split(listenOn, ":")
	if len(fields) == 0 {
		return listenOn
	}

	host := fields[0]
	if len(host) > 0 && host != allEths {
		return listenOn
	}

	ip := os.Getenv(envPodIp)
	if len(ip) == 0 {
		ip = netx.InternalIp()
	}
	if len(ip) == 0 {
		return listenOn
	}

	return strings.Join(append([]string{ip}, fields[1:]...), ":")
}

func ParseValue(value []byte) (Server, error) {
	server := Server{}
	if err := json.Unmarshal(value, &server); err != nil {
		return server, err
	}

	return server, nil
}

func SplitPath(path string) (Server, error) {
	server := Server{}
	strs := strings.Split(path, "/")
	if len(strs) == 0 {
		return server, errors.New("invalid path")
	}

	server.Addr = strs[len(strs)-1]

	return server, nil
}

// Exist helper function
func Exist(l []resolver.Address, addr resolver.Address) bool {
	for i := range l {
		if l[i].Addr == addr.Addr {
			return true
		}
	}

	return false
}

// Remove helper function
func Remove(s []resolver.Address, addr resolver.Address) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr.Addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}

func BuildResolverUrl(app string) string {
	return schema + ":///" + app
}
