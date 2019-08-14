module github.com/yuyangjack/docker-cli

go 1.12

require (
	cloud.google.com/go v0.34.0
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78
	github.com/Microsoft/go-winio v0.4.13-0.20190408173621-84b4ab48a507
	github.com/Microsoft/hcsshim v0.8.5
	github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412
	github.com/asaskevich/govalidator v0.0.0-20180720115003-f9ffefc3facf
	github.com/beorn7/perks v0.0.0-20190414131216-e7f67b54abbe
	github.com/containerd/console v0.0.0-20181022165439-0650fd9eeb50
	github.com/containerd/continuity v0.0.0-20190426062206-aaeac12a7ffc
	github.com/containerd/fifo v0.0.0-20190226154929-a9fb20d87448
	github.com/containerd/ttrpc v0.0.0-20190411181408-699c4e40d1e7
	github.com/containerd/typeurl v0.0.0-20190228175220-2a93cfde8c20
	github.com/coreos/etcd v3.3.12+incompatible
	github.com/cpuguy83/go-md2man v1.0.8
	github.com/creack/pty v1.1.7
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v2.6.0+incompatible
	github.com/docker/compose-on-kubernetes v0.4.23
	github.com/docker/docker-credential-helpers v0.6.3
	github.com/docker/go v0.0.0-20160303222718-d30aec9fd63c
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-events v0.0.0-20170721190031-9461782956ad
	github.com/docker/go-metrics v0.0.0-20170502235133-d466d4f6fd96
	github.com/docker/go-units v0.4.0
	github.com/docker/libtrust v0.0.0-20150526203908-9cbd2a1374f4
	github.com/docker/licensing v0.0.0-20190320170819-9781369abdb5
	github.com/docker/swarmkit v0.0.0-20190412192250-59163bf75df3
	github.com/evanphx/json-patch v4.1.0+incompatible
	github.com/gofrs/flock v0.7.0
	github.com/gogo/googleapis v1.2.0
	github.com/gogo/protobuf v1.2.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.2.0
	github.com/google/go-cmp v0.2.0
	github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf
	github.com/google/shlex v0.0.0-20181106134648-c34317bd91bf
	github.com/google/uuid v1.1.1
	github.com/googleapis/gnostic v0.2.0
	github.com/gorilla/mux v1.7.0
	github.com/grpc-ecosystem/grpc-gateway v0.0.0-20170714172803-1a03ca3bad1e
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/hashicorp/go-version v0.0.0-20180322230233-23480c066577
	github.com/hashicorp/golang-lru v0.5.0
	github.com/imdario/mergo v0.3.7
	github.com/inconshreveable/mousetrap v1.0.0
	github.com/json-iterator/go v1.1.6
	github.com/konsorten/go-windows-terminal-sequences v1.0.2
	github.com/mattn/go-shellwords v1.0.5
	github.com/matttproud/golang_protobuf_extensions v1.0.1
	github.com/miekg/pkcs11 v0.0.0-20190225171305-6120d95c0e95
	github.com/mitchellh/mapstructure v0.0.0-20180715050151-f15292f7a699
	github.com/moby/buildkit v0.0.0-20190513182223-f238f1efb04f
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742
	github.com/morikuni/aec v0.0.0-20170113033406-39771216ff4c
	github.com/opencontainers/go-digest v1.0.0-rc1
	github.com/opencontainers/image-spec v1.0.1
	github.com/opencontainers/runc v1.0.1-0.20190307181833-2b18fe1d885e
	github.com/opencontainers/runtime-spec v0.0.0-20190207185410-29686dbc5559
	github.com/opentracing/opentracing-go v0.0.0-20171003133519-1361b9cd60be
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v0.8.0
	github.com/prometheus/client_model v0.0.0-20170216185247-6f3806018612
	github.com/prometheus/common v0.0.0-20180518154759-7600349dcfe1
	github.com/prometheus/procfs v0.0.0-20180612222113-7d6f385de8be
	github.com/russross/blackfriday v0.0.0-20160531111224-1d6b8e9301e7
	github.com/shurcooL/sanitized_anchor_name v0.0.0-20151028001915-10ef21a441db
	github.com/sirupsen/logrus v1.4.1
	github.com/spf13/cobra v0.0.3
	github.com/syndtr/gocapability v0.0.0-20180916011248-d98352740cb2
	github.com/theupdateframework/notary v0.6.1
	github.com/tonistiigi/fsutil v0.0.0-20190327153851-3bbb99cdbd76
	github.com/tonistiigi/units v0.0.0-20180711220420-6950e57a87ea
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415
	github.com/xeipuuv/gojsonschema v0.0.0-20160323030313-93e72a773fad
	github.com/yuyangjack/distribution v0.0.0-20190814083944-f4e8d7ef141f
	golang.org/x/crypto v0.0.0-20190411191339-88737f569e3a
	golang.org/x/net v0.0.0-20190522155817-f3200d17e092
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys v0.0.0-20190602015325-4c4f7f33c9ed
	golang.org/x/text v0.3.0
	golang.org/x/time v0.0.0-20180412165947-fbb02b2291d2
	google.golang.org/grpc v1.20.1
	gopkg.in/inf.v0 v0.9.1
	gopkg.in/yaml.v2 v2.2.1
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.0.0-20190313235455-40a48860b5ab
	k8s.io/apimachinery v0.0.0-20190313205120-d7deff9243b1
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v0.2.0
	k8s.io/kube-openapi v0.0.0-20190320154901-5e45bb682580
	k8s.io/utils v0.0.0-20190308190857-21c4ce38f2a7
	sigs.k8s.io/yaml v1.1.0
	vbom.ml/util v0.0.0-20170409195630-256737ac55c4
)
