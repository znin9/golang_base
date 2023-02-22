# serveMux & Server

## 总结
```text
http包的服务端主要实现是Server,ServeMux实现了Handler接口(实现了ServeHTTP方法)。

Server中的Handler成员变量可以是自定义的Handler或者是http包提供的一个Handler。
也可以是ServeMux,如果是ServeMux，我们还可以在ServeMux中注册处理函数(HandlerFunc)以及对应的pattern,
pattern与HTTP请求中的URL进行匹配，选择最佳的一个处理函数进行处理本次发起的HTTP请求。

关于HandlerFunc只是一个函数类型(func (ResponseWriter,*Request))。
这个函数类型也实现了Handler接口的ServeHTTP(ResponseWriter,*Request)方法。
并且在该方法中实现了调用自身。这其实就是一个适配器。
func (f HandlerFunc) ServeHTTP(w ResponseWriter,r *Request){
    f(w,r)
}
注意：由于HandlerFunc也是一个Handler。因此也可以直接设置在Server的Handler字段中。
那么此时这个Server进来的所有请求都由这个HandlerFunc进行处理。
```
```text
HTTP底层也是使用的TCP协议。
1.首先创建了一个TCP的监听器,net.Listener。
2.在l.Accept()等待连接的时候，将每一个连接都使用Handler的ServeHTTP方法进行处理。
3.如果Handler是一个HTTP多路复用器(ServeMux),那么就会按照HTTP请求中的URL与多路复用器中保存的pattern进行匹配，选择一个最合适的Handler进行处理。
在多路复用器中存储的关于pattern记录的数据结构是 map[string]MuxEntry,MuxEntry中只有两个字段一个记录pattern,一个记录对应的handler。
```

## Handler
```go
// Handler响应 HTTP 请求。
type Handler interface{
	ServeHTTP(ResponseWriter,*Request)
}
```

## HandlerFunc
```go
// HandlerFunc 类型是一个适配器，允许将普通函数用作 HTTP 处理程序。
// 如果 f 是具有适当签名的函数，则 HandlerFunc(f) 是调用 f 的处理程序。
type HandlerFunc func(ResponseWriter,*Request)

// ServeHTTP 调用 f(w, r)。
func (f HandlerFunc) ServeHTTP(w ResponseWirter,r *Request)
```

##  ServeMux
```go
// ServeMux是一个请求多路复用器，将请求的URL与已注册的pattern进行匹配，并调用对应的处理程序。

// 请求中的URL匹配固定的pattern或具有根的pattern,例如:/favicon.ico,/images/ 。

// 较长的pattern优先匹配,例如有注册的pattern: /images/   /images/abc  如果请求是/images/abc则匹配，如果请求的URL是/images/其他 则匹配/images/
// 注意：pattern=/ 该pattern是有根子树的，因此任何请求的URL,如果都未注册，则都可以进行匹配并执行“/”对应的处理程序。

// 如果一个子树已经被注册并且收到了一个请求，该请求在没有尾部斜杠的情况下命名子树根，
// ServeMux 将该请求重定向到子树根（添加尾部斜杠）。可以使用没有尾部斜杠的单独注册路径来覆盖此行为。
// 例如，注册“/images/”会导致 ServeMux 将对“/images”的请求重定向到“/images/”，除非“/images”已单独注册。
type ServeMux struct{
	// ...
}


// 相关的方法或函数

// NewServeMux 分配并返回一个新的 ServeMux。
func NewServeMux() *ServeMux

// Handle  注册给定pattern的处理程序，如果pattern已经存在则会触发panic
func (mux *ServeMux) Handle(pattern string,handler Handler)

// HandleFunc 给指定的pattern注册处理函数
func (mux *ServeMux) HandleFunc(pattern string,handler func(ResponseWriter,*Request))

// Handler 返回用于给定请求的处理程序，参考 r.Method、r.Host 和 r.URL.Path。
// 它总是返回一个非零处理程序。如果路径不是规范形式，处理程序将是一个内部生成的处理程序，重定向到规范路径。
// 如果主机包含端口，则在匹配处理程序时会忽略它。
func (mux *ServeMux) Handler(r *Request) (h Handler,pattern string)

// ServeHTTP 将请求分派给其模式与请求 URL 最匹配的处理程序。
func (mux *ServeMux) ServeHTTP(w ResponseWriter,r *Request)


// 总结:
// 该结构体ServeMux,是一个请求多路复用器，可以注册Handler或者handler function.
// 该结构体提供的ServeHTTP方法将到达的请求匹配到URL合适的处理程序或者处理函数
```
## Server
```go
type Server struct {
	Addr string
	// Addr可以选择指定要侦听的服务器的TCP地址，格式为“host:port”。如果为空，则使用“：http”（端口80）。
	// 服务名称在 RFC 6335 中定义并由 IANA 分配。
	// 有关地址格式的详细信息，请参见 net.Dial。
	
	Handler Handler
	// 要调用的处理程序，如果为 nil，则为 http.DefaultServeMux
	
	DisableGeneralOptionsHandler bool
	TLSConfig *tls.Config
	ReadTimeout time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout time.Duration
	IdleTimeout time.Duration
	MaxHeaderBytes int
	TLSNextProto map[string]func(*Server,*tls.Conn,Handler)
	ConnState func(net.Conn,ConnState)
	ErrorLog *log.Logger
	BaseContext func(net.Listener) context.Context
	ConnContext func(ctx context.Context,c net.Conn) context.Context
}

// 相关的函数或方法

// Close 立即关闭所有活动的 net.Listeners 和状态 StateNew、StateActive 或 StateIdle 中的任何连接。
// 要正常关机，请使用 Shutdown。
// Close 不会尝试关闭（甚至不知道）任何被劫持的连接，例如 WebSockets。
// 关闭返回关闭服务器的底层监听器返回的任何错误。
func (srv *Server) Close() error

// ListenAndServe 侦听 TCP 网络地址 srv.Addr，然后调用 Serve 来处理传入连接的请求。接受的连接被配置为启用 TCP 保持活动。
// 如果 srv.Addr 为空，则使用“:http”。
// ListenAndServe 总是返回一个非零错误。 Shutdown 或 Close 后，返回的错误是 ErrServerClosed。
func (srv *Server) ListenAndServe() error

// ListenAndServeTLS 侦听 TCP 网络地址 srv.Addr，然后调用 ServeTLS 来处理传入 TLS 连接的请求。接受的连接被配置为启用 TCP 保持活动。
func (srv *Server) ListenAndServeTLS(certFile,keyFile string) error

// RegisterOnShutdown 注册一个函数以在关机时调用。
// 这可用于优雅地关闭经过 ALPN 协议升级或被劫持的连接。此函数应启动特定于协议的正常关机，但不应等待关机完成。
func (srv *Server) RegisterOnShutdown(f func())

// Serve 接受 Listener l 上的传入连接，为每个连接创建一个新的goroutine。
// goroutine 读取请求，然后调用 srv.Handler 来回复它们。
func (srv *Server) Serve(l net.Listener) error

func (srv *Server) ServeTLS(l net.Listener,certFile,keyFile string) error
func (srv *Server) SetKeepAlivesEnabled(v bool)
func (srv *Server) Shutdown(ctx context.Context) error
```

## DefaultServeMux
```go
// DefaultServeMux是Serve使用的默认ServeMux。
var DefaultServeMux = &defaultServeMux
var defaultServeMux ServeMux

// Handle 在 DefaultServeMux 中注册给定模式的处理程序。 ServeMux 的文档解释了模式是如何匹配的。
func Handle(pattern string, handler Handler)
// HandleFunc 在 DefaultServeMux 中注册给定模式的处理函数。 ServeMux 的文档解释了模式是如何匹配的。
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

// ListenAndServe 侦听 TCP 网络地址 addr，然后使用处理程序调用 Serve 来处理传入连接的请求。接受的连接被配置为启用 TCP 保持活动。
// 处理程序通常为 nil，在这种情况下使用 DefaultServeMux。
// ListenAndServe 总是返回一个非零错误。
func ListenAndServe(addr string, handler Handler) error
// ListenAndServeTLS 的行为与 ListenAndServe 相同，只是它需要 HTTPS 连接。此外，必须提供包含服务器证书和匹配私钥的文件。
// 如果证书由证书颁发机构签署，则 certFile 应该是服务器证书、任何中间证书和 CA 证书的串联。
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error


// Serve 在侦听器 l 上接受传入的 HTTP 连接，为每个连接创建一个新的服务goroutine。 服务goroutine 读取请求，然后调用处理程序来回复它们。
// 处理程序通常为 nil，在这种情况下使用 DefaultServeMux。
// 仅当侦听器返回 *tls.Conn 连接并且它们在 TLS Config.NextProtos 中配置为“h2”时，才启用 HTTP/2 支持。
// 服务总是返回一个非零错误。
func Serve(l net.Listener, handler Handler) error
func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)
func ServeFile(w ResponseWriter, r *Request, name string)
func ServeTLS(l net.Listener, handler Handler, certFile, keyFile string) error
```