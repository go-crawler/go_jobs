package pipeline

import (
	"sync"

	"github.com/go-crawler/lagou_jobs/model"
)

var (
	mutex sync.Mutex

	jobs []LgJob
)

type LgJob struct {
	// 城市名称
	City string
	// 地区
	District string

	// 公司简称
	CompanyShortName string
	// 公司全称
	CompanyFullName string
	// 公司标签
	CompanyLabelList string
	// 公司规模
	CompanySize string
	// 融资阶段
	FinanceStage string

	// 行业领域
	IndustryField string
	// 行业标签
	IndustryLables string

	// 职位名称
	PositionName string
	// 职位标签
	PositionLables string
	// 职位诱惑
	PositionAdvantage string
	// 工作年限
	WorkYear string
	// 学历要求
	Education string
	// 薪资范畴
	Salary string

	// 经度
	Longitude float64
	// 纬度
	Latitude float64
	// 附近的地铁
	Linestaion string

	// 发布时间
	CreateTime int64
	// 新增时间
	AddTime int64
}

func NewJobPipeline() *LgJob {
	return &LgJob{}
}

func (j *LgJob) Append(js []LgJob) {
	mutex.Lock()
	jobs = append(jobs, js...)
	mutex.Unlock()
}

func (j *LgJob) Get() []LgJob {
	return jobs
}

func (j *LgJob) Push() error {
	for _, v := range j.Get() {
		if err := model.DB.Create(v).Error; err != nil {
			return err
		}
	}

	return nil
}
