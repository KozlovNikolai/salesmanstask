```
package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	// Создаем файл для записи CPU профиля
	cpuFile, err := os.Create("cpu.out")
	if err != nil {
		fmt.Println("Error creating CPU profile:", err)
		return
	}
	defer cpuFile.Close()

	// Запускаем CPU профилирование
	err = pprof.StartCPUProfile(cpuFile)
	if err != nil {
		fmt.Println("Error starting CPU profiling:", err)
		return
	}
	defer pprof.StopCPUProfile() // Останавливаем CPU профилирование
	//#################################################################################################

	Приложение

	//###########################################################################################
	// Создаем файл для записи профиля использования памяти
	memFile, err := os.Create("mem.out")
	if err != nil {
		fmt.Println("Error creating memory profile:", err)
		return
	}
	defer memFile.Close()

	// Сохраняем профиль использования памяти
	runtime.GC() // Принудительный запуск сборщика мусора для точного профиля
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		fmt.Println("Error writing memory profile:", err)
		return
	}

	fmt.Println("Profiling completed. CPU profile saved to cpu.out, memory profile saved to mem.out.")

}
```

go tool pprof -http :8090  cpu.out 
go tool pprof -http :8090  mem.out 

