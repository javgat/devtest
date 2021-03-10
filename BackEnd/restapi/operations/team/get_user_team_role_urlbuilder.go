// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetUserTeamRoleURL generates an URL for the get user team role operation
type GetUserTeamRoleURL struct {
	Teamname string
	Username string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetUserTeamRoleURL) WithBasePath(bp string) *GetUserTeamRoleURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetUserTeamRoleURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetUserTeamRoleURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/users/{username}/teams/{teamname}/role"

	teamname := o.Teamname
	if teamname != "" {
		_path = strings.Replace(_path, "{teamname}", teamname, -1)
	} else {
		return nil, errors.New("teamname is required on GetUserTeamRoleURL")
	}

	username := o.Username
	if username != "" {
		_path = strings.Replace(_path, "{username}", username, -1)
	} else {
		return nil, errors.New("username is required on GetUserTeamRoleURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/DevTest"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetUserTeamRoleURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetUserTeamRoleURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetUserTeamRoleURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetUserTeamRoleURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetUserTeamRoleURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetUserTeamRoleURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}