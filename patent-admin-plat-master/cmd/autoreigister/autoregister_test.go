package autoreigister

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-admin/app/user-agent/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestRegisterExistedUser(t *testing.T) {
	url := "http://39.105.142.144:8000/api/v1/register"
	params := make(map[string]interface{})
	params["roleId"] = 2

	file, err := os.Open("accounts.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		userName := line[0]
		password := line[1]
		params["username"] = userName
		params["password"] = password
		params["nickname"] = userName
		bs, _ := json.Marshal(params)
		r := bytes.NewReader(bs)

		resp, err := http.Post(url, "application/json", r)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != 200 {
			fmt.Println("invalid status code", resp.StatusCode, "exit...")
			break
		}
		fmt.Println("register account", params["username"], "successful")
	}
}

func TestRegisterNewUser(t *testing.T) {
	url := "http://39.105.142.144:8000/api/v1/register"
	params := make(map[string]interface{})
	params["roleId"] = 2

	fileName := fmt.Sprintf("accounts_%s.csv", time.Now().Format("2006-01-02"))
	fmt.Println(fileName)

	// 打开文件
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建bufio.Writer
	writer := bufio.NewWriter(file)

	for i := 202301051; i <= 202301070; i++ {
		params["username"] = strconv.Itoa(i)
		params["password"] = utils.RandStringRunes(8)
		params["nickname"] = strconv.Itoa(i)
		bs, _ := json.Marshal(params)
		r := bytes.NewReader(bs)

		resp, err := http.Post(url, "application/json", r)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != 200 {
			fmt.Println("invalid status code", resp.StatusCode, "exit...")
			break
		}

		_, err = writer.WriteString(fmt.Sprintf("%s,%s\n", params["username"], params["password"]))
		assert.NoError(t, err)
		writer.Flush()

		fmt.Println("register account", params["username"], "successful")
	}
}
