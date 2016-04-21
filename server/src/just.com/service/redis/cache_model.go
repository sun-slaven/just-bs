package redis
import "just.com/common"

const (
	CACHE_MODULE_COURSE = "CPURSE"
	CACHE_MODULE_COMMENT = "COMMENT"
	CACHE_MODULE_FILE = "FILE"
	CACHE_MODULE_USER = "USER"
	CACHE_TYPE_ID = "ID"
	CACHE_TYPE_NAME = "NAME"
	CACHE_DELIMITER = "-"
)

type CacheKey struct {
	Module  string
	Type    string
	KeyData string
}

func NewCacheModel(module string, tp string, keyData string) (*CacheKey) {
	cache := new(CacheKey)
	cache.Module = module
	cache.Type = tp
	cache.KeyData = common.Base64Encoding(keyData)
	return cache
}

func (self *CacheKey) string() string {
	return self.Module + CACHE_DELIMITER + self.Type + self.KeyData
}