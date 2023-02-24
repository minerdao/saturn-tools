package main

import (
        "fmt"
//      "io/ioutil"
        "net/http"
        "strconv"
        "encoding/json"
        "time"
)

type Earnings struct {
        FilAmount float64
        Timestamp string
}

type Data struct {
        Earnings []Earnings
}

var filAddressLocationMap = map[string]string{
        "f1wf7lu7quwz5hgsl5qybnjf................": "US", //这里填写自己的地址和对应地址地理位置,可以添加多个地址
        "f1pyfpcscyfto7phqkv2ybx7................": "US",
        "f1qbqd757pul7b5dpttmcz2c................": "US",
}

func main() {
        filAddresses := []string{
                "f1wf7lu7quwz5hgsl........................", //这里对应的是上文中的地址
                "f1pyfpcscyfto7phq........................",
                "f1qbqd757pul7b5dp........................",
        }
        for _, filAddress := range filAddresses {
                if err := fetchAndProcessData(filAddress); err != nil {
                        fmt.Println(err)
                }
        }
        sumTotalFilAmount()
}

func fetchAndProcessData(filAddress string) error {
        currentTime := time.Now()
        currentTimestamp := currentTime.Unix()
        currentTimestampInt := int(currentTimestamp)
        currentTimestampStr := strconv.Itoa(currentTimestampInt)
        fmt.Printf("\nFetching data for Fil address: %s\n", filAddress)

        earnings, err := fetchEarningsForTimeRange(filAddress, currentTimestamp-86400, currentTimestamp, currentTimestampStr)
        if err != nil {
                return err
        }

        var output string
        var output1 string
        var totalFilAmount float64

        for _, v := range earnings {
                output += fmt.Sprintf("filAmount: \x1b[32m%f\x1b[0m, timestamp: %s\n", v.FilAmount, v.Timestamp)
                totalFilAmount += v.FilAmount
        }

        avgFilAmount := totalFilAmount / float64(len(earnings))

        output1 = fmt.Sprintf("avgFilAmount: \x1b[32m%f\x1b[0m, totalFilAmount: \x1b[32m%f\x1b[0m", avgFilAmount, totalFilAmount)

        fmt.Println(output)
        fmt.Println(output1)

        location, exists := filAddressLocationMap[filAddress]
        if exists {
                fmt.Printf("Location: %s\n", location)
        } else {
                fmt.Println("Location not found for Fil address: ", filAddress)
        }

        return nil
}

func fetchEarningsForTimeRange(filAddress string, startTimestamp, endTimestamp int64, currentTimestampStr string) ([]Earnings, error) {
    url := fmt.Sprintf("https://uc2x7t32m6qmbscsljxoauwoae0yeipw.lambda-url.us-west-2.on.aws/?filAddress=%s&startDate=%d000&endDate=%d000&step=hour&currentTimestamp=%s000", filAddress, startTimestamp, endTimestamp, currentTimestampStr)

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var data Data
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return nil, err
    }

    return data.Earnings, nil
}

func sumTotalFilAmount() {
        filAddresses := []string{
                "f1wf7lu7quwz5hgsl5qyb................", //这里对应的是上文中的地址
                "f1pyfpcscyfto7phqkv2y................",
                "f1qbqd757pul7b5dpttmc................",
        }

        totalFilAmount := 0.0
        for _, filAddress := range filAddresses {
                earnings, err := fetchEarningsForTimeRange(filAddress, time.Now().Unix()-86400, time.Now().Unix(), strconv.Itoa(int(time.Now().Unix())))
                if err != nil {
                        fmt.Println(err)
                        continue
                }

                var new Earnings
                for _, v := range earnings {
                        new.FilAmount += v.FilAmount
                        new.Timestamp = v.Timestamp
                }

                totalFilAmount += new.FilAmount
        }

    fmt.Printf("\x1b[31mTotal Fil Amount: %f\x1b[0m\n", totalFilAmount)
}
