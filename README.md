# YoGoCynical
我講的話你一定要聽，不然將來會吃虧

# Install
Prepare Golang environment before you install it.  
```shell
go get -d github.com/chihkaiyu/YoGoCynical
cd $GOPATH/bin
go build -o yogoc $GOPATH/src/github.com/chihkaiyu/YoGoCynical
./yogoc --help
```
or
```shell
go get github.com/chihkaiyu/YoGoCynical
./YoGoCynical --help
```

# Result
```
$ ./yogoc --help
Usage of ./yogoc:
        joke:
                get a joke
        song:
                get a song joke

$ ./yogoc 
「這個感覺不對」
「痾...請問是甚麼感覺不對？」「這個不夠2D」
我看著手上的2D美術圖，努力地嘗試達到老闆的境界
「不然你說，這個感覺你可以？」至少我確定它是2D的。「這個感覺不對」

$ ./yogoc joke
主管最近常常悶頭悶腦，不時嘆氣不時抱頭，對談中經常暗示並透漏自己技術已經不行。或許有冒牌者症候群的困擾。
其實我很想請他放輕鬆點，並對他說：不要給自己太大壓力，你並沒有冒牌，你真的就是那麼爛。

$ ./yogoc song
[這樣上班最爽]
#2 收割默默做事的同事
這世上就是有些人個性特別好，默默做事不會抱怨，這些人就是你收割的對象，是你的韭菜農場。
韭菜要怎麼收割呢？公司總有些雜事或鳥事是大家不想做的，主動接下來！想辦法推給韭菜們：對下說都是韭菜們的努力，對上說都是自己溝通良好。有權力的話，平時多農一些韭菜來放，時候到了直接成立一個原本這些韭菜就在做的事情的團隊，一次收割，大家只會記得是你創立了這個團隊做出貢獻，沒人會理韭菜們
。
不需要受良心譴責，一切只能怪他們為什麼這麼 M 呢。
```