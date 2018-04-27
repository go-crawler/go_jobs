package convert

import (
	"strconv"
	"strings"

	"github.com/go-crawler/lagou_jobs/downloader"
	"github.com/go-crawler/lagou_jobs/pipeline"
	"time"
)

func ToPipelineJobs(dJobs []downloader.Result) []pipeline.LgJob {
	var pJobs []pipeline.LgJob
	for _, v := range dJobs {
		longitude, _ := strconv.ParseFloat(v.Longitude, 64)
		latitude, _ := strconv.ParseFloat(v.Latitude, 64)
		pJobs = append(pJobs, pipeline.LgJob{
			City:     v.City,
			District: v.District,

			CompanyShortName: v.CompanyShortName,
			CompanyFullName:  v.CompanyFullName,
			CompanyLabelList: strings.Join(v.CompanyLabelList, ","),
			CompanySize:      v.CompanySize,
			FinanceStage:     v.FinanceStage,

			PositionName:      v.PositionName,
			PositionLables:    strings.Join(v.PositionLables, ","),
			PositionAdvantage: v.PositionAdvantage,
			WorkYear:          v.WorkYear,
			Education:         v.Education,
			Salary:            v.Salary,

			IndustryField:  v.IndustryField,
			IndustryLables: strings.Join(v.IndustryLables, ","),

			Longitude:  longitude,
			Latitude:   latitude,
			Linestaion: v.Linestaion,

			CreateTime: MustDateToUnix(v.CreateTime),
			AddTime:    time.Now().Unix(),
		})
	}

	return pJobs
}
