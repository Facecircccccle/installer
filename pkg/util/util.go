package util

import (
	"bytes"
	"fmt"
	"installer/pkg/version"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// IsIp check the input ip conforms to the format, e.g. "xxx.xxx.xxx.xxx".
func IsIp(ip string) (b bool) {
	if ip == "" {
		return true
	}
	if m, _ := regexp.MatchString("^(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)$", ip); !m {
		return false
	}
	return true
}

// IsIpWithSubnet check the input ip conforms to the format, e.g. "xxx.xxx.xxx.xxx/xx".
func IsIpWithSubnet(ip string) (b bool) {
	if ip == "" {
		return true
	}
	if m, _ := regexp.MatchString("^(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)/(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)$", ip); !m {
		return false
	}
	return true
}

// IsVipInSameSubnet check if VIP and master are on the same subnet.
func IsVipInSameSubnet(vip, target string, mask string) (b bool) {
	vipNet := strings.Split(vip, ".")
	targetNet := strings.Split(target, ".")
	maskNet := strings.Split(mask, ".")

	for i := 0; i < len(maskNet); i++ {
		m, _ := strconv.Atoi(maskNet[i])
		t, _ := strconv.Atoi(targetNet[i])
		v, _ := strconv.Atoi(vipNet[i])
		if (t & m) != (v & m) {
			return false
		}
	}
	return true
}
func GetNetmask(netcard string) string {
	mask := ExecShell("ifconfig " + netcard + " |sed -n 2p |awk -F ' ' '{print$4}'")
	return mask
}

// IsDigit check whether the string can be used as a digit.
func IsDigit(str string) bool {
	for _, x := range []rune(str) {
		if !unicode.IsDigit(x) {
			return false
		}
	}
	return true
}

// IsDigitWithStorage check whether the string can be used as a digit end with "Gi" or "Mi".
func IsDigitWithStorage(str string) bool {
	if (IsDigit(str[0 : len(str)-2])) == false {
		return false
	}
	if (str[len(str)-2:]) == "Gi" || (str[len(str)-2:]) == "Mi" {
		return true
	}
	return false
}

// IsDigitWithStorage check whether the string can be used as a digit with "%".
func IsDigitWithPercent(str string) bool {
	if (IsDigit(str[0 : len(str)-1])) == false {
		return false
	}
	if (str[len(str)-1:]) == "%" {
		return true
	}
	return false
}

// ExecShell run a shell command.
func ExecShell(s string) string {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	_ = cmd.Run()

	return out.String()
}

// StringIsInArray check whether a string array contains a string.
func StringIsInArray(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

// GetPauseVersion get pause version by given kubernetes version.
func GetPauseVersion(kubernetesVersion string) string {
	return version.GetComponentVersion()[kubernetesVersion].PauseVersion[0]
}

// GetCoreDNSVersion get CoreDNS version by given kubernetes version.
func GetCoreDNSVersion(kubernetesVersion string) string {
	return version.GetComponentVersion()[kubernetesVersion].CoreDNSVersion[0]
}

// GetEtcdVersion get Etcd version by given kubernetes version.
func GetEtcdVersion(kubernetesVersion string) string {
	return version.GetComponentVersion()[kubernetesVersion].EtcdVersion[0]
}

// StringAppend merge two strings together.
func StringAppend(s1 string, s2 string) string {
	var build strings.Builder
	build.WriteString(s1)
	build.WriteString(s2)
	return build.String()
}

// WriteToNewFile write the given string to target path.
func WriteToNewFile(path string, s string) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777) //linux 路径
	if err != nil {
		fmt.Printf("open err%s", err)
		return
	}
	defer f.Close()
	_, err2 := f.WriteString(s)
	if err2 != nil {
		fmt.Printf("write err:\n%s", err2)
		return
	}
}

// CheckNetCard check the net card exist on the remote server.
func CheckNetCard(ip string, card string) bool {
	ExecShell("rm -rf /etc/ansible/hosts")
	ExecShell("sh localScript/add_ansible_host.sh " + "netCardCheck " + ip)
	ExecShell("ansible-playbook localScript/netCardCheck.yaml")

	return strings.Contains(ExecShell("cat /home/"+ip+"/home/networkCard"), card)
}

// SshSuccess check the ssh key already configured properly.
func SshSuccess(ip string) bool {
	return strings.Contains(ExecShell("ansible "+ip+" -m ping"), "SUCCESS")
}
