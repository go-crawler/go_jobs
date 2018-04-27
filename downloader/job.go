package downloader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"fmt"
	"github.com/go-crawler/lagou_jobs/fake"
	"github.com/go-crawler/lagou_jobs/pkg/uuid"
	"net/url"
	"strings"
)

var (
	jobsApiUrl = "https://www.lagou.com/jobs/positionAjax.json?city=%s&needAddtionalResult=false"
)

type ListResult struct {
	Code    int
	Success bool
	Msg     string
	Content Content
}

type Content struct {
	PositionResult PositionResult
	PageNo         int
	PageSize       int
}

type PositionResult struct {
	Result     []Result
	TotalCount int
}

type Result struct {
	City              string
	BusinessZones     []string
	CompanyFullName   string
	CompanyLabelList  []string
	CompanyShortName  string
	CompanySize       string
	CreateTime        string
	District          string
	Education         string
	FinanceStage      string
	FirstType         string
	IndustryField     string
	IndustryLables    []string
	JobNature         string
	Latitude          string
	Longitude         string
	PositionAdvantage string
	PositionId        int32
	PositionLables    []string
	PositionName      string
	Salary            string
	SecondType        string
	Stationname       string
	Subwayline        string
	Linestaion        string
	WorkYear          string
}

type jobService struct {
	City string
}

func NewJobService(city string) *jobService {
	return &jobService{City: city}
}

func (l *jobService) GetUrl() string {
	req := fmt.Sprintf(jobsApiUrl, l.City)
	url, _ := url.Parse(req)
	query := url.Query()
	url.RawQuery = query.Encode()

	return url.String()
}

func (l *jobService) GetJobs(pn int, kd string) (*ListResult, error) {
	//client := fake.ProxyAuth{License: "", SecretKey: ""}.GetProxyClient()
	client := http.Client{}
	postReader := strings.NewReader(fmt.Sprintf("first=false&pn=%d&kd=%s", pn, kd))
	req, err := http.NewRequest("POST", l.GetUrl(), postReader)
	if err != nil {
		log.Printf("http.NewRequest err: %v", err)
	}

	//req.Header.Set("Proxy-Switch-Ip", "yes")

	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Languag", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Length", "25")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Cookie", "_ga=GA1.2.161331334.1522592243; "+
		"user_trace_token=20180401221723-"+uuid.GetUUID()+"; "+
		"LGUID=20180401221723-"+uuid.GetUUID()+"; "+
		"index_location_city=%E6%B7%B1%E5%9C%B3; "+
		"JSESSIONID="+uuid.GetUUID()+"; "+
		"_gid=GA1.2.1140631185.1523090450; "+
		"Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1522592243,1523090450; "+
		"TG-TRACK-CODE=index_search; _gat=1; "+
		"LGSID=20180407221340-"+uuid.GetUUID()+"; "+
		"PRE_UTM=; PRE_HOST=; PRE_SITE=https%3A%2F%2Fwww.lagou.com%2F; "+
		"PRE_LAND=https%3A%2F%2Fwww.lagou.com%2Fjobs%2Flist_golang%3FlabelWords%3D%26fromSearch%3Dtrue%26suginput%3D; "+
		"Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1523110425; "+
		"LGRID=20180407221344-"+uuid.GetUUID()+"; "+
		"SEARCH_ID="+uuid.GetUUID()+"")
	req.Header.Add("Host", "www.lagou.com")
	req.Header.Add("Origin", "https://www.lagou.com")
	req.Header.Add("Referer", "https://www.lagou.com/jobs/list_golang?labelWords=&fromSearch=true&suginput=")
	req.Header.Add("User-Agent", fake.GetUserAgent())

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var results ListResult
	err = json.Unmarshal([]byte(body), &results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}
