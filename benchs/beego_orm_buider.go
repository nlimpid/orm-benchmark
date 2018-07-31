package benchs

import (
	"fmt"
	"log"

	"github.com/astaxie/beego/orm"
)

var qb orm.QueryBuilder

// func Tinit() {
// 	st := NewSuite("orm_builder")
// 	st.InitF = func() {
// 		st.AddBenchmark("Insert", 2000*ORM_MULTI, BeegoOrmInsert)
// 		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, BeegoOrmInsertMulti)
// 		st.AddBenchmark("Update", 2000*ORM_MULTI, BeegoOrmUpdate)
// 		st.AddBenchmark("Read", 4000*ORM_MULTI, BeegoOrmBuilderRead)
// 		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, BeegoOrmBuilderReadSlice)

// 		orm.RegisterDataBase("default", "mysql", ORM_SOURCE, ORM_MAX_IDLE, ORM_MAX_CONN)
// 		orm.RegisterModel(new(Model))

// 		bo = orm.NewOrm()
// 		bo.Using("default")

// 	}
// }

func BeegoOrmBuilderRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		bo.Insert(m)
	})

	for i := 0; i < b.N; i++ {
		qb, _ = orm.NewQueryBuilder("mysql")
		qb.Select("name", "title", "fax", "web", "age", "rights", "counter").
			From("model").
			Where("id = ?")
		sql := qb.String()
		err := bo.Raw(sql, m.Id).QueryRow(m)
		if err != nil {
			log.Printf("err=%v\n", err)
			b.FailNow()
		}
	}
}

func BeegoOrmBuilderReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			if _, err := bo.Insert(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		qb, _ = orm.NewQueryBuilder("mysql")
		var models []*Model
		qb.Select("name", "title", "fax", "web", "age", "rights", "counter").From("model").
			Where("id > ?").Limit(100)
		sql := qb.String()
		if _, err := bo.Raw(sql, m.Id).QueryRows(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
