package logentries

import (
    "net"
    "errors"
    "crypto/tls"
    "crypto/x509"
)

type LogentriesLogger struct {
    Token string
}

// Write method to make LogentriesLogger implement the io/Writer interface
// http://golang.org/pkg/io/#Writer
func (w *LogentriesLogger) Write(p []byte) (n int, err error) {
    addr := "data.logentries.com:443"

    rootCerts := x509.NewCertPool()
    success := rootCerts.AppendCertsFromPEM([]byte(pem))

    var c net.Conn

    if success {
        c, err = tls.Dial(
            "tcp",
            addr,
            &tls.Config{
                RootCAs: rootCerts,
            },
        )
        if err != nil {
            return 0, err
        }
    } else {
        return 0, errors.New("Error appending root certificates to CertPool")
    }

    buff := append([]byte(w.Token + " "), p...)
    n, err = c.Write(buff)

    defer c.Close()
    return n, err
}
