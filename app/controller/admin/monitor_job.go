package admin

import (
	"gfast/app/model/admin/sys_job"
	"gfast/app/service/admin/dict_service"
	"gfast/app/service/admin/monitor_service"
	"gfast/app/service/admin/user_service"
	"gfast/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

type MonitorJob struct{}

//任务列表
func (c *MonitorJob) List(r *ghttp.Request) {
	var req *sys_job.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	total, page, list, err := monitor_service.JobListByPage(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	//获取相关选项
	jobStatus, err := dict_service.GetDictWithDataByType("sys_job_status", "", "全部")
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	jobGroup, err := dict_service.GetDictWithDataByType("sys_job_group", "", "全部")
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	result := g.Map{
		"currentPage":  page,
		"total":        total,
		"list":         list,
		"searchStatus": jobStatus,
		"searchGroup":  jobGroup,
	}
	response.SusJson(true, r, "任务列表", result)
}

//添加任务
func (c *MonitorJob) Add(r *ghttp.Request) {
	if r.Method == "POST" {
		var req *sys_job.ReqAdd
		//获取参数
		if err := r.Parse(&req); err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		userId := user_service.GetLoginID(r) //获取登陆用户id
		_, err := monitor_service.AddJob(req, userId)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "任务添加成功")
	}
	//获取相关选项
	jobStatus, err := dict_service.GetDictWithDataByType("sys_job_status", "", "")
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	jobGroup, err := dict_service.GetDictWithDataByType("sys_job_group", "", "")
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	res := g.Map{
		"jobStatus": jobStatus,
		"jobGroup":  jobGroup,
	}
	response.SusJson(true, r, "添加任务", res)
}

//修改任务
func (c *MonitorJob) Edit(r *ghttp.Request) {
	if r.Method == "POST" {
		var req *sys_job.ReqEdit
		//获取参数
		if err := r.Parse(&req); err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		userId := user_service.GetLoginID(r) //获取登陆用户id
		_, err := monitor_service.EditJob(req, userId)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "修改任务成功")
	}
	id := r.GetInt64("id")
	job, err := monitor_service.GetJobInfoById(id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	//获取相关选项
	jobStatus, err := dict_service.GetDictWithDataByType("sys_job_status", gconv.String(job.Status), "")
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	jobGroup, err := dict_service.GetDictWithDataByType("sys_job_group", job.JobGroup, "")
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	res := g.Map{
		"jobStatus": jobStatus,
		"jobGroup":  jobGroup,
		"jobInfo":   job,
	}
	response.SusJson(true, r, "添加任务", res)
}

//删除计划任务
func (c *MonitorJob) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	err := monitor_service.DeleteJobByIds(ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

//启动任务
func (c *MonitorJob) Start(r *ghttp.Request) {
	id := r.GetInt64("id")
	job, err := monitor_service.GetJobInfoById(id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	err = monitor_service.JobStart(job)
	if err != nil {
		response.FailJson(true, r, "定时任务管理启动"+err.Error())
	}
	response.SusJson(true, r, "定时任务管理启动成功")
}

//停止任务
func (c *MonitorJob) Stop(r *ghttp.Request) {
	id := r.GetInt64("id")
	job, err := monitor_service.GetJobInfoById(id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	err = monitor_service.JobStop(job)
	if err != nil {
		response.FailJson(true, r, "定时任务管理停止"+err.Error())
	}
	response.SusJson(true, r, "定时任务管理停止成功")
}
