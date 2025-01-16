# cpu-bound-prime-test

🌯🌯🌯 CPUバウンドな素数判定プログラムのパフォーマンスを向上させる！  

## 実行方法

DevContainerに入り、以下のコマンドを実行してください。  

```shell
go run ./src/main.go

# ---

go build -ldflags="-w -s" -o ./main.out ./src/main.go
chmod +x ./main.out
./main.out
```
