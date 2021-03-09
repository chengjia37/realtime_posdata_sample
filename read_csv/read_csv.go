package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/hpcloud/tail"
	"math/rand"
	"strconv"
	//"strings"
)

const deaultRegion = "ap-northeast-1"
const profileName = "kinesis"

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: read_csv [target file]")
		return
	}
	targetFile := args[0]

	t, err := tail.TailFile(targetFile, tail.Config{Follow: true, Poll: true})
	if err != nil {
		fmt.Println(err)
		return
	}

	mySession := session.Must(
		session.NewSessionWithOptions(session.Options{Profile: profileName}))
	service := kinesis.New(mySession, aws.NewConfig().WithRegion(deaultRegion))

	for line := range t.Lines {
		// _ = line
		// enc := base64.StdEncoding.EncodeToString([]byte("Hello"))
		key := rand.Intn(50)
		fmt.Println(key)
		record := &kinesis.PutRecordInput{
			Data:         []byte(line.Text),
			PartitionKey: aws.String(strconv.Itoa(key)),
			StreamName:   aws.String("realtime_pos"),
		}
		kinesisReesp, err := service.PutRecord(record)
		if err != nil {
			panic(err)
		}
		fmt.Println(line.Text)
		fmt.Println(kinesisReesp)
	}
	// watcher, err := fsnotify.NewWatcher()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer watcher.Close()

	// done := make(chan bool)

	// go processDirChange(watcher, done)

	// err = watcher.Add(targetDir)

	// // wait complete
	// <-done
}

// func processDirChange(watcher *fsnotify.Watcher, done chan bool) {
// 	for {
// 		select {
// 		case event := <-watcher.Events:
// 			fmt.Println("events: ", event)
// 			switch {
// 			case event.Op&fsnotify.Create == fsnotify.Create:
// 				fmt.Println("Created file:", event.Name)
// 				// time.Sleep(time.Millisecond * 1500)
// 				go processFile(event.Name)
// 			}
// 		case err := <-watcher.Errors:
// 			fmt.Println("error: ", err)
// 		}
// 	}

// }

// func processFile(targetFile string) {
// 	time.Sleep(time.Millisecond * 1500)
// 	lines, err := readLine(targetFile)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, l := range lines {
// 		fmt.Println(l)
// 	}
// }

// func readLine(filename string) ([]string, error) {
// 	fmt.Println("Opening " + filename)
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var lineItems []string
// 	reader := bufio.NewReader(file)
// 	var buff string = ""
// 	for {
// 		line, isPrefix, err := reader.ReadLine()
// 		if err == io.EOF {
// 			break
// 		}
// 		buff = buff + string(line)
// 		if !isPrefix {
// 			lineItems = append(lineItems, buff)
// 			buff = ""
// 		}
// 	}
// 	return lineItems, nil
// }
