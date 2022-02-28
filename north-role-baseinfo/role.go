package north_role_baseinfo

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type Role struct {
	Id         int64  `json:"id,omitempty"`
	Status     int    `json:"status,omitempty"`
	CreateTime int    `json:"create_time,omitempty"`
	UpdateTime int    `json:"update_time,omitempty"`
	ListOrder  int8   `json:"list_order,omitempty"`
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
}

func SelectRoleById(DB *sql.DB, id int) (*Role, error) {
	var pe Role
	err := DB.QueryRow("SELECT id,status,list_order,name FROM role WHERE id = ?", id).Scan(&pe.Id, &pe.Status, &pe.ListOrder, &pe.Name)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	marshal, _ := json.Marshal(pe)
	fmt.Println(string(marshal))
	return &pe, nil
}

func SelectAllRole(DB *sql.DB) (*[]Role, error) {
	result, err := DB.Query("SELECT id,status,list_order,name FROM role WHERE status = 1")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	roles := make([]Role, 0)
	for result.Next() {
		var role Role
		err := result.Scan(&role.Id, &role.Status, &role.ListOrder, &role.Name)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		roles = append(roles, role)
	}
	fmt.Println(roles)
	return &roles, nil
}

//func RoleAdd(DB *sql.DB,status int,order int,name string)  {
//	result
//}