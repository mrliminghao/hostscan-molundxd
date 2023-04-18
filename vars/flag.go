package vars

import "flag"

var (
	Version  = flag.Bool("v", false, "Show hostscan version")
	Host     = flag.String("d", "tst.qq.com", "Host to test")
	Ip       = flag.String("i", "", "Nginx IP")
	Timeout  = flag.Int("t", 3, "Timeout for Http connection.")
	Thread   = flag.Int("T", 100, "Thread for Http connection.")
	HostFile = flag.String("D", "", "Hosts in file to test")
	IpFile   = flag.String("I", "", "Nginx Ip in file to test")
	OutFile  = flag.String("O", "result.txt", "Output File")
	Flag     = flag.String("x", "contact O2000", "Host返回的内容包含该字符串则认为存在漏洞")
	IsRandUA = flag.Bool("U", false, "Open to send random UserAgent to avoid bot detection.")
)
