// OCP - Open-Closed Principle
package main

import "fmt"

// Specification, a requirement
// for example a user should be able to filter by a certain criteria

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	// ...
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// this interferes with the open close principle
// we shouldn't have to add more methods to Filter to create new filters
// we shouldn't have to modify Filter to add new filters
//
// we shouldn't modify something that already exists and has already been tested
// we should have an extendable setup
//
// func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
// 	result := make([]*Product, 0)
// 	for i, v := range products {
// 		if v.size == size {
// 			result = append(result, &products[i])
// 		}
// 	}
// 	return result
// }
//
// func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
// 	result := make([]*Product, 0)
// 	for i, v := range products {
// 		if v.size == size && v.color == color {
// 			result = append(result, &products[i])
// 		}
// 	}
// 	return result
// }

// This is better because you are very unlikely to modify Specification or BetterFilter
// You will just create new Specifications for new filters
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

// creating a filter for color and size, or any 2 specifications
type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

type BetterFilter struct {
	// ...
}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	car := Product{"Car", blue, medium}

	products := []Product{apple, tree, car}
	fmt.Println("Green products (old):")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf("  -  %s is green\n", v.name)
	}

	fmt.Println("Green products (new):")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf("  -  %s is green\n", v.name)
	}

	fmt.Println("Large green products (new):")
	largeSpec := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpec, largeSpec}
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf("  -  %s is large and green\n", v.name)
	}
}
