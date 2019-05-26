package main

import (
	"fmt"
	"math"
)

func main() {
	dijkstrasAlgorithm();
}

func dijkstrasAlgorithm() {
	// Строим взвешенный ацикличный граф
	graph := make(map[string]map[string]int)
	graph["start"] = map[string]int{}
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = map[string]int{}
	graph["a"]["finish"] = 1

	graph["b"] = map[string]int{}
	graph["b"]["a"] = 3
	graph["b"]["finish"] = 5

	graph["finish"] = map[string]int{} 	// Конечная точка не должна иметь соседей

	// Стоимость переходов
	costs := make(map[string]int)
	costs["a"] = 6
	costs["b"] = 2
	costs["finish"] = math.MaxInt32

	// Родители
	parents := make(map[string]string)
	parents["a"] = "start"
	parents["b"] = "start"
	parents["finish"] = ""

	// Обработанные узлы
	processed := make(map[string]bool)

	lowestCostNode := findLowestCostNode(costs, processed)
	for ; lowestCostNode != "" ; {
		// Считает цену для всех соседей узла
		for node, cost := range graph[lowestCostNode] {
			newCost := costs[lowestCostNode] + cost
			// Если цена меньше предыдущей, то обновим
			if newCost < costs[node] {
				costs[node] = newCost
				parents[node] = lowestCostNode
			}
		}

		processed[lowestCostNode] = true
		lowestCostNode = findLowestCostNode(costs, processed)
	}

	fmt.Println(costs)
	fmt.Println(parents)
}

// Поиск элемента с минимальной ценой
func findLowestCostNode(costs map[string]int, processed map[string]bool) string {
	lowestCost := math.MaxInt32
	lowestCostNode := ""
	for node, cost := range costs {
		// Можно инициализировать переменные в if
		if _, inProcessed := processed[node]; cost < lowestCost && !inProcessed {
			lowestCost = cost
			lowestCostNode = node
		}
	}

	return lowestCostNode
}
