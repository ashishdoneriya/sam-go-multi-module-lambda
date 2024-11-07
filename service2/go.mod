module my-multi-module-project/service2

go 1.23.2

require my-multi-module-project/common v0.0.0
replace my-multi-module-project/common => ../common
require github.com/aws/aws-lambda-go v1.36.1
replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8

