// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "context"
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "github.com/apache/thrift/lib/go/thrift"
        "zgobandRPC"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  i8 putChess(string player1, string player2, i32 deskID, i8 seatID, i8 row, i8 column)")
  fmt.Fprintln(os.Stderr, "  bool takeBackReq(string account, string otherSide, i8 seatID)")
  fmt.Fprintln(os.Stderr, "  bool takeBackRespond(string player1, string player2, i8 seatID, bool resp)")
  fmt.Fprintln(os.Stderr, "  void loseReq(string player1, string player2, i32 deskID, i8 seatID)")
  fmt.Fprintln(os.Stderr, "  void drawReq(string account, string otherSide, i8 seatID)")
  fmt.Fprintln(os.Stderr, "  void drawResponse(string player1, string player2, i32 deskID, i8 seatID, bool resp)")
  fmt.Fprintln(os.Stderr, "  void sendChatText(string toAccount, string account, string text)")
  fmt.Fprintln(os.Stderr, "  void saveGame(string account)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := zgobandRPC.NewGameOperatorClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "putChess":
    if flag.NArg() - 1 != 6 {
      fmt.Fprintln(os.Stderr, "PutChess requires 6 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    tmp2, err67 := (strconv.Atoi(flag.Arg(3)))
    if err67 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    tmp3, err68 := (strconv.Atoi(flag.Arg(4)))
    if err68 != nil {
      Usage()
      return
    }
    argvalue3 := int8(tmp3)
    value3 := argvalue3
    tmp4, err69 := (strconv.Atoi(flag.Arg(5)))
    if err69 != nil {
      Usage()
      return
    }
    argvalue4 := int8(tmp4)
    value4 := argvalue4
    tmp5, err70 := (strconv.Atoi(flag.Arg(6)))
    if err70 != nil {
      Usage()
      return
    }
    argvalue5 := int8(tmp5)
    value5 := argvalue5
    fmt.Print(client.PutChess(context.Background(), value0, value1, value2, value3, value4, value5))
    fmt.Print("\n")
    break
  case "takeBackReq":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "TakeBackReq requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    tmp2, err73 := (strconv.Atoi(flag.Arg(3)))
    if err73 != nil {
      Usage()
      return
    }
    argvalue2 := int8(tmp2)
    value2 := argvalue2
    fmt.Print(client.TakeBackReq(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "takeBackRespond":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "TakeBackRespond requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    tmp2, err76 := (strconv.Atoi(flag.Arg(3)))
    if err76 != nil {
      Usage()
      return
    }
    argvalue2 := int8(tmp2)
    value2 := argvalue2
    argvalue3 := flag.Arg(4) == "true"
    value3 := argvalue3
    fmt.Print(client.TakeBackRespond(context.Background(), value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "loseReq":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "LoseReq requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    tmp2, err80 := (strconv.Atoi(flag.Arg(3)))
    if err80 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    tmp3, err81 := (strconv.Atoi(flag.Arg(4)))
    if err81 != nil {
      Usage()
      return
    }
    argvalue3 := int8(tmp3)
    value3 := argvalue3
    fmt.Print(client.LoseReq(context.Background(), value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "drawReq":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "DrawReq requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    tmp2, err84 := (strconv.Atoi(flag.Arg(3)))
    if err84 != nil {
      Usage()
      return
    }
    argvalue2 := int8(tmp2)
    value2 := argvalue2
    fmt.Print(client.DrawReq(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "drawResponse":
    if flag.NArg() - 1 != 5 {
      fmt.Fprintln(os.Stderr, "DrawResponse requires 5 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    tmp2, err87 := (strconv.Atoi(flag.Arg(3)))
    if err87 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    tmp3, err88 := (strconv.Atoi(flag.Arg(4)))
    if err88 != nil {
      Usage()
      return
    }
    argvalue3 := int8(tmp3)
    value3 := argvalue3
    argvalue4 := flag.Arg(5) == "true"
    value4 := argvalue4
    fmt.Print(client.DrawResponse(context.Background(), value0, value1, value2, value3, value4))
    fmt.Print("\n")
    break
  case "sendChatText":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "SendChatText requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.SendChatText(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "saveGame":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SaveGame requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.SaveGame(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
