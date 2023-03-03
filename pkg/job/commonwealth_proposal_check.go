package job

import (
	"sync"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
)

type commonwealthProposalData struct {
	entity  *entities.Entity
	service *service.Service
	log     logger.Logger
}

func NewCommonwealthProposalData(e *entities.Entity, svc *service.Service, l logger.Logger) Job {
	return &commonwealthProposalData{
		entity:  e,
		service: svc,
		log:     l,
	}
}

func (job *commonwealthProposalData) Run() error {
	communities, err := job.entity.GetAllCommonwealthData()
	if err != nil {
		job.log.Errorf(err, "[entity.GetAllCommonwealthData] failed")
		return err
	}
	wg := sync.WaitGroup{}
	// get threads from every community
	for _, data := range communities {
		// get threads from community
		cmwData, err := job.entity.GetAllCommonwealthThreads(data.CommunityID)
		if err != nil {
			job.log.Fields(logger.Fields{"community_id": data.CommunityID}).Errorf(err, "failed to get Commonwealth thread")
			continue
		}
		// get thread create after latest db time
		newThreads := []response.CommonwealthDiscussion{}
		for _, t := range *cmwData.Result.Threads {
			if t.CreatedAt.After(data.LatestAt) {
				newThreads = append(newThreads, t)
			}
		}
		if len(newThreads) == 0 {
			continue
		}
		// update commonwealth latest data
		job.entity.UpdateCommonwealthData(model.CommonwealthLatestData{
			CommunityID: data.CommunityID,
			Name:        data.Name,
			Description: data.Description,
			IconURL:     data.IconURL,
			Website:     data.Website,
			PostCount:   data.PostCount,
			LatestAt:    newThreads[0].CreatedAt,
		})
		// get matching config
		configs, err := job.entity.GetAllDaoTrackerBySpaceAndSource(data.CommunityID, "commonwealth")
		if err != nil {
			job.log.Fields(logger.Fields{"community_id": data.CommunityID}).Errorf(err, "failed to get dao tracker configs")
			continue
		}

		// send thread to channel
		for _, thr := range newThreads {
			for _, cfg := range configs {
				wg.Add(1)
				go func(channelId string, discussion response.CommonwealthDiscussion, community model.CommonwealthLatestData) {
					defer wg.Done()
					req := request.NewCommonwealthDiscussionRequest{
						ChannelID:  channelId,
						Discussion: discussion,
						Community:  community,
					}
					job.service.Discord.NotifyNewCommonwealthDiscussion(req)
				}(cfg.ChannelID, thr, data)
			}
		}
	}
	wg.Wait()
	return nil
}
