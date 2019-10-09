package main

import (
	"io/ioutil"
  	"encoding/json"
  	"fmt"
  	"os"
  	"path/filepath"
  	"bytes"
)

type ImageSize struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type Box struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
}

type BoundingBox struct {
	DataID int64 `json:"data_id"`
	Dots   []Box `json:"dots"`
	Image ImageSize `json:"image"`
}


func main() {
	dirname := "./data" + string(filepath.Separator)
	fmt.Println(dirname)
	
	d, err := os.Open(dirname)
	if err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}

	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".json" {
				//fmt.Println(dirname+file.Name())
				// Open our jsonFile
				jsonFile, err := os.Open(dirname+file.Name())
				// if we os.Open returns an error then handle it
				if err != nil {
				    fmt.Println(err)
				}
				fmt.Println("Successfully Opened", dirname+file.Name())

				defer jsonFile.Close()
				// parsing
				byteValue, _ := ioutil.ReadAll(jsonFile)
				byteValue = bytes.TrimPrefix(byteValue, []byte("\xef\xbb\xbf"))
				//fmt.Println(string(byteValue))
			    var result BoundingBox
			    //var orgSize ImageSize

			    if err := json.Unmarshal(byteValue, &result); err != nil {
			    	panic(err)
			    }
			    fmt.Println("======================")
			    fmt.Println(result.DataID)
			    fmt.Println(result.Image.Height)
				fmt.Println(result.Image.Width)
				fmt.Println("========BBOX========")
				fmt.Println(result.Dots[0].Height)
				fmt.Println(result.Dots[0].Width)
				fmt.Println(result.Dots[0].X)
				fmt.Println(result.Dots[0].Y)
				fmt.Println("======================")
			}
		}
	}
}
