package storage

import "github.com/garyburd/redigo/redis"

// Redis a manager for redis connections
type Redis struct {
	Connections chan redis.Conn // connection pool to redis server
}

// Get performes a get command to the redis server
func (r *Redis) Get(s string) ([]byte, error) {
	c := <-r.Connections
	b, err := c.Do("GET", s)
	if err != nil {
		r.Connections <- c
		return b, err
	}
	r.Connections <- c
	return []byte(b), err

}

// Set performes a set command to the redis server
func (r *Redis) Set(s string, b []byte) error {
	c := <-r.Connections
	_, err := c.Do("SET", s, b)
	r.Connections < -c
	return err
}

// NewRedis create a connection pool and return an initialized redis providor
func NewRedis(host string, poolSize int) (*Redis, error) {
	r := &Redis{Connections: make(chan redis.Conn, poolSize)}
	// fill the pool
	for i := 0; i < poolSize; i++ {
		c, err := redis.Dial("tcp", host)
		if err != nil {
			return nil, err
		}
		r.Connections <- c
	}
	return r, nil
}
