package main

import (
	"document_service_system/interval/service"
	"fmt"
)

func main() {
	versionmanager := service.NewVersionManager(5)
	fmt.Println(versionmanager)
	versionmanager.AddVersion("version1")
	versionmanager.AddVersion("version2")
	versionmanager.AddVersion("version3")

	versionmanager.GetCurrentVersion() //3
	versionmanager.Undo()


	versionmanager.GetCurrentVersion() //2
	versionmanager.Undo()

	versionmanager.GetCurrentVersion() //1
	versionmanager.Redo()

	versionmanager.GetCurrentVersion() //2
	versionmanager.AddVersion("version4")

	versionmanager.GetCurrentVersion() //4
	versionmanager.Undo()

	versionmanager.GetCurrentVersion() //2

}
