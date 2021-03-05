package model

import (
	"log"
	"shuwen-redis/src/chapter01/common"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v7"
)

type Article interface {
	ArticleVote(string, string)
	PostArticle(string, string, string) string
	GetArticles(int64, string) []map[string]string
	AddRemoveGroups(string, []string, []string)
	GetGroupArticles(string, string, int64, []map[string]string)
	Reset()
}

type ArticleRepo struct {
	Conn *redis.Client
}

func NewArticleRepo(conn *redis.Client) *ArticleRepo {
	return &ArticleRepo{Conn: conn}
}

func (r *ArticleRepo) ArticleVote(article, user string) {
	cutoff := time.Now().Unix() - common.OneWeekInSeconds
	if r.Conn.ZScore("time", article).Val() < float64(cutoff) {
		return
	}
	articleId := strings.Split(article, ":")[1]
	if r.Conn.SAdd("voted:"+articleId, user).Val() != 0 {
		r.Conn.ZIncrBy("score", common.VoteScore, article)
		r.Conn.HIncrByFloat(article, "votes", 1)
	}
}

// 发布并获取文章
func (r *ArticleRepo) PostArticle(user, title, link string) string {
	articleId := strconv.Itoa(int(r.Conn.Incr("article:").Val())) // 生成新的文章ID
	voted := "voted:" + articleId
	r.Conn.SAdd(voted, user)                                  // 将文章发布者的ID放入集合中
	r.Conn.Expire(voted, common.OneWeekInSeconds*time.Second) // 设置过期时间：一周

	now := time.Now().Unix()
	article := "article:" + articleId
	r.Conn.HMSet(article, map[string]interface{}{ // 将文章信息存储到散列里面
		"title":  title,
		"link":   link,
		"poster": user,
		"time":   now,
		"votes":  1,
	})
	// 将文章添加到根据发布时间排序的有序集合和根据评分排序的有序集合里面
	r.Conn.ZAdd("score:", &redis.Z{Score: float64(now + common.VoteScore), Member: article})
	r.Conn.ZAdd("time", &redis.Z{Score: float64(now), Member: article})
	return articleId
}

// GetArticles 获取文章
func (r *ArticleRepo) GetArticles(page int64, order string) []map[string]string {
	if order == "" {
		order = "score:"
	}
	start := (page - 1) * common.ArticlesPerPage     // 获取文章的起始索引
	end := start + common.ArticlesPerPage - 1        // 获取文章的结束索引
	ids := r.Conn.ZRevRange(order, start, end).Val() // 获取多个文章ID
	articles := []map[string]string{}                // 根据文章的ID获取文章的详细信息
	for _, id := range ids {
		articleData := r.Conn.HGetAll(id).Val()
		articleData["id"] = id
		articles = append(articles, articleData)
	}
	return articles
}

// AddRemoveGroups 对文章进行分组
func (r *ArticleRepo) AddRemoveGroups(articleId string, toAdd, toRemove []string) {
	article := "article:" + articleId // 构建存储文章信息的键名
	for _, group := range toAdd {     // 将文章添加到它所属的群组里面
		r.Conn.SAdd("group:"+group, article)
	}
	for _, group := range toRemove { // 从群组里面移除文章
		r.Conn.SRem("group:"+group, article)
	}
}

// GetGroupArticles 从群组里面获取整页文章
func (r *ArticleRepo) GetGroupArticles(group, order string, page int64) []map[string]string {
	if order == "" {
		order = "score"
	}
	key := order + group               // 为每个群组的每种排列顺序都创建一个键
	if r.Conn.Exists(key).Val() == 0 { // 检查是否有已缓存的排序结果，如果没有的话就现在进行排序
		res := r.Conn.ZInterStore(key, &redis.ZStore{Aggregate: "MAX", Keys: []string{"group:" + group, order}}).Val()
		if res <= 0 {
			log.Println("ZInterStore")
		}
	}
	r.Conn.Expire(key, 60*time.Second)
	return r.GetArticles(page, key)
}

// Flushall
func (r *ArticleRepo) Reset() {
	r.Conn.FlushDB()
}
