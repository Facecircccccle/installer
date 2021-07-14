package util

import (
	"bytes"
	"fmt"
	"installer/pkg/version"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

func IsIp(ip string) (b bool) {
	if ip == ""{
		return true
	}
	if m, _ := regexp.MatchString("^(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)$", ip); !m {
		return false
	}
	return true
}

func IsIpWithSubnet(ip string) (b bool) {
	if ip == ""{
		return true
	}
	if m, _ := regexp.MatchString("^(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)/(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)$", ip); !m {
		return false
	}
	return true
}

func IsDigit(str string) bool {
	for _, x := range []rune(str) {
		if !unicode.IsDigit(x) {
			return false
		}
	}
	return true
}

func IsDigitWithStorage(str string) bool {
	if(IsDigit(str[0:len(str)-2])) == false{
		return false
	}
	if(str[len(str)-2:]) == "Gi" || (str[len(str)-2:]) == "Mi"{
		return true
	}
	return false
}

func IsDigitWithPercent(str string) bool {
	if(IsDigit(str[0:len(str)-1])) == false{
		return false
	}
	if(str[len(str)-1:]) == "%"{
		return true
	}
	return false
}

func ExecShell(s string) string {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	_ = cmd.Run()

	return out.String()
}

func StringIsInArray(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

func GetPauseVersion(kubernetesVersion string) string {
	return version.GetComponentVersion()[kubernetesVersion].PauseVersion[0]
}

func GetCoreDNSVersion(kubernetesVersion string) string {
	return version.GetComponentVersion()[kubernetesVersion].CoreDNSVersion[0]
}

func GetEtcdVersion(kubernetesVersion string) string {
	return version.GetComponentVersion()[kubernetesVersion].EtcdVersion[0]
}

func StringAppend(s1 string, s2 string) string {
	var build strings.Builder
	build.WriteString(s1)
	build.WriteString(s2)
	return build.String()
}

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

func CheckNetCard(ip string, card string) bool {
	ExecShell("rm -rf /etc/ansible/hosts")
	ExecShell("sh localScript/add_ansible_host.sh " + "netCardCheck " + ip)
	ExecShell("ansible-playbook localScript/netCardCheck.yaml")

	return strings.Contains(ExecShell("cat /home/" + ip + "/home/networkCard"), card)
}

func SshSuccess(ip string) bool {
	return strings.Contains(ExecShell("ansible " + ip + " -m ping"), "SUCCESS")
}