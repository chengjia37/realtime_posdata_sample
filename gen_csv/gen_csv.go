package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// tx_date, shop_id, pos_id, tx_id, item_id, ammount, cardnum
func main() {
	// 引数のParse
	var flagShopID = flag.Int("shopid", 0, "Shop ID(0-3000)")
	var flagPosID = flag.Int("posid", 0, "POS ID(0-5)")
	flag.Parse()

	shopID := strconv.Itoa(*flagShopID)
	posID := strconv.Itoa(*flagPosID)
	multiShopWait := 0
	if *flagShopID < 0 {
		multiShopWait = -(*flagShopID)
	}
	// fmt.Println("shopID=" + shopID)
	// fmt.Println("multiShopWait=" + strconv.Itoa(multiShopWait))

	var linenumArray = []int{1, 1, 1, 2, 2, 2, 3, 3, 4, 5}
	var itemIDsArray = []int{
		2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
		5, 5, 5, 5, 5, 5,
		9, 9, 9, 9, 9,
		4, 4, 4,
		3, 3,
		1, 6, 7, 8, 10,
	}
	var ammountsArray = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 3}
	var maxCardnum = 1000000

	// 乱数の初期化
	rand.Seed(time.Now().UnixNano())

	for {
		// 一回のトランザクション
		for i := 0; i < linenumArray[rand.Intn(len(linenumArray))]; i++ {
			if multiShopWait > 0 {
				shopID = strconv.Itoa(1 + rand.Intn(3000))
			}
			txDate := time.Now().Format("2006/01/02 15:04:05")
			txuuid, _ := uuid.NewUUID()
			txID := txuuid.String()
			// fmt.Println("  ", uuidObj.String())
			itemID := strconv.Itoa(itemIDsArray[rand.Intn(len(itemIDsArray))])
			ammount := strconv.Itoa(ammountsArray[rand.Intn(len(ammountsArray))])
			cardnum := strconv.Itoa(rand.Intn(maxCardnum))
			items := []string{txDate, shopID, posID, txID, itemID, ammount, cardnum}
			fmt.Println(strings.Join(items, ","))
		}

		// sleepLength := 0
		if multiShopWait > 0 {
			sleepLength := multiShopWait
			time.Sleep(time.Duration(sleepLength) * time.Millisecond)
		} else {
			// sleep 1 - 3 sec
			sleepLength := 1 + rand.Intn(3)
			time.Sleep(time.Duration(sleepLength) * time.Second)
		}
		// fmt.Println(sleepLength)
	}
}
