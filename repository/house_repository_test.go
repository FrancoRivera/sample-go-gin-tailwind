// testing suite for house_repository.go

package repository

import "testing"

// Test the function that fetches all

func TestFetchAllHouses(t *testing.T) {
	hlist := HouseRepository.FetchAll()

	// Check that the length of the list of houses returned is the
	// same as the length of the global variable holding the list
	if len(hlist) != len(houseList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range hlist {
		if v.Content != houseList[i].Content ||
			v.ID != houseList[i].ID ||
			v.Title != houseList[i].Title {

			t.Fail()
			break
		}
	}
}
