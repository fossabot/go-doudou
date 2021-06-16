package query

import (
	"fmt"
	"github.com/unionj-cloud/go-doudou/ddl/sortenum"
)

func ExampleCriteria() {

	query := C().Col("name").Eq(Literal("wubin")).Or(C().Col("school").Eq(Literal("havard"))).And(C().Col("age").Eq(Literal(18)))
	fmt.Println(query.Sql())

	query = C().Col("name").Eq(Literal("wubin")).Or(C().Col("school").Eq(Literal("havard"))).And(C().Col("delete_at").IsNotNull())
	fmt.Println(query.Sql())

	query = C().Col("name").Eq(Literal("wubin")).Or(C().Col("school").In(Literal("havard"))).And(C().Col("delete_at").IsNotNull())
	fmt.Println(query.Sql())

	query = C().Col("name").Eq(Literal("wubin")).Or(C().Col("school").In(Literal([]string{"havard", "beijing unv"}))).And(C().Col("delete_at").IsNotNull())
	fmt.Println(query.Sql())

	var d int
	var e int
	d = 10
	e = 5

	query = C().Col("name").Eq(Literal("wubin")).Or(C().Col("age").In(Literal([]*int{&d, &e}))).And(C().Col("delete_at").IsNotNull())
	fmt.Println(query.Sql())

	query = C().Col("name").Ne(Literal("wubin")).Or(C().Col("create_at").Lt(Func("now()")))
	fmt.Println(query.Sql())

	query = C().Col("name").Ne(Literal("wubin")).Or(C().Col("create_at").Lte(Func("now()")))
	fmt.Println(query.Sql())

	query = C().Col("name").Ne(Literal("wubin")).Or(C().Col("create_at").Gt(Func("now()")))
	fmt.Println(query.Sql())

	query = C().Col("name").Ne(Literal("wubin")).Or(C().Col("create_at").Gte(Func("now()")))
	fmt.Println(query.Sql())

	page := Page{
		Orders: []Order{
			{
				Col:  "create_at",
				Sort: sortenum.Desc,
			},
		},
		Offset: 20,
		Size:   10,
	}
	page = page.Order(Order{
		Col:  "score",
		Sort: sortenum.Asc,
	})
	page = page.Limit(30, 5)
	fmt.Println(page.Sql())
	pageRet := NewPageRet(page)
	fmt.Println(pageRet.PageNo)

	fmt.Println(P().Order(Order{
		Col:  "score",
		Sort: sortenum.Asc,
	}).Limit(20, 10).Sql())

	query = C().Col("name").Eq(Literal("wubin")).Or(C().Col("school").Eq(Literal("havard"))).
		And(C().Col("age").Eq(Literal(18))).
		Or(C().Col("score").Gte(Literal(90)))
	fmt.Println(query.Sql())

	// Output:
	// ((`name` = 'wubin' or `school` = 'havard') and `age` = '18')
	// ((`name` = 'wubin' or `school` = 'havard') and `delete_at` is not null)
	// ((`name` = 'wubin' or `school` in ('havard')) and `delete_at` is not null)
	// ((`name` = 'wubin' or `school` in ('havard','beijing unv')) and `delete_at` is not null)
	// ((`name` = 'wubin' or `age` in ('10','5')) and `delete_at` is not null)
	// (`name` != 'wubin' or `create_at` < now())
	// (`name` != 'wubin' or `create_at` <= now())
	// (`name` != 'wubin' or `create_at` > now())
	// (`name` != 'wubin' or `create_at` >= now())
	// order by create_at desc,score asc limit 30,5
	// 7
	// order by score asc limit 20,10
	// (((`name` = 'wubin' or `school` = 'havard') and `age` = '18') or `score` >= '90')
}
