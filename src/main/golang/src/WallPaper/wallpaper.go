package WallPaper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// BING API
const (
	BingApiTemplate = "https://global.bing.com/HPImageArchive.aspx?format=js&idx=0&n=9&pid=hp&FORM=BEHPTB&uhd=1&uhdwidth=3840&uhdheight=2160&setmkt=%s&setlang=en"
	BingApi         = "https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=10&nc=1612409408851&pid=hp&FORM=BEHPTB&uhd=1&uhdwidth=3840&uhdheight=2160"
	BingUrl         = "https://cn.bing.com"
)

type BingApiInfo struct {
	Images []struct {
		Startdate     string        `json:"startdate"`
		Fullstartdate string        `json:"fullstartdate"`
		Enddate       string        `json:"enddate"`
		URL           string        `json:"url"`
		Urlbase       string        `json:"urlbase"`
		Copyright     string        `json:"copyright"`
		Copyrightlink string        `json:"copyrightlink"`
		Title         string        `json:"title"`
		Quiz          string        `json:"quiz"`
		Wp            bool          `json:"wp"`
		Hsh           string        `json:"hsh"`
		Drk           int           `json:"drk"`
		Top           int           `json:"top"`
		Bot           int           `json:"bot"`
		Hs            []interface{} `json:"hs"`
	} `json:"images"`
	Tooltips struct {
		Loading  string `json:"loading"`
		Previous string `json:"previous"`
		Next     string `json:"next"`
		Walle    string `json:"walle"`
		Walls    string `json:"walls"`
	} `json:"tooltips"`
}
type WallPaper struct {
}

func initReqClient(url string) (*http.Request, *http.Client, error) {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	req.Header.Add("ect", "4g")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("priority", "u=0, i")
	req.Header.Add("sec-ch-ua-arch", "\"x86\"")
	req.Header.Add("sec-ch-ua-bitness", "\"64\"")
	req.Header.Add("sec-ch-ua-full-version", "\"130.0.6723.116\"")
	req.Header.Add("sec-ch-ua-full-version-list", "\"Chromium\";v=\"130.0.6723.116\", \"Google Chrome\";v=\"130.0.6723.116\", \"Not?A_Brand\";v=\"99.0.0.0\"")
	req.Header.Add("sec-ch-ua-model", "")
	req.Header.Add("sec-ch-ua-platform-version", "\"6.8.0\"")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("Cookie", "MUID=2E3C72F3287C68A0262567D9291269EB; SRCHD=AF; SRCHUID=V; SRCHUSR=DOB; _SS=SID; MUIDB=2E3C72F3287C68A0262567D9291269EB; SRCHHPGUSR=SRCHLANG; _EDGE_S=SID; SNRHOP=TS")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36")
	return req, client, nil
}
func (p *WallPaper) Spider() {
	url := fmt.Sprintf(BingApiTemplate, "zh-CN")
	req, client, err := initReqClient(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	info := BingApiInfo{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, image := range info.Images {
		url = BingUrl + image.URL
		// 下载并保存图片
		err = p.DownloadImage(url, fmt.Sprintf("WallPaper/%s.jpg", image.Enddate))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

// DownloadImage 下载并保存图片
func (p *WallPaper) DownloadImage(url, filePath string) error {
	// 发起 GET 请求下载图片
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	// 创建文件用于保存图片
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// 将图片内容从响应体写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}

	return nil
}
