package logrus_scribe

import (
    "fmt"
    "io/ioutil"
    "net"
    "os"

    "github.com/Sirupsen/logrus"
    "github.com/samuel/go-thrift/examples/scribe"
    "github.com/samuel/go-thrift/thrift"
)

// ScribeHook to send logs via scribe.
type ScribeHook struct {
    Writer                 scribe.ScribeClient
    LogStreamName          string
    DisableStdoutStderr    bool
}

// Creates a hook to be added to an instance of logger.
func NewScribeHook(logStreamName string, DisableStdoutStderr bool, ScribeIpAndPort string) (*ScribeHook, error) {
    Writer, err := InitializeScribeConnection(ScribeIpAndPort)
    if DisableStdoutStderr {
        logrus.SetOutput(ioutil.Discard)
    }

    return &ScribeHook{Writer, logStreamName, DisableStdoutStderr}, err
}

func (hook *ScribeHook) Fire(entry *logrus.Entry) error {
    line, err := entry.String()

    if err != nil {
        if hook.DisableStdoutStderr == false {
            fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
        }
        return err
    }

    _, err = hook.Writer.Log([]*scribe.LogEntry{{hook.LogStreamName, line}})
    return err
}

func (hook *ScribeHook) Levels() []logrus.Level {
    return logrus.AllLevels
}


func InitializeScribeConnection(scribeIpAndPort string) (scribe.ScribeClient, error) {
    conn, err := net.Dial("tcp", scribeIpAndPort)
    t := thrift.NewTransport(thrift.NewFramedReadWriteCloser(conn, 0), thrift.BinaryProtocol)
    client := thrift.NewClient(t, false)

    return scribe.ScribeClient{Client: client}, err
}

