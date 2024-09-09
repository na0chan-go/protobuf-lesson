# protobuf-lesson

## protoc コマンド

`protoc`コマンドは、Protocol Buffers（protobuf）ファイルをコンパイルするために使用されます。このコマンドを使用することで、.proto ファイルから言語特有のコードを生成できます。

### 使用方法

```bash
protoc -I=インポートするディレクトリ --go_out=. ファイル名.proto
```

### 主なオプション

- `--go_out=出力先`: Go 用のコードを生成します。
- `-I ディレクトリ`: インポートするディレクトリを指定します。

### 例

```bash
protoc -I. --go_out=. proto/employee.proto proto/date.proto
```
