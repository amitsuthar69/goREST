### A Go REST API which reads and writes to a `json` file.

```
 goREST
├── bin
│   └── goREST (executable file)
├── cmd
│   ├── apis.go (route handlers)
│   ├── fileio.go (IO handlers)
│   ├── main.go (main Go file)
│   └── models.go (types)
└── db
    └── posts.json (fake NoSQL db)
```

---

**To run the server, use the command :**

```bash
make run
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
