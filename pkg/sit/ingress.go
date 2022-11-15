package sit

import (
	"cd_platform/common"
	"cd_platform/util"
	"context"
)

func (s *Service) InsertSitIngressRule(ctx context.Context, application string, rule *common.IngressRule) error {
	return s.Exec.InsertIngressRule(ctx, util.ToSit(application), rule)
}

func (s *Service) DeleteSitIngressRule(ctx context.Context, application string) error {
	return s.Exec.DeleteIngressRule(ctx, util.ToSit(application), util.ToSit(application))
}
