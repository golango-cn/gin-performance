package core

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	rd "github.com/go-redis/redis/v7"
)

type Employee struct {
	EmployeeNo int
	BirthDate  time.Time
	FirstName  string
	LastName   string
	Gender     string
	HireDate   time.Time
}

var locker sync.Locker

func GetEmployees() ([]*Employee, error) {

	var employees []*Employee
	result, err := redis.Get("employees").Result()

	if err != rd.Nil {
		json.Unmarshal([]byte(result), &employees)
		return employees, nil
	}

	locker.Lock()
	defer locker.Unlock()

	// , first_name, last_name , birth_date, gender, hire_date
	rows, err := db.Query(`
							select emp_no, first_name
							from employees where first_name = ? limit 1 `, `Georgi`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()


	for rows.Next() {

		//,
		//&employee.FirstName,
		//	&employee.LastName,
		//	&employee.BirthDate,
		//	&employee.Gender,
		//	&employee.HireDate

		employee := &Employee{}
		if err = rows.Scan(&employee.EmployeeNo, &employee.FirstName); err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	bytes, _ := json.Marshal(&employees)

	cmd := redis.Set("employees", string(bytes), 0)
	fmt.Println(cmd.Result())

	return employees, nil

}
