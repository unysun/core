package script

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expr"
	"github.com/project-flogo/core/data/resolve"
	"github.com/stretchr/testify/assert"
)

var resolver = resolve.NewCompositeResolver(map[string]resolve.Resolver{"static": &TestStaticResolver{}, ".": &TestResolver{}})
var factory = NewExprFactory(resolver)

func TestLitExprInt(t *testing.T) {
	expr, err := factory.NewExpr(`123`)
	assert.Nil(t, err)
	assert.NotNil(t, expr)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 123, v)
}

func TestLitExprFloat(t *testing.T) {
	expr, err := factory.NewExpr(`123.5`)
	assert.Nil(t, err)
	assert.NotNil(t, expr)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 123.5, v)
}

func TestLitExprBool(t *testing.T) {
	expr, err := factory.NewExpr(`true`)
	assert.Nil(t, err)
	assert.NotNil(t, expr)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)
}

func TestLitExprStringSQ(t *testing.T) {
	expr, err := factory.NewExpr(`'foo bar'`)
	assert.Nil(t, err)
	assert.NotNil(t, expr)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, "foo bar", v)
}

func TestLitExprStringDQ(t *testing.T) {
	expr, err := factory.NewExpr(`"foo bar"`)
	assert.Nil(t, err)
	assert.NotNil(t, expr)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, "foo bar", v)
}

func TestLitExprNil(t *testing.T) {
	expr, err := factory.NewExpr(`nil`)
	assert.Nil(t, err)
	assert.NotNil(t, expr)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Nil(t, v)

	expr, err = factory.NewExpr(`null`)
	assert.Nil(t, err)
	assert.NotNil(t, expr)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Nil(t, v)
}

func TestLitExprRef(t *testing.T) {

	expr, err := factory.NewExpr(`$.foo`)
	assert.Nil(t, err)
	assert.NotNil(t, expr)

	scope := newScope(map[string]interface{}{"foo": "bar"})
	v, err := expr.Eval(scope)
	assert.Nil(t, err)
	assert.Equal(t, "bar", v)
}

func TestLitExprStaticRef(t *testing.T) {

	expr, err := factory.NewExpr(`$static.foo`)
	assert.Nil(t, err)
	assert.NotNil(t, expr)

	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, "bar", v)
}

func TestCmpExprEq(t *testing.T) {
	expr, err := factory.NewExpr(`123==123`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`123==321`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`123==123.0`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`123==123.5`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`"foo"=="foo"`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`"foo"=='foo'`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`"foo"=="bar"`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)
}

func TestCmpExprNotEq(t *testing.T) {
	expr, err := factory.NewExpr(`123!=123`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`123!=321`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`123!=123.0`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`123!=123.5`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`"foo"!="foo"`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`"foo"!='foo'`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`"foo"!="bar"`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)
}

func TestCmpExprLt(t *testing.T) {
	expr, err := factory.NewExpr(`123<123`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`123<321`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`123<123.0`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`123<123.5`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`123.5<123`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`"ab"<"ac"`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)
}

func TestCmpExprLtEq(t *testing.T) {
	expr, err := factory.NewExpr(`123<=123`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`123<=321`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`123<=123.0`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`123<=123.5`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`123.5<=123`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`"ab"<="ac"`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)
}

func TestCmpExprGt(t *testing.T) {
	expr, err := factory.NewExpr(`123>123`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`123>321`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`123>123.0`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`123>123.5`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`123.5>123`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`"ab">"ac"`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)
}

func TestCmpExprGtEq(t *testing.T) {
	expr, err := factory.NewExpr(`123>=123`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`123>=321`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`123>=123.0`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`123>=123.5`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`123.5>=123`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`"ab">="ac"`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)
}

func TestArithExprAdd(t *testing.T) {
	expr, err := factory.NewExpr(`12+13`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 25, v)

	expr, err = factory.NewExpr(`12+13.5`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 25.5, v)

	expr, err = factory.NewExpr(`"foo"+'bar'`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, "foobar", v)
}

func TestArithExprSub(t *testing.T) {
	expr, err := factory.NewExpr(`13-12`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 1, v)

	expr, err = factory.NewExpr(`12-13`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, -1, v)

	expr, err = factory.NewExpr(`13.5-12`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 1.5, v)
}

func TestArithExprMul(t *testing.T) {
	expr, err := factory.NewExpr(`2*5`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 10, v)

	//expr, err = factory.NewExpr(`2*-5`)
	//assert.Nil(t, err)
	//v, err = expr.Eval(nil)
	//assert.Nil(t, err)
	//assert.Equal(t, -10, v)

	expr, err = factory.NewExpr(`2*.1`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, .2, v)
}

func TestArithExprDiv(t *testing.T) {
	expr, err := factory.NewExpr(`10/2`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 5, v)

	expr, err = factory.NewExpr(`2/10`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 0, v)

	expr, err = factory.NewExpr(`2/.5`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 4.0, v)
}

func TestArithExprMod(t *testing.T) {
	expr, err := factory.NewExpr(`10%2`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 0, v)

	expr, err = factory.NewExpr(`10%3`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 1, v)

	expr, err = factory.NewExpr(`10.5%2`) //todo should we throw an error in this case?
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 0, v)
}

func TestBoolExprOr(t *testing.T) {
	expr, err := factory.NewExpr(`true || false`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`false || false`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`1 || 0`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`0 || 0`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)
}

func TestBoolExprAnd(t *testing.T) {
	expr, err := factory.NewExpr(`true && true`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`true && false`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`1 && 1`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	expr, err = factory.NewExpr(`1 && 0`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)
}

func TestArithPrecedence(t *testing.T) {
	expr, err := factory.NewExpr(`1+5*2`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 11, v)

	expr, err = factory.NewExpr(`1+5/2.0`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 3.5, v)

	expr, err = factory.NewExpr(`6/2+1*2`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 5, v)

	expr, err = factory.NewExpr(`1+5%2`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 2, v)
}

func TestArithParen(t *testing.T) {
	expr, err := factory.NewExpr(`(1+5)*2`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 12, v)

	expr, err = factory.NewExpr(`10/(5-3)`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 5, v)

	expr, err = factory.NewExpr(`11/(5-3.0)`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 5.5, v)
}

func TestUnaryExpr(t *testing.T) {
	expr, err := factory.NewExpr(`-1`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, -1, v)

	expr, err = factory.NewExpr(`-1.5`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, -1.5, v)

	expr, err = factory.NewExpr(`!true`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`!false`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)
}

func TestUnaryExprComplex(t *testing.T) {
	expr, err := factory.NewExpr(`-(1+2)`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, -3, v)

	expr, err = factory.NewExpr(`-1*2`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, -2, v)

	expr, err = factory.NewExpr(`2*-1`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, -2, v)

	expr, err = factory.NewExpr(`!(false||true)`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	expr, err = factory.NewExpr(`!(false&&true)`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)
}

func TestTernaryExpr(t *testing.T) {
	expr, err := factory.NewExpr(` 1<2 ? 10 : 20`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 10, v)

	expr, err = factory.NewExpr(`4>3 ? 40 :30`)
	assert.Nil(t, err)
	v, err = expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 40, v)
}

var result interface{}

func BenchmarkLit(b *testing.B) {
	var r interface{}

	expr, _ := factory.NewExpr(`123`)

	for n := 0; n < b.N; n++ {

		r, _ = expr.Eval(nil)
	}
	result = r
}

func BenchmarkLit2(b *testing.B) {
	var r interface{}

	f := expression.NewExprFactory(nil, nil)
	expr, _ := f.NewExpr(`123`)

	for n := 0; n < b.N; n++ {

		r, _ = expr.Eval(nil)
	}
	result = r
}

/////////////////////////
// Resolver Helpers

func newScope(values map[string]interface{}) data.Scope {
	return &TestScope{values: values}
}

type TestScope struct {
	values map[string]interface{}
}

func (s *TestScope) GetValue(name string) (value interface{}, exists bool) {
	value, exists = s.values[name]
	return
}

func (TestScope) SetValue(name string, value interface{}) error {
	//ignore
	return nil
}

type TestResolver struct {
}

func (*TestResolver) GetResolverInfo() *data.ResolverInfo {
	return data.NewResolverInfo(false, false)
}

func (*TestResolver) Resolve(scope data.Scope, item string, field string) (interface{}, error) {

	value, exists := scope.GetValue(field)
	if !exists {
		err := fmt.Errorf("failed to resolve variable: '%s', not in scope", field)
		return "", err
	}

	return value, nil
}

type TestStaticResolver struct {
}

func (*TestStaticResolver) GetResolverInfo() *data.ResolverInfo {
	return data.NewResolverInfo(true, false)
}

func (*TestStaticResolver) Resolve(scope data.Scope, item string, field string) (interface{}, error) {

	if field == "foo" {
		return "bar", nil
	}

	if field == "bar" {
		return "for", nil
	}

	return nil, fmt.Errorf("failed to resolve variable: '%s'", field)
}