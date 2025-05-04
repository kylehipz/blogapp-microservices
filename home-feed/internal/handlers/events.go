package handlers

import (
	"fmt"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/constants"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/utils"

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
	fmt.Println("Blog created event handled")
	return nil
}

func (h *HomeFeedEventsHandler) blogUpdated(payload *types.Blog) error {
	fmt.Println("Blog updated event handled")
	return nil
}

func (h *HomeFeedEventsHandler) blogDeleted(payload *types.Blog) error {
	fmt.Println("Blog deleted event handled")
	return nil
}

func (h *HomeFeedEventsHandler) StartListener() {
	messages := h.homeFeedService.ListenToEvents(h.events)

	for message := range messages {
		payload := utils.UnmarshalBlog(message.Payload.(map[string]interface{}))
		switch message.Event {
		case constants.BLOG_CREATED:
			h.blogCreated(payload)
		case constants.BLOG_UPDATED:
			h.blogUpdated(payload)
		case constants.BLOG_DELETED:
			h.blogDeleted(payload)
		}
	}
}
