package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	//LOG
	logFile, err := os.OpenFile("solist.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	checkErr(err)

	defer logFile.Close()

	log.SetOutput(logFile)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//API

	getAllSolanaAPI()

	runSolana()

}

func getAllSolanaAPI() {

	apiFile, err := os.Open("./apis/alchemySolanaAPI.txt")
	checkErr(err)

	defer apiFile.Close()

	apiFileInfo, err := apiFile.Stat()
	checkErr(err)

	if apiFileInfo.Size() > 0 {

		scanner := bufio.NewScanner(apiFile)

		for scanner.Scan() {
			SolanaAPI = append(SolanaAPI, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	} else {
		fmt.Println("Файл с API : " + apiFileInfo.Name() + " - пустой, проверять нечем\n")
		fmt.Println("можно взять тут: https://www.alchemy.com/")
	}
}

func getActualSolanaAPI() string {

	result := SolanaAPI[IndexSolanaAPI]

	IndexSolanaAPI++

	if IndexSolanaAPI == len(SolanaAPI) {

		IndexSolanaAPI = 0

	}

	return result
}

func getSolanaBalance(address string) float64 {

	url := "https://solana-mainnet.g.alchemy.com/v2/" + getActualSolanaAPI()

	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"getBalance\",\"params\":[\"" + address + "\"]}")

	req, err := http.NewRequest("POST", url, payload)
	checkErr(err)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	checkErr(err)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var myJson SolStruct
	json.Unmarshal(body, &myJson)

	if myJson.Result.Value != 0 {

		Value := float64(myJson.Result.Value) / 1000000000
		return Value

	} else {

		return 0

	}
}

func runSolana() {

	addressesFile, err := os.Open("./addresses/addressesSolana.txt")
	checkErr(err)

	defer addressesFile.Close()

	balancesFile, err := os.OpenFile("./balances/balancesSolana.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	checkErr(err)

	defer balancesFile.Close()

	_, err = balancesFile.Seek(0, io.SeekStart)
	checkErr(err)

	err = os.Truncate(balancesFile.Name(), 0)
	checkErr(err)

	scanner := bufio.NewScanner(addressesFile)

	for scanner.Scan() {

		SolanaAmount = getSolanaBalance(scanner.Text())
		balancesFile.WriteString("\n" + "Amount " + scanner.Text() + " Amount " + strconv.FormatFloat(SolanaAmount, 'f', -1, 64))
		SolanaSum = SolanaSum + SolanaAmount

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	balancesFile.WriteString("\n" + "Общий баланс аккаунтов в сети Solana в SOL: " + strconv.FormatFloat(SolanaSum, 'f', -1, 64) + "\n")
}
