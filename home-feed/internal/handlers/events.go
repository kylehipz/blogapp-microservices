package handlers

import (
	"context"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/constants"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/utils"
	"github.com/labstack/gommon/log"

	"github.com/kylehipz/blogapp-microservices/home-feed/internal/services"
)

type HomeFeedEventsHandler struct {
	homeFeedService *services.HomeFeedService
	events          []string
}

func NewHomeFeedEventsHandler(homeFeedService *services.HomeFeedService) *HomeFeedEventsHandler {
	return &HomeFeedEventsHandler{
		homeFeedService: homeFeedService,
		events: []string{
			constants.BLOG_CREATED,
			constants.BLOG_UPDATED,
			constants.BLOG_DELETED,
		},
	}
}

func (h *HomeFeedEventsHandler) blogCreated(payload *types.Blog) error {
	log.Info("Received event blog.created")
	if err := h.homeFeedService.AddToCacheOfFollowers(context.Background(), payload); err != nil {
		return err
	}
	return nil
}

func (h *HomeFeedEventsHandler) blogUpdated(payload *types.Blog) error {
	log.Info("Received event blog.updated")
	if err := h.homeFeedService.UpdateInCacheOfFollowers(context.Background(), payload); err != nil {
		return err
	}
	return nil
}

func (h *HomeFeedEventsHandler) blogDeleted(payload *types.Blog) error {
	log.Info("Received event blog.deleted")
	if err := h.homeFeedService.DeleteInCacheOfFollowers(context.Background(), payload); err != nil {
		return err
	}
	return nil
}

func (h *HomeFeedEventsHandler) StartListener() {
	messages := h.homeFeedService.ListenToEvents(h.events)

	for message := range messages {
		go func() {
			payload := utils.UnmarshalBlog(message.Payload.(map[string]interface{}))
			switch message.Event {
			case constants.BLOG_CREATED:
				h.blogCreated(payload)
			case constants.BLOG_UPDATED:
				h.blogUpdated(payload)
			case constants.BLOG_DELETED:
				h.blogDeleted(payload)
			}
		}()
	}
}
