package watch

import (
	"cd_platform/common"
	"cd_platform/util"
	"context"
	"encoding/base64"
)

func (s *Service) GetApplicationConfigList(ctx context.Context, project string, application string) (*common.ConfigList, error) {
	cfm, err := s.Mid.K8sclient.CMLister.ConfigMaps(util.ProjectToNS(project)).Get(application)
	if err != nil {
		util.Logger.Errorf("watch.GetApplicationConfigList err: %s", err)
		return nil, err
	}

	ret := make([]string, 0)
	for k := range cfm.Data {
		ret = append(ret, k)
	}

	return &common.ConfigList{
		ConfigNameList: ret,
	}, nil
}

func (s *Service) GetApplicationConfigDetail(ctx context.Context, project string, application string) ([]*common.ConfigDetail, error) {
	cfm, err := s.Mid.K8sclient.CMLister.ConfigMaps(util.ProjectToNS(project)).Get(application)
	if err != nil {
		util.Logger.Errorf("watch.GetApplicationConfigDetail err: %s", err)
		return nil, err
	}

	ret := make([]*common.ConfigDetail, 0)
	for k, v := range cfm.Data {
		s, err := base64.RawStdEncoding.DecodeString(v)
		if err != nil {
			util.Logger.Errorf("watch.GetApplicationConfigmapDetail err: %s", err)
			continue
		}
		tmp := &common.ConfigDetail{
			ConfigName:  k,
			ConfigValue: s,
		}
		ret = append(ret, tmp)
	}

	return ret, nil
}
