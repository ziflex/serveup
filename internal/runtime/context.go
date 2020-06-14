package runtime

type (
	Execution struct {
		Result string
		Error  string
	}

	Context struct {
		Env  interface{}
		Req  interface{}
		Exec Execution
	}
)

func NewContext(env interface{}, req interface{}) Context {
	return Context{
		Env:  env,
		Req:  req,
		Exec: Execution{},
	}
}
