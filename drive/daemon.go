package drive

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	dae "github.com/sevlyar/go-daemon"
)

type Daemon struct {
	pidName string
	srvAddr string
	engine  *gin.Engine

	instance *Instance
	shutdown func()
}

func (d *Daemon) Add(c *gin.Context) {
	ctx := c.Request.Context()
	k := c.Request.URL.Query().Get("key")
	p := c.Request.URL.Query().Get("path")

	if err := d.instance.Add(ctx, k, p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Key":  k,
			"Path": p,
			"Err":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Operation": "Add",
		"Key":       k,
		"Path":      p,
		"Status":    "OK",
	})
}

func (d *Daemon) Get(c *gin.Context) {

}

func (d *Daemon) List(c *gin.Context) {
	ctx := c.Request.Context()
	r, err := d.instance.List(ctx, "")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Operation": "List",
			"Error":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Result": r,
	})
}

func (d *Daemon) Stop(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Operation": "Close",
		"Timestamp": time.Now(),
		"Status":    "OK",
	})

	go func() {
		time.Sleep(1 * time.Second)
		d.Close()
	}()
}

func (d *Daemon) Close() {
	if d.shutdown != nil {
		d.shutdown()
	}
}

func (d *Daemon) Start() error {
	cntxt := &dae.Context{
		PidFileName: "",
		PidFilePerm: 0644,
		LogFileName: "",
		LogFilePerm: 0644,
		WorkDir:     "/",
		Umask:       027,
		Args:        []string{},
	}

	child, err := cntxt.Reborn()
	if err != nil {
		return err
	}
	if child != nil {
		return nil
	}
	defer cntxt.Release()

	srv := http.Server{
		Addr:    d.srvAddr,
		Handler: d.engine,
	}

	d.shutdown = func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server is forced to shutdown:", err)
		}
		log.Println("Server exiting")
	}

	return srv.ListenAndServe()
}

func NewDaemon(srvAddr, dbAddr string) (*Daemon, error) {
	r := gin.Default()

	ins, err := New(context.Background(), dbAddr)
	if err != nil {
		return nil, err
	}

	d := &Daemon{
		engine:   r,
		instance: ins,
		srvAddr:  srvAddr,
	}

	r.GET("/add", d.Add)

	r.GET("/get", d.Get)

	r.GET("/list", d.List)

	r.GET("/stop", d.Stop)

	return d, nil
}

func httpHandler(handler func(w http.ResponseWriter, r *http.Request)) func(c *gin.Context) {
	return func(c *gin.Context) {
		handler(c.Writer, c.Request)
	}
}
