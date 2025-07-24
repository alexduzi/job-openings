package handler

import (
	"fmt"

	"github.com/alexduzi/job-openings/model"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:role`
	Company  string `json:company`
	Location string `json:location`
	Remote   *bool  `json:remote`
	Link     string `json:link`
	Salary   int64  `json:salary`
}

func (c *CreateOpeningRequest) Validate() error {
	if c.Role == "" && c.Company == "" && c.Location == "" && c.Remote == nil && c.Salary <= 0 {
		return fmt.Errorf("request body is empty")
	}
	if c.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if c.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if c.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if c.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if c.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if c.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	return nil
}

func (c *CreateOpeningRequest) Merge(op *model.Opening) {
	if c.Role != "" && (c.Role != op.Role) {
		op.Role = c.Role
	}
	if c.Company != "" && (c.Company != op.Company) {
		op.Company = c.Company
	}
	if c.Location != "" && (c.Location != op.Location) {
		op.Location = c.Location
	}
	if c.Link != "" && (c.Link != op.Link) {
		op.Link = c.Link
	}
	if c.Remote != nil {
		op.Remote = *c.Remote
	}
}
