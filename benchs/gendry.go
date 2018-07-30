package benchs

import (
	"database/sql"
	"fmt"

	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
	"github.com/sirupsen/logrus"
)

func init() {
	st := NewSuite("gendry")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, GendryInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, GendryInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, GendryUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, GendryRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, GendryReadSlice)

		raw, _ = sql.Open("mysql", ORM_SOURCE)
	}
}

func GendryInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	mp1 := map[string]interface{}{
		"name":    m.Name,
		"title":   m.Title,
		"fax":     m.Fax,
		"web":     m.Web,
		"age":     m.Age,
		"rights":  m.Rights,
		"counter": m.Counter,
	}

	var cond string
	var vals []interface{}
	var err error
	mpList := []map[string]interface{}{mp1}

	for i := 0; i < b.N; i++ {

		cond, vals, err = builder.BuildInsert("model", mpList)
		if err != nil {
			logrus.Errorf("GendryInsert build err=%v\n", err)
			b.FailNow()
		}
		res, err := raw.Exec(cond, vals...)
		if err != nil {
			logrus.Errorf("GendryInsert Exec err=%v\n", err)
			b.FailNow()
		}
		id, err := res.LastInsertId()
		m.Id = int(id)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GendryInsertMulti(b *B) {
	var ms []*Model
	wrapExecute(b, func() {
		initDB()

		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}
	})

	mpList := make([]map[string]interface{}, 0)
	for _, val := range ms {
		mp1 := map[string]interface{}{
			"name":    val.Name,
			"title":   val.Title,
			"fax":     val.Fax,
			"web":     val.Web,
			"age":     val.Age,
			"rights":  val.Rights,
			"counter": val.Counter,
		}
		mpList = append(mpList, mp1)
	}

	var cond string
	var vals []interface{}
	var err error

	for i := 0; i < b.N; i++ {
		cond, vals, err = builder.BuildInsert("model", mpList)
		if err != nil {
			logrus.Errorf("multiple insert err=%v\n", err)
			b.FailNow()
		}
		_, err = raw.Exec(cond, vals...)
		if err != nil {
			logrus.Errorf("multiple exec err=%v\n", err)
			b.FailNow()
		}
	}
}

func GendryUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		rawInsert(m)
	})

	where := map[string]interface{}{
		"id": m.Id,
	}
	update := map[string]interface{}{
		"name":    m.Name,
		"title":   m.Title,
		"fax":     m.Fax,
		"web":     m.Web,
		"age":     m.Age,
		"rights":  m.Rights,
		"counter": m.Counter,
	}
	var cond string
	var vals []interface{}
	var err error

	for i := 0; i < b.N; i++ {
		cond, vals, err = builder.BuildUpdate("model", where, update)
		if err != nil {
			b.FailNow()
		}
		_, err = raw.Exec(cond, vals...)
		if err != nil {
			b.FailNow()
		}
	}
}

func GendryRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		rawInsert(m)
	})
	var mout Model
	where := map[string]interface{}{
		"id": m.Id,
	}
	for i := 0; i < b.N; i++ {
		cond, vals, err := builder.BuildSelect("model", where, []string{"name", "title", "title", "fax", "web", "age", "rights", "counter"})
		if err != nil {
			b.FailNow()
		}
		rows, err := raw.Query(cond, vals...)
		if err != nil {
			b.FailNow()
		}
		defer rows.Close()
		err = scanner.Scan(rows, &mout)
		if err != nil {
			b.FailNow()
		}
	}
}

func GendryReadSlice2(b *B) {
	panic(fmt.Errorf("Not support multi insert"))
}

func GendryReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		var err error
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			err = rawInsert(m)
			if err != nil {
				logrus.Errorf("build insert err=%v\n", err)
				b.FailNow()
			}
		}
	})

	models := make([]Model, 100)
	where := map[string]interface{}{
		"id > ":  0,
		"_limit": []uint{0, 100},
	}

	for i := 0; i < b.N; i++ {
		cond, vals, err := builder.BuildSelect("model", where, []string{"name", "title", "fax", "web", "age", "rights", "counter"})
		if err != nil {
			logrus.Errorf("buildselect err=%v\n", err)
			b.FailNow()
		}
		rows, err := raw.Query(cond, vals...)
		if err != nil {
			logrus.Errorf("build query err=%v\n", err)
			b.FailNow()
		}
		defer rows.Close()
		err = scanner.Scan(rows, &models)
		if err != nil {
			logrus.Errorf("build scan err=%v\n", err)
			b.FailNow()
		}
	}
}
