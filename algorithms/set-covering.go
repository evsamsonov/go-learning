package main

import "fmt"

func main() {
	// Эмуляция set
	requiredStates := map[string]bool{
		"mt": true,
		"wa": true,
		"or": true,
		"id": true,
		"nv": true,
		"ut": true,
		"ca": true,
		"az": true,
	}

	stations := make(map[string]map[string]bool)
	stations["kone"] = map[string]bool{"id": true, "nv": true, "ut": true}
	stations["ktwo"] = map[string]bool{"wa": true, "id": true, "mt": true}
	stations["kthree"] = map[string]bool{"or": true, "nv": true, "ca": true}
	stations["kfour"] = map[string]bool{"nv": true, "ut": true}
	stations["kfive"] = map[string]bool{"ca": true, "az": true}

	optimalCovering := findOptimalCovering(stations, requiredStates)
	fmt.Println(optimalCovering)
}

// Поиск оптимального покрытия станциями
// с использованием жадного алгоритма
func findOptimalCovering(stations map[string]map[string]bool, requiredStates map[string]bool) map[string]bool {
	optimalCovering := make(map[string]bool)

	for ; len(requiredStates) > 0 ; {
		bestStation := ""
		statesCovered := make(map[string]bool)
		for station, states := range stations {
			// Сколько покрывает текущая?
			covered := intersectionStates(states, requiredStates)
			if len(covered) > len(statesCovered) {
				statesCovered = covered
				bestStation = station
			}
		}

		// Удаляем штаты, которые покрыты этой станцией
		for state, _ := range statesCovered {
			delete(requiredStates, state)
		}

		optimalCovering[bestStation] = true
	}

	return optimalCovering
}

// Находит пересечение всех переданных множеств
func intersectionStates(statesCollection ...map[string]bool) map[string]bool {
	if len(statesCollection) == 1 {
		// Если получили одно множество, считаем, что это и есть пересечение
		return statesCollection[0]
	}

	if len(statesCollection) == 2 {
		intersection := make(map[string]bool)
		for state, _ := range statesCollection[0] {
			if _, hasState := statesCollection[1][state]; hasState {
				intersection[state] = true
			}
		}

		return intersection
	}

	return intersectionStates(statesCollection[0], intersectionStates(statesCollection[1:]...))
}
