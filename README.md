### A Go REST API which reads and writes to a `json` file.

```
 goREST
├── cmd
│   ├── apis.go (route handlers)
│   ├── fileio.go (IO handlers)
│   ├── main.go (main Go file)
│   └── models.go (types)
├── db
│    └── posts.json (fake NoSQL db)
└── tmp
     └── main (executable file)
```

---

**I'm using [air](https://github.com/cosmtrek/air) package for live server reloading. Check the `.air.toml` config.**

**To run the server, use the command in the root of project :**

```bash
air
```

---

**Use Postman or curl :**

```bash
curl "http://localhost:3000/posts"
```

```bash
curl http://localhost:3000/posts \
--include \
--request "POST" \
--data '{"id":"69", "title":"your_title","authorName":"name"}'
```

```bash
curl http://localhost:3000/posts/69 \
--include \
--request "PUT" \
--data '{"id":"6969", "title":"new-title","authorName":"new-name"}'
```

```bash
curl http://localhost:3000/posts/69 \
--include \
--request "DELETE" \
--data '{"id":"6969", "title":"new-title","authorName":"new-name"}'
```
