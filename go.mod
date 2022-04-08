module aftermath.link/repo/am-bot-stats

go 1.17

// replace github.com/andersfylling/disgord => github.com/byvko-dev/disgord v0.35.3
replace github.com/andersfylling/disgord => ../../disgord

replace github.com/byvko-dev/am-core => ../am-core

require (
	github.com/andersfylling/disgord v0.35.1
	github.com/byvko-dev/am-core v1.2.6
	github.com/byvko-dev/am-types v1.1.13
	github.com/joho/godotenv v1.4.0
	github.com/sirupsen/logrus v1.8.1
)

require (
	github.com/andersfylling/snowflake/v5 v5.0.1 // indirect
	github.com/klauspost/compress v1.15.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	go.mongodb.org/mongo-driver v1.8.4 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064 // indirect
	golang.org/x/net v0.0.0-20220325170049-de3da57026de // indirect
	golang.org/x/sys v0.0.0-20220327210214-530d0810a4d0 // indirect
	nhooyr.io/websocket v1.8.7 // indirect
)
