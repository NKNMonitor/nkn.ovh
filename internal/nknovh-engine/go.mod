module github.com/nkn.ovh/internal/nknovh-engine

go 1.22.5

replace templater v1.0.0 => ../templater

require (
	github.com/go-sql-driver/mysql v1.8.1
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/gobwas/ws v1.4.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/nknorg/nkn-sdk-go v1.4.8
	templater v1.0.0
)

require go.uber.org/multierr v1.10.0 // indirect

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.0.0 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/itchyny/base58-go v0.0.5 // indirect
	github.com/joho/godotenv v1.5.1
	github.com/nknorg/ncp-go v1.0.5 // indirect
	github.com/nknorg/nkn/v2 v2.2.1-0.20240509224846-24ade56074a3 // indirect
	github.com/nknorg/nkngomobile v0.0.0-20220615081414-671ad1afdfa9 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pbnjay/memory v0.0.0-20190104145345-974d429e7ae4 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pion/datachannel v1.5.6 // indirect
	github.com/pion/dtls/v2 v2.2.10 // indirect
	github.com/pion/ice/v3 v3.0.6 // indirect
	github.com/pion/interceptor v0.1.29 // indirect
	github.com/pion/logging v0.2.2 // indirect
	github.com/pion/mdns/v2 v2.0.7 // indirect
	github.com/pion/randutil v0.1.0 // indirect
	github.com/pion/rtcp v1.2.14 // indirect
	github.com/pion/rtp v1.8.5 // indirect
	github.com/pion/sctp v1.8.16 // indirect
	github.com/pion/sdp/v3 v3.0.9 // indirect
	github.com/pion/srtp/v3 v3.0.1 // indirect
	github.com/pion/stun/v2 v2.0.0 // indirect
	github.com/pion/transport/v2 v2.2.4 // indirect
	github.com/pion/transport/v3 v3.0.2 // indirect
	github.com/pion/turn/v3 v3.0.2 // indirect
	github.com/pion/webrtc/v4 v4.0.0-beta.17 // indirect
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	go.uber.org/zap v1.27.0
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/mobile v0.0.0-20230301163155-e0f57694e12c // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
