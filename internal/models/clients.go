package models

type Client struct {
	ID       uint
	Name     *string
	Url      string
	Versions Versions
}
