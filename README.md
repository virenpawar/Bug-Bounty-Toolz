# Bug-Bounty-Toolz
Collection of my made tools used for bug bounty  😌

---
### `forwardx`

- Installation:
```
go get github.com/virenpawar/Bug-Bounty-Toolz/forwardx
```

- Forwards request at provided url in parameter. Usage:
```
Request:
curl 0.0.0.0:88/?url=https://domain.com/path/you/want/to/visit

Response:
HTTP/1.1 302 Found
...
Location: https://domain.com/path/you/want/to/visit
...

<a href="https://domain.com/path/you/want/to/visit">Found</a>.
```

- Use-case:
    1. Use this with the burp's intruder to request multiple urls and read source code of every request
    2. You can filter, sort and even grep any specific words in response.
    3. No more visiting each domain and doing view-source
    4. Can be used to work as [`ffuf`](https://github.com/ffuf/ffuf) and/or [`meg`](https://github.com/tomnomnom/meg)
    5. Using intruder's cluster-bomb, multiple domain's multiple path can be visited and logged at same time.

- Idea credit:
    1. [Zseano's tweet](https://twitter.com/zseano/status/1288873926102310914?s=20) regarding this methodology.
    
- Suggestions and feedbacks are welcomed.