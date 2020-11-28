package database

import "github.com/jphacks/F_2002_1/go/domain/entity"

var Plants = entity.Plants{
	entity.Plant{
		ID:          1,
		Name:        "じゃがいも",
		NickName:    "ジャガーくん",
		Price:       1200,
		Period:      90,
		Difficulty:  1,
		Description: "僕はおいしさ満点のジャガイモだいも！！！育てばホクホクフィーバーできるいもよ！！！！",
		KitName:     "初めてでも安心！収穫キット",
		SeasonFrom:  1,
		SeasonTo:    12,
	},
	entity.Plant{
		ID:          2,
		Name:        "いちご",
		NickName:    "ベリーちゃん",
		Price:       1200,
		Period:      210,
		Difficulty:  3,
		Description: "ハァイ、私はベリーちゃん。甘いけど、私を育てるのは甘くないわよ。試してみる？",
		KitName:     "初めてでも安心！収穫キット",
		SeasonFrom:  1,
		SeasonTo:    12,
	},
	entity.Plant{
		ID:          3,
		Name:        "なす",
		NickName:    "なっくん",
		Price:       1200,
		Period:      120,
		Difficulty:  4,
		Description: "やぁ！！！！！！！ナスだよ！！！！！田楽、焼き浸し、なんでも美味しいよ！！！！！僕を選んでおくれ！！！！！",
		KitName:     "初めてでも安心！収穫キット",
		SeasonFrom:  1,
		SeasonTo:    12,
	},
	entity.Plant{
		ID:          4,
		Name:        "たまねぎ",
		NickName:    "たまちゃん",
		Price:       1200,
		Period:      80,
		Difficulty:  3,
		Description: "こんにちは。体によくて、サラダでもカレーでも何にでも使える玉ねぎ、育てたことある？",
		KitName:     "初めてでも安心！収穫キット",
		SeasonFrom:  1,
		SeasonTo:    12,
	},
	entity.Plant{
		ID:          5,
		Name:        "きゅうり",
		NickName:    "キューちゃん",
		Price:       1200,
		Period:      90,
		Difficulty:  3,
		Description: "きゅきゅ！！僕は漬物にしてよし、サラダにしてよしのジューシーきゅうりだキュ！試してみるきゅ？",
		KitName:     "初めてでも安心！収穫キット",
		SeasonFrom:  1,
		SeasonTo:    12,
	},
	entity.Plant{
		ID:          6,
		Name:        "にんじん",
		NickName:    "ジンさん",
		Price:       1200,
		Period:      120,
		Difficulty:  3,
		Description: "人参、俺の仮の名前さ……甘くて芯があるのさ……君だけの名前を、僕にくれるかい……？",
		KitName:     "初めてでも安心！収穫キット",
		SeasonFrom:  1,
		SeasonTo:    12,
	},
}


// ToDo: LowerValue はしきい値が確定次第修正する

var PlantTemperatures = entity.PlantTemperatures{
	entity.PlantTemperature{
		ID:    1,
		PlantID: 1,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantTemperature{
		ID:    2,
		PlantID: 1,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantTemperature{
		ID:    3,
		PlantID: 2,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantTemperature{
		ID:    4,
		PlantID: 2,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantTemperature{
		ID:    5,
		PlantID: 3,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantTemperature{
		ID:    6,
		PlantID: 3,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantTemperature{
		ID:    7,
		PlantID: 4,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantTemperature{
		ID:    8,
		PlantID: 4,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantTemperature{
		ID:    9,
		PlantID: 5,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantTemperature{
		ID:    10,
		PlantID: 5,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantTemperature{
		ID:    11,
		PlantID: 6,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantTemperature{
		ID:    12,
		PlantID: 6,
		Status: "high",
		LowerValue: 30.0,
	},
}

var PlantPressures = entity.PlantPressures{
	entity.PlantPressure{
		ID:    1,
		PlantID: 1,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantPressure{
		ID:    2,
		PlantID: 1,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantPressure{
		ID:    3,
		PlantID: 2,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantPressure{
		ID:    4,
		PlantID: 2,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantPressure{
		ID:    5,
		PlantID: 3,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantPressure{
		ID:    6,
		PlantID: 3,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantPressure{
		ID:    7,
		PlantID: 4,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantPressure{
		ID:    8,
		PlantID: 4,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantPressure{
		ID:    9,
		PlantID: 5,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantPressure{
		ID:    10,
		PlantID: 5,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantPressure{
		ID:    11,
		PlantID: 6,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantPressure{
		ID:    12,
		PlantID: 6,
		Status: "high",
		LowerValue: 30.0,
	},
}


var PlantIlluminances = entity.PlantIlluminances{
	entity.PlantIlluminance{
		ID:    1,
		PlantID: 1,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantIlluminance{
		ID:    2,
		PlantID: 1,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantIlluminance{
		ID:    3,
		PlantID: 2,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantIlluminance{
		ID:    4,
		PlantID: 2,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantIlluminance{
		ID:    5,
		PlantID: 3,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantIlluminance{
		ID:    6,
		PlantID: 3,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantIlluminance{
		ID:    7,
		PlantID: 4,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantIlluminance{
		ID:    8,
		PlantID: 4,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantIlluminance{
		ID:    9,
		PlantID: 5,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantIlluminance{
		ID:    10,
		PlantID: 5,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantIlluminance{
		ID:    11,
		PlantID: 6,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantIlluminance{
		ID:    12,
		PlantID: 6,
		Status: "high",
		LowerValue: 30.0,
	},
}

var PlantHumidities = entity.PlantHumidities{
	entity.PlantHumidity{
		ID:    1,
		PlantID: 1,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantHumidity{
		ID:    2,
		PlantID: 1,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantHumidity{
		ID:    3,
		PlantID: 2,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantHumidity{
		ID:    4,
		PlantID: 2,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantHumidity{
		ID:    5,
		PlantID: 3,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantHumidity{
		ID:    6,
		PlantID: 3,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantHumidity{
		ID:    7,
		PlantID: 4,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantHumidity{
		ID:    8,
		PlantID: 4,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantHumidity{
		ID:    9,
		PlantID: 5,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantHumidity{
		ID:    10,
		PlantID: 5,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantHumidity{
		ID:    11,
		PlantID: 6,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantHumidity{
		ID:    12,
		PlantID: 6,
		Status: "high",
		LowerValue: 30.0,
	},
}


var PlantSoilMoistures = entity.PlantSoilMoistures{
	entity.PlantSoilMoisture{
		ID:    1,
		PlantID: 1,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantSoilMoisture{
		ID:    2,
		PlantID: 1,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantSoilMoisture{
		ID:    3,
		PlantID: 2,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantSoilMoisture{
		ID:    4,
		PlantID: 2,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantSoilMoisture{
		ID:    5,
		PlantID: 3,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantSoilMoisture{
		ID:    6,
		PlantID: 3,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantSoilMoisture{
		ID:    7,
		PlantID: 4,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantSoilMoisture{
		ID:    8,
		PlantID: 4,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantSoilMoisture{
		ID:    9,
		PlantID: 5,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantSoilMoisture{
		ID:    10,
		PlantID: 5,
		Status: "high",
		LowerValue: 30.0,
	},
	entity.PlantSoilMoisture{
		ID:    11,
		PlantID: 6,
		Status: "low",
		LowerValue: 0.0,
	},
	entity.PlantSoilMoisture{
		ID:    12,
		PlantID: 6,
		Status: "high",
		LowerValue: 30.0,
	},
}
