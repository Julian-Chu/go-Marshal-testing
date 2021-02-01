package go_Marshal_testing

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Benchmark_MarshalIndent_MapToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var items []interface{} = getMapInterface()
		responseMap := map[string]interface{}{
			"data": map[string]interface{}{
				"items": items,
			},
		}
		json.MarshalIndent(responseMap, "", "  ")
	}
}

func Benchmark_MarshalIndent_StructToBytes_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var items []interface{} = getArticleToInterface()
		responseMap := map[string]interface{}{
			"data": map[string]interface{}{
				"items": items,
			},
		}
		json.MarshalIndent(responseMap, "", "  ")
	}
}

func Benchmark_MarshalIndent_StructToBytes_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var items []Article = getArticles()
		responseMap := map[string]interface{}{
			"data": map[string][]Article{
				"items": items,
			},
		}
		json.MarshalIndent(responseMap, "", "  ")
	}
}

func Benchmark_MarshalIndent_StructToBytes_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var items []Article = getArticles()
		response := struct {
			Data struct {
				Items []Article `json:"items"`
			} `json:"data"`
		}{
			Data: struct {
				Items []Article `json:"items"`
			}{
				Items: items,
			},
		}
		json.MarshalIndent(response, "", "  ")
	}
}

func Benchmark_Marshal_MapToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var items []interface{} = getMapInterface()
		responseMap := map[string]interface{}{
			"data": map[string]interface{}{
				"items": items,
			},
		}
		json.Marshal(responseMap)
	}
}

func Benchmark_Marshal_StructToBytes_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var items []interface{} = getArticleToInterface()
		responseMap := map[string]interface{}{
			"data": map[string]interface{}{
				"items": items,
			},
		}
		json.Marshal(responseMap)
	}
}

func Benchmark_Marshal_StructToBytes_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var items []Article = getArticles()
		responseMap := map[string]interface{}{
			"data": map[string][]Article{
				"items": items,
			},
		}
		json.Marshal(responseMap)
	}
}

func Benchmark_Marshal_StructToBytes_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var items []Article = getArticles()
		response := struct {
			Data struct {
				Items []Article `json:"items"`
			} `json:"data"`
		}{
			Data: struct {
				Items []Article `json:"items"`
			}{
				Items: items,
			},
		}

		json.Marshal(response)
	}
}

func Benchmark_Unmarshal_ToMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		response := make(map[string]interface{})
		json.Unmarshal([]byte(jsonStr), &response)
	}
}

func Benchmark_Unmarshal_ToStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		response := struct {
			Data struct {
				Items []Article `json:"items"`
			} `json:"data"`
		}{}
		json.Unmarshal([]byte(jsonStr), &response)
	}
}

type Article struct {
	Filename  string `json:"filename"`
	Modified  string `json:"modified_time"`
	Recommend string `json:"recommend"`
	Date      string `json:"post_date"`
	Title     string `json:"title"`
	Money     string `json:"money"`
	Owner     string `json:"owner"`
}

var articles = []Article{
	{
		Filename:  "test",
		Modified:  "test",
		Recommend: "test",
		Date:      "test",
		Title:     "test",
		Money:     "test",
		Owner:     "test",
	},
	{
		Filename:  "test",
		Modified:  "test",
		Recommend: "test",
		Date:      "test",
		Title:     "test",
		Money:     "test",
		Owner:     "test",
	},
	{
		Filename:  "test",
		Modified:  "test",
		Recommend: "test",
		Date:      "test",
		Title:     "test",
		Money:     "test",
		Owner:     "test",
	},
	{
		Filename:  "test",
		Modified:  "test",
		Recommend: "test",
		Date:      "test",
		Title:     "test",
		Money:     "test",
		Owner:     "test",
	},
	{
		Filename:  "test",
		Modified:  "test",
		Recommend: "test",
		Date:      "test",
		Title:     "test",
		Money:     "test",
		Owner:     "test",
	},
}

func getArticleToInterface() []interface{} {
	var items []interface{}
	for _, f := range articles {
		m := Article{
			Filename:  f.Filename,
			Modified:  f.Modified,
			Recommend: f.Recommend,
			Date:      f.Date,
			Title:     f.Title,
			Money:     f.Money,
			Owner:     f.Owner,
		}
		items = append(items, m)
	}
	return items
}

func getArticles() []Article {
	var items []Article
	for _, f := range articles {
		m := Article{
			Filename:  f.Filename,
			Modified:  f.Modified,
			Recommend: f.Recommend,
			Date:      f.Date,
			Title:     f.Title,
			Money:     f.Money,
			Owner:     f.Owner,
		}
		items = append(items, m)
	}
	return items
}

func getMapInterface() []interface{} {
	var items []interface{}
	for _, f := range articles {
		m := map[string]interface{}{
			"filename": f.Filename,
			// Bug(pichu): f.Modified time will be 0 when file is vote
			"modified_time":   f.Modified,
			"recommend_count": f.Recommend,
			"post_date":       f.Date,
			"title":           f.Title,
			"money":           fmt.Sprintf("%v", f.Money),
			"owner":           f.Owner,
			// "aid": ""
			"url": "url" + f.Filename,
		}
		items = append(items, m)
	}
	return items
}

var jsonStr = `{
"data": {
"items": [
{
"filename": "test",
"modified_time": "test",
"money": "test",
"owner": "test",
"post_date": "test",
"recommend_count": "test",
"title": "test",
"url": "urltest"
},
{
"filename": "test",
"modified_time": "test",
"money": "test",
"owner": "test",
"post_date": "test",
"recommend_count": "test",
"title": "test",
"url": "urltest"
},
{
"filename": "test",
"modified_time": "test",
"money": "test",
"owner": "test",
"post_date": "test",
"recommend_count": "test",
"title": "test",
"url": "urltest"
},
{
"filename": "test",
"modified_time": "test",
"money": "test",
"owner": "test",
"post_date": "test",
"recommend_count": "test",
"title": "test",
"url": "urltest"
},
{
"filename": "test",
"modified_time": "test",
"money": "test",
"owner": "test",
"post_date": "test",
"recommend_count": "test",
"title": "test",
"url": "urltest"
}
]
}
}`
