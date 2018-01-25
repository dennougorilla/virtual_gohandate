package main

import (
	"math/rand"
	"time"
)

type Geocoding struct {
	Results []Geometry `json:"results"`
}

type Geometry struct {
	GeoRes Location `json:"geometry"`
}

type Location struct {
	Location locations `json:"location"`
}

type locations struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Locs struct {
	Lat float64
	Lng float64
}

type buttonTemp struct {
	key     string
	image   string
	title   string
	label   string
	select1 string
}

type buttonTemp2 struct {
	key     string
	image   string
	title   string
	label   string
	select1 string
	select2 string
}

type buttonTemp4 struct {
	key     string
	title   string
	image   string
	label   string
	select1 string
	select2 string
	select3 string
	select4 string
}

type TalkRes struct {
	text  string
	image string
}

var (
	button  map[string]buttonTemp
	button2 map[string]buttonTemp2
	button4 map[string]buttonTemp4
	talk    map[int]buttonTemp4
	talkres map[string]TalkRes
	word    map[string]string
)

func init() {
	rand.Seed(time.Now().UnixNano())
	//using button template 1 of select or nothing select
	button = map[string]buttonTemp{
		"いただきます！": buttonTemp{
			key:     " ",
			image:   "https://i.imgur.com/97XRjTa.png",
			title:   "いただきます♪",
			label:   " ",
			select1: " ",
		},
		"so-su": buttonTemp{
			key:     "menu1",
			image:   "https://img.retty.me/img_repo/l/01/11709105.jpg",
			title:   "味の里　しおえ",
			label:   "ソースカツ丼",
			select1: "これにする!",
		},
		"miso": buttonTemp{
			key:     "menu2",
			image:   "https://i.imgur.com/9Oam9dS.jpg",
			title:   "味の里　しおえ",
			label:   "味噌チャーシュー",
			select1: "これにする!",
		},
	}
	//using button template 2 of select
	button2 = map[string]buttonTemp2{
		"ご飯行かない？": buttonTemp2{
			key:     "meshi",
			image:   "https://i.imgur.com/tNxL35o.png",
			title:   "しおえってお店行ってみたいなぁ♪",
			label:   " ",
			select1: "行く",
			select2: "行かない",
		},
		"location": buttonTemp2{
			key:     "osusume",
			image:   "https://i.imgur.com/AZ9L8d6.png",
			title:   "待ってたよ〜〜、お腹すいちゃった😂",
			label:   "オススメを聞く？聞かない？",
			select1: "聞く",
			select2: "聞かない",
		},
	}
	//using button template 4 of select
	button4 = map[string]buttonTemp4{
		"ごちそうさま！": buttonTemp4{
			key:     "review",
			title:   "美味しかったね！\n味はどうだった？",
			image:   "https://i.imgur.com/oxoKeI5.png",
			label:   " ",
			select1: "★☆☆☆",
			select2: "★★☆☆",
			select3: "★★★☆",
			select4: "★★★★",
		},
	}
	//talk theme response
	talk = map[int]buttonTemp4{
		0: buttonTemp4{
			key:     "shiro",
			title:   "なんのお話にする？",
			image:   "https://i.imgur.com/iazlG5a.png",
			label:   " ",
			select1: "大阪城？",
			select2: "鶴ヶ城？",
			select3: "名古屋城？",
			select4: "カリオストロの城？",
		},
		1: buttonTemp4{
			key:     "kyodo",
			title:   "なんの話にする？",
			image:   "https://i.imgur.com/iazlG5a.png",
			label:   " ",
			select1: "ちゃんちゃんやき", //北海道の郷土料理だけどあんまわかんないや
			select2: "かにまき汁",    // 宮崎県の郷土料理だけどあんまわかんないや
			select3: "イノシシカレー",  // 山梨県の郷土料理だけどあんまわかんないや
			select4: "こづゆ",
		},
		2: buttonTemp4{
			key:     "men",
			title:   "なんの話にする？",
			image:   "https://i.imgur.com/iazlG5a.png",
			label:   " ",
			select1: "喜多方ラーメン",
			select2: "白河ラーメン",
			select3: "博多ラーメン",
			select4: "札幌ラーメン",
		},
	}

	//various word map
	word = map[string]string{
		"location": "嘘つき！\n全然違う場所じゃない！！",
		"meshi1":   "やったぁ！\n着いたら教えてね♪",
		"meshi2":   "そっかぁ...残念。。。",
		"osusume2": "了解♪",
		"menu1":    "おぉ！良いね♪\n私もソースカツ丼にしよう！",
		"menu2":    "おぉ！良いね♪\n私も味噌チャーシューにしよう！",
		"review":   "また来ようね！",
	}

	//response of talk postback
	talkres = map[string]TalkRes{
		"shiro1": TalkRes{
			text:  "それはあんま興味ないなぁ〜〜...",
			image: "",
		},
		"shiro2": TalkRes{
			text:  "peco 鶴ヶ城には詳しいんだぁ〜！\n**********\n福島県会津若松市追手町にあった日本の城で、地元では鶴ヶ城（つるがじょう）と言うが、同名の城が他にあるため、地元以外では会津若松城と呼ばれることが多い。文献では旧称である黒川城（くろかわじょう）、または単に会津城とされることもある。国の史跡としては、若松城跡（わかまつじょうあと）の名称で指定されている。\n**********",
			image: "https://i.imgur.com/nPejtHV.jpg",
		},
		"shiro3": TalkRes{
			text:  "名古屋城かぁ〜名古屋城はあんまり詳しくないんだぁ〜",
			image: "",
		},
		"shiro4": TalkRes{
			text:  "ルパ〜ン3世...だね！！！",
			image: "",
		},
		"kyodo1": TalkRes{
			text:  "北海道の郷土料理だけどあんまわかんないや",
			image: "",
		},
		"kyodo2": TalkRes{
			text:  "宮崎県の郷土料理だけどあんまわかんないや",
			image: "",
		},
		"kyodo3": TalkRes{
			text:  "山梨県の郷土料理だけどあんまわかんないや",
			image: "",
		},
		"kyodo4": TalkRes{
			text:  "こづゆは知ってるよ！\n**********\n内陸の会津地方でも入手が可能な、海産物の乾物を素材とした汁物である。江戸時代後期から明治初期にかけて会津藩の武家料理や庶民のごちそうとして広まり、現在でも正月や冠婚葬祭などハレの席で、必ず振る舞われる郷土料理である。\n**********",
			image: "https://i.imgur.com/uUWeU5G.jpg",
		},
		"men1": TalkRes{
			text:  "peco 喜多方ラーメン大好きなんだぁ！\n**********\n喜多方ラーメン（きたかたラーメン）とは福島県喜多方市発祥のご当地ラーメン（ご当地グルメ）で、2006年（平成18年）1月の市町村合併前の旧喜多方市では人口37,000人あまりに対し120軒ほどのラーメン店があり、対人口比の店舗数では日本一であった。札幌ラーメン、博多ラーメンと並んで日本三大ラーメンの一つに数えられている。\n**********",
			image: "https://i.imgur.com/w6kws4W.png",
		},
		"men2": TalkRes{
			text:  "喜多方ラーメンの話ししようよ〜！！",
			image: "",
		},
		"men3": TalkRes{
			text:  "喜多方ラーメンの話ししようよ〜！！",
			image: "",
		},
		"men4": TalkRes{
			text:  "喜多方ラーメンの話ししようよ〜！！",
			image: "",
		},
	}

}
