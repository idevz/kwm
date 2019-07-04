module github.com/idevz/kwm

go 1.12

require (
	github.com/idevz/kwm/pkg HEAD
	github.com/spf13/cobra v0.0.5
	k8s.io/apimachinery v0.0.0-20190703205208-4cfb76a8bf76
	k8s.io/client-go v0.0.0-20190703205746-5aceab658876
	k8s.io/klog v0.3.1
)

replace github.com/idevz/kwm/pkg => /Users/idevz/code/k8s-x/kwm/pkg
