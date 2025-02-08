# TITLE
## a
### a
#### a
**bold**


(https://m-takagi.github.io/aosa-ja/aosa.pdf p53)
> Unix のシェルは、ユーザーと OS との間のコマンドによるインターフェイスを提供する。
しかし、シェルはまた、リッチなプログラミング言語でもある。

本質的な部分：AST（ツリー構造）
理論上はASTで全ての言語の処理ができるかも。

時代の主流は、JIT（Go言語でつくるインタプリタ，117）。
Goでシェルを書いてみる
Goの特徴
１、テストツールがしっかりしている
２、gofmtは可読性をあげてくれる。
３、gc(garbage collecter)はmallocの追跡を自動でしてくれる
> 再帰的にASTを評価するtree-walkingインタプリタはおそらく全てのアプローチの中でもっとも遅い。






