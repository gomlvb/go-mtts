package main

import (
	"fmt"
	"go-mtts/common/mysql"
)

type video_repair_t struct {
	client     string
	video_name string
}

func main() {

	mysqlConf := mysql.MysqlConfig{
		"root", "zxcvbnm,./", "127.0.0.1", 3306, "dr_vrs_data", 12, 10, true,
	}

	mysql.Init(mysqlConf)

	db := mysql.GetConn()
	defer mysql.Close()

	results := make([]video_repair_t, 0)
	rows, _ := db.Raw("select client, video_name from t_video_repair_info").Rows()
	defer rows.Close()
	for rows.Next() {
		var vrt video_repair_t
		rows.Scan(&vrt.client, &vrt.video_name)
		results = append(results, vrt)
		fmt.Printf("vrt.client: %s, vrt.video_name: %s\n", vrt.client, vrt.video_name)
	}

	err := db.Exec("replace into dr_vrs_data.t_video_repair_info(client,video_name,original_filename,src_ip,src_oss_path,des_oss_path,repair_status,oper_type,video_format,encoder,resolution,size,duration,insert_mq_at,get_mq_at,repair_complete_at,create_at,update_at) values ('1234567888','c4ca4238a0b923820dcc509a6f75849b.mp4','1.mp4','192.168.1.118','repairit/20200831/1234567888/c4ca4238a0b923820dcc509a6f75849b.mp4','',-1,-1,'','','',0,0,now(),now(),now(),now(),now());").Error
	if err != nil {
		fmt.Println("db.Table err:", err)
	}
}
