package stats

import (
	types "github.com/ssharifzoda/kun/v2"
	"reflect"
	"testing"
)

func TestCategoriesTotal(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "fun", Amount: 5_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 800000,
		"food": 200000,
		"fun":  500000,
	}
	result := CategoriesTotal(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result, expected %v, actual %v", expected, result)
	}
}
func TestFilterByCategory_notFound(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}
func TestFilterByCategory_oneFound(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	expected := []types.Payment{
		{ID: 2, Category: "food"},
	}
	result := FilterByCategory(payments, "food")

	if reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}
func TestFilterByCategory_MultiFound(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	expected := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
	}
	result := FilterByCategory(payments, "auto")

	if reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}
