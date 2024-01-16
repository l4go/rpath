# l4go/rpath ライブラリ

ディレクトリを区別して、http URLのpath部用のパス操作ライブラリです。  
要は、末尾に`/`がある場合はディレクトリとして扱い、勝手に末尾の`/`を消さないようにしつつ、
pathモジュールと同じような操作を提供するライブラリです。

## 末尾の`/`の問題

URLのpath部の文字列を操作するためのライブラリです。  
以下は、RFC3986(URIのRFC)からの引用ですが、この引用の`/over/there`の部分がpath部です。

```
 URI         = scheme ":" hier-part [ "?" query ] [ "#" fragment ]

 hier-part   = "//" authority path-abempty
             / path-absolute
             / path-rootless
             / path-empty

   foo://example.com:8042/over/there?name=ferret#nose
   \_/   \______________/\_________/ \_________/ \__/
    |           |            |            |        |
 scheme     authority       path        query   fragment
    |   _____________________|__
   / \ /                        \
   urn:example:animal:ferret:nose
```

このURLのpath部を厳密に扱う場合、以下の2つは、異なるURLとして扱う必要があります。
(実際に、同じページを表示するか、違うページを表示するかはサーバ次第です。)

- `http://example.com/foo/bar`
- `http://example.com/foo/bar/`

これらから、path部だけ取り出したのが以下のものです。

- `/foo/bar`
- `/foo/bar/`

この2つのパスを意味する文字列に対して、
Go標準のpathパッケージの`path.Clear("/foo/bar")`を実行すると、末尾の`/`が消えて、
どちらのパスも`/foo/bar`になってしまい、区別がつかなくなるのです。

URLのpath部を厳密扱うコードを書こうとした場合、標準のpathパッケージはかなり使いづらいので、
そのために作られたのが、このライブラリです。

# 関数の仕様

pathライブラリーにない独自関数と、pathライブラリー類似の関数に分類して説明します。

## pathライブラリーにない独自関数

### `func IsDir(p string) bool`

ディレクトリを示すパスかどうかを判定します。  
文字列の末尾が`/`であるかを判断します。

### `func SetFile(p string) string`

ファイルを示すパスに変更します。
文字列の末尾が`/`があれば消します。

### `func SetDir(p string) string`

ディレクトリを示すパスに変更します。
文字列の末尾が`/`が無ければ、足します。

### `func SetType(p string, is_dir bool) string`

`is_dir`がtureなら`SetDir`の動作を、`is_dir`がfalseなら`SetFile`の動作をします。

### `func Dir(p string) string`

末尾の`/`を消さないように、 標準ライブラリの`path.Clean()`の動作をします。

## pathライブラリ類似関数

### `func Base(p string) string`

末尾の`/`を消さないように、 標準ライブラリの`path.Base()`の動作をします。

### `func Clean(p string) string`

末尾の`/`を消さないように、 標準ライブラリの`path.Clean()`の動作をします。

### `func Ext(p string) string`

末尾に`/`があるときは空文字(`""`)を返します。それ以外では、 標準ライブラリの`path.Ext()`の動作をします。

### `func Split(p string) (string, string)`

末尾の`/`を消さないように、 標準ライブラリの`path.Split()`の動作をします。

### `func Match(pat, n string) (bool, error)`

`Clean()`を実行した上で、 標準ライブラリの`path.Match()`の動作をします。

### `func Join(pe ...string) string`

末尾の`/`を消さないように、 標準ライブラリの`path.Join()`の動作をします。
