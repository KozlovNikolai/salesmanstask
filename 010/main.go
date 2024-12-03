package main

import (
	"fmt"
	"salesmanstask/010/app"
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

	app.Run(store)

	output(store)
	fmt.Println("##########################################################################")
	app.Run(store)

	output(store)
}

func output(store *app.Store) {
	fmt.Println("Все узлы:")
	for key, value := range store.Tree {
		// fmt.Printf("Key: %d, Node: %+v\n", key, value)
		fmt.Printf(
			"Key: %d, NodeID: %d, Name: %s, ParID: %d, (%d,%d) W: %d\n",
			key,
			value.ID,
			value.Name,
			value.ParentID,
			value.Out,
			value.In,
			value.W,
		)
	}
	fmt.Println()

	fmt.Println("Только листья:")
	fmt.Printf("индекс узла с минимальным весом %d: %d\n",
		store.Leaves.MinWeight,
		store.Leaves.MinWeightID)
	for idx := range store.Leaves.NodeIDs {
		fmt.Printf("index: %d,\n", idx)
	}

	fmt.Printf("Текущий узел: %+v\n", store.Tree[store.CurrentNodeID])

	for i, v := range store.Tree {
		fmt.Printf("Node ID:%d, Name:%s, (%d,%d), W:%d, ParID:%d\n", i, v.Name, v.Out, v.In, v.W, v.ParentID)
		app.PrintMatrix(v.MX)
	}
}