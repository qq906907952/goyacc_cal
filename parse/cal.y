%{
package parse

%}

%union{
 Op string
 Cal Cal
 val float64
}

%token <Op> '+' '-' '*' '/'
%token '(' ')'

%token <val> NUM
%type <Cal>  term factor expr
%%

end: factor{
	yylex.(*lex).setResult($1)
}

factor:term
|factor '+' term{
	$$ = &Expr{
		Left: $1,
		Op: $2,
		Right: $3,
	}
}
| factor '-' term{
	$$ = &Expr{
		Left: $1,
		Op: $2,
		Right: $3,
	}
}

term: expr
|  term '*' expr{
	$$ = &Expr{
		Left: $1,
		Op: $2,
		Right: $3,
	}
}
| term '/' expr{
	$$ = &Expr{
		Left: $1,
		Op: $2,
		Right: $3,
	}
}

expr: NUM{
	$$ = &Factor{
		Val: $1,
	}
}
 | '(' factor ')'{
	$$ = $2
}


%%

type Cal interface {
	GetVal() float64
}

type Factor struct {
	Val float64
}

func (f *Factor) GetVal() float64 {
	return f.Val
}

type Expr struct {
	Left  Cal
	Right Cal
	Op    string
}

func (e *Expr) GetVal() float64 {
	switch e.Op {
	case "+":
		return e.Left.GetVal() + e.Right.GetVal()
	case "-":
		return e.Left.GetVal() - e.Right.GetVal()
	case "*":
		return e.Left.GetVal() * e.Right.GetVal()
	case "/":
		return e.Left.GetVal() / e.Right.GetVal()
	default:
		panic("unknown op")
	}
}