package utils

import "syscall"

func IntArrtoString(arr [65]int8) string {
	var s string
	for _, r := range arr {
		if r != 0 {
			s += string(r)
		}
	}
	return s
}

// UtsString is used to
// https://golang.org/pkg/syscall/#Utsname
func UtsString() (string, error) {
	utsname := syscall.Utsname{}
	var s string
	err := syscall.Uname(&utsname)

	if err == nil {
		s += IntArrtoString(utsname.Sysname) + " " 
		s += IntArrtoString(utsname.Nodename) + " "
		s += IntArrtoString(utsname.Release) + " "
		s += IntArrtoString(utsname.Version) + " "
		s += IntArrtoString(utsname.Machine)
		return s, nil
	}
	return "", err
}