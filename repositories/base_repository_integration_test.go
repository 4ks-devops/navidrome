package repositories

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/deluan/gosonic/tests"
	"fmt"
)

type TestEntity struct {
	Id string
	Name string
}

func shouldBeEqual(actualStruct interface{}, expectedStruct ...interface{}) string {
	actual := fmt.Sprintf("%#v", actualStruct)
	expected := fmt.Sprintf("%#v", expectedStruct[0])
	return ShouldEqual(actual, expected)
}

func TestIntegrationBaseRepository(t *testing.T) {
	tests.Init(t, true)
	Convey("Subject: saveOrUpdate", t, func() {

		Convey("Given an empty DB", func() {
			dropDb()
			repo := &BaseRepository{table: "test"}

			Convey("When I save a new entity", func() {
				entity := &TestEntity{"123", "My Name"}
				err := repo.saveOrUpdate("123", entity)

				Convey("Then the method shouldn't return any errors", func() {
					So(err, ShouldBeNil)
				})

				Convey("Then the number of entities should be 1", func() {
					count, _ := repo.CountAll()
					So(count, ShouldEqual, 1)
				})

				Convey("And this entity should be equal to the the saved one", func() {
					actualEntity := &TestEntity{}
					repo.loadEntity("123", actualEntity)
					So(actualEntity, shouldBeEqual, entity)
				})

			})

		})

		Convey("Given a table with one entity", func() {
			dropDb()
			repo := &BaseRepository{table: "test"}
			entity := &TestEntity{"111", "One Name"}
			repo.saveOrUpdate(entity.Id, entity)

			Convey("When I save an entity with a different Id", func() {
				newEntity := &TestEntity{"222", "Another Name"}
				repo.saveOrUpdate(newEntity.Id, newEntity)

				Convey("Then the number of entities should be 2", func() {
					count, _ := repo.CountAll()
					So(count, ShouldEqual, 2)
				})

			})

			Convey("When I save an entity with the same Id", func() {
				newEntity := &TestEntity{"111", "New Name"}
				repo.saveOrUpdate(newEntity.Id, newEntity)

				Convey("Then the number of entities should be 1", func() {
					count, _ := repo.CountAll()
					So(count, ShouldEqual, 1)
				})

				Convey("And the entity should be updated", func() {
					actualEntity := &TestEntity{}
					repo.loadEntity("111", actualEntity)
					So(actualEntity.Name, ShouldEqual, newEntity.Name)
				})

			})

		})

	})
}