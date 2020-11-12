package cli

import (
	"context"
)

// Context denotes a context passing to a cli command.
type Context struct {
	ctx   context.Context
	shell *Shell
	args  []string
}

// Context returns the instance of context.Context.
func (c *Context) Context() context.Context {
	return c.ctx
}

// Args returns the arguments of the context.
func (c *Context) Args() []string {
	return c.args
}

// Shell returns the shell instance of the context.
func (c *Context) Shell() *Shell {
	return c.shell
}

// Get retrieves value from the Extra field of the shell.
func (c *Context) Get(key string) (interface{}, bool) {
	return c.shell.Get(key)
}

// GetString gets a string value from the Extra field of the shell.
// If the key does not exist, it will return with an empty string.
func (c *Context) GetString(key string) string {
	v, ok := c.Get(key)
	if !ok {
		return ""
	}
	return v.(string)
}

// GetBool gets a bool value from the Extra field of the shell.
// If the key does not exist, it will return false.
func (c *Context) GetBool(key string) bool {
	v, ok := c.Get(key)
	if !ok {
		return false
	}
	return v.(bool)
}

// GetInt gets a int value from the Extra field of the shell.
// If the key does not exist, it will return with a zero value.
func (c *Context) GetInt(key string) int {
	v, ok := c.Get(key)
	if !ok {
		return 0
	}
	return v.(int)
}
