
# Header Guard

**Header Guard** is a middleware plugin for [Traefik](https://traefik.io) that filters HTTP requests based on the value of a specified HTTP header.

---

## 📦 Features

✅ Checks for the presence of a specific HTTP header (default: `X-Auth-Request-Groups`)  
✅ Compares header values against a list of allowed values  
✅ Allows or blocks the request (returns `403 Forbidden` if not allowed)  
✅ Supports custom separators (`|`, `,`, etc.)

---

## ⚙️ Installation

1️⃣ Add the plugin configuration to your `traefik.yml`:

```yaml
experimental:
  plugins:
    headerguard:
      moduleName: github.com/bigouden/headerguard
      version: v0.0.1
```

2️⃣ Configure the middleware in your dynamic configuration file (`dynamic.yml`):

```yaml
http:
  middlewares:
    my-headerguard:
      plugin:
        headerguard:
          header: "X-Auth-Request-Groups"
          allow:
            - "admin"
            - "user"
          separator: "|"
```

3️⃣ Apply the middleware to your routers:

```yaml
http:
  routers:
    my-router:
      rule: "Host(`myapp.example.com`)"
      entryPoints:
        - web
      middlewares:
        - my-headerguard
      service: my-service
```

---

## 🔍 Example Usage

If a request arrives like this:

```http
GET /some/path HTTP/1.1
Host: myapp.example.com
X-Auth-Request-Groups: admin|dev
```

and either `admin` or `dev` is in the list of allowed values, the request is allowed. Otherwise, the plugin returns a `403 Forbidden` response.


---

## 👤 Author

- Thomas GUIRRIEC
- GitHub: [https://github.com/Bigouden](https://github.com/Bigouden)

---

## 📄 License

Distributed under the MIT License. See [LICENSE](./LICENSE) for more details.

---

## 🤝 Contributions

Contributions, issues, and feature requests are welcome!

