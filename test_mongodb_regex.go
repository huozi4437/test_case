//mongoDB用正则表达式做查找条件
package main

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var PhonemeProgressDao *mgo.Collection

type PhonemeProgress struct {
	PhonemeId       string `bson:"_id" gorm:"column:id;primary_key"`
	OffsetSize      int64  `bson:"offsetSize" gorm:"column:offsetSize"`
	TotalSize       int64  `bson:"totalSize" gorm:"column:totalSize"`
	ExistPhoneme    bool   `bson:"existPhoneme" gorm:"column:existPhoneme"`
	CreateTimeStamp int64  `bson:"create_timestamp" gorm:"column:create_timestamp"`
	UpdateTimeStamp int64  `bson:"update_timestamp" gorm:"column:update_timestamp"`
	PhonemeNotFound bool   `bson:"phoneme_not_found" gorm:"column:phoneme_not_found"`
	Available       int    `bson:"available" gorm:"column:available"`
	HideTimeStamp   int64  `bson:"hideTimeStamp" gorm:"column:hideTimeStamp"`
	Finish          bool   `bson:"finish" gorm:"column:finish"`
	DataVersion     int64  `bson:"dataversion" gorm:"column:data_version" desc:"data syn time"`
	SelectBeginTime int    `bson:"select_begin_time" gorm:"column:select_begin_time"` // millisecond
	SelectEndTime   int    `bson:"select_end_time" gorm:"column:select_end_time"`
}

func initDb() error {
	sess, err := mgo.Dial("mongodb://speakin:speakin@192.168.1.215:27017")
	if err != nil {
		fmt.Println("initDb err:", err)
		return err
	}
	PhonemeProgressDao = sess.DB("identify_system_offline").C("phoneme_progress")

	return nil
}

func Find(param bson.M) error {
	retList := make([]*PhonemeProgress, 0)
	err := PhonemeProgressDao.Find(param).All(&retList)
	if err != nil {
		fmt.Println("Find err:", err)
		return err
	}

	for _, progress := range retList {
		fmt.Println(progress)
	}

	return nil
}

func main() {
	err := initDb()
	if err != nil {
		panic("initDb")
	}
	// _id:{$regex:"file20190815131619_c2dc2bda5c8840a8b2ad0ed62fd70cb9",$options:"$*"}
	param := bson.M{
		"_id": bson.M{
			"$regex":   "file20190815131619_c2dc2bda5c8840a8b2ad0ed62fd70cb9",
			"$options": "$*",
		},
	}
	Find(param)
}
