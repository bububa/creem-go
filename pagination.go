package creem

// Pagination details for the list
type Pagination struct {
	// TotalRecords Total number of records in the list
	TotalRecords int64 `json:"total_records,omitempty"`
	// TotalPages Total number of pages available
	TotalPages int64 `json:"total_pages,omitempty"`
	// CurrentPage The current page number
	CurrentPage int64 `json:"current_page,omitempty"`
	// NextPage The next page number, or null if there is no next page
	NextPage int64 `json:"next_page,omitempty"`
	// PrevPage The previous page number, or null if there is no previous page
	PrevPage int64 `json:"prev_page,omitempty"`
}
