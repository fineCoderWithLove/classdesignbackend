package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	populationSize     = 100  // 种群大小
	generationCount    = 1000 // 迭代次数
	crossoverRate      = 0.8  // 交叉率
	mutationRate       = 0.05 // 变异率
	tournamentSize     = 5    // 锦标赛选择的参与者数量
	elitismCount       = 1    // 精英主义数量
	numCities          = 10   // 城市数量
	maxCoordinateValue = 100  // 坐标的最大值
)

// City 表示一个城市
type City struct {
	x int
	y int
}

// Individual 表示一个个体（路径）
type Individual struct {
	path     []int
	fitness  float64
	distance float64
}

// 初始化种群
func initializePopulation() []Individual {
	population := make([]Individual, populationSize)

	for i := range population {
		population[i] = Individual{
			path:     generateRandomPath(),
			fitness:  0,
			distance: 0,
		}
	}

	return population
}

// 生成随机路径
func generateRandomPath() []int {
	path := make([]int, numCities)

	for i := 0; i < numCities; i++ {
		path[i] = i
	}

	// Fisher-Yates 洗牌算法
	rand.Seed(time.Now().UnixNano())
	for i := numCities - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		path[i], path[j] = path[j], path[i]
	}

	return path
}

// 计算路径的总距离
func calculateDistance(path []int, cities []City) float64 {
	distance := 0.0

	for i := 0; i < numCities-1; i++ {
		city1 := cities[path[i]]
		city2 := cities[path[i+1]]
		distance += calculateEuclideanDistance(city1, city2)
	}

	// 回到起点
	distance += calculateEuclideanDistance(cities[path[numCities-1]], cities[path[0]])

	return distance
}

// 计算两个城市之间的欧几里得距离
func calculateEuclideanDistance(city1 City, city2 City) float64 {
	dx := float64(city1.x - city2.x)
	dy := float64(city1.y - city2.y)
	return float64((dx*dx + dy*dy))
}

// 计算个体的适应度
func calculateFitness(individual *Individual) {
	individual.fitness = 1 / individual.distance
}

// 主函数
func main() {
	// 生成城市坐标
	cities := generateCities()

	// 初始化种群
	population := initializePopulation()

	// 迭代进化
	for g := 0; g < generationCount; g++ {
		// 计算每个个体的距离和适应度
		for i := range population {
			population[i].distance = calculateDistance(population[i].path, cities)
			calculateFitness(&population[i])
		}

		// 排序种群，按适应度降序排序
		sortPopulation(population)

		// 打印每代最优个体的距离
		fmt.Printf("Generation %d: Best Distance = %.2f\n", g+1, population[0].distance)

		// 选择精英个体
		elites := population[:elitismCount]

		// 生成下一代种群
		nextPopulation := make([]Individual, populationSize)

		// 复制精英个体到下一代种群
		copy(nextPopulation, elites)

		// 生成子代个体
		for i := elitismCount; i < populationSize; i++ {
			parent1 := selectParent(population)
			parent2 := selectParent(population)

			child := crossover(parent1, parent2)
			child = mutate(child)

			nextPopulation[i] = child
		}

		// 更新当前种群为下一代种群
		population = nextPopulation
	}

	// 打印最优解
	fmt.Printf("Best Path: ")
	printPath(population[0].path)
}

// 生成城市坐标
func generateCities() []City {
	cities := make([]City, numCities)

	rand.Seed(time.Now().UnixNano())
	for i := range cities {
		cities[i] = City{
			x: rand.Intn(maxCoordinateValue),
			y: rand.Intn(maxCoordinateValue),
		}
	}

	return cities
}

// 排序种群，按适应度降序排序
func sortPopulation(population []Individual) {
	sort.Slice(population, func(i, j int) bool {
		return population[i].fitness > population[j].fitness
	})
}

// 选择父代个体
func selectParent(population []Individual) Individual {
	tournament := make([]Individual, tournamentSize)

	// 随机选择参与者
	for i := 0; i < tournamentSize; i++ {
		randomIndex := rand.Intn(populationSize)
		tournament[i] = population[randomIndex]
	}

	// 找到参与者中适应度最高的个体
	winner := tournament[0]
	for i := 1; i < tournamentSize; i++ {
		if tournament[i].fitness > winner.fitness {
			winner = tournament[i]
		}
	}

	return winner
}

// 交叉操作
func crossover(parent1 Individual, parent2 Individual) Individual {
	child := Individual{
		path:     make([]int, numCities),
		fitness:  0,
		distance: 0,
	}

	// 随机选择交叉点
	crossoverPoint := rand.Intn(numCities)

	// 从父代1中复制部分路径到子代
	copy(child.path[:crossoverPoint], parent1.path[:crossoverPoint])

	// 从父代2中复制剩余部分路径到子代
	childIndex := crossoverPoint
	for i := 0; i < numCities; i++ {
		if !containsCity(child.path, parent2.path[i]) {
			child.path[childIndex] = parent2.path[i]
			childIndex++
		}
	}

	return child
}

// 变异操作
func mutate(individual Individual) Individual {
	if rand.Float64() < mutationRate {
		// 随机选择两个位置进行交换
		index1 := rand.Intn(numCities)
		index2 := rand.Intn(numCities)

		individual.path[index1], individual.path[index2] = individual.path[index2], individual.path[index1]
	}

	return individual
}

// 检查路径中是否包含指定城市
func containsCity(path []int, city int) bool {
	for _, c := range path {
		if c == city {
			return true
		}
	}

	return false
}

// 打印路径
func printPath(path []int) {
	for _, city := range path {
		fmt.Printf("%d ", city)
	}
	fmt.Println()
}
