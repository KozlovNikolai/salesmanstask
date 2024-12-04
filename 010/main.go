package main

import (
	"fmt"
	"salesmanstask/010/app"
	"slices"
)

func main() {
	// именуем строки и столбцы
	app.SetNaming(app.Matrixes[0])
	app.PrintMatrix(app.GetRootMatrix())

	// находим нижнюю границу
	lb := app.GetLB(app.GetRootMatrix())

	// создаем хранилище узлов
	store := app.NewStore(app.GetRootMatrix(), lb)
	output(store)
	for i := 0; i < 6; i++ {
		fmt.Printf("%d ##########################################################################\n", i)
		app.Run(store)
		output(store)
	}
}

func output(store *app.Store) {
	fmt.Println("Все узлы:")
	for key := 0; key < len(store.Tree); key++ {
		fmt.Printf(
			"Key: %d, NodeID: %d, Name: %s, ParID: %d, (%d,%d) W: %d\n",
			key,
			store.Tree[key].ID,
			store.Tree[key].Name,
			store.Tree[key].ParentID,
			store.Tree[key].Out,
			store.Tree[key].In,
			store.Tree[key].W,
		)
	}
	// for key, value := range store.Tree {
	// 	// fmt.Printf("Key: %d, Node: %+v\n", key, value)
	// 	fmt.Printf(
	// 		"Key: %d, NodeID: %d, Name: %s, ParID: %d, (%d,%d) W: %d\n",
	// 		key,
	// 		value.ID,
	// 		value.Name,
	// 		value.ParentID,
	// 		value.Out,
	// 		value.In,
	// 		value.W,
	// 	)
	// }
	fmt.Println()

	fmt.Println("Только листья:")
	fmt.Printf("индекс узла с минимальным весом %d: %d\n",
		store.Leaves.MinWeight,
		store.Leaves.MinWeightID)
	var leaves []int

	for key := range store.Leaves.NodeIDs {
		leaves = append(leaves, key)
		// fmt.Printf("ID: %d,\n", key)
	}
	slices.Sort(leaves)
	for _, id := range leaves {
		v := store.Tree[id]
		fmt.Printf("ID: %d, (%d,%d) W:%d\n", id, v.Out, v.In, v.W)
	}

	fmt.Printf("Текущий узел: %+v\n", store.Tree[store.CurrentNodeID])

	// for i, v := range store.Tree {
	// 	fmt.Printf("Node ID:%d, Name:%s, (%d,%d), W:%d, ParID:%d\n", i, v.Name, v.Out, v.In, v.W, v.ParentID)
	// 	app.PrintMatrix(v.MX)
	// }
	for i := 0; i < len(store.Tree); i++ {
		v := store.Tree[i]
		fmt.Printf("Node ID:%d, Name:%s, (%d,%d), W:%d, ParID:%d\n", i, v.Name, v.Out, v.In, v.W, v.ParentID)
		app.PrintMatrix(v.MX)
	}
}
