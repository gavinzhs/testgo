package main

import (
    "bytes"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "golang.org/x/text/encoding/simplifiedchinese"
    "golang.org/x/text/transform"
    "io"
    "io/ioutil"
    "log"
    "net"
    "os"
    "strings"
    "sync"
    "testing"
    "time"
    "net/http"
    "github.com/go-martini/martini"
    "encoding/xml"
    "code.google.com/p/go-charset/charset")

func TestSessionId(t *testing.T) {
    log.Println("start")
    b := make([]byte, 32)
    if _, err := io.ReadFull(rand.Reader, b); err != nil {
        log.Printf("io.ReadFull err :%+v", err)
        return
    }
    log.Printf("session id :%s", base64.URLEncoding.EncodeToString(b))
}

func TestRangeString(t *testing.T) {
    s := "你好"
    s = "我也不错"
    log.Println(s)
    for i, v := range s {
        log.Printf("第%d个是%s", i, string(v))
    }
}

func TestIp(t *testing.T) {
    log.Println("ip start")
    //    if len(os.Args) != 2 {
    //        fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
    //        os.Exit(1)
    //    }
    //    name := os.Args[1]
    addr := net.ParseIP("192.168.2.47")
    for i := 0; i < len(addr); i++ {
        log.Printf("第%d个是%04x\n", i, addr[i])
    }
    //    if addr == nil {
    //        fmt.Println("Invalid address")
    //    } else {
    //        fmt.Printf("The address is : %v",  addr)
    //    }
    //    os.Exit(0)
}

func TestFile(t *testing.T) {
    //创建file
    //    f, err := os.Create("/Users/zhengsai/Documents/file.txt")
    //    os.NewFile()
    //    f, err := os.Open("/Users/zhengsai/Documents/file.txt")
    f, err := os.OpenFile("/Users/zhengsai/Documents/file1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("create err : %v", err)
    }
    //关闭file
    defer f.Close()
    log.Printf("file name: %s", f.Name())
    //        srcb, _ := ioutil.ReadAll(f)
    //打开并写入file
    src := "郑赛啊"
    //    log.Println([]byte(src))
    //    srcbyte, _ := ioutil.ReadAll(transform.NewReader(strings.NewReader(src), simplifiedchinese.GBK.NewEncoder()))
    //    log.Println([]byte(srcbyte))
    //    log.Printf("src encoder : %s", string(srcbyte))
    //    distbyte, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader(srcbyte), simplifiedchinese.GBK.NewDecoder()))
    //    log.Println([]byte(distbyte))
    //    log.Printf("src decoder dist : %s", string(distbyte))
    //    dist, num, err := transform.String(simplifiedchinese.GBK.NewEncoder(), src)
    //    if err != nil{
    //        log.Fatalf("transform err : %v", err)
    //    }
    //    log.Printf("transform dist: %s, num : %d", dist, num)
    n, err := f.WriteString(src)
    //    n, err := f.WriteAt([]byte(dist), int64(len(srcb)))
    if err != nil {
        log.Fatalf("write err : %v", err)
    }
    log.Printf("write num: %d", n)

    //读取file
    //    srcb, _ := ioutil.ReadAll(f)
    //    aaa := transform.NewReader(f, simplifiedchinese.GBK.NewDecoder())
    //////    b , _ := ioutil.ReadAll(f)
    ////    log.Printf("file content: %s", string(b))
    //
    //    buf := bytes.NewBuffer(make([]byte, 0))
    //    i := 0
    //    for {
    //        oneByte := make([]byte, 1)
    //        n, err := aaa.Read(oneByte)
    //        log.Println(oneByte)
    //        if err == io.EOF{
    //            break
    //        }
    //        if err != nil{
    //            log.Fatalf("read errd : %+v", err)
    //        }
    //        log.Printf("第%d次读取了byte个数: %d", i, n)
    ////        m, err := buf.Read(oneByte)
    ////        if err != nil{
    ////            log.Fatalf("buf errd : %+v", err)
    ////        }
    ////        log.Println(m)
    //        buf.Write(oneByte)
    //        i++
    //    }
    //    log.Println(buf.Bytes())
    //
    //    log.Println("result:" + string(buf.Bytes()))

    //不同编码的读与写

}

func TestEncoding(t *testing.T) {
    src := "编码转换内容"
    fmt.Println([]byte(src)) //byte
    //    data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder()))
    data, _ := ioutil.ReadAll(transform.NewReader(strings.NewReader(src), simplifiedchinese.GBK.NewEncoder()))
    for a := 0; a < len(data); a++ {
        fmt.Printf("%o\n", data[a]) //byte
    }
    fmt.Println(string(data)) //打印为乱码

    gbkData, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(data), simplifiedchinese.HZGB2312.NewDecoder()))
    if err != nil {
        log.Fatalf("transform err : %+v", err)
    }

    fmt.Println(data)            //byte
    fmt.Println(string(gbkData)) //打印为乱码
}

func TestIntToByte(t *testing.T) {
    i := "中国"
    b := []byte(i)
    for a := 0; a < len(b); a++ {
        log.Printf("第%d个是%s", a, ByteToBinaryString(b[a]))
    }
}

func ByteToBinaryString(data byte) (str string) {
    var a byte
    for i := 0; i < 8; i++ {
        a = data
        data <<= 1
        data >>= 1

        switch a {
            case data:
            str += "0"
            default:
            str += "1"
        }

        data <<= 1
    }
    return str
}

func TestDefer(t *testing.T) {
    i := 0
    defer func() {
        fmt.Printf("inner : %d", i)
    }()
    i++
    fmt.Printf("out: %d", i)
}

type MM struct {
    sync.Mutex
}

func TestMutex(t *testing.T) {
    mm := &MM{}
    mm.Lock()
    go func() {
        mm.Lock()
        fmt.Println("wo sleep 1")
        //        time.Sleep(time.Second * 1)
        mm.Unlock()
    }()

    //    time.Sleep(time.Second * 1)

    go func() {
        mm.Lock()
        fmt.Println("wo sleep 2")
        //        time.Sleep(time.Second * 1)
        mm.Unlock()
    }()

    //    time.Sleep(time.Second * 1)

    go func() {
        mm.Lock()
        fmt.Println("wo sleep 3")
        //        time.Sleep(time.Second * 1)
        mm.Unlock()
    }()

    time.Sleep(time.Second * 1)
    mm.Unlock()

    time.Sleep(time.Second * 10)
}

func TestOperator(t *testing.T) {
    str := `go "10"`
    log.Println(string(str))
}

//test 时间单位/utc与local的区别     utc是世界标准时间       local是北京时间      北京时间比世界标准时早8个小时     所以用世界标准时间减去8个小时就是北京时间
//时间单位有秒   毫秒     微秒   纳秒       差1000呗        unix时间为秒     js中返回的时间是毫秒数    好多地方记录的时候使用毫秒     包支持秒以及纳秒   所以用纳秒除以10*1e6   会得到一个整数int64位的
func TestTimeUnit(tt *testing.T) {
    t := time.Now()
    log.Println()
    log.Println(t.UnixNano())
    log.Println(t.Local().Unix())
    log.Println(t.UTC().Unix())
}

func TestMartini(t *testing.T) {
    m := martini.Classic()
    m.Get("/m", testHandler)
    m.Get("/m/main", testMainHandler)
    http.Handle("/m/", m)

    http.ListenAndServe("localhost:3000", nil)
}

type Email struct {
    Where string `xml:"where,attr"`
    Addr  string
}
type Address struct {
    City, State string
}
type Result struct {
    XMLName xml.Name `xml:"Person"`
    Name    string   `xml:"FullName"`
    Phone   string
    Email   []Email
    Groups  []string `xml:"Group>Value"`
    Address
}

func TestDemoXML(t *testing.T) {

    v := Result{Name: "none", Phone: "none"}

    data := `
        <?xml version="1.0" encoding="gb2312"?>
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
	`
    err := xml.Unmarshal([]byte(data), &v)
    if err != nil {
        fmt.Printf("error: %v\n", err)
        return
    }
    fmt.Printf("XMLName: %#v\n", v.XMLName)
    fmt.Printf("Name: %q\n", v.Name)
    fmt.Printf("Phone: %q\n", v.Phone)
    fmt.Printf("Email: %v\n", v.Email)
    fmt.Printf("Groups: %v\n", v.Groups)
    fmt.Printf("Address: %v\n", v.Address)
}





type result struct {
    XMLName xml.Name `xml:"property"`
    Returncode    string   `xml:"returncode"`
    Original string `xml:"original"`
    Key string `xml:"key"`
}

func TestXML(t *testing.T) {
    log.Println("test xml start")
    res, err := http.Get("http://panda.www.net.cn/cgi-bin/check.cgi?area_domain=baidu.com")
    //    res, err := http.Get("http://localhost:3000/m/")
    if err != nil {
        log.Fatalf("get info err : %v", err)
    }

    defer res.Body.Close()

    b, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatalf("read res data err : %v", err)
    }

    var r result

    //    log.Printf("infos : %s", string(b))
    data := strings.Replace(string(b), "gb2312", "utf-8", 1)
    //    data = strings.Replace(data, "200", "你好", 1)
    //    log.Printf("转化后的数据: %s", data)
    //    data = `<?xml version="1.0" encoding="utf-8"?>
    //<property>
    //<returncode>200</returncode>
    //<key>baidu.com</key>
    //<original>211 : Domain exists</original>
    //</property>`
//    err = xml.Unmarshal(b, &r)
        err = xml.Unmarshal([]byte(data), &r)
        decoder := xml.NewDecoder(bytes.NewReader(b))
        decoder.CharsetReader = charset.NewReader
        charset.CharsetDir = "/Users/zhengsai/Documents/go/src/code.google.com/p/go-charset/datafiles"
        err = decoder.Decode(&r)
    if err != nil {
        log.Fatalf("unmarshal data err : %v\n", err)
    }

    log.Printf("r : %+v", r)
    //    log.Printf("charset : %s", procInst("encoding", data))


}

func procInst(param, s string) string {
    param = param + "="
    idx := strings.Index(s, param)
    if idx == -1 {
        return ""
    }
    v := s[idx+len(param):]
    if v == "" {
        return ""
    }
    if v[0] != '\'' && v[0] != '"' {
        return ""
    }
    idx = strings.IndexRune(v[1:], rune(v[0]))
    if idx == -1 {
        return ""
    }
    return v[1 : idx+1]
}

func TestFileServer(t *testing.T){

//    addr := net.TCPAddr{"127.0.0.1", 8080}
//    fmt.Printf("TCP Addr: %s\n", addr)

    log.Println("file server start")
    http.Handle("/", http.FileServer(http.Dir("/Users/zhengsai/Downloads/")))
    http.ListenAndServe(":5000", nil)


}
