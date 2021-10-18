package communication

import "fmt"

type Review struct {
	ReviewID uint64 `json:"review_id"`
	Product  uint64 `json:"item_id"`
	User     uint64 `json:"user_id"`
	Text     string `json:"text"`
	Rating   int    `json:"rating"`
}

func (r *Review) String() string {
	return fmt.Sprintf(`Review-%d (Product: %d; User: %d)`, r.ReviewID, r.Product, r.User)
}
