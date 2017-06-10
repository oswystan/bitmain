//===============================================================================
//                      Copyright (C) 2017 wystan
//
//        filename: student.go
//     description:
//         created: 2017-06-10 16:13:58
//          author: wystan
//
//===============================================================================
package store

import (
	"fmt"

	"github.com/oswystan/bitmain/model"
)

var sqlSum = `select sum(score) as total from students where class_number = ?`

type StudentStore struct {
}

func NewStudentStore() *StudentStore {
	return &StudentStore{}
}

func (st *StudentStore) Post(data *model.Student) error {
	// insert first otherwise update
	err := db.Insert(data)
	if err == nil {
		return nil
	}
	err = db.Update(data)
	return err
}

func (st *StudentStore) Get(data *model.Student) error {
	err := db.Select(data)
	if err != nil {
		return err
	}
	return nil
}

func (st *StudentStore) GetTotalScore(data *model.Student) (*model.TotalScore, error) {
	err := st.Get(data)
	if err != nil {
		return nil, fmt.Errorf("student-not-found")
	}

	score := &model.TotalScore{}
	_, err = db.QueryOne(score, sqlSum, data.ClassNumber)
	if err != nil {
		return nil, err
	}

	return score, nil
}

//==================================== END ======================================
