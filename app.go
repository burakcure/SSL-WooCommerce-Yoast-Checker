package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var output []string
var stillinCode bool
var outputFile string

func upload(c echo.Context) error {
	if stillinCode == false {
		stillinCode = true
		output = output[:0]
		output = append(output, "<div>---------------------------------WEBSITE-------------------SSL      </div>	")

		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		contents, _ := ioutil.ReadAll(src)

		arrayofWebsites := strings.Split(string(contents), "\r\n")
		if len(arrayofWebsites) > 0 {
			for i := 0; i < len(arrayofWebsites); i++ {
				fmt.Println(arrayofWebsites[i])

				handled := false
				output = append(output, "<div style=\"color:red;font-size:150%;margin-left:7%;margin-right:70%;margin-top:1%;border-style:solid;\">"+arrayofWebsites[i]+"		")

				_, err := http.Get("https://" + arrayofWebsites[i] + "/")
				if err != nil {
					handled = true
					if strings.Contains(err.Error(), "no such host") == true {
						//	fmt.Println("This site does not exists")
						output = append(output, "DOES NOT EXIST</div>")

					} else {
						//		fmt.Println("This site is not using (maybe invalid) SSL certificate")
						output = append(output, "NO</div>")
					}

				}
				if handled == false {

					//	fmt.Println("This site is using a SSL certificate.")
					output[len(output)-1] = strings.Replace(output[len(output)-1], "red", "blue", 1)
					output = append(output, "YES</div>")
				}
			}

			outputFile = strings.Join(output, "")

		}
	}
	stillinCode = false
	return c.HTML(http.StatusOK, string(outputFile))

}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")
	e.POST("/upload", upload)
	e.POST("/uploadseo", uploadseo)
	e.Logger.Fatal(e.Start(":1323"))
}

func uploadseo(c echo.Context) error {
	if stillinCode == false {
		stillinCode = true
		output = output[:0]
		output = append(output, "<div>---------------------------------WEBSITE------SEO-----WOOCOMMERCE      </div>	")

		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		contents, _ := ioutil.ReadAll(src)

		arrayofWebsites := strings.Split(string(contents), "\r\n")
		if len(arrayofWebsites) > 0 {
			for i := 0; i < len(arrayofWebsites); i++ {
				fmt.Println(arrayofWebsites[i])

				output = append(output, "<div style=\"color:red;font-size:150%;margin-left:7%;margin-right:70%;margin-top:1%;border-style:solid;\">"+arrayofWebsites[i]+"		")

				inside, _ := http.Get("https://" + arrayofWebsites[i] + "/")
				defer inside.Body.Close()
				html, _ := ioutil.ReadAll(inside.Body)
				seo1 := strings.Contains(string(html), "SEO")
				seo2 := strings.Contains(string(html), "Seo")

				if seo1 == true || seo2 == true {
					output[len(output)-1] = strings.Replace(output[len(output)-1], "red", "#339b21", 1)
					output = append(output, "YES   ")
				} else {
					output = append(output, "<span style:\"color:red\">NO</span>    ")

				}
				woo1 := strings.Contains(string(html), "WooCommerce")
				woo2 := strings.Contains(string(html), "woocommerce")
				woo3 := strings.Contains(string(html), "1.5.4")
				if woo1 == true || woo2 == true || woo3 == true {
					output[len(output)-1] = strings.Replace(output[len(output)-1], "red", "#339b21", 1)
					output = append(output, "YES</div>")
				} else {
					output = append(output, "<span style:\"font-	color:red\">NO</span></div>")

				}

			}
		}

		outputFile = strings.Join(output, "")

	}
	stillinCode = false
	return c.HTML(http.StatusOK, string(outputFile))
}
