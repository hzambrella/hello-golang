package router
import(
	"gopkg.in/redis.v4"
	"log"
)

var redisClient *redis.Client

func init(){
	redisClient=CreateRedis()
}

func CreateRedis()(*redis.Client){
	redisAddr:="6379"
	client:=redis.NewClient(&redis.Options{
		Addr:"localhost:"+redisAddr,
		Password:"",
		DB:0,
	})

	pong,err:=client.Ping().Result()
	if err!=nil{
		panic("redis fail"+err.Error())
	}else{
		log.Println("redis success",pong)
	}

	return client
}
