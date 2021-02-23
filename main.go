package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	jokesPath   = "src/github.com/chihkaiyu/YoGoCynical/yogodata/jokes.json"
	songsPath   = "src/github.com/chihkaiyu/YoGoCynical/yogodata/songs.json"
	defaultJoke = "KR1：\n坐在團隊旁，達成的天數比例達到 80% (五天有四天)\n- 參與日常任務開發中的討論 => 讓團隊對 po 有信任感\n- 上午坐在 A 團隊旁邊，下午坐在 B 團隊旁邊\n- 紀錄討論發生的次數、時間\n以上就是我偉大的 PO 主管花了幾個月的時間，終於訂出來的個人 OKR。\n之前就懷疑他被排擠，現在看起來他是真的很寂寞。"
	defaultSong = "[這樣上班最爽]\n#5 掌握朦朧美的說話技巧\n為了掩飾自己的無能，含糊其辭是你職涯必備的絕學。解釋東西的時候加上「大概八成對」，對於不熟的人就是對，對於熟悉的人還有兩成空間，進可功退可守，反正不會有人無聊到去深究到底對不對。\n開會或報告的結尾，隨意丟個困難問題出來然後說「這個我留給大家回去想一下」，好像你已經知道答案，簡直就是全公司最聰明的人一樣。\n講錯話千萬別認錯，「其實我跟你並不衝突，只是不同角度看同一件事」，你又跟他一樣聰明了呢！"
	helpMsg     = `Usage of ./yogoc:
	joke:
		get a joke
	song:
		get a song joke`
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	args := os.Args[1:]
	target := ""
	if len(args) == 0 {
		target = "joke"
	} else if strings.Contains(args[0], "help") {
		fmt.Println(helpMsg)
	}

	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		switch target {
		case "joke":
			fmt.Println(defaultJoke)
			return
		case "song":
			fmt.Println(defaultSong)
			return
		default:
			fmt.Println(defaultJoke)
			return
		}
	}

	jokes := []string{}
	path := ""
	switch target {
	case "joke":
		path = jokesPath
	case "song":
		path = songsPath
	}

	data, err := ioutil.ReadFile(filepath.Join(goPath, path))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &jokes); err != nil {
		panic(err)
	}

	joke := jokes[rand.Intn(len(jokes))]
	if strings.HasPrefix(joke, "https") {
		cmd := exec.Command("open", joke)
		if err := cmd.Run(); err != nil {
			panic(err)
		}
		return
	}

	fmt.Println(joke)
}
