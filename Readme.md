<p align="center">
  <img src="https://github.com/alikrugl/snippetbox/blob/main/snippetbox.png?raw=true" width="100" />
</p>
<p align="center">
    <h1 align="center">SNIPPETBOX</h1>
</p>

A simple web application built while following the [Let's Go](https://lets-go.alexedwards.net/) book by Alex Edwards. This project is designed to help you get familiar with building web servers in Go, covering key concepts such as routing, templating, database interaction, authentication and authorization, HTTPS, and testing using Go's standard library.

### ⚙️ Some technical details
- [Slog package](https://pkg.go.dev/log/slog) is used to add structured logging.
- [Command-line flags](https://pkg.go.dev/flag) are used to manage configuration settings.
  ```go
  go run ./cmd/web -addr=":80"
  ```
- Dynamic HTML templates with [html/template](https://pkg.go.dev/html/template). In-memory caching is used to read and parse the relevant template files only once.

- **Content-Security-Policy** headers are set to prevent a variety of cross-site scripting, clickjacking, and other code-injection attacks.

- The session manager [alexedwards/scs](https://pkg.go.dev/github.com/alexedwards/scs/v2@v2.8.0) was considered as its allows to store session data server-side and comparing to [gorilla/sessions](https://pkg.go.dev/github.com/gorilla/sessions) it was a better choice due to the ability to renew session IDs.

- [justinas/nosurf](https://pkg.go.dev/github.com/justinas/nosurf) is used to mitigate the risk of CSRF

- [justinas/alice](https://github.com/justinas/alice) package is used to manage middleware/handler chains and make them composable.




### 🛠 Project
  Overview:
  <img src="https://github.com/alikrugl/snippetbox/blob/main/snippetOverview.png?raw=true" />

  Create form:
  <img src="https://github.com/alikrugl/snippetbox/blob/main/newSnippet.png?raw=true" />

  Snippet show page after creation:
  <img src="https://github.com/alikrugl/snippetbox/blob/main/createdSnippet.png?raw=true" />