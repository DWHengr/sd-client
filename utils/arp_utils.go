package utils

import (
	"bufio"
	"net"
	"net/netip"
	"os/exec"
	"runtime"
	"sd-client/logger"
	"strings"

	"github.com/mdlayher/arp"
)

func SendGratutiousArp() {
	if runtime.GOOS == "windows" {
		return
	}
	interfaces, err := net.Interfaces()
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	for _, ifi := range interfaces {
		if ifi.Flags&net.FlagLoopback != 0 || ifi.Flags&net.FlagUp == 0 {
			continue
		}
		// 创建ARP客户端
		client, err := arp.Dial(&ifi)
		if err != nil {
			logger.Logger.Error(err)
			return
		}
		// 发送ARP请求
		addr, err := netip.ParseAddr("255.255.255.255")
		err = client.Request(addr)
		if err != nil {
			logger.Logger.Error(err)
			return
		}
		logger.Logger.Info(ifi.Name + ":ARP请求已发送")
		client.Close()
	}
}

func GetARPTables() map[string]string {
	arpTable := make(map[string]string)

	var cmd *exec.Cmd
	var output []byte
	var err error

	// 根据不同操作系统执行不同的命令
	if runtime.GOOS == "windows" {
		cmd = exec.Command("arp", "-a")
	} else {
		cmd = exec.Command("arp", "-n")
	}

	// 执行命令并捕获输出
	if output, err = cmd.Output(); err != nil {
		println("Error executing command:", err)
		return arpTable
	}

	// 解析命令输出
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			if runtime.GOOS == "windows" {
				arpTable[strings.ToLower(fields[1])] = fields[0]
			} else {
				arpTable[strings.ToLower(fields[2])] = fields[0]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		println("Error scanning command output:", err)
	}
	return arpTable
}
