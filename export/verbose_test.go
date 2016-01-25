package export

import (
	"testing"

	"github.com/robertkrimen/otto/parser"
	"github.com/robertkrimen/otto/walk"
	"strings"
)

func displayString(s string) string {
	r := strings.NewReplacer(" ", "_", "\n", "\\n", "\t", "+++")
	return r.Replace(s)
}

func TestVerbose(t *testing.T) {
	tests := []struct {
		src      string
		expected string
	}{
		{`c = a + b - (1 + 2)`, `c = ((a + b) - (1 + 2));
`},
		{`function fun(a) {alert("hej");c =1 + 2}`, `function fun(a) {
   alert("hej");
   c = (1 + 2);
}
`},
		{`for(i = 0 ; i < 10 ; i++) {a = b + 2}`, `for(i = 0 ; (i < 10) ; (i)++) {
   a = (b + 2);
}
`},
		{`while(true) {a = b + 2}`, `while(true) {
   a = (b + 2);
}
`},
		{`a = [1, 2, 3]`, `a = [1, 2, 3];
`},
		{`a[0] + a[1]`, `(a[0] + a[1]);
`},
		{`fun(a)`, `fun(a);
`},
		{`a = (b ? 1 : 2)`, `a = (b ? 1 : 2);
`},
		{`debugger`, `debugger`},
		{`a.b.c`, `a.b.c;
`},
		{`do{a++}while(true)`, `do {
   (a)++;
} while(true);
`},
		{`for(a in b) {a = 1}`, `for(a in b) {
   a = 1;
}
`},
		{`if(true){a++}else{b++}`, `if(true) {
   (a)++;
} else {
   (b)++;
}
`},
		{`if(true){a++}`, `if(true) {
   (a)++;
}
`},
		//{`mylabel:`, `mylabel:`},
		{`obj = new MyObject()`, `obj = new MyObject();
`},
		{`obj = new MyObject(a, b, 1+2)`, `obj = new MyObject(a, b, (1 + 2));
`},
		{`a = null`, `a = null;
`},
		{`obj = {a: "b", c: 4, e: {x: 1, y: 2, z: 3}}`, `obj = {"a": "b", "c": 4, "e": {"x": 1, "y": 2, "z": 3}};
`},
		{`function f() {return 1}`, `function f() {
   return 1;
}
`},
		{`function f() {return 1 + 2}`, `function f() {
   return (1 + 2);
}
`},
		{`function f(a, b, c) {}`, `function f(a, b, c) {
}
`},
		{`re = /ab+c/g`, `re = /ab+c/g;
`},
		{`a = (1, 2, 3, 4)`, `a = (1, 2, 3, 4);
`},
		{`switch(t) {}`, `switch(t) {
}
`},
		{`switch(t) {case "a":case 2:a = 2;break;default:}`, `switch(t) {
case "a":
case 2:
   a = 2;
   break;
default:
}
`},
		{`switch(t) {case "a":break;}`, `switch(t) {
case "a":
   break;
}
`},
		{`a = this.val + 2`, `a = (this.val + 2);
`},
		{`this.val = 2`, `this.val = 2;
`},
		{`throw new Exception(val)`, `throw new Exception(val);
`},
		{`try{fun(1)}catch(e){fun2()}finally{fun3()}`, `try {
   fun(1);
} catch(e) {
   fun2();
} finally {
   fun3();
}
`},
		{`try{fun(1)}catch(e){fun2()}`, `try {
   fun(1);
} catch(e) {
   fun2();
}
`},
		{`try{fun(1)}finally{fun3()}`, `try {
   fun(1);
} finally {
   fun3();
}
`},
		{`a = b++`, `a = (b)++;
`},
		{`a = ++b`, `a = ++(b);
`},
		{`var x = 2`, `var x = 2;
`},
		{`var x, y = 2`, `var x, y = 2;
`},
		{`while(true){a++}`, `while(true) {
   (a)++;
}
`},
		{`with(Math){a = PI * r * r}`, `with(Math) {
   a = ((PI * r) * r);
}
`},
		{`(function(){alert("test")}());`, `(function() {
   alert("test");
}());
`},
		{`var a, b = 2, c, d = 3, e`, `var a, b = 2, c, d = 3, e;
`},
		{`var a;
var b;
var c;`, `var a;
var b;
var c;
`},
		{`var f = function() {saymyname()}`, `var f = function() {
   saymyname();
}
`},
		{`(function(){var a;var b; var c;}())`, `(function() {
   var a;
   var b;
   var c;
}());
`},
		{`System.debug(error,"Message","thing",info);`, `System.debug(error, "Message", "thing", info);
`},
		{`o = {æ: 0}`, `o = {"\u00e6": 0};
`},
		{`(function(){(function(){function sub(a, b){}; k = sub(1, 2);}())}())`, `(function() {
   (function() {
      k = sub(1, 2);
      function sub(a, b) {
      }
   }());
}());
`},
		{`new Object(a, b, c)`, `new Object(a, b, c);
`},
		{`if(!yes && (m = map["test"])){}`, `if((!(yes) && (m = map["test"]))) {
}
`},
		{`for(;;) {}`, `for( ;  ; ) {
}
`},
		{`typeof(somevar)`, `typeof(somevar);
`},
		{`function f() {for(;;){};while(true){};do{}while(true);switch(a){}}`, `function f() {
   for( ;  ; ) {
   }
   while(true) {
   }
   do {
   } while(true);
   switch(a) {
   }
}
`},
		{`a += 1`, `a += 1;
`},
		{`for(var f = 0 ; ;) {}`, `for(var f = 0 ;  ; ) {
}
`},
		{`for(f = 0 ; ;) {}`, `for(f = 0 ;  ; ) {
}
`},
		{`x = 2`, `x = 2;
`},
		{`var x = 2; var y; z = 2`, `var x = 2;
var y;
z = 2;
`},
		{`var x, y, z; var a, b, c = 2`, `var x, y, z;
var a, b, c = 2;
`},
		{`for(var i = 1, t ; t = 1;){}`, `for(var i = 1, t ; t = 1 ; ) {
}
`},
		{`a = [a, b, c, , ,]`, `a = [a, b, c, , , ];
`},
	}

	for i, test := range tests {
		program, err := parser.ParseFile(nil, "", test.src, 0)
		if err != nil {
			t.Errorf("[%v] Failed, %v\n%v", i, err, test.src)
		}

		verbose := &Verbose{}
		walker := walk.Walker{verbose}
		walker.Begin(program)
		if test.expected != verbose.ToString() {
			t.Errorf("[%v] Export failed!\nExpected: \"%v\"\nActual  : \"%v\"", i, displayString(test.expected), displayString(verbose.ToString()))
		}
	}
}
