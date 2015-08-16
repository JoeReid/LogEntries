LogEntries
==========

Overview
--------
A Golang Log-Entries library that supports encrypted connections to send logs.


Usage
-----
The LogentriesLogger implements the [io/Writer interface](http://golang.org/pkg/io/#Writer).
The simplest way to use the library is to use that Writer interface.

```go
writer := logentries.LogentriesLogger{
    "Your Logentries Token Here",
    logentries.ENCRYPTED
}

_, err := writer.Write([]byte("Time to write\n"))
if err != nil {
    panic(err)
}
```

Alternatively you could provide the Writer interface to the standard [log library](http://golang.org/pkg/log/).
The library then affords all the extended functionality that golang's log library gives you.

```go
writer := logentries.LogentriesLogger{
    "Your Logentries Token Here",
    logentries.ENCRYPTED
}

logger := log.New(&writer, "", log.LstdFlags)

logger.Println("Hello from logger")
logger.Fatal("Goodbye from logger")

```

Unencrypted logs
----------------
The sending of unencrypted logs (transmited over plain tcp) is supported. To
enable this behaviour, initialise the logentries.LogentriesLogger struct with
the value `logentries.PLAINTEXT`

```go
writer := logentries.LogentriesLogger{
    "Your Logentries Token Here",
    logentries.PLAINTEXT
}
```

Code Attribution
----------------
The const pem in the [pem.go file](https://github.com/JoeReid/LogEntries/blob/master/pem.go)
was lifted from Mirco Zeiss' [papertrail project](https://github.com/zemirco/papertrail/blob/master/pem.go)
