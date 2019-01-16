package app

// Client struct
type Client struct {
	ItemSrv ItemService
}

// Item struct
type Item struct {
}

// ItemService interface
type ItemService interface {
	List() ([]ItemData, error)
}

// ItemData struct
type ItemData struct {
	ID   int
	Name string
}

// List get Client list
func (item *Item) List() ([]ItemData, error) {
	return item.List()
}

// ItemList item list
func (c *Client) ItemList() ([]ItemData, error) {
	return c.ItemSrv.List()
}
