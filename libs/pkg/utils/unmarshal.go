package utils

import (
	"encoding/json"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
)

func UnmarshalBlog(raw map[string]interface{}) *types.Blog {
	bytes, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}

	blog := &types.Blog{}

	err = json.Unmarshal(bytes, blog)
	if err != nil {
		panic(err)
	}

	return blog
}
