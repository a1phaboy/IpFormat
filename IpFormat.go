package Format

import (
	"strconv"
	"strings"
)

func StringToOct(str string)int{
	num:=0
	k:=1
	for i := len(str) - 1; i >= 0; i-- {
		num += int(str[i] - '0') * k
		k *= 2
	}
	return num
}

func Increase(num string)string {
	stringbyte := []byte(num)
	stringbyte[31]++
	for i := 30; i >= 0; i-- {
		if stringbyte[i+1] == '2' {
			stringbyte[i]++
			stringbyte[i+1] = '0'
		}
	}
	return string(stringbyte)
}

func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func DecToBin(num int)string{
	var str string
	for ;num!=0; {
		str+=strconv.Itoa(num % 2)
		num /= 2
	}
	for i := len(str); i < 8; i++{
		str +="0"
	}
	return ReverseString(str)
}

func IpTransform(str string) []string {
	var list [] string
	ip:= [4]int{ 0,0,0,0 }
	var ipstring string
	if strings.Contains(str,"/") {
		substring := strings.Split(str,"/")
		num, _ := strconv.Atoi(substring[1])
		ipaddress := strings.Split(substring[0],".")
		for i:=0;i<len(ipaddress);i++ {
			ipnum,_ := strconv.Atoi(ipaddress[i])
			ipstring+=DecToBin(ipnum)
		}
		ipstringbyte := []byte(ipstring)
		for i := num; i < 32; i++{
			ipstringbyte[i] = '0'
		}
		ipstring = string(ipstringbyte)
		for i := 0; i < 1<<(32-num)-2; i++{
			ipstring=Increase(ipstring)
			var s string
			for j := 0; j < 4; j++{
				ip[j] = StringToOct(ipstring[j*8:j*8+8])
				s+=strconv.Itoa(ip[j])
				if j != 3 {
					s+="."
				}
			}
			list = append(list,s)
		}
	} else if strings.Index(str,"/") != -1{
		substring := strings.Split(str,"-")
		var bstring string
		var estring string
		ipaddress := strings.Split(strings.Trim(substring[0]," "),".")
		for i:=0;i<len(ipaddress);i++ {
			ipnum,_ := strconv.Atoi(ipaddress[i])
			bstring+=DecToBin(ipnum)
		}
		ipaddress = strings.Split(strings.Trim(substring[1]," "),".")
		for i:=0;i<len(ipaddress);i++ {
			ipnum,_ := strconv.Atoi(ipaddress[i])
			estring+=DecToBin(ipnum)
		}
		for i := bstring;; i=Increase(i) {
			var s string
			for j := 0; j < 4; j++{

				ip[j] = StringToOct(i[j*8: j*8+8])
				s+= strconv.Itoa(ip[j])
				if j != 3{
					s+= "."
				}
			}
			list = append(list,s)
			if i==estring {
				break
			}
		}
	} else{
		substring := strings.Split(str,".")
		if len(substring) == 4 {
			ip1, _ := strconv.Atoi(substring[0])
			ip2, _ := strconv.Atoi(substring[1])
			ip3, _ := strconv.Atoi(substring[2])
			ip4, _ := strconv.Atoi(substring[3])
			if ip1 > 255 || ip2 > 255 || ip3 > 255 || ip4 > 255 {
				return nil
			}else{
				list = append(list, str)
			}
		}else{
			return nil
		}
	}
	return list
}
