//===============================================================================
//                      Copyright (C) 2017 wystan
//
//        filename: class.go
//     description:
//         created: 2017-06-10 16:57:26
//          author: wystan
//
//===============================================================================

package store

import "github.com/oswystan/bitmain/model"

var sqlMaxScore = `select  * from students as s where s.score = (select max(score) from students) order by id limit 1`

type ClassStore struct {
	/* data */
}

func NewClassStore() *ClassStore {
	return &ClassStore{}
}

func (st *ClassStore) Get(data *model.Class) error {
	err := db.Select(data)
	return err
}

func (st *ClassStore) Post(data *model.Class) error {
	err := db.Insert(data)
	if err == nil {
		return nil
	}
	err = db.Update(data)
	return err
}
func (st *ClassStore) GetTopTeacher() (*model.TopTeacher, error) {
	student := &model.Student{}
	_, err := db.QueryOne(student, sqlMaxScore)
	if err != nil {
		return nil, err
	}
	class := &model.Class{ClassNumber: student.ClassNumber}
	err = db.Select(class)
	if err != nil {
		return nil, err
	}
	return &model.TopTeacher{Teacher: class.Teacher}, nil
}

//==================================== END ======================================
