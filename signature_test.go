package signature

import "testing"

func TestSortBody(t *testing.T) {
	body := `{
		"string": "teststring", 
		"int": 8923, 
		"map": {
			"name": "hello"
		},
		"int64": 89089, 
		"float64": 9.999, 
		"slice": [
			1,
			2,
			3
		],
		"sliceMap": [
			{
				"name": "hello"
			},
			{
				"name": "world"
			}
		],
		"boolean": true,
		"sliceSlice":[
			[
				1
			],
			[
				2
			]
		]
	}`
	body1 := `{
		"boolean": true,
		"string": "teststring", 
		"int": 8923, 
		"map": {
			"name": "hello"
		},
		"int64": 89089, 
		"float64": 9.999, 
		"slice": [
			1,
			2,
			3
		],
		"sliceSlice":[
			[
				1
			],
			[
				2
			]
		],
		"sliceMap": [
			{
				"name": "hello"
			},
			{
				"name": "world"
			}
		]
	}`
	str, err := SignatureJson([]byte(body))
	if err != nil {
		t.Fatalf("error: [%s]", err.Error())
	}
	if str != "{boolean:true,float64:9.999,int:8923,int64:89089,map:{name:hello},slice:[1,2,3],sliceMap:[{name:hello},{name:world}],sliceSlice:[[1],[2]],string:teststring}" {
		t.Fatalf("生成字符串与期望不符合")
	}
	str1, err1 := SignatureJson([]byte(body1))
	if err1 != nil {
		t.Fatalf("error: [%s]", err.Error())
	}
	if str1 != str {
		t.Fatalf("参数顺序交换， 生成签名字符串不一致")
	}
}
