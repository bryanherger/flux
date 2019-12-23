package runtime

import (
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/codes"
	"github.com/influxdata/flux/internal/errors"
	"github.com/influxdata/flux/values"
)

const versionFuncName = "version"

var errBuildInfoNotPresent = errors.New(codes.NotFound, "build info is not present")

func init() {
	flux.RegisterPackageValue("runtime", versionFuncName, values.NewFunction(
		versionFuncName,
		flux.LookupBuiltInType("runtime", versionFuncName),
		Version,
		false,
	))
}
