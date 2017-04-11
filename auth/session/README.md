# Auth/Session for Go-Kit

Provides abstraction from **gorilla/sessions** for **go-kit**. The sessions will be available through context. It means, everytime we need to use the session we just need to get the value from the current context.
The session itself stored to the context using the **http.ServerBefore** from transportation module on **go-kit**

## Usage Examples:

```
import (
	"github.com/IndonesiaX/go-kit-toolkit/auth/session"
)

func (s *srv) SetSession(ctx context.Context) error {
	session := ctx.Value("session").(session.Session)
	session.Set("name", "IndonesiaX")
	session.Save()
	return nil
}

func (s *srv) GetSession(ctx context.Context) string {
	session := ctx.Value("session").(session.Session)
	return session.Get("name").(string)
}
```

The session is store by the key of "session" on context, we ain't provide the custom name for the moment.
`session.Save()` command was not actually save anything to the store when it executed, the real save job is on the transportation layer **http.ServerAfter**.

## How to use

- Create a new store, using cookiestore from gorilla (other store will be covered soon)
	```
	var store = sessions.NewCookieStore([]byte("session_store"))
	```
- Add new options on your http handler
	```
	import (
			...
			"github.com/IndonesiaX/go-kit-toolkit/auth/session"
			gokithttp "github.com/go-kit/kit/transport/http"
			...
	)

	options := []gokithttp.ServerOption{
		gokithttp.ServerErrorLogger(logger),
		gokithttp.ServerErrorEncoder(encodeError),
		gokithttp.ServerBefore(session.ToHTTPContext(store, "your key name")),
		gokithttp.ServerAfter(session.FromHTTPContext(logger)),
	}
	```

With the two above setup, you are ready to use the session from your service.

## TODO
- [ ] Add Inline Code Documentation
- [ ] Real examples of implementation
- [ ] More store implemenatation, such as: redis, memcached, mysql, etc
