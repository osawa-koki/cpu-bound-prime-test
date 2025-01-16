# cpu-bound-prime-test

🌯🌯🌯 CPUバウンドな素数判定プログラムのパフォーマンスを向上させる！  

## 実行方法

DevContainerに入り、以下のコマンドを実行してください。  

```shell
go run $(find ./src/ -name "*.go" -not -name "*_test.go")

# ---

go build -ldflags="-w -s" -o ./main.out $(find ./src/ -name "*.go" -not -name "*_test.go")
chmod +x ./main.out
./main.out
```
