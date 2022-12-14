package controllers

import (
	"context"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"go.mongodb.org/mongo-driver/bson"
	"ks8-demo/connectdb"
	"ks8-demo/models"
	"math/rand"
)

type ChooseOne struct {
	beego.Controller
}

func (a *ChooseOne) Get() {
	client := connectdb.GetMgoCli()
	collection := client.Database("newdb").Collection("runood")
	res, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println("1", err)
	}
	var result models.ChooseOne
	var alldb []models.Xl

	//result = new(models.Showall)

	res2 := res.All(context.TODO(), &alldb)
	if res2 != nil {
		fmt.Println(res2)
	}
	l := len(alldb)
	l = rand.Intn(l)

	//fmt.Println(l)
	//l = int(ele) * l * 100

	//fmt.Println(" l ", l)
	/*for i := range alldb {
		fmt.Println(alldb[i])
	}*/
	result.D = alldb[l].Name
	result.Message3 = "选择幸运儿成功！"
	//fmt.Println(result.D)
	a.Data["json"] = &result
	a.ServeJSON()
	result.D = nil

	//fmt.Println("end")
}
