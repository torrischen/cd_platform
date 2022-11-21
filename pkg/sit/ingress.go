package sit

import (
	"cd_platform/common"
	"cd_platform/util"
	"context"
)

func (s *Service) InsertSitIngressRule(ctx context.Context, rule *common.IngressRule) error {
	rule.Project = util.ToSit(rule.Project)
	return s.Exec.InsertIngressRule(ctx, rule)
}

func (s *Service) DeleteSitIngressRule(ctx context.Context, application string) error {
	return s.Exec.DeleteIngressRule(ctx, util.ToSit(application), util.ToSit(application))
}

func (s *Service) DeleteSpecifiedSitIngressRule(ctx context.Context, path string) error {
	return s.Exec.DeleteSpecifiedIngressRule(ctx, path)
}
