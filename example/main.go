package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/toxyl/filejanitor"
	"github.com/toxyl/flo"
)

func main() {
	c, err := filejanitor.ConfigFromFile("config.yaml")
	if err != nil {
		fmt.Printf("Could not load config: %s", err)
		return
	}
	fmt.Printf("Hi there, FileJanitor here, I'm gonna watch the %d paths you gave me.\n", len(c.Policies))
	wg := &sync.WaitGroup{}
	for _, policy := range c.Policies {
		_ = os.MkdirAll(policy.Path, 0755)
		wg.Add(1)
		go func(path, extension string) {
			defer wg.Done()
			for {
				f := time.Now().Format(time.UnixDate)
				if extension != "" {
					f += "." + extension
				}
				_ = flo.Dir(path).File(f).WriteString(time.Now().Format(time.RFC3339))
				time.Sleep(10 * time.Second)
			}
		}(policy.Path, policy.Extension)
	}
	_ = filejanitor.Run(c, func(errors []error) {
		for _, err := range errors {
			fmt.Printf("Encountered error: %s\n", err)
		}
	})
	wg.Wait()
}
