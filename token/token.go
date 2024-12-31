package token

type TokenType string //* 別名

type Token struct {
	Type TokenType
	Literal string  //* Literalは文字だから今回はコードの文字全部
 }


 const (
ILLEGAL = "ILLEGAL" //* トークンや文字が道であることを表す
EOF = "EOF"  //* ファイルが終わるところを示す end of file
//* 識別子　リテラル
IDENT = "IDENT" //* add,foobar,x,y 変数とか
INT= "INT" //* 数字

//* 演算子
ASSIGN = "="
PLUS = "+"

//* デリミタ カッコとか

COMMA = ","
SEMICOLON = ";"

LPAREN = "("
RPAREN = ")"
LBRACE = "{"
RBRACE = "}"

//* キーワード 言語組み込み機能
FUNCTION = "FUNCTION"
LET = "LET"



 )