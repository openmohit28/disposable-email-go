package disposableemail

import (
	"bufio"
	"embed"
	"strings"
)

//go:embed domains.txt
var domainsFile embed.FS

var disposableDomains map[string]struct{}

func init() {
	disposableDomains = make(map[string]struct{})
	
	file, err := domainsFile.Open("domains.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		disposableDomains[domain] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// IsDisposable checks if the given email address uses a disposable domain
func IsDisposable(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	
	domain := strings.ToLower(parts[1])
	_, exists := disposableDomains[domain]
	return exists
}