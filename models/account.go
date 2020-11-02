package models

import (
	randstr "github.com/0xVERS/go-randstr"
)

func (old *Account) Copy() Account {
	return Account{
		Id:          old.Id,
		Name:        old.Name,
		Password:    old.Password,
		ServiceName: old.ServiceName,
	}
}

const (
	accountIDSize = 8
	genIDTrials   = 8
)

type accountModel struct {
	accountMap   map[string]*Account
	accountOrder []string
}

func NewAccountModel() *accountModel {
	return &accountModel{
		accountMap: map[string]*Account{},
	}
}

func (m *accountModel) Create(name, password, serviceName string) error {
	var id string
	for i := 0; i < genIDTrials; i++ {
		id = randstr.Gen(accountIDSize)
		if _, exists := m.GetByID(id); !exists {
			goto SuccessGenID
		}
	}
	return ErrIDExhaustion

SuccessGenID:
	m.accountMap[id] = &Account{
		Id:          id,
		Name:        name,
		Password:    password,
		ServiceName: serviceName,
	}
	return nil
}

func (m *accountModel) Save(a Account) {
	m.accountMap[a.Id] = &a
}

func (m *accountModel) Delete(id string) error {
	if _, exists := m.GetByID(id); !exists {
		return ErrNotExists
	}
	delete(m.accountMap, id)
	return nil
}

func (m *accountModel) Update(id string, new Account) error {
	old, exists := m.GetByID(id)
	if !exists {
		return ErrNotExists
	}
	if old.Id != new.Id {
		if _, exists = m.GetByID(new.Id); exists {
			return ErrAlreadyExists
		}
		if err := m.Delete(id); err != nil {
			return err
		}
	}

	m.accountMap[new.Id] = &new
	return nil
}

func (m *accountModel) Len() int {
	return int(len(m.accountMap))
}

func (m *accountModel) GetByID(id string) (Account, bool) {
	a, exists := m.accountMap[id]
	return a.Copy(), exists
}

func (m *accountModel) Find(filter func(Account) bool) (*accountModel, int) {
	new := NewAccountModel()
	if m.Len() <= 0 {
		return new, 0
	}
	for _, a := range m.accountMap {
		if filter(*a) {
			new.Save(a.Copy())
		}
	}
	return new, new.Len()
}

func (m *accountModel) FindByName(name string) (*accountModel, int) {
	return m.Find(func(a Account) bool {
		return a.Name == name
	})
}

func (m *accountModel) FindByServiceName(serviceName string) (*accountModel, int) {
	return m.Find(func(a Account) bool {
		return a.ServiceName == serviceName
	})
}

func (m *accountModel) List() []Account {
	var result []Account
	if len(m.accountOrder) > 0 {
		result = make([]Account, len(m.accountOrder))
		for i, id := range m.accountOrder {
			if a, exists := m.GetByID(id); exists {
				result[i] = a
			}
		}
	} else {
		result = make([]Account, len(m.accountMap))
		for _, a := range m.accountMap {
			result = append(result, a.Copy())
		}
	}
	return result
}
