package macmap

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func RefreshDB() {
	downloadLinks := make(map[string]string)
	downloadLinks["MAS"] = "http://standards-oui.ieee.org/oui/oui.csv"
	downloadLinks["MAM"] = "http://standards-oui.ieee.org/oui28/mam.csv"
	downloadLinks["MAL"] = "http://standards-oui.ieee.org/oui36/oui36.csv"

	ch := make(chan string)
	var fileCount int = 0

	for name, link := range downloadLinks {
		go func(name, link string) {
			res, err := http.Get(link)
			defer res.Body.Close()
			if err != nil {
				log.Fatalln(err)
			}

			localOut, err := os.Create(fmt.Sprintf("./%s.csv", name))
			if err != nil {
				log.Fatalln(err)
			}
			defer localOut.Close()

			_, err = io.Copy(localOut, res.Body)
			if err != nil {
				log.Fatalln(err)
			}

			fileCount++
			ch <- localOut.Name()

		}(name, link)
	}

	timeout := time.After(600 * time.Second)
	for fileCount < len(downloadLinks) {
		select {
		case res := <-ch:
			fmt.Printf("[%s] Finish download %s\n", time.Now().Format("2006-01-02 15:04:05"), res)
		case <-timeout:
			fmt.Printf("[%s] Timeout... \n", time.Now().Format("2006-01-02 15:04:05"))
			break
		}
	}

	fmt.Printf("[%s] All downloads are finished... \n", time.Now().Format("2006-01-02 15:04:05"))
}
