package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
//	"regexp"
	"net"
)

func checkPi(e error) {
	if(e != nil) {
		fmt.Println("Error: Given paramater(s) should be numeric!");
		os.Exit(1)
	}
}
/*
func checkME(e error) {
        if(e != nil) {
                fmt.Println("Error: Unknown error!");
                os.Exit(1)
        }
}*/

type Range struct {
	start, end int
}

func pingD(d string, port Range) {
	for s:= port.start; s <= port.end; s++ {
        	var addr string = fmt.Sprintf("%s:%d", d, s);
	        c, e := net.Dial("tcp", addr);

	        if e == nil {
                	fmt.Printf("Found %s\n", addr);
        	        c.Close();
	        }
}
}

func main() {
	var e error

	var port Range

	var address [4]Range
	var addressr string

	var am int

	var headers [][]string

	os.Args = os.Args[1:]
	for i := 0; i < len(os.Args); i++ {
		var arg string = os.Args[i]
		switch arg {
			case "-h", "--header":
				for _,hs := range strings.Split(os.Args[i+1], ":") {
					headers = append(headers,strings.Split(hs, "="))
				}
			case "-dr", "--drange":
				if am != 0 {
					fmt.Println("domain/ip range already set!")
					os.Exit(1)
				}
				am = 1
				var parts []string = strings.Split(os.Args[i+1], ":")
                                if len(parts) > 1 {
					var ports, porte int
                                        var pr []string = strings.Split(parts[1], "-")
                                        ports, e = strconv.Atoi(pr[0])
                                        checkPi(e)
                                        porte = ports
                                        if len(pr) > 1 {
                                                porte, e = strconv.Atoi(pr[1])
                                                checkPi(e)
                                        }
					port.start = ports
					port.end = porte
                                }
				addressr = parts[0]
			case "-r", "--range":
				if am != 0 {
                                        fmt.Println("domain/ip range already set!")
                                        os.Exit(1)
                                }
				am = 2
				var parts []string = strings.Split(os.Args[i+1], ":")
				if len(parts) > 1 {
					var ports, porte int
					var pr []string = strings.Split(parts[1], "-")
					ports, e = strconv.Atoi(pr[0])
					checkPi(e)
					porte = ports
					if len(pr) > 1 {
						porte, e = strconv.Atoi(pr[1])
						checkPi(e)
					}
					port.start = ports
                                        port.end = porte
				}
				
				
					for ii,ad := range strings.Split(parts[0], ".") {
						var as, ae int
						var ar []string = strings.Split(ad, "-")
                        	                as, e = strconv.Atoi(ar[0])
                	                        checkPi(e)
        	                                ae = as
	                                        if len(ar) > 1 {
                                        	        ae, e = strconv.Atoi(ar[1])
                                	                checkPi(e)
                        	                }
						address[ii] = Range{as, ae}
					}
					if len(address) != 4 {
						fmt.Println("write the range correctly")
						fmt.Println("	eg. \"127.0.0.1-10\"")
						fmt.Println("   eg. \"127.0.0.1:80-85\"")
						fmt.Println("   eg. \":8080\"")
					}
				
		}
	}

	if am == 0 {
		am = 2
		address = [4]Range{Range {127, 127},Range {0, 0},Range {0, 0},Range {1, 1}}
	}

	if am == 1 {
		pingD(addressr, port)
	}
	if am == 2 {
		for s1:= address[0].start; s1 <= address[0].end; s1++ {
			for s2:= address[1].start; s2 <= address[1].end; s2++ {
				for s3:= address[2].start; s3 <= address[2].end; s3++ {
					for s4:= address[3].start; s4 <= address[3].end; s4++ {
						pingD(fmt.Sprintf("%d.%d.%d.%d", s1, s2, s3, s4), port);
					}
				}
			}
		}
	}
}

