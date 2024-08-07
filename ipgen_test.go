package main

import (
	"net"
	"reflect"
	"testing"
)

func TestGenerateIPv4Representations(t *testing.T) {
	ip := net.ParseIP("192.168.1.1")

	tests := []struct {
		name    string
		formats []string
		mixed   bool
		pad     int
		want    []string
	}{
		{
			name:    "decimal only",
			formats: []string{"dec"},
			mixed:   false,
			pad:     0,
			want: []string{
				"192.168.1.1",
				"192.168.257",
				"192.11010305",
				"3232235777",
			},
		},
		{
			name:    "octal only",
			formats: []string{"oct"},
			mixed:   false,
			pad:     0,
			want: []string{
				"0300.0250.01.01",
				"0300.0250.0401",
				"0300.052000401",
				"030052000401",
			},
		},
		{
			name:    "hexadecimal only",
			formats: []string{"hex"},
			mixed:   false,
			pad:     0,
			want: []string{
				"0xc0.0xa8.0x1.0x1",
				"0xc0.0xa8.0x101",
				"0xc0.0xa80101",
				"0xc0a80101",
			},
		},
		{
			name:    "all formats",
			formats: []string{"dec", "oct", "hex"},
			mixed:   false,
			pad:     0,
			want: []string{
				"192.168.1.1",
				"192.168.257",
				"192.11010305",
				"3232235777",
				"0300.0250.01.01",
				"0300.0250.0401",
				"0300.052000401",
				"030052000401",
				"0xc0.0xa8.0x1.0x1",
				"0xc0.0xa8.0x101",
				"0xc0.0xa80101",
				"0xc0a80101",
			},
		},
		{
			name:    "mixed formats",
			formats: []string{"dec", "oct", "hex"},
			mixed:   true,
			pad:     0,
			want: []string{
				"192.168.1.1",
				"192.168.257",
				"192.11010305",
				"3232235777",
				"0300.0250.01.01",
				"0300.0250.0401",
				"0300.052000401",
				"030052000401",
				"0xc0.0xa8.0x1.0x1",
				"0xc0.0xa8.0x101",
				"0xc0.0xa80101",
				"0xc0a80101",
				"192.168.1.1",
				"192.168.1.01",
				"192.168.1.0x1",
				"192.168.01.1",
				"192.168.01.01",
				"192.168.01.0x1",
				"192.168.0x1.1",
				"192.168.0x1.01",
				"192.168.0x1.0x1",
				"192.0250.1.1",
				"192.0250.1.01",
				"192.0250.1.0x1",
				"192.0250.01.1",
				"192.0250.01.01",
				"192.0250.01.0x1",
				"192.0250.0x1.1",
				"192.0250.0x1.01",
				"192.0250.0x1.0x1",
				"192.0xa8.1.1",
				"192.0xa8.1.01",
				"192.0xa8.1.0x1",
				"192.0xa8.01.1",
				"192.0xa8.01.01",
				"192.0xa8.01.0x1",
				"192.0xa8.0x1.1",
				"192.0xa8.0x1.01",
				"192.0xa8.0x1.0x1",
				"0300.168.1.1",
				"0300.168.1.01",
				"0300.168.1.0x1",
				"0300.168.01.1",
				"0300.168.01.01",
				"0300.168.01.0x1",
				"0300.168.0x1.1",
				"0300.168.0x1.01",
				"0300.168.0x1.0x1",
				"0300.0250.1.1",
				"0300.0250.1.01",
				"0300.0250.1.0x1",
				"0300.0250.01.1",
				"0300.0250.01.01",
				"0300.0250.01.0x1",
				"0300.0250.0x1.1",
				"0300.0250.0x1.01",
				"0300.0250.0x1.0x1",
				"0300.0xa8.1.1",
				"0300.0xa8.1.01",
				"0300.0xa8.1.0x1",
				"0300.0xa8.01.1",
				"0300.0xa8.01.01",
				"0300.0xa8.01.0x1",
				"0300.0xa8.0x1.1",
				"0300.0xa8.0x1.01",
				"0300.0xa8.0x1.0x1",
				"0xc0.168.1.1",
				"0xc0.168.1.01",
				"0xc0.168.1.0x1",
				"0xc0.168.01.1",
				"0xc0.168.01.01",
				"0xc0.168.01.0x1",
				"0xc0.168.0x1.1",
				"0xc0.168.0x1.01",
				"0xc0.168.0x1.0x1",
				"0xc0.0250.1.1",
				"0xc0.0250.1.01",
				"0xc0.0250.1.0x1",
				"0xc0.0250.01.1",
				"0xc0.0250.01.01",
				"0xc0.0250.01.0x1",
				"0xc0.0250.0x1.1",
				"0xc0.0250.0x1.01",
				"0xc0.0250.0x1.0x1",
				"0xc0.0xa8.1.1",
				"0xc0.0xa8.1.01",
				"0xc0.0xa8.1.0x1",
				"0xc0.0xa8.01.1",
				"0xc0.0xa8.01.01",
				"0xc0.0xa8.01.0x1",
				"0xc0.0xa8.0x1.1",
				"0xc0.0xa8.0x1.01",
				"0xc0.0xa8.0x1.0x1",
				"192.168.257",
				"192.168.0401",
				"192.168.0x101",
				"192.0250.257",
				"192.0250.0401",
				"192.0250.0x101",
				"192.0xa8.257",
				"192.0xa8.0401",
				"192.0xa8.0x101",
				"0300.168.257",
				"0300.168.0401",
				"0300.168.0x101",
				"0300.0250.257",
				"0300.0250.0401",
				"0300.0250.0x101",
				"0300.0xa8.257",
				"0300.0xa8.0401",
				"0300.0xa8.0x101",
				"0xc0.168.257",
				"0xc0.168.0401",
				"0xc0.168.0x101",
				"0xc0.0250.257",
				"0xc0.0250.0401",
				"0xc0.0250.0x101",
				"0xc0.0xa8.257",
				"0xc0.0xa8.0401",
				"0xc0.0xa8.0x101",
				"192.11010305",
				"192.052000401",
				"192.0xa80101",
				"0300.11010305",
				"0300.052000401",
				"0300.0xa80101",
				"0xc0.11010305",
				"0xc0.052000401",
				"0xc0.0xa80101",
			},
		},
		{
			name:    "octal only pad 1",
			formats: []string{"oct"},
			mixed:   false,
			pad:     1,
			want: []string{
				"00300.00250.001.001",
				"00300.00250.00401",
				"00300.0052000401",
				"0030052000401",
			},
		},
		{
			name:    "hexadecimal only pad 1",
			formats: []string{"hex"},
			mixed:   false,
			pad:     1,
			want: []string{
				"0x0c0.0x0a8.0x01.0x01",
				"0x0c0.0x0a8.0x0101",
				"0x0c0.0x0a80101",
				"0x0c0a80101",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateIPv4Representations(ip, tt.formats, tt.mixed, tt.pad)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateIPv4Representations() = %v, want %v", got, tt.want)
			}
		})
	}
}
