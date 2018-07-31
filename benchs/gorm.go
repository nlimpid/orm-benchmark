package benchs

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var gormDB *gorm.DB

func init() {
	st := NewSuite("gorm")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, GormInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, GormInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, GormUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, GormRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, GormReadSlice)

		gormDB, _ = gorm.Open("mysql", ORM_SOURCE)
	}
}

func GormInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if err := gormDB.Create(m).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GormInsertMulti(b *B) {
	panic(fmt.Errorf("Not support multi insert"))
}

func GormUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		gormDB.Create(m)
	})

	for i := 0; i < b.N; i++ {
		if err := gormDB.Save(m).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GormRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		gormDB.Create(m)
	})

	for i := 0; i < b.N; i++ {
		if err := gormDB.Find(m).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GormReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			if err := gormDB.Create(m).Error; err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		if err := gormDB.Where("id > ?", 0).Limit(100).Find(&models).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
