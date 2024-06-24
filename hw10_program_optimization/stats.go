package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	stat := make(DomainStat)
	if domain == "" {
		return stat, nil
	}
	scaner := bufio.NewScanner(r)
	if err := scaner.Err(); err != nil {
		return stat, fmt.Errorf("scanner error: %w", err)
	}
	for scaner.Scan() {
		email := jsoniter.Get(scaner.Bytes(), "Email").ToString()
		dom := strings.Split(email, "@")
		if strings.Split(dom[1], ".")[1] != domain || len(dom) < 2 {
			continue
		}
		stat[strings.ToLower(dom[1])]++
	}
	return stat, nil
}
