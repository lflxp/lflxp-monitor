package main

import (
	"flag"

	"github.com/lflxp/lflxp-monitor/pkg"
)

func init() {
	flag.String("i", "1", "时间间隔 默认1秒")
	flag.Int("C", 0, "运行时间 默认无限")
	flag.Bool("t", false, "打印当前时间")
	flag.Bool("nocolor", false, "不显示颜色")
	flag.Bool("l", false, "打印Load info")
	flag.Bool("c", false, "打印Cpu info")
	flag.Bool("s", false, "打印swap info")
	flag.Bool("d", false, "打印Disk info")
	flag.Bool("n", false, "打印net info")
	flag.Bool("slave", false, "打印Slave info")
	flag.String("u", "root", "mysql用户名")
	flag.String("p", "system", "mysql密码")
	flag.String("H", "127.0.0.1", "Mysql连接主机，默认127.0.0.1")
	flag.String("P", "3306", "Mysql连接端口,默认3306")
	flag.String("S", "/tmp/mysql.sock", "mysql socket连接文件地址")
	flag.Bool("com", false, "Print MySQL Status(Com_select,Com_insert,Com_update,Com_delete).")
	flag.Bool("hit", false, "Print Innodb Hit%.")
	flag.Bool("innodb_rows", false, "Print Innodb Rows Status(Innodb_rows_inserted/updated/deleted/read).")
	flag.Bool("innodb_pages", false, "Print Innodb Buffer Pool Pages Status(Innodb_buffer_pool_pages_data/free/dirty/flushed)")
	flag.Bool("innodb_data", false, "Print Innodb Data Status(Innodb_data_reads/writes/read/written)")
	flag.Bool("innodb_log", false, "Print Innodb Log  Status(Innodb_os_log_fsyncs/written)")
	flag.Bool("innodb_status", false, "Print Innodb Status from Command: 'Show Engine Innodb Status'")
	flag.Bool("T", false, "Print Threads Status(Threads_running,Threads_connected,Threads_created,Threads_cached).")
	flag.Bool("rt", false, "Print MySQL DB RT(us).")
	flag.Bool("B", false, "Print Bytes received from/send to MySQL(Bytes_received,Bytes_sent).")
	flag.Bool("mysql", false, "Print MySQLInfo (include -t,-com,-hit,-T,-B).")
	flag.Bool("innodb", false, "Print InnodbInfo(include -t,-innodb_pages,-innodb_data,-innodb_log,-innodb_status)")
	flag.Bool("sys", false, "Print SysInfo   (include -t,-l,-c,-s).")
	flag.Bool("lazy", false, "Print Info  (include -t,-l,-c,-s,-com,-hit).")
	flag.Bool("semi", false, "半同步监控")
	flag.String("L", "none", "Print to Logfile.")
	flag.Bool("logfile_by_day", false, "One day a logfile,the suffix of logfile is 'yyyy-mm-dd';")

	flag.Parse()
}

func main() {
	pkg.Run("monitor -lazy -d -n")
}
