package main

import (
	"github.com/myfantasy/mdp"
	"github.com/myfantasy/myfdb/istorage"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetLevel(log.DebugLevel)

	c, err := mdp.ConnectionGetFromJSON([]byte(`{"server":"http://localhost:9170/"}`))
	log.Debugln(c)
	if err != nil {
		log.Fatalln(err)
	}

	serviceStructTest(c)

	structGetQuery(c)

	structSetQuery(c)

	itemGet(c)

}

func itemGet(c *mdp.Connection) {

	igq := mdp.ItemsGetQuery{
		TableName: "test_i",
		IKey:      4,
	}

	ig := c.ItemsRawGet(igq)

	igqOut(ig)

	// ------------------

	ii := mdp.ItemInt{
		Key: 4,
	}

	isq := mdp.ItemsSetQuery{
		TableName: "test_i",
		IItem:     &ii,
	}

	ig = c.ItemsRawSet(isq)

	igqOut(ig)

	// ------------------

	ig = c.ItemsRawGet(igq)

	igqOut(ig)

	// ------------------

	isq.IItem.IsRemoved = true
	//isq.IItem.IsRemoved = false

	ig = c.ItemsRawSet(isq)

	igqOut(ig)
}

func igqOut(ig mdp.ItemsGet) {
	if ig.InternalErr != nil {
		log.Errorln(ig.InternalErr)
	}
	if ig.ParamsErr != nil {
		log.Errorln(ig.ParamsErr)
	}
	log.Infoln("ig: ", ig)
}

func structSetQuery(c *mdp.Connection) {

	var ssq mdp.StructSetQuery

	ssq.CreateTable = &mdp.TableDefinition{
		TableName: "test_i",
		TableType: istorage.IntMapLocalTableTableType,
		KeyType:   mdp.KeyTypeInt,
	}

	ssg := c.StructRawSet(ssq)

	if ssg.InternalErr != nil {
		log.Errorln(ssg.InternalErr)
	}
	if ssg.ParamsErr != nil {
		log.Errorln(ssg.ParamsErr)
	}
	log.Infoln("ssg: ", ssg)

	var sgq mdp.StructGetQuery

	sgq.TableName = "test_i"

	sg := c.StructRawGet(sgq)

	if sg.InternalErr != nil {
		log.Fatalln(sg.InternalErr)
	}
	if sg.ParamsErr != nil {
		log.Fatalln(sg.ParamsErr)
	}
	log.Infoln("sg: ", sg)

	sgq.LoadInternalInfo = true

	sg = c.StructRawGet(sgq)

	if sg.InternalErr != nil {
		log.Fatalln(sg.InternalErr)
	}
	if sg.ParamsErr != nil {
		log.Fatalln(sg.ParamsErr)
	}
	log.Infoln("sg: ", sg)
}

func structGetQuery(c *mdp.Connection) {
	var sgq mdp.StructGetQuery

	sg := c.StructRawGet(sgq)

	if sg.InternalErr != nil {
		log.Fatalln(sg.InternalErr)
	}
	if sg.ParamsErr != nil {
		log.Fatalln(sg.ParamsErr)
	}
	log.Infoln("sg: ", sg)

}

func serviceStructTest(c *mdp.Connection) {

	var ssgq mdp.StructStorageGetQuery

	ssgq.Name = true

	ssg := c.StructStorageRawGet(ssgq)

	if ssg.InternalErr != nil {
		log.Fatalln(ssg.InternalErr)
	}
	if ssg.ParamsErr != nil {
		log.Fatalln(ssg.ParamsErr)
	}

	log.Infoln("ssg.Storage.StorageName: ", ssg.Storage.StorageName)

	var sssq mdp.StructStorageSetQuery

	sssq.Name = "s2"

	ssg = c.StructStorageRawSet(sssq)

	if ssg.InternalErr != nil {
		log.Fatalln(ssg.InternalErr)
	}
	if ssg.ParamsErr != nil {
		log.Fatalln(ssg.ParamsErr)
	}

	log.Infoln("ssg.Storage.StorageName: ", ssg.Storage.StorageName)

	sssq.Name = "s1"

	ssg = c.StructStorageRawSet(sssq)

	if ssg.InternalErr != nil {
		log.Fatalln(ssg.InternalErr)
	}
	if ssg.ParamsErr != nil {
		log.Fatalln(ssg.ParamsErr)
	}

	log.Infoln("ssg.Storage.StorageName: ", ssg.Storage.StorageName)

	// Stop service
	if false {
		sssq.Stop = true
		sssq.Name = ""

		ssg = c.StructStorageRawSet(sssq)

		if ssg.InternalErr != nil {
			log.Fatalln(ssg.InternalErr)
		}
		if ssg.ParamsErr != nil {
			log.Fatalln(ssg.ParamsErr)
		}

		log.Infoln("ssg.Storage.StorageName: ", ssg.Storage.StorageName)
	}
}
