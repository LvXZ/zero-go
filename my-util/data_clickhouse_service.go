package my_util

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/mailru/go-clickhouse"
)

// @Author: lvxiaozheng
// @Date: 2021/4/6 17:47
// @Description: titan clickhouse bitmap 计算日期数量

// DataClickhouseService 数据处理通知服务类
type DataClickhouseService struct {
	sql      string
	date     string
	idType   int8
	allCount int64
}

func (d *DataClickhouseService) querySyncCount() {
	connect, err := sql.Open("clickhouse", "http://bdp:bdp2020@olap.bilibili.co:80/new_titan")
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		log.Fatal(err)
	}

	rows, err := connect.Query(`
select groupBitmapOr(idx) as allQuantity from (select groupBitmapOrState(mid_idx) as idx from new_titan.ads_prty_entity_vip_mid_profile_full_lables_bitmap where log_date = '2021-04-03' SETTINGS distributed_group_by_no_merge = 1);
`)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var allQuantity int64
		if err := rows.Scan(&allQuantity); err != nil {
			log.Fatal(err)
		}
		log.Printf("allQuantity: %d", allQuantity)
	}

	_ = connect.Close()

}

func (d *DataClickhouseService) queryNormal() {
	connect, err := sql.Open("clickhouse", "http://default:123456@127.0.0.1:8123/new_titan")
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		log.Fatal(err)
	}

	rows, err := connect.Query(`
		SELECT
			tag,
			log_date
		FROM
			ads_prty_entity_buvid_profile_full_lables_local`)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			tag int64
			//buvid_idx           []int64
			log_date time.Time
		)
		if err := rows.Scan(
			&tag,
			//&buvid_idx,
			&log_date,
		); err != nil {
			log.Fatal(err)
		}
		log.Printf("tag: %d, action_time: %s",
			tag, log_date,
		)
	}

	_ = connect.Close()

}

//创建并查询
func (d *DataClickhouseService) createAndQuery() {

	connect, err := sql.Open("clickhouse", "http://default:123456@127.0.0.1:8123/new_titan")
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = connect.Exec(`
		CREATE TABLE IF NOT EXISTS example (
			country_code FixedString(2),
			os_id        UInt8,
			browser_id   UInt8,
			categories   Array(Int16),
			action_day   Date,
			action_time  DateTime
		) engine=Memory
	`)

	if err != nil {
		log.Fatal(err)
	}

	tx, err := connect.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(`
		INSERT INTO example (
			country_code,
			os_id,
			browser_id,
			categories,
			action_day,
			action_time
		) VALUES (
			?, ?, ?, ?, ?, ?
		)`)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 100; i++ {
		if _, err := stmt.Exec(
			"RU",
			10+i,
			100+i,
			clickhouse.Array([]int16{1, 2, 3}),
			clickhouse.Date(time.Now()),
			time.Now(),
		); err != nil {
			log.Fatal(err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	rows, err := connect.Query(`
		SELECT
			country_code,
			os_id,
			browser_id,
			categories,
			action_day,
			action_time
		FROM
			example`)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			country               string
			os, browser           uint8
			categories            []int16
			actionDay, actionTime time.Time
		)
		if err := rows.Scan(
			&country,
			&os,
			&browser,
			&categories,
			&actionDay,
			&actionTime,
		); err != nil {
			log.Fatal(err)
		}
		log.Printf("country: %s, os: %d, browser: %d, categories: %v, action_day: %s, action_time: %s",
			country, os, browser, categories, actionDay, actionTime,
		)
	}

	ctx := context.Background()
	rows, err1 := connect.QueryContext(context.WithValue(ctx, clickhouse.QueryID, "dummy-query-id"), `
		SELECT
			country_code,
			os_id,
			browser_id,
			categories,
			action_day,
			action_time
		FROM
			example`)

	if err1 != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			country               string
			os, browser           uint8
			categories            []int16
			actionDay, actionTime time.Time
		)
		if err := rows.Scan(
			&country,
			&os,
			&browser,
			&categories,
			&actionDay,
			&actionTime,
		); err != nil {
			log.Fatal(err)
		}
		log.Printf("country: %s, os: %d, browser: %d, categories: %v, action_day: %s, action_time: %s",
			country, os, browser, categories, actionDay, actionTime,
		)
	}
	_ = connect.Close()

}
