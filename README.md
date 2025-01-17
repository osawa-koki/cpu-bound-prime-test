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

```result
----- ----- ----- ----- -----
SequentialPrimeChecker
From 1 to 1000000
Prime count: 78498
Execution time: 18.429458ms
----- ----- ----- ----- -----
----- ----- ----- ----- -----
GoroutinePrimeChecker
From 1 to 1000000
Prime count: 78498
Execution time: 7.228208ms
----- ----- ----- ----- -----
```

## おおさわメモ

- 並列処理と並行処理
  - 並列処理 (Parallel)
    - 複数のタスクを物理的に同時に実行する。
    - 複数のCPUコアやプロセッサを利用する。
    - 複数の処理を同時に実行すること。
    - ハードウェアの特性を利用する。
  - 並行処理 (Concurrent)
    - 複数のタスクを論理的に同時に実行する。
    - 1つのCPUがタスクリストを迅速に切り替える。
    - 複数の処理を独立に実行できる構成のこと。
    - プログラミングパターンとして。
- Goroutine vs スレッド
  - Goroutine
    - Goランタイムによって管理される。 (ユーザースレッド)
    - 生成コストが小さい。
    - Goランタイムが協調的スケジューリングを使用する。 (M:Nスケジューリングモデル)
      - Goランタイムが大量のユーザスレッド(Goroutine)を少数のカーネルスレッド上で管理。
  - スレッド
    - OSによって管理される。 (カーネルスレッド)
    - 生成コストが大きい。
    - OSがプリエンプティブスケジューリングを使用する。
