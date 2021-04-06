package my_util

import "testing"

// @Author: lvxiaozheng
// @Date: 2021/4/6 18:03
// @Description:

func TestCreateAndQueryService(t *testing.T) {
	s := new(DataClickhouseService)
	s.createAndQuery()
}

func TestQueryNormalService(t *testing.T) {
	s := new(DataClickhouseService)
	s.queryNormal()
}

func TestQuerySyncJobCountService(t *testing.T) {
	s := new(DataClickhouseService)
	s.querySyncCount()
}
