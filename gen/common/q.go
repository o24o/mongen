package common

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Q[T any] struct {
	collection *mongo.Collection
	ctx        context.Context
	filter     []QueryCondition
}

func (q *Q[T]) Collection(collection *mongo.Collection) *Q[T] {
	q.collection = collection
	return q
}

// WithContext 设置查询上下文
func (q *Q[T]) WithContext(ctx context.Context) *Q[T] {
	q.ctx = ctx
	return q
}

// Where 添加过滤条件
func (q *Q[T]) Where(condition ...QueryCondition) *Q[T] {
	q.filter = append(q.filter, condition...)
	return q
}

// First 返回第一个匹配的结果
func (q *Q[T]) First() (*T, error) {
	var result T
	filter := bson.D{}
	for _, c := range q.filter {
		filter = append(filter, bson.E{Key: c.Field, Value: bson.D{{c.Op, c.Value}}})
	}
	err := q.collection.FindOne(q.ctx, filter).Decode(&result)
	return &result, err
}

// Find 返回所有匹配的结果
func (q *Q[T]) Find() ([]*T, error) {
	filter := bson.D{}
	for _, c := range q.filter {
		filter = append(filter, bson.E{Key: c.Field, Value: bson.D{{c.Op, c.Value}}})
	}
	cursor, err := q.collection.Find(q.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(q.ctx)

	var results []*T
	if err = cursor.All(q.ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}
