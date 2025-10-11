## Tests
1. go test ./...
2. shortenertest_v2-darwin-arm64 -test.v -test.run=^TestIteration1$ -binary-path=cmd/shortener/shortener
3. shortenertest_v2-darwin-arm64 -test.v -test.run=^TestIteration2$ -source-path=.
4. shortenertest_v2-darwin-arm64 -test.v -test.run=^TestIteration3$ -source-path=.
5. shortenertest_v2-darwin-arm64 -test.v -test.run=^TestIteration4$ -binary-path=cmd/shortener/shortener -server-port=8088

## Misc
- go build -o shortener *.go
- lsof -nP -i4TCP:8888 | grep LISTEN
- curl -X POST -d "url=https://ya.ru" 127.0.0.1:8888 -v