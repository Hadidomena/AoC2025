package day8

import (
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X, Y, Z int
}

type Edge struct {
	P1, P2     int
	DistanceSq int64
}

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{parent: parent, size: size}
}

func (d *DSU) Find(i int) int {
	if d.parent[i] == i {
		return i
	}
	d.parent[i] = d.Find(d.parent[i])
	return d.parent[i]
}

func (d *DSU) Union(i, j int) {
	rootI := d.Find(i)
	rootJ := d.Find(j)
	if rootI != rootJ {
		if d.size[rootI] < d.size[rootJ] {
			rootI, rootJ = rootJ, rootI
		}
		d.parent[rootJ] = rootI
		d.size[rootI] += d.size[rootJ]
	}
}

func parsePoints(lines []string) []Point {
	points := make([]Point, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		z, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
		points[i] = Point{X: x, Y: y, Z: z}
	}
	return points
}

func distanceSq(p1, p2 Point) int64 {
	dx := int64(p1.X - p2.X)
	dy := int64(p1.Y - p2.Y)
	dz := int64(p1.Z - p2.Z)
	return dx*dx + dy*dy + dz*dz
}

func SolvePart1(lines []string, connectionsToMake int) int {
	points := parsePoints(lines)
	numPoints := len(points)

	var edges []Edge
	for i := 0; i < numPoints; i++ {
		for j := i + 1; j < numPoints; j++ {
			dist := distanceSq(points[i], points[j])
			edges = append(edges, Edge{P1: i, P2: j, DistanceSq: dist})
		}
	}

	sort.SliceStable(edges, func(i, j int) bool {
		return edges[i].DistanceSq < edges[j].DistanceSq
	})

	dsu := NewDSU(numPoints)
	for i := 0; i < connectionsToMake && i < len(edges); i++ {
		edge := edges[i]
		dsu.Union(edge.P1, edge.P2)
	}

	componentSizes := make(map[int]int)
	for i := 0; i < numPoints; i++ {
		root := dsu.Find(i)
		componentSizes[root] = dsu.size[root]
	}

	var sizes []int
	for _, size := range componentSizes {
		sizes = append(sizes, size)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	if len(sizes) < 3 {
		result := 1
		for _, s := range sizes {
			result *= s
		}
		return result
	}

	return sizes[0] * sizes[1] * sizes[2]
}

func SolveSecondPart(lines []string) int64 {
	return 0
}
