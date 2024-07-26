package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func GenerateIPv4Representations(ip net.IP, formats []string, mixed bool, pad int) []string {
	var results []string

	p := strings.Repeat("0", pad)
	// Extract individual bytes
	b1, b2, b3, b4 := ip[12], ip[13], ip[14], ip[15]

	formatMap := map[string][]string{
		"dec": {
			fmt.Sprintf("%d.%d.%d.%d", b1, b2, b3, b4),
			fmt.Sprintf("%d.%d.%d", b1, b2, uint32(b3)<<8|uint32(b4)),
			fmt.Sprintf("%d.%d", b1, uint32(b2)<<16|uint32(b3)<<8|uint32(b4)),
			fmt.Sprintf("%d", uint32(b1)<<24|uint32(b2)<<16|uint32(b3)<<8|uint32(b4)),
		},
		"oct": {
			fmt.Sprintf("0%s%o.0%s%o.0%s%o.0%s%o", p, b1, p, b2, p, b3, p, b4),
			fmt.Sprintf("0%s%o.0%s%o.0%s%o", p, b1, p, b2, p, uint32(b3)<<8|uint32(b4)),
			fmt.Sprintf("0%s%o.0%s%o", p, b1, p, uint32(b2)<<16|uint32(b3)<<8|uint32(b4)),
			fmt.Sprintf("0%s%o", p, uint32(b1)<<24|uint32(b2)<<16|uint32(b3)<<8|uint32(b4)),
		},
		"hex": {
			fmt.Sprintf("0x%s%x.0x%s%x.0x%s%x.0x%s%x", p, b1, p, b2, p, b3, p, b4),
			fmt.Sprintf("0x%s%x.0x%s%x.0x%s%x", p, b1, p, b2, p, uint32(b3)<<8|uint32(b4)),
			fmt.Sprintf("0x%s%x.0x%s%x", p, b1, p, uint32(b2)<<16|uint32(b3)<<8|uint32(b4)),
			fmt.Sprintf("0x%s%x", p, uint32(b1)<<24|uint32(b2)<<16|uint32(b3)<<8|uint32(b4)),
		},
	}

	for _, format := range formats {
		if val, ok := formatMap[format]; ok {
			results = append(results, val...)
		}
	}

	if !mixed {
		return results
	}

	// Generate all combinations of mixed formats
	fourOctetCombos := [][]string{
		{fmt.Sprintf("%d", b1), fmt.Sprintf("0%s%o", p, b1), fmt.Sprintf("0x%s%x", p, b1)},
		{fmt.Sprintf("%d", b2), fmt.Sprintf("0%s%o", p, b2), fmt.Sprintf("0x%s%x", p, b2)},
		{fmt.Sprintf("%d", b3), fmt.Sprintf("0%s%o", p, b3), fmt.Sprintf("0x%s%x", p, b3)},
		{fmt.Sprintf("%d", b4), fmt.Sprintf("0%s%o", p, b4), fmt.Sprintf("0x%s%x", p, b4)},
	}

	// Generate permutations
	for _, f1 := range fourOctetCombos[0] {
		for _, f2 := range fourOctetCombos[1] {
			for _, f3 := range fourOctetCombos[2] {
				for _, f4 := range fourOctetCombos[3] {
					mixedFormat := fmt.Sprintf("%s.%s.%s.%s", f1, f2, f3, f4)
					results = append(results, mixedFormat)
				}
			}
		}
	}

	threeOctetCombos := [][]string{
		{fmt.Sprintf("%d", b1), fmt.Sprintf("0%s%o", p, b1), fmt.Sprintf("0x%s%x", p, b1)},
		{fmt.Sprintf("%d", b2), fmt.Sprintf("0%s%o", p, b2), fmt.Sprintf("0x%s%x", p, b2)},
		{
			fmt.Sprintf("%d", uint32(b3)<<8|uint32(b4)),
			fmt.Sprintf("0%s%o", p, uint32(b3)<<8|uint32(b4)),
			fmt.Sprintf("0x%s%x", p, uint32(b3)<<8|uint32(b4)),
		},
	}

	// Generate permutations
	for _, f1 := range threeOctetCombos[0] {
		for _, f2 := range threeOctetCombos[1] {
			for _, f3 := range threeOctetCombos[2] {
				mixedFormat := fmt.Sprintf("%s.%s.%s", f1, f2, f3)
				results = append(results, mixedFormat)
			}
		}
	}

	twoOctetCombos := [][]string{
		{fmt.Sprintf("%d", b1), fmt.Sprintf("0%s%o", p, b1), fmt.Sprintf("0x%s%x", p, b1)},
		{
			fmt.Sprintf("%d", uint32(b2)<<16|uint32(b3)<<8|uint32(b4)),
			fmt.Sprintf("0%s%o", p, uint32(b2)<<16|uint32(b3)<<8|uint32(b4)),
			fmt.Sprintf("0x%s%x", p, uint32(b2)<<16|uint32(b3)<<8|uint32(b4)),
		},
	}

	// Generate permutations
	for _, f1 := range twoOctetCombos[0] {
		for _, f2 := range twoOctetCombos[1] {
			mixedFormat := fmt.Sprintf("%s.%s", f1, f2)
			results = append(results, mixedFormat)
		}
	}

	return results
}

func main() {
	var mixFlag bool
	var padFlag int
	var formatFlag string

	flag.Usage = func() {
		fmt.Println(`
▪   ▄▄▄· ▄▄ • ▄▄▄ . ▐ ▄ 
██ ▐█ ▄█▐█ ▀ ▪▀▄.▀·•█▌▐█
▐█· ██▀·▄█ ▀█▄▐▀▀▪▄▐█▐▐▌
▐█▌▐█▪·•▐█▄▪▐█▐█▄▄▌██▐█▌
▀▀▀.▀   ·▀▀▀▀  ▀▀▀ ▀▀ █▪`)
		fmt.Printf("\nUsage\n  %s [options] <IPv4 address>\n\nOptions:\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	flag.BoolVar(&mixFlag, "mix", false, "all mixed combinations")
	flag.IntVar(&padFlag, "pad", 0, "number of 0s to pad hex and oct numbers")
	flag.StringVar(&formatFlag, "format", "all", "specify formats: dec, oct, hex, all")
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	input := flag.Arg(0)
	ip := net.ParseIP(input)
	if ip == nil || ip.To4() == nil {
		fmt.Println("Address needs to be IPv4")
		os.Exit(1)
	}

	validFormats := map[string]bool{
		"dec": true,
		"oct": true,
		"hex": true,
		"all": true,
	}

	formats := strings.Split(formatFlag, ",")
	for _, format := range formats {
		if !validFormats[format] {
			log.Fatalf("Invalid format: %s. Valid formats are dec, oct, hex, all.", format)
		}
	}

	if formatFlag == "all" {
		formats = []string{"dec", "oct", "hex"}
	}

	result := GenerateIPv4Representations(ip, formats, mixFlag, padFlag)
	for _, format := range result {
		fmt.Println(format)
	}
}
